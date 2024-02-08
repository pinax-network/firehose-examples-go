package main

import (
	"context"
	"encoding/hex"
	"fmt"
	pbcosmos "github.com/graphprotocol/proto-cosmos/pb/sf/cosmos/type/v1"
	pbantelope "github.com/pinax-network/firehose-antelope/types/pb/sf/antelope/type/v1"
	"github.com/streamingfast/dgrpc"
	pbbtc "github.com/streamingfast/firehose-bitcoin/pb/sf/bitcoin/type/v1"
	pbeth "github.com/streamingfast/firehose-ethereum/types/pb/sf/ethereum/type/v2"
	pbfirehose "github.com/streamingfast/pbgo/sf/firehose/v2"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"os"
)

const FirehoseETH = "eth.firehose.pinax.network:443"

func main() {
	apiKey := os.Getenv("SUBSTREAMS_API_KEY")
	if apiKey == "" {
		panic("SUBSTREAMS_API_KEY env variable must be set")
	}

	conn, err := dgrpc.NewExternalClient(FirehoseETH)
	if err != nil {
		log.Panicf("failed to create external gRPC client: %s", err)
	}
	defer conn.Close()

	client := pbfirehose.NewStreamClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-api-key", apiKey)

	blocks, err := client.Blocks(ctx, &pbfirehose.Request{
		StartBlockNum:   -1,
		StopBlockNum:    0,
		FinalBlocksOnly: false,
	})
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
		var _ pbbtc.Block
		var _ pbcosmos.Block

		err = block.Block.UnmarshalTo(&ethBlock)
		if err != nil {
			log.Panicf("failed to decode ETH block: %s", err)
		}

		fmt.Printf("received block: %d, blocktime: %s, hash: %s, trxs: %d\n",
			ethBlock.Number,
			ethBlock.Header.Timestamp.AsTime(),
			hex.EncodeToString(ethBlock.Hash),
			len(ethBlock.TransactionTraces),
		)
	}
}
