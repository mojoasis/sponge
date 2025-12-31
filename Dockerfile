# --- 第一阶段：编译 (Builder) ---
FROM golang:1.21-alpine AS builder

# 设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

WORKDIR /build

# 先拷贝依赖文件，利用 Docker 缓存层
COPY go.mod go.sum ./
RUN go mod download

# 拷贝全量源码
COPY . .

# 生成 wire 和 swagger (容器内构建不需要 air)
RUN go install github.com/google/wire/cmd/wire@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest && \
    wire ./cmd && \
    swag init -g cmd/main.go --parseDependency --parseInternal

# 编译
RUN go build -o sponge-app ./cmd/main.go ./cmd/wire_gen.go

# --- 第二阶段：运行 (Runner) ---
FROM alpine:latest

# 安装基础库（如时区数据）
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Shanghai

WORKDIR /app

# 从编译阶段拷贝成果物
COPY --from=builder /build/sponge-app .
COPY --from=builder /build/config.yaml .
# 如果有 docs 或者静态资源也需要拷贝
COPY --from=builder /build/docs ./docs

# 暴露端口 (对应你的 server.port)
EXPOSE 8080

# 启动程序
ENTRYPOINT ["./sponge-app"]