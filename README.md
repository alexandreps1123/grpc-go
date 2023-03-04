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
$ protoc -Isum/proto --go_out=. --go_opt=module=github.com/alexandreps1123/grpc-go --go-grpc_out=. --go-grpc_opt=module=github.com/alexandreps1123/grpc-go sum/proto/helloworld.proto