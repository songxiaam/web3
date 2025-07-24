import { useState, useEffect } from 'react'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useAccount, useSignMessage } from 'wagmi'
import Navbar from '../components/Navbar'
import { login } from '../api/user'
import { getTokenList } from '../api/token'
import TokenInputCard, { TokenInfo } from '../components/TokenInputCard'

function TokenSelectModal({ open, onClose, tokens, onSelect, title }: {
  open: boolean
  onClose: () => void
  tokens: { address: string; symbol: string; name: string; logoURI?: string }[]
  onSelect: (token: { address: string; symbol: string; name: string; logoURI?: string }) => void
  title: string
}) {
  if (!open) return null
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-40">
      <div className="bg-white rounded-2xl shadow-2xl w-full max-w-md p-6">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-lg font-bold text-gray-900">{title}</h3>
          <button onClick={onClose} className="text-gray-400 hover:text-gray-700 text-2xl leading-none">×</button>
        </div>
        <div className="max-h-80 overflow-y-auto divide-y divide-gray-100 custom-scrollbar">
          {tokens.map(token => (
            <button
              key={token.address}
              className="w-full flex items-center px-3 py-3 hover:bg-blue-50 focus:bg-blue-100 rounded-xl transition-all mb-1"
              onClick={() => { onSelect(token); onClose() }}
              type="button"
            >
              {token.logoURI ? (
                <img src={token.logoURI} className="w-7 h-7 rounded-full mr-3 border border-gray-200" alt={token.symbol} />
              ) : (
                <div className="w-7 h-7 rounded-full bg-gray-200 mr-3" />
              )}
              <span className="font-semibold text-blue-700 mr-2">{token.symbol}</span>
              <span className="text-gray-600 text-sm truncate max-w-[100px]">{token.name}</span>
              <span className="ml-auto text-xs text-gray-400 font-mono">{token.address.slice(0, 6)}...{token.address.slice(-4)}</span>
            </button>
          ))}
        </div>
      </div>
      <style jsx>{`
        .custom-scrollbar::-webkit-scrollbar {
          width: 6px;
        }
        .custom-scrollbar::-webkit-scrollbar-thumb {
          background: #e0e7ef;
          border-radius: 3px;
        }
      `}</style>
    </div>
  )
}

