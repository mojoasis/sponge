.PHONY: all build wire swag run clean help

# 变量定义
BINARY_NAME=main
BINARY_DIR=./tmp
MAIN_FILE=./cmd/main.go
WIRE_GEN_FILE=./cmd/wire_gen.go

# 默认目标
all: swag wire build

## wire: 重新生成依赖注入代码
wire:
	@echo "==> Generating wire_gen.go..."
	@wire ./cmd

## swag: 重新生成 Swagger 文档
swag:
	@echo "==> Generating swagger docs..."
	@swag init -g $(MAIN_FILE)

## build: 编译项目
build:
	@echo "==> Building binary..."
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_FILE) $(WIRE_GEN_FILE)

## run: 直接运行项目
run: all
	@$(BINARY_DIR)/$(BINARY_NAME)

## clean: 清理临时文件
clean:
	@echo "==> Cleaning..."
	@if [ -d "$(BINARY_DIR)" ]; then rm -rf $(BINARY_DIR); fi
	@echo "Clean done."

## help: 显示帮助信息
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

## docker-build: 构建 Docker 镜像
docker-build:
	docker build -t sponge-service .

## deploy: 一键部署脚本执行
deploy:
	chmod +x deploy.sh
	./deploy.sh
