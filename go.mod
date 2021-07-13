module github.com/Tsumida/my-cmdb

go 1.16

require (
	github.com/google/go-cmp v0.5.6 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1

	my-cmdb v0.0.1
)

replace my-cmdb => ./
