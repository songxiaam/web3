#!/bin/bash

# Smart Route 数据库停止脚本

echo "停止 Smart Route 数据库服务..."

# 检查 Docker 是否运行
if ! docker info > /dev/null 2>&1; then
    echo "错误: Docker 未运行"
    exit 1
fi

# 停止服务
echo "停止 PostgreSQL 和 Redis..."
docker-compose down

echo "✅ 服务已停止"

# 可选：删除数据卷（谨慎使用）
if [ "$1" = "--clean" ]; then
    echo "⚠️  删除数据卷..."
    docker-compose down -v
    echo "✅ 数据卷已删除"
fi

echo ""
echo "服务已停止。重新启动:"
echo "  ./scripts/start-db.sh" 