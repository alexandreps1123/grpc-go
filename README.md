install golang
install vscode

install protobuf
$ sudo apt install -y protobuf-compiler
$ protoc --version

install the protocol compiler plugins for Go
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2