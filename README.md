install golang
install vscode

install protobuf
$ sudo apt install -y protobuf-compiler
$ protoc --version

install the protocol compiler plugins for Go
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

$ go mod init github.com/alexandreps1123/grpc-go

generate gRPC code
$ protoc -Ihelloworld/proto --go_out=. --go_opt=module=github.com/alexandreps1123/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/alexandreps1123/grpc-go helloworld/proto/helloworld.proto

go build server
$ go build -o bin/helloworld/server ./helloworld/server

go build client
$ go build -o bin/helloworld/client ./helloworld/client
