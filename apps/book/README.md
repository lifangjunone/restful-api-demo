## Book management system   
#### code generate command

```shell
protoc -I=. -I=./common/pb   --go_out=. --go_opt=module="github.com/lifangjunone/restful-api-demo" --go-grpc_out=. --go-grpc_opt=module="github.com/lifangjunone/restful-api-demo"  apps/*/pb/*proto  common/pb/*/*proto
```

#### generate validator code command

```shell
protoc-go-inject-tag  -input=apps/*/*.pb.go
```