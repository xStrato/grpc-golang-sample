# grpc-golang-sample

### dependencies:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### To generate .proto files:
```
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```