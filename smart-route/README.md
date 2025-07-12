# smart-route 聚合器+最优路径 Go 后端

## 目录结构

- `cmd/smart-route/main.go`：程序入口，负责启动服务
- `internal/aggregator/`：聚合器核心逻辑
- `internal/pathfinder/`：最优路径查找核心逻辑
- `pkg/api/`：HTTP handler 路由层
- `configs/config.yaml`：基础配置文件

## 启动方式

```bash
cd cmd/smart-route
# 推荐设置端口环境变量，否则默认 8080
export SMART_ROUTE_PORT=8080
go run main.go
```

访问 http://localhost:8080 可看到 "smart-route backend running" 响应。 