package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/mostynb/go-grpc-compression/zstd"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"

	pbantelope "buf.build/gen/go/pinax/firehose-antelope/protocolbuffers/go/sf/antelope/type/v1"
	pbarweave "buf.build/gen/go/pinax/firehose-arweave/protocolbuffers/go/sf/arweave/type/v1"
	pbbeacon "buf.build/gen/go/pinax/firehose-beacon/protocolbuffers/go/sf/beacon/type/v1"
	pbbtc "buf.build/gen/go/streamingfast/firehose-bitcoin/protocolbuffers/go/sf/bitcoin/type/v1"
	pbeth "buf.build/gen/go/streamingfast/firehose-ethereum/protocolbuffers/go/sf/ethereum/type/v2"
	pbnear "buf.build/gen/go/streamingfast/firehose-near/protocolbuffers/go/sf/near/type/v1"
	pbcosmos "github.com/graphprotocol/proto-cosmos/pb/sf/cosmos/type/v1"
	pbfirehose "github.com/streamingfast/pbgo/sf/firehose/v2"

	"github.com/streamingfast/firehose-core/firehose/client"
)

const FirehoseETH = "eth.firehose.pinax.network:443"

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
	fhClient, closeFunc, callOpts, err := client.NewFirehoseClient(FirehoseETH, "", apiKey, false, false)
	if err != nil {
		log.Panicf("failed to create Firehose client: %s", err)
	}
	defer closeFunc()

	// Optionally you can enable gRPC compression
	callOpts = append(callOpts, grpc.UseCompressor(zstd.Name))

	// This sends the request to stream blocks, adapt the block parameters to your needs.
	blocks, err := fhClient.Blocks(context.Background(), &pbfirehose.Request{
		StartBlockNum:   -1,
		StopBlockNum:    0,
		FinalBlocksOnly: false,
		// Instead of sending a start and stop block number, you can also send the cursor you'll receive on each block
		// to continue exactly where you left off (this is useful for reconnects).
		// Cursor: "",
	}, callOpts...)
	if err != nil {
		log.Panicf("failed to stream blocks: %s", err)
	}

	for {
		block, err := blocks.Recv()
		if err == io.EOF {
			return
		} else if err != nil {
			log.Panicf("failed to receive block: %s", err)
		}

		// Here we unmarshal the block payload from the generic block wrapper into the chain-specific block type.
		// Use one of the block types below depending on the endpoint you are connecting to. If you are retrieving
		// blocks from bitcoin.firehose.pinax.network for example, then you need to unmarshal the payload into a
		// pbbtc.Block here instead.

		var ethBlock pbeth.Block
		var _ pbantelope.Block
		var _ pbarweave.Block
		var _ pbbtc.Block
		var _ pbbeacon.Block
		var _ pbcosmos.Block
		var _ pbnear.Block

		err = block.Block.UnmarshalTo(&ethBlock)
		if err != nil {
			log.Panicf("failed to decode ETH block: %s", err)
		}

		fmt.Printf("received block: %d, blocktime: %s, hash: %s, trxs: %d, cursor: %s\n",
			ethBlock.Number,
			ethBlock.Header.Timestamp.AsTime(),
			hex.EncodeToString(ethBlock.Hash),
			len(ethBlock.TransactionTraces),
			block.Cursor,
		)
	}
}
