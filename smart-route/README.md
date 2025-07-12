# smart-route 聚合器+最优路径 Go 后端

## 目录结构

- `cmd/smart-route/main.go`：程序入口，负责启动服务
- `internal/aggregator/`：聚合器核心逻辑
- `internal/pathfinder/`：最优路径查找核心逻辑
- `internal/auth/`：钱包登录和 JWT 鉴权
- `internal/config/`：配置管理
- `pkg/api/`：HTTP handler 路由层
- `configs/config.yaml`：配置文件
- `docs/`：Swagger API 文档

## 配置说明

项目使用 `configs/config.yaml` 进行配置，包含：

- **Server**: 服务端口配置
- **Database**: PostgreSQL 数据库配置
- **Redis**: Redis 缓存配置  
- **JWT**: JWT 密钥和过期时间配置

## 启动方式

```bash
cd cmd/smart-route
# 推荐设置端口环境变量，否则使用配置文件中的端口
export SMART_ROUTE_PORT=8080
go run main.go
```

## API 文档

启动服务后，访问以下地址查看 Swagger API 文档：

- **Swagger UI**: http://localhost:8080/swagger/index.html
- **API 文档**: http://localhost:8080/swagger/doc.json

## API 接口

### 健康检查
- `GET /ping` - 返回 `{"message":"pong"}`

### 钱包登录
- `GET /auth/nonce` - 获取随机 nonce
- `POST /auth/login` - 钱包登录，需要提供 address、message、signature

### 需要鉴权的接口
- `GET /api/profile` - 获取用户信息（需要在 Header 中携带 `Authorization: Bearer <token>`）

## 钱包登录流程

1. 调用 `GET /auth/nonce` 获取随机 nonce
2. 使用钱包对 nonce 进行签名
3. 调用 `POST /auth/login` 提交地址、消息和签名
4. 获得 JWT token 后，在后续请求的 Header 中携带 `Authorization: Bearer <token>`

访问 http://localhost:8080/ping 可看到 "pong" 响应。 