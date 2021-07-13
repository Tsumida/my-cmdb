INPUT_DIR=.
GO_OUT_DIR=./gen/go

protoc -I=$INPUT_DIR \
    --go_out $GO_OUT_DIR --go_opt paths=source_relative \
    --go-grpc_out $GO_OUT_DIR  --go-grpc_opt paths=source_relative \
    host.proto