package main

import (
	"context"
	"fmt"
	pbantelope "github.com/pinax-network/firehose-antelope/types/pb/sf/antelope/type/v1"
	"github.com/streamingfast/dgrpc"
	pbfirehose "github.com/streamingfast/pbgo/sf/firehose/v2"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"
	"io"
	"log"
	"os"
)

const endpoint = "eos.firehose.pinax.network:443"

func main() {
	jwt := os.Getenv("SUBSTREAMS_API_TOKEN")
	if jwt == "" {
		panic("SUBSTREAMS_API_TOKEN env variable must be set")
	}

	conn, err := dgrpc.NewExternalClient(endpoint, grpc.WithPerRPCCredentials(oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: jwt})}))
	if err != nil {
		log.Panicf("failed to create external gRPC client: %s", err)
	}
	defer conn.Close()

	client := pbfirehose.NewStreamClient(conn)
	blocks, err := client.Blocks(context.Background(), &pbfirehose.Request{
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

		var antelopeBlock pbantelope.Block
		err = block.Block.UnmarshalTo(&antelopeBlock)
		if err != nil {
			log.Panicf("failed to decode to Antelope block: %s", err)
		}

		fmt.Printf("received block: %d\n", antelopeBlock.Number)
	}
}
