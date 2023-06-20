go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

apply evans CLI https://github.com/ktr0731/evans#from-github-releases
-> cmd: evans --host localhost --port 50051 --reflection repl