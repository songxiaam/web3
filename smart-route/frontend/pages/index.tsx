import { useState } from 'react'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useAccount, useContractRead, useContractWrite, usePrepareContractWrite } from 'wagmi'
import { ethers } from 'ethers'

export default function Home() {
  const { address, isConnected } = useAccount()
  const [tokenIn, setTokenIn] = useState('')
  const [tokenOut, setTokenOut] = useState('')
  const [amountIn, setAmountIn] = useState('')
  const [amountOut, setAmountOut] = useState('')

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Header */}
        <div className="flex justify-between items-center py-6">
          <h1 className="text-3xl font-bold text-gray-900">Smart Route</h1>
          <ConnectButton />
        </div>

        {/* Main Content */}
        <div className="mt-8">
          <div className="bg-white shadow rounded-lg p-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-6">
              智能路由聚合器
            </h2>

            {!isConnected ? (
              <div className="text-center py-12">
                <p className="text-gray-500 mb-4">请先连接钱包</p>
                <ConnectButton />
              </div>
            ) : (
              <div className="space-y-6">
                {/* Token Input */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      输入代币地址
                    </label>
                    <input
                      type="text"
                      value={tokenIn}
                      onChange={(e) => setTokenIn(e.target.value)}
                      placeholder="0x..."
                      className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      输出代币地址
                    </label>
                    <input
                      type="text"
                      value={tokenOut}
                      onChange={(e) => setTokenOut(e.target.value)}
                      placeholder="0x..."
                      className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    />
                  </div>
                </div>

                {/* Amount Input */}
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      输入金额
                    </label>
                    <input
                      type="number"
                      value={amountIn}
                      onChange={(e) => setAmountIn(e.target.value)}
                      placeholder="0.0"
                      className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-gray-700 mb-2">
                      预期输出金额
                    </label>
                    <input
                      type="number"
                      value={amountOut}
                      onChange={(e) => setAmountOut(e.target.value)}
                      placeholder="0.0"
                      className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
                    />
                  </div>
                </div>

                {/* Action Buttons */}
                <div className="flex space-x-4">
                  <button
                    className="flex-1 bg-primary-600 text-white px-4 py-2 rounded-md hover:bg-primary-700 focus:outline-none focus:ring-2 focus:ring-primary-500"
                    onClick={() => {
                      // TODO: 实现获取最优路径逻辑
                      console.log('获取最优路径')
                    }}
                  >
                    获取最优路径
                  </button>
                  <button
                    className="flex-1 bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500"
                    onClick={() => {
                      // TODO: 实现执行交换逻辑
                      console.log('执行交换')
                    }}
                  >
                    执行交换
                  </button>
                </div>

                {/* Route Information */}
                <div className="mt-6 p-4 bg-gray-50 rounded-md">
                  <h3 className="text-lg font-medium text-gray-900 mb-2">路由信息</h3>
                  <div className="text-sm text-gray-600">
                    <p>连接地址: {address}</p>
                    <p>最优路径: 待计算</p>
                    <p>预期收益: 待计算</p>
                  </div>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
} 