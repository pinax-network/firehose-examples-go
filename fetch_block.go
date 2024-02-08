package main

import (
	"context"
	"fmt"
	pbcosmos "github.com/graphprotocol/proto-cosmos/pb/sf/cosmos/type/v1"
	pbantelope "github.com/pinax-network/firehose-antelope/types/pb/sf/antelope/type/v1"
	"github.com/streamingfast/dgrpc"
	pbbtc "github.com/streamingfast/firehose-bitcoin/pb/sf/bitcoin/type/v1"
	pbeth "github.com/streamingfast/firehose-ethereum/types/pb/sf/ethereum/type/v2"
	pbfirehose "github.com/streamingfast/pbgo/sf/firehose/v2"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"time"
)

const FirehoseBTC = "bitcoin.firehose.pinax.network:443"

func main() {
	apiKey := os.Getenv("SUBSTREAMS_API_KEY")
	if apiKey == "" {
		panic("SUBSTREAMS_API_KEY env variable must be set")
	}

	conn, err := dgrpc.NewExternalClient(FirehoseBTC)
	if err != nil {
		log.Panicf("failed to create external gRPC client: %s", err)
	}
	defer conn.Close()

	client := pbfirehose.NewFetchClient(conn)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "x-api-key", apiKey)

	block, err := client.Block(ctx, &pbfirehose.SingleBlockRequest{
		Reference: &pbfirehose.SingleBlockRequest_BlockNumber_{
			BlockNumber: &pbfirehose.SingleBlockRequest_BlockNumber{Num: 800_000},
		},
	})
	if err != nil {
		log.Panicf("failed to fetch block: %s", err)
	}

	// Here we unmarshal the block payload from the generic block wrapper into the chain-specific block type.
	// Use one of the block types below depending on the endpoint you are connecting to. If you are retrieving
	// blocks from eth.firehose.pinax.network for example, then you need to unmarshal the payload into a pbeth.Block
	// here instead.

	var btcBlock pbbtc.Block
	var _ pbantelope.Block
	var _ pbcosmos.Block
	var _ pbeth.Block

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
