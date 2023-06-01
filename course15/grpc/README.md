```
https://colobu.com/2017/04/06/dive-into-gRPC-streaming/
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/hello.proto
    

protoc --go_out=. --go_opt=paths=source_relative \
    proto/hello.proto
    
protoc --java_out=. \
    --java-grpc_out=.  \
    proto/hello.proto
    
protoc --plugin=protoc-gen-grpc-java=/Users/gaozhe/IdeaProjects/open-api/target/protoc-plugins/protoc-gen-grpc-java-1.19.0-osx-x86_64.exe \
  --java_out=. \
  --grpc-java_out=.  \
   proto/hello.proto
```