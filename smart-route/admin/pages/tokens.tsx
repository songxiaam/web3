import { useState } from 'react'
import Layout from '../components/Layout'
import { fetchTokenList, Token } from '../api/token'
import { useQuery } from 'react-query'

export default function TokenListPage() {
  const [page, setPage] = useState(1)
  const [pageSize] = useState(20)
  const [filters, setFilters] = useState({ symbol: '', name: '', chain: '', chainId: '' })

  const { data, isLoading, refetch } = useQuery([
    'tokens', page, pageSize, filters
  ], () => fetchTokenList({
    page,
    pageSize,
    symbol: filters.symbol || undefined,
    name: filters.name || undefined,
    chain: filters.chain || undefined,
    chainId: filters.chainId && !isNaN(Number(filters.chainId)) ? Number(filters.chainId) : undefined,
  }), { keepPreviousData: true })

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFilters(prev => ({
      ...prev,
      [name]: name === 'chainId' ? (value.replace(/\D/g, '')) : value
    }))
  }

  const handleSearch = () => {
    setPage(1)
    refetch()
  }

  return (
    <Layout>
      <div className="space-y-6">
        <h1 className="text-2xl font-bold text-gray-900 mb-4">支持的代币种类</h1>
        {/* 筛选表单 */}
        <div className="admin-card mb-4">
          <div className="flex flex-wrap gap-4 items-end">
            <div>
              <label className="block text-sm font-medium mb-1">Symbol</label>
              <input name="symbol" value={filters.symbol} onChange={handleInputChange} className="px-3 py-2 border border-gray-300 rounded-md" placeholder="如 USDT" />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Name</label>
              <input name="name" value={filters.name} onChange={handleInputChange} className="px-3 py-2 border border-gray-300 rounded-md" placeholder="如 Tether" />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Chain</label>
              <input name="chain" value={filters.chain} onChange={handleInputChange} className="px-3 py-2 border border-gray-300 rounded-md" placeholder="如 Ethereum" />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">ChainId</label>
              <input name="chainId" value={filters.chainId} onChange={handleInputChange} className="px-3 py-2 border border-gray-300 rounded-md" placeholder="如 1" />
            </div>
            <button onClick={handleSearch} className="px-4 py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700">搜索</button>
          </div>
        </div>
        {/* 列表表格 */}
        <div className="admin-card">
          <table className="admin-table">
            <thead>
              <tr>
                <th>Logo</th>
                <th>Symbol</th>
                <th>Name</th>
                <th>Chain</th>
                <th>ChainId</th>
                <th>Address</th>
                <th>Decimals</th>
                <th>原生币</th>
                <th>稳定币</th>
              </tr>
            </thead>
            <tbody>
              {isLoading ? (
                <tr><td colSpan={9} className="text-center">加载中...</td></tr>
              ) : data && data.tokens.length > 0 ? (
                data.tokens.map((token: Token) => (
                  <tr key={token.id}>
                    <td>{token.logo ? <img src={token.logo} alt={token.symbol} className="h-6 w-6" /> : '-'}</td>
                    <td>{token.symbol}</td>
                    <td>{token.name}</td>
                    <td>{token.chain}</td>
                    <td>{token.chainId}</td>
                    <td className="font-mono text-xs break-all">{token.address}</td>
                    <td>{token.decimals}</td>
                    <td>{token.isNative ? '是' : '否'}</td>
                    <td>{token.isStable ? '是' : '否'}</td>
                  </tr>
                ))
              ) : (
                <tr><td colSpan={9} className="text-center">暂无数据</td></tr>
              )}
            </tbody>
          </table>
          {/* 分页 */}
          {data && data.total > pageSize && (
            <div className="flex justify-end items-center mt-4 space-x-2">
              <button
                className="px-3 py-1 border rounded disabled:opacity-50"
                onClick={() => setPage(page - 1)}
                disabled={page === 1}
              >上一页</button>
              <span>第 {page} / {Math.ceil(data.total / pageSize)} 页</span>
              <button
                className="px-3 py-1 border rounded disabled:opacity-50"
                onClick={() => setPage(page + 1)}
                disabled={page >= Math.ceil(data.total / pageSize)}
              >下一页</button>
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
} 