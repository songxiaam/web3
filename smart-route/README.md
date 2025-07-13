# smart-route 聚合器+最优路径项目

## 项目结构

```
smart-route/
├── backend/                 # Go 后端
│   ├── cmd/smart-route/    # 程序入口
│   ├── internal/           # 核心逻辑
│   │   ├── aggregator/    # 聚合器逻辑
│   │   ├── pathfinder/    # 最优路径查找
│   │   ├── auth/          # 钱包登录和 JWT 鉴权
│   │   └── config/        # 配置管理
│   ├── pkg/api/           # HTTP handler 路由层
│   ├── configs/           # 配置文件
│   ├── docs/              # Swagger API 文档
│   ├── go.mod             # Go 模块文件
│   └── go.sum             # Go 依赖锁定文件
├── contracts/              # 智能合约
│   ├── contracts/          # Solidity 合约源码
│   │   ├── SmartRoute.sol           # 主合约（可升级）
│   │   ├── SmartRouteV2.sol        # 升级版本示例
│   │   ├── SmartRouteProxy.sol     # 代理合约
│   │   └── SmartRouteProxyAdmin.sol # 代理管理员
│   ├── scripts/            # 部署脚本
│   │   ├── deploy.js       # 部署脚本
│   │   └── upgrade.js      # 升级脚本
│   ├── test/               # 合约测试
│   ├── package.json        # 合约项目配置
│   └── hardhat.config.js   # Hardhat 配置
├── frontend/               # 前端应用
│   ├── pages/              # Next.js 页面
│   ├── components/         # React 组件
│   ├── styles/             # 样式文件
│   ├── package.json        # 前端项目配置
│   └── next.config.js      # Next.js 配置
├── admin/                  # 后台管理系统
│   ├── pages/              # Next.js 页面
│   ├── components/         # React 组件
│   ├── styles/             # 样式文件
│   ├── package.json        # 管理后台配置
│   └── next.config.js      # Next.js 配置
├── scripts/                # 数据库脚本
│   ├── init.sql            # 数据库初始化脚本
│   ├── start-db.sh         # 数据库启动脚本
│   └── stop-db.sh          # 数据库停止脚本
└── docker-compose.yml      # Docker 编排文件
```

## 数据库服务

### 服务配置

项目使用 Docker Compose 管理数据库服务：

- **PostgreSQL 15** - 主数据库
  - 端口: 5433
  - 数据库: smartroute
  - 用户名: smartroute
  - 密码: 12345678

- **Redis 7** - 缓存服务
  - 端口: 6379
  - 无密码

### 启动数据库

```bash
# 启动数据库服务
./scripts/start-db.sh

# 或者直接使用 docker-compose
docker-compose up -d
```

### 停止数据库

```bash
# 停止数据库服务
./scripts/stop-db.sh

# 或者直接使用 docker-compose
docker-compose down
```

### 数据库表结构

初始化脚本会自动创建以下表：

- **users** - 用户表
- **routes** - 路由表
- **transactions** - 交易记录表
- **system_config** - 系统配置表

### 连接信息

```
PostgreSQL:
  Host: localhost
  Port: 5433
  Database: smartroute
  Username: smartroute
  Password: 12345678

Redis:
  Host: localhost
  Port: 6379
```

## 后端 (Go)

### 目录结构

- `backend/cmd/smart-route/main.go`：程序入口，负责启动服务
- `backend/internal/aggregator/`：聚合器核心逻辑
- `backend/internal/pathfinder/`：最优路径查找核心逻辑
- `backend/internal/auth/`：钱包登录和 JWT 鉴权
- `backend/internal/config/`：配置管理
- `backend/pkg/api/`：HTTP handler 路由层
- `backend/configs/config.yaml`：配置文件
- `backend/docs/`：Swagger API 文档

### 配置说明

项目使用 `backend/configs/config.yaml` 进行配置，包含：

- **Server**: 服务端口配置
- **Database**: PostgreSQL 数据库配置
- **Redis**: Redis 缓存配置  
- **JWT**: JWT 密钥和过期时间配置

### 启动方式

```bash
cd backend/cmd/smart-route
# 推荐设置端口环境变量，否则使用配置文件中的端口
export SMART_ROUTE_PORT=8080
go run main.go
```

### API 文档

启动服务后，访问以下地址查看 Swagger API 文档：

- **Swagger UI**: http://localhost:8080/swagger/index.html
- **API 文档**: http://localhost:8080/swagger/doc.json

### API 接口

