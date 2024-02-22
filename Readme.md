## Firehose Client Examples (Golang)

This is an exemplary Golang application on how to stream or fetch blocks from Firehose and transform the chain agnostic
block wrapper into a chain-specific block type.

### Prerequisite

Before you can connect to our infrastructure, you need to sign up on https://app.pinax.network and create an API key for
yourself. You will find all Firehose endpoints on the Pinax App as well. We offer a generous free tier plan which should
be sufficient for most use cases.

### Fetch Firehose Block

To fetch a single block from firehose just run the example in `fetch_block.go`:

```bash
$ export SUBSTREAMS_API_KEY=<your_api_key>
$ go run fetch_block.go
received block: 800000, blocktime: 2023-07-24 05:17:09 +0200 CEST, hash: 00000000000000000002a7c4c1e48d76c5a37902165a270156b7a8d72728a054, trxs: 3721
```

### Stream Firehose Blocks

To stream blocks from a configurable start block, run:

```bash
$ export SUBSTREAMS_API_KEY=<your_api_key>
$ go run stream_blocks_eth.go
received block: 19183642, blocktime: 2024-02-08 13:07:23 +0000 UTC, hash: bbd1f26f3c68458e8d764d6a9a381258cb4f3b5defaadcf0c5e447c76e3f5026, trxs: 143
received block: 19183643, blocktime: 2024-02-08 13:07:35 +0000 UTC, hash: b7a587953892fe5042be072a4e8ffeceab77c25f0e27f043eaa2464dc8f3f780, trxs: 92
received block: 19183644, blocktime: 2024-02-08 13:07:47 +0000 UTC, hash: 23c3ccfd2f591590ed4838efac0d9285be36cdf87e5290515dee5f0c3bec2d22, trxs: 124
received block: 19183645, blocktime: 2024-02-08 13:07:59 +0000 UTC, hash: 4b2491a32ceeb470628de53ef7d66faf16afc77f2b03b1b5188bac6f8791ee99, trxs: 137
```

### Block Definitions

To consume other chains, you will need their respective block definitions. Note that chains that are based on the same
blockchain technology will not provide dedicated block definitions (for example, the BSC firehose endpoints do not come
with their own types but will use the Ethereum ones).

| Blockchain Technology | Go Module                                                             |
|-----------------------|-----------------------------------------------------------------------|
| Antelope              | `buf.build/gen/go/pinax/firehose-antelope/protocolbuffers/go`         |
| Arweave               | `buf.build/gen/go/pinax/firehose-arweave/protocolbuffers/go`          |
| Beacon                | `buf.build/gen/go/pinax/firehose-beacon/protocolbuffers/go`           |
| Bitcoin               | `buf.build/gen/go/streamingfast/firehose-bitcoin/protocolbuffers/go`  |
| Cosmos                | `github.com/graphprotocol/proto-cosmos`                               |
| Ethereum              | `buf.build/gen/go/streamingfast/firehose-ethereum/protocolbuffers/go` |
| Near                  | `buf.build/gen/go/streamingfast/firehose-near/protocolbuffers/go`     |
