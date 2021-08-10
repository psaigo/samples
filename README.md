# samples

## demo-grpc
Basic working client and server in Go that leverages GRPC.

### Instructions:

#### Generate the ~pb.go:
cd $GOPATH/src/github.com/psaigo/samples/demo-grpc/pb

run on terminal: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/demo.proto

#### run the server:
cd $GOPATH/src/github.com/psaigo/samples/demo-grpc/server
go run main.go

#### run the client:
cd $GOPATH/src/github.com/psaigo/samples/demo-grpc/client
go run main.go
