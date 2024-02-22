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
	"github.com/streamingfast/dgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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
	}, grpc.UseCompressor(zstd.Name))
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
