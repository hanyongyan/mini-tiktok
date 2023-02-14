FROM golang:1.19.3-alpine
WORKDIR /server
COPY .. .

# #构建后端和安装环境
RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o app main.go

EXPOSE 9294

CMD ./app
