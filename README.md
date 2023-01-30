
# Start
#### Preparation
Startup middleware
- nacos 
- jaeger
- opentelemetry
- redis

Startup necessary services
- ftp server(eg. [FileZilla](https://www.filezilla.cn/))
- nginx

#### Startup
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
