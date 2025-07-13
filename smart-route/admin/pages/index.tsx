import { useState } from 'react'
import Layout from '../components/Layout'
import { 
  CurrencyDollarIcon, 
  UserGroupIcon, 
  ChartBarIcon,
  ArrowTrendingUpIcon 
} from '@heroicons/react/24/outline'

const stats = [
  { name: '总交易量', value: '$2.4M', change: '+12%', changeType: 'positive' },
  { name: '活跃用户', value: '1,234', change: '+8%', changeType: 'positive' },
  { name: '路由数量', value: '45', change: '+3', changeType: 'positive' },
  { name: '成功率', value: '99.2%', change: '+0.5%', changeType: 'positive' },
]

const recentTransactions = [
  { id: 1, user: '0x742d35...', amount: '1.5 ETH', route: 'Uniswap V3', time: '2分钟前' },
  { id: 2, user: '0x8f3a1b...', amount: '0.8 ETH', route: 'SushiSwap', time: '5分钟前' },
  { id: 3, user: '0x1a2b3c...', amount: '2.1 ETH', route: 'Balancer', time: '8分钟前' },
]

export default function Dashboard() {
  return (
    <Layout>
      <div className="space-y-6">
        {/* Stats */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          {stats.map((stat) => (
            <div key={stat.name} className="admin-card">
              <div className="flex items-center">
                <div className="flex-1">
                  <p className="text-sm font-medium text-gray-600">{stat.name}</p>
                  <p className="text-2xl font-bold text-gray-900">{stat.value}</p>
                </div>
                <div className={`flex items-center text-sm ${
                  stat.changeType === 'positive' ? 'text-green-600' : 'text-red-600'
                }`}>
                  <ArrowTrendingUpIcon className="h-4 w-4 mr-1" />
                  {stat.change}
                </div>
              </div>
            </div>
          ))}
        </div>

        {/* Charts and Tables */}
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Recent Transactions */}
          <div className="admin-card">
            <h3 className="text-lg font-medium text-gray-900 mb-4">最近交易</h3>
            <div className="space-y-3">
              {recentTransactions.map((transaction) => (
                <div key={transaction.id} className="flex items-center justify-between p-3 bg-gray-50 rounded-lg">
                  <div>
                    <p className="text-sm font-medium text-gray-900">{transaction.user}</p>
                    <p className="text-xs text-gray-500">{transaction.route}</p>
                  </div>
                  <div className="text-right">
                    <p className="text-sm font-medium text-gray-900">{transaction.amount}</p>
                    <p className="text-xs text-gray-500">{transaction.time}</p>
                  </div>
                </div>
              ))}
            </div>
          </div>

          {/* Quick Actions */}
          <div className="admin-card">
            <h3 className="text-lg font-medium text-gray-900 mb-4">快速操作</h3>
            <div className="space-y-3">
              <button className="w-full flex items-center justify-center px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700">
                <ChartBarIcon className="h-5 w-5 mr-2" />
                添加新路由
              </button>
              <button className="w-full flex items-center justify-center px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700">
                <UserGroupIcon className="h-5 w-5 mr-2" />
                管理用户
              </button>
              <button className="w-full flex items-center justify-center px-4 py-2 bg-yellow-600 text-white rounded-md hover:bg-yellow-700">
                <CurrencyDollarIcon className="h-5 w-5 mr-2" />
                查看报告
              </button>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  )
} 