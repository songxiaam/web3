#!/bin/bash

# Smart Route 数据库启动脚本

echo "启动 Smart Route 数据库服务..."

# 检查 Docker 是否运行
if ! docker info > /dev/null 2>&1; then
    echo "错误: Docker 未运行，请先启动 Docker"
    exit 1
fi

# 启动服务
echo "启动 PostgreSQL 和 Redis..."
docker-compose up -d

# 等待服务启动
echo "等待服务启动..."
sleep 10

# 检查服务状态
echo "检查服务状态..."
docker-compose ps

# 检查数据库连接
echo "检查数据库连接..."
docker-compose exec postgres pg_isready -U smartroute

if [ $? -eq 0 ]; then
    echo "✅ PostgreSQL 启动成功"
    echo "数据库连接信息:"
    echo "  Host: localhost"
    echo "  Port: 5433"
    echo "  Database: smartroute"
    echo "  Username: smartroute"
    echo "  Password: 12345678"
else
    echo "❌ PostgreSQL 启动失败"
    exit 1
fi

# 检查 Redis 连接
echo "检查 Redis 连接..."
docker-compose exec redis redis-cli ping

if [ $? -eq 0 ]; then
    echo "✅ Redis 启动成功"
    echo "Redis 连接信息:"
    echo "  Host: localhost"
    echo "  Port: 6379"
else
    echo "❌ Redis 启动失败"
    exit 1
fi

echo ""
echo "🎉 所有服务启动成功！"
echo "可以使用以下命令查看日志:"
echo "  docker-compose logs -f"
echo ""
echo "停止服务:"
echo "  docker-compose down" 