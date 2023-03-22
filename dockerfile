# 使用 golang 官方提供的 alpine 镜像作为基础镜像
FROM golang:alpine AS builder

# 将工作目录设置为 /app
WORKDIR /app

# 将本地文件拷贝到容器中
COPY . .

RUN go env -w GOPROXY=https://goproxy.io,direct
# 使用 go mod 下载依赖
RUN go mod tidy

# 编译 Go 程序
RUN go build -o main .

# 使用一个更小的基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 复制二进制文件到新的镜像中
COPY --from=builder /app/main .
RUN mkdir store

# 设置环境变量
ENV PORT=8080

# 暴露端口
EXPOSE $PORT

# 运行 Go 程序
CMD ["./main"]







