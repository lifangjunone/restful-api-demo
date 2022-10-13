## Book management system   
#### code generate commaon

```shell
protoc -I=. -I=./common/pb   --go_out=. --go_opt=module="restful-api-demo" --go-grpc_out=. --go-grpc_opt=module="restful-api-demo"  apps/*/pb/*proto
```