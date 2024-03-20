package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pbantelope "buf.build/gen/go/pinax/firehose-antelope/protocolbuffers/go/sf/antelope/type/v1"
	pbarweave "buf.build/gen/go/pinax/firehose-arweave/protocolbuffers/go/sf/arweave/type/v1"
	pbbeacon "buf.build/gen/go/pinax/firehose-beacon/protocolbuffers/go/sf/beacon/type/v1"
	pbbtc "buf.build/gen/go/streamingfast/firehose-bitcoin/protocolbuffers/go/sf/bitcoin/type/v1"
	pbeth "buf.build/gen/go/streamingfast/firehose-ethereum/protocolbuffers/go/sf/ethereum/type/v2"
	pbnear "buf.build/gen/go/streamingfast/firehose-near/protocolbuffers/go/sf/near/type/v1"
	pbcosmos "github.com/graphprotocol/proto-cosmos/pb/sf/cosmos/type/v1"
	pbfirehose "github.com/streamingfast/pbgo/sf/firehose/v2"

	"github.com/mostynb/go-grpc-compression/zstd"
	"github.com/streamingfast/firehose-core/firehose/client"
	"google.golang.org/grpc"
)

const FirehoseBTC = "bitcoin.firehose.pinax.network:443"

func main() {
	apiKey := os.Getenv("SUBSTREAMS_API_KEY")
	if apiKey == "" {
		panic("SUBSTREAMS_API_KEY env variable must be set")
	}

	// Create a new Firehose stream client to connect to the infrastructure. The parameters set here are set for our
	// public endpoints.
	//
	// In case you are running a Firehose node yourself, you might want to set useInsecureTLSConnection or use
	// PlainTextConnection depending on whether you are using self-signed TLS certificates or non-TLS connections.
	fhClient, closeFunc, callOpts, err := client.NewFirehoseFetchClient(FirehoseBTC, "", apiKey, false, false)
	if err != nil {
		log.Panicf("failed to create Firehose client: %s", err)
	}
	defer closeFunc()

	// Optionally you can enable gRPC compression
	callOpts = append(callOpts, grpc.UseCompressor(zstd.Name))

	block, err := fhClient.Block(context.Background(), &pbfirehose.SingleBlockRequest{
		// Request a block by its block number
		Reference: &pbfirehose.SingleBlockRequest_BlockNumber_{
			BlockNumber: &pbfirehose.SingleBlockRequest_BlockNumber{Num: 800_000},
		},
		// Alternatively you can ensure a block hash additionally to the block number
		//Reference: &pbfirehose.SingleBlockRequest_BlockHashAndNumber_{
		//	BlockHashAndNumber: &pbfirehose.SingleBlockRequest_BlockHashAndNumber{
		//		Hash: "00000000000000000002a7c4c1e48d76c5a37902165a270156b7a8d72728a054",
		//		Num:  800_000,
		//	},
		//},
	}, callOpts...)
	if err != nil {
		log.Panicf("failed to fetch block: %s", err)
	}

	// Here we unmarshal the block payload from the generic block wrapper into the chain-specific block type.
	// Use one of the block types below depending on the endpoint you are connecting to. If you are retrieving
	// blocks from eth.firehose.pinax.network for example, then you need to unmarshal the payload into a pbeth.Block
	// here instead.

	var btcBlock pbbtc.Block
	var _ pbantelope.Block
	var _ pbarweave.Block
	var _ pbbeacon.Block
	var _ pbcosmos.Block
	var _ pbeth.Block
	var _ pbnear.Block

	err = block.Block.UnmarshalTo(&btcBlock)
	if err != nil {
		log.Panicf("failed to decode to Bitcoin block: %s", err)
	}

	fmt.Printf("received block: %d, blocktime: %s, hash: %s, trxs: %d\n",
		btcBlock.Height,
		time.Unix(btcBlock.Time, 0),
		btcBlock.Hash,
		len(btcBlock.Tx),
	)
}
