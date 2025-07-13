import { useState } from 'react'
import Layout from '../components/Layout'
import { PlusIcon, PencilIcon, TrashIcon } from '@heroicons/react/24/outline'

const routes = [
  { id: 1, name: 'Uniswap V3', aggregator: '0x1234...', status: 'active', fee: '0.3%', volume: '$1.2M' },
  { id: 2, name: 'SushiSwap', aggregator: '0x5678...', status: 'active', fee: '0.25%', volume: '$800K' },
  { id: 3, name: 'Balancer', aggregator: '0x9abc...', status: 'inactive', fee: '0.4%', volume: '$500K' },
]

export default function Routes() {
  const [isAddModalOpen, setIsAddModalOpen] = useState(false)

  return (
    <Layout>
      <div className="space-y-6">
        {/* Header */}
        <div className="flex justify-between items-center">
          <h1 className="text-2xl font-bold text-gray-900">路由管理</h1>
          <button
            onClick={() => setIsAddModalOpen(true)}
            className="flex items-center px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700"
          >
            <PlusIcon className="h-5 w-5 mr-2" />
            添加路由
          </button>
        </div>

        {/* Routes Table */}
        <div className="admin-card">
          <table className="admin-table">
            <thead>
              <tr>
                <th>路由名称</th>
                <th>聚合器地址</th>
                <th>状态</th>
                <th>手续费</th>
                <th>交易量</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              {routes.map((route) => (
                <tr key={route.id}>
                  <td className="font-medium">{route.name}</td>
                  <td className="font-mono text-sm">{route.aggregator}</td>
                  <td>
                    <span className={`px-2 py-1 text-xs rounded-full ${
                      route.status === 'active' 
                        ? 'bg-green-100 text-green-800' 
                        : 'bg-red-100 text-red-800'
                    }`}>
                      {route.status === 'active' ? '活跃' : '停用'}
                    </span>
                  </td>
                  <td>{route.fee}</td>
                  <td>{route.volume}</td>
                  <td>
                    <div className="flex space-x-2">
                      <button className="p-1 text-blue-600 hover:text-blue-800">
                        <PencilIcon className="h-4 w-4" />
                      </button>
                      <button className="p-1 text-red-600 hover:text-red-800">
                        <TrashIcon className="h-4 w-4" />
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* Add Route Modal */}
        {isAddModalOpen && (
          <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
            <div className="bg-white rounded-lg p-6 w-full max-w-md">
              <h3 className="text-lg font-medium text-gray-900 mb-4">添加新路由</h3>
              <form className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    路由名称
                  </label>
                  <input
                    type="text"
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    placeholder="输入路由名称"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    聚合器地址
                  </label>
                  <input
                    type="text"
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    placeholder="0x..."
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 mb-1">
                    手续费
                  </label>
                  <input
                    type="number"
                    step="0.01"
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    placeholder="0.3"
                  />
                </div>
                <div className="flex space-x-3">
                  <button
                    type="button"
                    onClick={() => setIsAddModalOpen(false)}
                    className="flex-1 px-4 py-2 border border-gray-300 text-gray-700 rounded-md hover:bg-gray-50"
                  >
                    取消
                  </button>
                  <button
                    type="submit"
                    className="flex-1 px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700"
                  >
                    添加
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </Layout>
  )
} 