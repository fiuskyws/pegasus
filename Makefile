main=cmd/main.go

test:
# the '@' suppress the command echo
	@go test ./...

test-v:
# the '@' suppress the command echo
	@go test -v ./...

run:
# the '@' suppress the command echo
	@go run $(main)

proto:
	protoc --go_out=./src/ \
    --go-grpc_out=./src/ \
    .proto/pegasus.proto
