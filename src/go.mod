module github.com/GetElastech/flow-observer/m/v2

go 1.16

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/libp2p/go-libp2p-core v0.11.0
	github.com/libp2p/go-libp2p-pubsub v0.6.0
	github.com/onflow/flow-go v0.18.0
	github.com/onflow/flow-go/crypto v0.24.3
	github.com/onflow/flow/protobuf/go/flow v0.2.5
	github.com/rs/zerolog v1.19.0
	github.com/spf13/pflag v1.0.5
	google.golang.org/grpc v1.44.0
)

replace github.com/onflow/flow-go => ../dep

replace mellium.im/sasl => github.com/mellium/sasl v0.2.1
