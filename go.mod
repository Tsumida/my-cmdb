module github.com/Tsumida/my-cmdb

go 1.16

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	my-cmdb v0.0.1
)

replace my-cmdb => ./
