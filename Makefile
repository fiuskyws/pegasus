main=cmd/main.go

run:
	# the '@' suppress the command echo
	@go run $(main)

proto:
	protoc --go_out=./src/ \
    --go-grpc_out=./src/ \
    .proto/pegasus.proto
