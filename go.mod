module firehose-example-go

go 1.22

toolchain go1.22.0

require (
	buf.build/gen/go/pinax/firehose-antelope/protocolbuffers/go v1.32.0-20240129214250-3af26456175d.1
	buf.build/gen/go/pinax/firehose-arweave/protocolbuffers/go v1.32.0-20240129212333-eeea46c6211b.1
	buf.build/gen/go/pinax/firehose-beacon/protocolbuffers/go v1.32.0-20240222164713-07555fcd06be.1
	buf.build/gen/go/streamingfast/firehose-bitcoin/protocolbuffers/go v1.32.0-20231205205020-0d8ce32fe714.1
	buf.build/gen/go/streamingfast/firehose-ethereum/protocolbuffers/go v1.32.0-20240117171201-d869cb39aae9.1
	buf.build/gen/go/streamingfast/firehose-near/protocolbuffers/go v1.32.0-20230712201405-0b7e4efe1b9f.1
	github.com/graphprotocol/proto-cosmos v0.1.4
	github.com/mostynb/go-grpc-compression v1.2.2
	github.com/streamingfast/firehose-core v1.2.1
	github.com/streamingfast/pbgo v0.0.6-0.20240131193313-6b88bc7139db
	google.golang.org/grpc v1.62.0
)

require (
	cloud.google.com/go/compute v1.24.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blendle/zapdriver v1.3.2-0.20200203083823-9200777f8a3d // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cncf/xds/go v0.0.0-20231128003011-0fa0005c9caa // indirect
	github.com/cosmos/cosmos-proto v1.0.0-beta.4 // indirect
	github.com/envoyproxy/go-control-plane v0.12.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/prometheus/client_golang v1.18.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/sercand/kuberesolver/v5 v5.1.1 // indirect
	github.com/streamingfast/dgrpc v0.0.0-20240222213940-b9f324ff4d5c // indirect
	github.com/streamingfast/logging v0.0.0-20230608130331-f22c91403091 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.48.0 // indirect
	go.opentelemetry.io/otel v1.23.1 // indirect
	go.opentelemetry.io/otel/metric v1.23.1 // indirect
	go.opentelemetry.io/otel/trace v1.23.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/oauth2 v0.16.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/term v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240221002015-b0ce06bbee7c // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/streamingfast/firehose-core => ../firehose-core
