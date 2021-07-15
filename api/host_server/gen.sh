INPUT_DIR=.
DEPENDENCIES_PROTO_DIR=../proto_dependencies/
GO_OUT_DIR=./
OPENAPI_OUT_DIR=./openapiv2/

protoc \
    -I=$INPUT_DIR \
    -I=$DEPENDENCIES_PROTO_DIR \
    -I=$GOPATH/src \
    --go_out=$GO_OUT_DIR \
    --go_opt paths=source_relative \
    --go-grpc_out=$GO_OUT_DIR \
    --go-grpc_opt paths=source_relative \
    --grpc-gateway_out=$GO_OUT_DIR \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --openapiv2_out=$OPENAPI_OUT_DIR \
    --openapiv2_opt logtostderr=true \
    host.proto

echo "code generating done"