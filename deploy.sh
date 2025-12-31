#!/bin/bash

# 配置变量
APP_NAME="sponge-service"
VERSION=$(date +%Y%m%d%H%M%S)
PORT=8080

echo "开始部署项目: $APP_NAME ..."

# 1. 停止并删除旧容器
echo "清理旧容器..."
docker stop $APP_NAME 2>/dev/null
docker rm $APP_NAME 2>/dev/null

# 2. 构建镜像
echo "正在构建 Docker 镜像..."
docker build -t $APP_NAME:$VERSION .
docker tag $APP_NAME:$VERSION $APP_NAME:latest

# 3. 运行新容器
echo "启动容器..."
# 这里挂载本地 config.yaml 方便修改后重启生效，不需要重新打镜像
docker run -d \
  --name $APP_NAME \
  -p $PORT:$PORT \
  -v $(pwd)/config.yaml:/app/config.yaml \
  --restart always \
  $APP_NAME:latest

echo "部署完成！"
echo "API 地址: http://localhost:$PORT/swagger/index.html"