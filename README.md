# learn grpc by chatGPT

## 1. 安装protoc
```shell 
brew install protobuf
```

## 2. 安装protoc-gen-go
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

## 3. 编译proto文件
```shell
protoc --go_out=. --go-grpc_out=. *.proto
```

## 4. 编写服务端代码
```go
package main
```

## 5. 编写客户端代码
```go
package main
```

## 6. 运行服务端
```shell
go run main.go
```

## 7. 运行客户端
```shell
go run cmd/client/client.go
go run cmd/client/client.go zhangsan
```


## 附录——初始化本项目
```shell
go mod init server
```