#### 健康检查
- `GET /ping` - 返回 `{"message":"pong"}`

#### 钱包登录
- `GET /auth/nonce` - 获取随机 nonce
- `POST /auth/login` - 钱包登录，需要提供 address、message、signature

#### 需要鉴权的接口
- `GET /api/profile` - 获取用户信息（需要在 Header 中携带 `Authorization: Bearer <token>`）

## 智能合约

### 目录结构

- `contracts/contracts/SmartRoute.sol` - 主合约（可升级）
- `contracts/contracts/SmartRouteV2.sol` - 升级版本示例
- `contracts/contracts/SmartRouteProxy.sol` - 代理合约
- `contracts/contracts/SmartRouteProxyAdmin.sol` - 代理管理员
- `contracts/scripts/deploy.js` - 部署脚本
- `contracts/scripts/upgrade.js` - 升级脚本
- `contracts/hardhat.config.js` - Hardhat 配置

### 功能特性

- **可升级合约架构** - 使用 OpenZeppelin 透明代理模式
- **智能路由聚合器** - 多 DEX 聚合
- **手续费管理** - 协议手续费和路由手续费
- **安全防护** - ReentrancyGuard、Pausable
- **用户统计** - V2 版本新增用户交易统计
- **批量查询** - V2 版本新增批量路由查询

### 可升级特性

#### 代理模式
- **透明代理** - 使用 OpenZeppelin 的 TransparentUpgradeableProxy
- **代理管理员** - 独立的代理管理员合约
- **UUPS 升级** - 支持 UUPS (Universal Upgradeable Proxy Standard)

#### 升级流程
1. 部署新的实现合约
2. 通过代理管理员升级代理合约
3. 验证升级后的功能

### 部署方式

```bash
cd contracts
npm install
npx hardhat compile

# 部署合约
npm run deploy

# 启动本地区块链
npm run node

# 升级合约（需要先修改 upgrade.js 中的地址）
npm run upgrade
```

### 合约版本

#### V1 功能
- 基础路由聚合
- 手续费管理
- 安全防护

#### V2 功能（升级版本）
- 继承 V1 所有功能
- 用户交易统计
- 批量路由查询
- 增强的事件记录

## 前端应用

### 目录结构

- `frontend/pages/` - Next.js 页面
- `frontend/components/` - React 组件
- `frontend/styles/` - 样式文件

### 技术栈

- **Next.js** - React 框架
- **Wagmi** - Ethereum hooks
- **RainbowKit** - 钱包连接
- **Tailwind CSS** - 样式框架

### 启动方式

```bash
cd frontend
npm install
npm run dev
```

访问 http://localhost:3000 查看前端应用。

## 后台管理系统

### 目录结构

- `admin/pages/` - Next.js 页面
- `admin/components/` - React 组件
- `admin/styles/` - 样式文件

### 功能特性

- **仪表板** - 实时数据统计和监控
- **路由管理** - 添加/编辑/删除路由
- **用户管理** - 用户权限和交易记录
- **交易记录** - 实时交易监控和分析
- **合约管理** - 合约升级和参数配置
- **系统设置** - 管理员权限和系统配置

### 技术栈

- **Next.js** - React 框架
- **TypeScript** - 类型安全
- **Tailwind CSS** - 样式框架
- **Wagmi** - Ethereum hooks
- **RainbowKit** - 钱包连接
- **React Query** - 数据管理
- **Recharts** - 图表组件

### 启动方式

```bash
cd admin
npm install
npm run dev
```

访问 http://localhost:3001 查看管理后台。

## 钱包登录流程

1. 调用 `GET /auth/nonce` 获取随机 nonce
2. 使用钱包对 nonce 进行签名
3. 调用 `POST /auth/login` 提交地址、消息和签名
4. 获得 JWT token 后，在后续请求的 Header 中携带 `Authorization: Bearer <token>`

## 开发环境

### 数据库服务
```bash
# 启动数据库
./scripts/start-db.sh

# 停止数据库
./scripts/stop-db.sh
```

### 后端开发
```bash
# 启动后端服务
cd backend/cmd/smart-route
go run main.go
```

### 智能合约开发
```bash
# 启动本地区块链
cd contracts
npm run node

# 部署合约
npm run deploy

# 升级合约
npm run upgrade
```

### 前端开发
```bash
# 启动前端开发服务器
cd frontend
npm run dev
```

### 后台管理系统开发
```bash
# 启动管理后台
cd admin
npm run dev
```

访问 http://localhost:8080/ping 可看到 "pong" 响应。 