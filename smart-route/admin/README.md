# Smart Route 后台管理系统

## 功能特性

### 仪表板
- 实时数据统计
- 交易量趋势图
- 用户活跃度监控
- 路由性能分析

### 路由管理
- 添加/编辑/删除路由
- 路由状态监控
- 手续费设置
- 聚合器地址管理

### 用户管理
- 用户列表查看
- 用户权限管理
- 交易历史记录
- 用户统计分析

### 交易记录
- 实时交易监控
- 交易详情查看
- 失败交易分析
- 交易报告导出

### 合约管理
- 合约升级管理
- 参数配置
- 安全设置
- 事件监控

### 系统设置
- 管理员权限
- 系统参数配置
- 日志查看
- 备份恢复

## 技术栈

- **Next.js** - React 框架
- **TypeScript** - 类型安全
- **Tailwind CSS** - 样式框架
- **Wagmi** - Ethereum hooks
- **RainbowKit** - 钱包连接
- **React Query** - 数据管理
- **Recharts** - 图表组件

## 启动方式

```bash
cd admin
npm install
npm run dev
```

访问 http://localhost:3001 查看管理后台。

## 环境变量

```bash
# .env.local
NEXT_PUBLIC_API_URL=http://localhost:8080
NEXT_PUBLIC_CONTRACT_ADDRESS=0x...
NEXT_PUBLIC_ADMIN_API_URL=http://localhost:8080/admin
```

## 开发指南

### 添加新页面
1. 在 `pages/` 目录创建新页面
2. 在 `components/Layout.tsx` 中添加导航项
3. 实现页面功能和样式

### 添加新组件
1. 在 `components/` 目录创建组件
2. 使用 TypeScript 定义接口
3. 添加必要的样式和功能

### 数据获取
使用 React Query 进行数据管理：
```typescript
import { useQuery } from 'react-query'

const { data, isLoading } = useQuery('routes', fetchRoutes)
```

### 状态管理
使用 React hooks 进行状态管理：
```typescript
const [isModalOpen, setIsModalOpen] = useState(false)
```

## 部署

### 构建
```bash
npm run build
```

### 启动生产服务
```bash
npm start
```

## 安全注意事项

1. **权限控制** - 确保只有管理员可以访问
2. **数据验证** - 所有输入都需要验证
3. **错误处理** - 完善的错误处理机制
4. **日志记录** - 记录所有管理操作
5. **备份策略** - 定期备份重要数据 