export default function Home() {
  const [mounted, setMounted] = useState(false)
  const { address, isConnected } = useAccount()
  const [tokenIn, setTokenIn] = useState('')
  const [tokenOut, setTokenOut] = useState('')
  const [amountIn, setAmountIn] = useState('')
  const [amountOut, setAmountOut] = useState('')
  const [loginLoading, setLoginLoading] = useState(false)
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const { signMessageAsync } = useSignMessage()
  const [tokens, setTokens] = useState<{ address: string; symbol: string; name: string }[]>([])
  const [showTokenInModal, setShowTokenInModal] = useState(false)
  const [showTokenOutModal, setShowTokenOutModal] = useState(false)

  // 计算当前选中的 token 对象
  const tokenInObj = tokens.find(t => t.address === tokenIn)
  const tokenOutObj = tokens.find(t => t.address === tokenOut)

  useEffect(() => {
    setMounted(true)
  }, [])

  // 获取代币列表
  useEffect(() => {
    const fetchTokens = async () => {
      try {
        const res = await getTokenList({ pageSize: 1000 })
        let arr: { address: string; symbol: string; name: string }[] = []
        if (res && res.data && Array.isArray(res.data.tokens)) {
          arr = res.data.tokens
        } else if (Array.isArray(res)) {
          arr = res
        }
        setTokens(arr)
      } catch (e) {
        setTokens([])
      }
    }
    fetchTokens()
  }, [])

  // 钱包登录逻辑（如需登录才显示主操作区可保留，否则可移除 isLoggedIn 判断）
  useEffect(() => {
    const walletLogin = async () => {
      if (isConnected && address) {
        setLoginLoading(true)
        try {
          const message = `Login to Smart Route at ${new Date().toISOString()} for address: ${address}`
          const signature = await signMessageAsync({ message })
          const res = await login({ address, message, signature })
          if (res && res.data && res.data.token) {
            localStorage.setItem('jwt', res.data.token)
            setIsLoggedIn(true)
          }
        } catch (e) {
          setIsLoggedIn(false)
        } finally {
          setLoginLoading(false)
        }
      } else {
        setIsLoggedIn(false)
      }
    }
    walletLogin()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isConnected, address])

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-gray-900">
      <Navbar />
      {mounted && (
        <div className="max-w-md mx-auto py-10">
          <div className="bg-white/10 backdrop-blur-lg rounded-2xl p-6 border border-white/20 shadow-2xl">
            <h2 className="text-3xl font-bold text-white mb-8 text-center">
              智能路由聚合器
            </h2>
            <div className="space-y-6">
              {/* 输入代币卡片 */}
              <div className="bg-white/20 rounded-2xl p-4 flex flex-col space-y-2">
                <div className="flex justify-between items-center mb-2">
                  <span className="text-xs text-gray-300">输入代币</span>
                  <span className="text-xs text-gray-400">余额: {tokenInObj && typeof (tokenInObj as any).balance === 'string' ? (tokenInObj as any).balance : '--'}</span>
                </div>
                <button
                  className="flex items-center px-3 py-2 bg-white/10 rounded-xl hover:bg-white/20 mb-3 w-fit"
                  onClick={() => setShowTokenInModal(true)}
                  type="button"
                >
                  {'logoURI' in (tokenInObj || {}) && (tokenInObj as any).logoURI ? (
                    <img src={(tokenInObj as any).logoURI} className="w-6 h-6 rounded-full mr-2" alt={tokenInObj?.symbol} />
                  ) : (
                    <div className="w-6 h-6 rounded-full bg-gray-300 mr-2" />
                  )}
                  <span className="font-bold text-white mr-1">{tokenInObj?.symbol || '选择代币'}</span>
                  <svg className="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" strokeWidth={2} viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>
                <div className="flex items-center space-x-3">
                  <input
                    type="number"
                    className="bg-transparent text-2xl text-white flex-1 outline-none placeholder-gray-400"
                    placeholder="0.0"
                    value={amountIn}
                    onChange={e => setAmountIn(e.target.value)}
                  />
                  <button
                    className="text-xs text-blue-400 px-2 py-1 rounded hover:bg-blue-100/20"
                    onClick={() => setAmountIn('999')}
                    type="button"
                  >
                    MAX
                  </button>
                </div>
              </div>
              {/* 箭头 */}
              <div className="flex justify-center -my-2">
                <div className="w-8 h-8 rounded-full bg-white/20 flex items-center justify-center shadow">
                  <svg className="w-5 h-5 text-white" fill="none" stroke="currentColor" strokeWidth={2} viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
                  </svg>
                </div>
              </div>
              {/* 输出代币卡片 */}
              <div className="bg-white/20 rounded-2xl p-4 flex flex-col space-y-2">
                <div className="flex justify-between items-center mb-2">
                  <span className="text-xs text-gray-300">输出代币</span>
                  <span className="text-xs text-gray-400">余额: {tokenOutObj && typeof (tokenOutObj as any).balance === 'string' ? (tokenOutObj as any).balance : '--'}</span>
                </div>
                <button
                  className="flex items-center px-3 py-2 bg-white/10 rounded-xl hover:bg-white/20 mb-3 w-fit"
                  onClick={() => setShowTokenOutModal(true)}
                  type="button"
                >
                  {'logoURI' in (tokenOutObj || {}) && (tokenOutObj as any).logoURI ? (
                    <img src={(tokenOutObj as any).logoURI} className="w-6 h-6 rounded-full mr-2" alt={tokenOutObj?.symbol} />
                  ) : (
                    <div className="w-6 h-6 rounded-full bg-gray-300 mr-2" />
                  )}
                  <span className="font-bold text-white mr-1">{tokenOutObj?.symbol || '选择代币'}</span>
                  <svg className="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" strokeWidth={2} viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
                  </svg>
                </button>
                <div className="flex items-center space-x-3">
                  <input
                    type="number"
                    className="bg-transparent text-2xl text-white flex-1 outline-none placeholder-gray-400"
                    placeholder="0.0"
                    value={amountOut}
                    onChange={e => setAmountOut(e.target.value)}
                    disabled
                  />
                </div>
              </div>
            </div>
            {/* 代币选择弹窗 */}
            <TokenSelectModal
              open={showTokenInModal}
              onClose={() => setShowTokenInModal(false)}
              tokens={tokens}
              onSelect={token => setTokenIn(token.address)}
              title="选择输入代币"
            />
            <TokenSelectModal
              open={showTokenOutModal}
              onClose={() => setShowTokenOutModal(false)}
              tokens={tokens}
              onSelect={token => setTokenOut(token.address)}
              title="选择输出代币"
            />
            {/* 钱包登录和主操作区（如需登录才可操作可保留 isLoggedIn 判断） */}
            {!isConnected ? (
              <div className="text-center py-12">
                <div className="mb-6 flex flex-col items-center justify-center">
                  <div className="w-16 h-16 bg-gradient-to-r from-blue-500 to-purple-600 rounded-full flex items-center justify-center mx-auto mb-4">
                    <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                    </svg>
                  </div>
                  <p className="text-gray-300 mb-6 text-lg">如需交易请先连接钱包</p>
                  <ConnectButton />
                </div>
              </div>
            ) : loginLoading ? (
              <div className="flex flex-col items-center justify-center py-12">
                <div className="mb-4 text-white text-lg">正在登录中，请在钱包中签名...</div>
                <div className="animate-spin rounded-full h-10 w-10 border-b-2 border-white"></div>
              </div>
            ) : !isLoggedIn ? (
              <div className="flex flex-col items-center justify-center py-12">
                <div className="mb-4 text-red-300 text-lg">登录失败，请刷新页面重试</div>
              </div>
            ) : (
              <div className="flex space-x-4 pt-4">
                <button
                  className="flex-1 bg-gradient-to-r from-blue-600 to-purple-600 text-white px-6 py-3 rounded-xl hover:from-blue-700 hover:to-purple-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-gray-900 transition-all duration-200 font-medium shadow-lg"
                  onClick={() => {
                    // TODO: 实现获取最优路径逻辑
                    console.log('获取最优路径')
                  }}
                >
                  获取最优路径
                </button>
                <button
                  className="flex-1 bg-gradient-to-r from-green-600 to-emerald-600 text-white px-6 py-3 rounded-xl hover:from-green-700 hover:to-emerald-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 focus:ring-offset-gray-900 transition-all duration-200 font-medium shadow-lg"
                  onClick={() => {
                    // TODO: 实现执行交换逻辑
                    console.log('执行交换')
                  }}
                >
                  执行交换
                </button>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  )
} 