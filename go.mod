module github.com/walteh/cloudstack-proxy

go 1.24.2

replace github.com/bufbuild/protovalidate-go => github.com/walteh/protovalidate-go v0.0.0-20250410121243-8ed37c2d5a44

require (
	connectrpc.com/connect v1.18.1
	github.com/google/cel-go v0.24.1
	google.golang.org/grpc v1.71.1
	google.golang.org/protobuf v1.36.6
)

require (
	cel.dev/expr v0.19.1 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20240325151524-a685a6edb6d8 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
