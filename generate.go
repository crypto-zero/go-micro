package micro

//go:generate ./.github/generate.sh
//go:generate protoc --go-grpc_out=paths=source_relative:. transport/grpc/proto/transport.proto
