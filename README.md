https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/
https://www.docker.com/blog/containerize-your-go-developer-environment-part-2/
https://www.docker.com/blog/tag/go-env-series/

https://medium.com/vptech/complexity-is-the-bane-of-every-software-engineer-e2878d0ad45a
https://github.com/abronan/todo-grpc/blob

brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc -I=. --proto_path=. --go_out=. ./todo.proto
