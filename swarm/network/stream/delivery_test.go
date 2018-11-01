// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package stream

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/ShyftNetwork/go-empyrean/node"
	"github.com/ShyftNetwork/go-empyrean/p2p"
	"github.com/ShyftNetwork/go-empyrean/p2p/simulations/adapters"
	p2ptest "github.com/ShyftNetwork/go-empyrean/p2p/testing"
	"github.com/ShyftNetwork/go-empyrean/swarm/log"
	"github.com/ShyftNetwork/go-empyrean/swarm/network"
	"github.com/ShyftNetwork/go-empyrean/swarm/network/simulation"
	"github.com/ShyftNetwork/go-empyrean/swarm/state"
	"github.com/ShyftNetwork/go-empyrean/swarm/storage"
)

func TestStreamerRetrieveRequest(t *testing.T) {
	tester, streamer, _, teardown, err := newStreamerTester(t, nil)
	defer teardown()
	if err != nil {
		t.Fatal(err)
	}

	node := tester.Nodes[0]

	ctx := context.Background()
	req := network.NewRequest(
		storage.Address(hash0[:]),
		true,
		&sync.Map{},
	)
	streamer.delivery.RequestFromPeers(ctx, req)

	err = tester.TestExchanges(p2ptest.Exchange{
		Label: "RetrieveRequestMsg",
		Expects: []p2ptest.Expect{
			{
				Code: 5,
				Msg: &RetrieveRequestMsg{
					Addr:      hash0[:],
					SkipCheck: true,
				},
				Peer: node.ID(),
			},
		},
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestStreamerUpstreamRetrieveRequestMsgExchangeWithoutStore(t *testing.T) {
	tester, streamer, _, teardown, err := newStreamerTester(t, &RegistryOptions{
		DoServeRetrieve: true,
	})
	defer teardown()
	if err != nil {
		t.Fatal(err)
	}

	node := tester.Nodes[0]

	chunk := storage.NewChunk(storage.Address(hash0[:]), nil)

	peer := streamer.getPeer(node.ID())

	peer.handleSubscribeMsg(context.TODO(), &SubscribeMsg{
		Stream:   NewStream(swarmChunkServerStreamName, "", true),
		History:  nil,
		Priority: Top,
	})

	err = tester.TestExchanges(p2ptest.Exchange{
		Label: "RetrieveRequestMsg",
		Triggers: []p2ptest.Trigger{
			{
				Code: 5,
				Msg: &RetrieveRequestMsg{
					Addr: chunk.Address()[:],
				},
				Peer: node.ID(),
			},
		},
		Expects: []p2ptest.Expect{
			{
				Code: 1,
				Msg: &OfferedHashesMsg{
					HandoverProof: nil,
					Hashes:        nil,
					From:          0,
					To:            0,
				},
				Peer: node.ID(),
			},
		},
	})

	expectedError := `exchange #0 "RetrieveRequestMsg": timed out`
	if err == nil || err.Error() != expectedError {
		t.Fatalf("Expected error %v, got %v", expectedError, err)
	}
}

// upstream request server receives a retrieve Request and responds with
// offered hashes or delivery if skipHash is set to true
func TestStreamerUpstreamRetrieveRequestMsgExchange(t *testing.T) {
	tester, streamer, localStore, teardown, err := newStreamerTester(t, &RegistryOptions{
		DoServeRetrieve: true,
	})
	defer teardown()
	if err != nil {
		t.Fatal(err)
	}

	node := tester.Nodes[0]
	peer := streamer.getPeer(node.ID())

	stream := NewStream(swarmChunkServerStreamName, "", true)

	peer.handleSubscribeMsg(context.TODO(), &SubscribeMsg{
		Stream:   stream,
		History:  nil,
		Priority: Top,
	})

	hash := storage.Address(hash0[:])
	chunk := storage.NewChunk(hash, hash)
	err = localStore.Put(context.TODO(), chunk)
	if err != nil {
		t.Fatalf("Expected no err got %v", err)
	}

	err = tester.TestExchanges(p2ptest.Exchange{
		Label: "RetrieveRequestMsg",
		Triggers: []p2ptest.Trigger{
			{
				Code: 5,
				Msg: &RetrieveRequestMsg{
					Addr: hash,
				},
				Peer: node.ID(),
			},
		},
		Expects: []p2ptest.Expect{
			{
				Code: 1,
				Msg: &OfferedHashesMsg{
					HandoverProof: &HandoverProof{
						Handover: &Handover{},
					},
					Hashes: hash,
					From:   0,
					// TODO: why is this 32???
					To:     32,
					Stream: stream,
				},
				Peer: node.ID(),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	hash = storage.Address(hash1[:])
	chunk = storage.NewChunk(hash, hash1[:])
	err = localStore.Put(context.TODO(), chunk)
	if err != nil {
		t.Fatalf("Expected no err got %v", err)
	}

	err = tester.TestExchanges(p2ptest.Exchange{
		Label: "RetrieveRequestMsg",
		Triggers: []p2ptest.Trigger{
			{
				Code: 5,
				Msg: &RetrieveRequestMsg{
					Addr:      hash,
					SkipCheck: true,
				},
				Peer: node.ID(),
			},
		},
		Expects: []p2ptest.Expect{
			{
				Code: 6,
				Msg: &ChunkDeliveryMsg{
					Addr:  hash,
					SData: hash,
				},
				Peer: node.ID(),
			},
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamerDownstreamChunkDeliveryMsgExchange(t *testing.T) {
	tester, streamer, localStore, teardown, err := newStreamerTester(t, &RegistryOptions{
		DoServeRetrieve: true,
	})
	defer teardown()
	if err != nil {
		t.Fatal(err)
	}

	streamer.RegisterClientFunc("foo", func(p *Peer, t string, live bool) (Client, error) {
		return &testClient{
			t: t,
		}, nil
	})

	node := tester.Nodes[0]

	stream := NewStream("foo", "", true)
	err = streamer.Subscribe(node.ID(), stream, NewRange(5, 8), Top)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	chunkKey := hash0[:]
	chunkData := hash1[:]

	err = tester.TestExchanges(p2ptest.Exchange{
		Label: "Subscribe message",
		Expects: []p2ptest.Expect{
			{
				Code: 4,
				Msg: &SubscribeMsg{
					Stream:   stream,
					History:  NewRange(5, 8),
					Priority: Top,
				},
				Peer: node.ID(),
			},
		},
	},
		p2ptest.Exchange{
			Label: "ChunkDelivery message",
			Triggers: []p2ptest.Trigger{
				{
					Code: 6,
					Msg: &ChunkDeliveryMsg{
						Addr:  chunkKey,
						SData: chunkData,
					},
					Peer: node.ID(),
				},
			},
		})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// wait for the chunk to get stored
	storedChunk, err := localStore.Get(ctx, chunkKey)
	for err != nil {
		select {
		case <-ctx.Done():
			t.Fatalf("Chunk is not in localstore after timeout, err: %v", err)
		default:
		}
		storedChunk, err = localStore.Get(ctx, chunkKey)
		time.Sleep(50 * time.Millisecond)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !bytes.Equal(storedChunk.Data(), chunkData) {
		t.Fatal("Retrieved chunk has different data than original")
	}

}

func TestDeliveryFromNodes(t *testing.T) {
	testDeliveryFromNodes(t, 2, 1, dataChunkCount, true)
	testDeliveryFromNodes(t, 2, 1, dataChunkCount, false)
	testDeliveryFromNodes(t, 4, 1, dataChunkCount, true)
	testDeliveryFromNodes(t, 4, 1, dataChunkCount, false)
	testDeliveryFromNodes(t, 8, 1, dataChunkCount, true)
	testDeliveryFromNodes(t, 8, 1, dataChunkCount, false)
	testDeliveryFromNodes(t, 16, 1, dataChunkCount, true)
	testDeliveryFromNodes(t, 16, 1, dataChunkCount, false)
}

func testDeliveryFromNodes(t *testing.T, nodes, conns, chunkCount int, skipCheck bool) {
	sim := simulation.New(map[string]simulation.ServiceFunc{
		"streamer": func(ctx *adapters.ServiceContext, bucket *sync.Map) (s node.Service, cleanup func(), err error) {
			node := ctx.Config.Node()
			addr := network.NewAddr(node)
			store, datadir, err := createTestLocalStorageForID(node.ID(), addr)
			if err != nil {
				return nil, nil, err
			}
			bucket.Store(bucketKeyStore, store)
			cleanup = func() {
				os.RemoveAll(datadir)
				store.Close()
			}
			localStore := store.(*storage.LocalStore)
			netStore, err := storage.NewNetStore(localStore, nil)
			if err != nil {
				return nil, nil, err
			}

			kad := network.NewKademlia(addr.Over(), network.NewKadParams())
			delivery := NewDelivery(kad, netStore)
			netStore.NewNetFetcherFunc = network.NewFetcherFactory(delivery.RequestFromPeers, true).New

			r := NewRegistry(addr.ID(), delivery, netStore, state.NewInmemoryStore(), &RegistryOptions{
				SkipCheck:       skipCheck,
				DoServeRetrieve: true,
			})
			bucket.Store(bucketKeyRegistry, r)

			fileStore := storage.NewFileStore(netStore, storage.NewFileStoreParams())
			bucket.Store(bucketKeyFileStore, fileStore)

			return r, cleanup, nil

		},
	})
	defer sim.Close()

	log.Info("Adding nodes to simulation")
	_, err := sim.AddNodesAndConnectChain(nodes)
	if err != nil {
		t.Fatal(err)
	}

	log.Info("Starting simulation")
	ctx := context.Background()
	result := sim.Run(ctx, func(ctx context.Context, sim *simulation.Simulation) error {
		nodeIDs := sim.UpNodeIDs()
		//determine the pivot node to be the first node of the simulation
		sim.SetPivotNode(nodeIDs[0])
		//distribute chunks of a random file into Stores of nodes 1 to nodes
		//we will do this by creating a file store with an underlying round-robin store:
		//the file store will create a hash for the uploaded file, but every chunk will be
		//distributed to different nodes via round-robin scheduling
		log.Debug("Writing file to round-robin file store")
		//to do this, we create an array for chunkstores (length minus one, the pivot node)
		stores := make([]storage.ChunkStore, len(nodeIDs)-1)
		//we then need to get all stores from the sim....
		lStores := sim.NodesItems(bucketKeyStore)
		i := 0
		//...iterate the buckets...
		for id, bucketVal := range lStores {
			//...and remove the one which is the pivot node
			if id == *sim.PivotNodeID() {
				continue
			}
			//the other ones are added to the array...
			stores[i] = bucketVal.(storage.ChunkStore)
			i++
		}
		//...which then gets passed to the round-robin file store
		roundRobinFileStore := storage.NewFileStore(newRoundRobinStore(stores...), storage.NewFileStoreParams())
		//now we can actually upload a (random) file to the round-robin store
		size := chunkCount * chunkSize
		log.Debug("Storing data to file store")
		fileHash, wait, err := roundRobinFileStore.Store(ctx, io.LimitReader(crand.Reader, int64(size)), int64(size), false)
		// wait until all chunks stored
		if err != nil {
			return err
		}
		err = wait(ctx)
		if err != nil {
			return err
		}

		log.Debug("Waiting for kademlia")
		if _, err := sim.WaitTillHealthy(ctx, 2); err != nil {
			return err
		}

		//each of the nodes (except pivot node) subscribes to the stream of the next node
		for j, node := range nodeIDs[0 : nodes-1] {
			sid := nodeIDs[j+1]
			item, ok := sim.NodeItem(node, bucketKeyRegistry)
			if !ok {
				return fmt.Errorf("No registry")
			}
			registry := item.(*Registry)
			err = registry.Subscribe(sid, NewStream(swarmChunkServerStreamName, "", true), nil, Top)
			if err != nil {
				return err
			}
		}

		//get the pivot node's filestore
		item, ok := sim.NodeItem(*sim.PivotNodeID(), bucketKeyFileStore)
		if !ok {
			return fmt.Errorf("No filestore")
		}
		pivotFileStore := item.(*storage.FileStore)
		log.Debug("Starting retrieval routine")
		go func() {
			// start the retrieval on the pivot node - this will spawn retrieve requests for missing chunks
			// we must wait for the peer connections to have started before requesting
			n, err := readAll(pivotFileStore, fileHash)
			log.Info(fmt.Sprintf("retrieved %v", fileHash), "read", n, "err", err)
			if err != nil {
				t.Fatalf("requesting chunks action error: %v", err)
			}
		}()

		log.Debug("Watching for disconnections")
		disconnections := sim.PeerEvents(
			context.Background(),
			sim.NodeIDs(),
			simulation.NewPeerEventsFilter().Type(p2p.PeerEventTypeDrop),
		)

		go func() {
			for d := range disconnections {
				if d.Error != nil {
					log.Error("peer drop", "node", d.NodeID, "peer", d.Event.Peer)
					t.Fatal(d.Error)
				}
			}
		}()

		//finally check that the pivot node gets all chunks via the root hash
		log.Debug("Check retrieval")
		success := true
		var total int64
		total, err = readAll(pivotFileStore, fileHash)
		if err != nil {
			return err
		}
		log.Info(fmt.Sprintf("check if %08x is available locally: number of bytes read %v/%v (error: %v)", fileHash, total, size, err))
		if err != nil || total != int64(size) {
			success = false
		}

		if !success {
			return fmt.Errorf("Test failed, chunks not available on all nodes")
		}
		log.Debug("Test terminated successfully")
		return nil
	})
	if result.Error != nil {
		t.Fatal(result.Error)
	}
}

func BenchmarkDeliveryFromNodesWithoutCheck(b *testing.B) {
	for chunks := 32; chunks <= 128; chunks *= 2 {
		for i := 2; i < 32; i *= 2 {
			b.Run(
				fmt.Sprintf("nodes=%v,chunks=%v", i, chunks),
				func(b *testing.B) {
					benchmarkDeliveryFromNodes(b, i, 1, chunks, true)
				},
			)
		}
	}
}

func BenchmarkDeliveryFromNodesWithCheck(b *testing.B) {
	for chunks := 32; chunks <= 128; chunks *= 2 {
		for i := 2; i < 32; i *= 2 {
			b.Run(
				fmt.Sprintf("nodes=%v,chunks=%v", i, chunks),
				func(b *testing.B) {
					benchmarkDeliveryFromNodes(b, i, 1, chunks, false)
				},
			)
		}
	}
}

func benchmarkDeliveryFromNodes(b *testing.B, nodes, conns, chunkCount int, skipCheck bool) {
	sim := simulation.New(map[string]simulation.ServiceFunc{
		"streamer": func(ctx *adapters.ServiceContext, bucket *sync.Map) (s node.Service, cleanup func(), err error) {
			node := ctx.Config.Node()
			addr := network.NewAddr(node)
			store, datadir, err := createTestLocalStorageForID(node.ID(), addr)
			if err != nil {
				return nil, nil, err
			}
			bucket.Store(bucketKeyStore, store)
			cleanup = func() {
				os.RemoveAll(datadir)
				store.Close()
			}
			localStore := store.(*storage.LocalStore)
			netStore, err := storage.NewNetStore(localStore, nil)
			if err != nil {
				return nil, nil, err
			}
			kad := network.NewKademlia(addr.Over(), network.NewKadParams())
			delivery := NewDelivery(kad, netStore)
			netStore.NewNetFetcherFunc = network.NewFetcherFactory(delivery.RequestFromPeers, true).New

			r := NewRegistry(addr.ID(), delivery, netStore, state.NewInmemoryStore(), &RegistryOptions{
				SkipCheck:       skipCheck,
				DoSync:          true,
				SyncUpdateDelay: 0,
			})

			fileStore := storage.NewFileStore(netStore, storage.NewFileStoreParams())
			bucket.Store(bucketKeyFileStore, fileStore)

			return r, cleanup, nil

		},
	})
	defer sim.Close()

	log.Info("Initializing test config")
	_, err := sim.AddNodesAndConnectChain(nodes)
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()
	result := sim.Run(ctx, func(ctx context.Context, sim *simulation.Simulation) error {
		nodeIDs := sim.UpNodeIDs()
		node := nodeIDs[len(nodeIDs)-1]

		item, ok := sim.NodeItem(node, bucketKeyFileStore)
		if !ok {
			b.Fatal("No filestore")
		}
		remoteFileStore := item.(*storage.FileStore)

		pivotNode := nodeIDs[0]
		item, ok = sim.NodeItem(pivotNode, bucketKeyNetStore)
		if !ok {
			b.Fatal("No filestore")
		}
		netStore := item.(*storage.NetStore)

		if _, err := sim.WaitTillHealthy(ctx, 2); err != nil {
			return err
		}

		disconnections := sim.PeerEvents(
			context.Background(),
			sim.NodeIDs(),
			simulation.NewPeerEventsFilter().Type(p2p.PeerEventTypeDrop),
		)

		go func() {
			for d := range disconnections {
				if d.Error != nil {
					log.Error("peer drop", "node", d.NodeID, "peer", d.Event.Peer)
					b.Fatal(d.Error)
				}
			}
		}()
		// benchmark loop
		b.ResetTimer()
		b.StopTimer()
	Loop:
		for i := 0; i < b.N; i++ {
			// uploading chunkCount random chunks to the last node
			hashes := make([]storage.Address, chunkCount)
			for i := 0; i < chunkCount; i++ {
				// create actual size real chunks
				ctx := context.TODO()
				hash, wait, err := remoteFileStore.Store(ctx, io.LimitReader(crand.Reader, int64(chunkSize)), int64(chunkSize), false)
				if err != nil {
					b.Fatalf("expected no error. got %v", err)
				}
				// wait until all chunks stored
				err = wait(ctx)
				if err != nil {
					b.Fatalf("expected no error. got %v", err)
				}
				// collect the hashes
				hashes[i] = hash
			}
			// now benchmark the actual retrieval
			// netstore.Get is called for each hash in a go routine and errors are collected
			b.StartTimer()
			errs := make(chan error)
			for _, hash := range hashes {
				go func(h storage.Address) {
					_, err := netStore.Get(ctx, h)
					log.Warn("test check netstore get", "hash", h, "err", err)
					errs <- err
				}(hash)
			}
			// count and report retrieval errors
			// if there are misses then chunk timeout is too low for the distance and volume (?)
			var total, misses int
			for err := range errs {
				if err != nil {
					log.Warn(err.Error())
					misses++
				}
				total++
				if total == chunkCount {
					break
				}
			}
			b.StopTimer()

			if misses > 0 {
				err = fmt.Errorf("%v chunk not found out of %v", misses, total)
				break Loop
			}
		}
		if err != nil {
			b.Fatal(err)
		}
		return nil
	})
	if result.Error != nil {
		b.Fatal(result.Error)
	}

}
