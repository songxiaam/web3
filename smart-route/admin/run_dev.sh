#!/bin/bash

# 一键启动 Smart Route Admin 前端
cd "$(dirname "$0")"

if [ ! -d "node_modules" ]; then
  echo "[INFO] 安装依赖..."
  npm install
fi

echo "[INFO] 启动开发服务器..."
npm run dev 