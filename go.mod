module github.com/shynome/v2ray-apiproxy

go 1.12

require (
	github.com/golang/protobuf v1.3.1
	github.com/google/uuid v1.1.1
	github.com/sirupsen/logrus v1.4.1 // indirect
	golang.org/x/net v0.0.0-20190206173232-65e2d4e15006
	golang.org/x/tools v0.0.0-20180928181343-b3c0be4c978b // indirect
	google.golang.org/grpc v1.18.0
	v2ray.com/core v4.14.2+incompatible
)

replace v2ray.com/core v4.14.2+incompatible => github.com/shynome/v2ray-go-grpc-sdk v0.0.0-20190213051928-a726c184a649
