
# Start
Preparation
Startup nacos and opentelemetry
then
```shell
go mod download
```

User Service
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

Video Service
```shell
cd cmd/video
sh build.sh
sh output/bootstrap.sh
```

Gateway
```shell
cd cmd/api
go run .
```
