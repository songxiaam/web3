import React from 'react'

export interface TokenInfo {
  address: string
  symbol: string
  name: string
  logoURI?: string
  balance?: string
}

interface TokenInputCardProps {
  label: string
  token?: TokenInfo
  amount: string
  onAmountChange: (v: string) => void
  onSelectToken: () => void
  onMax?: () => void
  balance?: string
  disabled?: boolean
}

export default function TokenInputCard({
  label,
  token,
  amount,
  onAmountChange,
  onSelectToken,
  onMax,
  balance,
  disabled,
}: TokenInputCardProps) {
  return (
    <div className="bg-white/20 rounded-2xl p-4 flex flex-col space-y-2">
      <div className="flex justify-between items-center">
        <span className="text-xs text-gray-300">{label}</span>
        <span className="text-xs text-gray-400">余额: {balance ?? '--'}</span>
      </div>
      <div className="flex items-center space-x-3">
        <input
          type="number"
          className="bg-transparent text-2xl text-white flex-1 outline-none placeholder-gray-400"
          placeholder="0.0"
          value={amount}
          onChange={e => onAmountChange(e.target.value)}
          disabled={disabled}
        />
        {onMax && (
          <button
            className="text-xs text-blue-400 px-2 py-1 rounded hover:bg-blue-100/20"
            onClick={onMax}
            type="button"
          >
            MAX
          </button>
        )}
        <button
          className="flex items-center px-3 py-2 bg-white/10 rounded-xl hover:bg-white/20"
          onClick={onSelectToken}
          type="button"
        >
          {token?.logoURI ? (
            <img src={token.logoURI} className="w-6 h-6 rounded-full mr-2" alt={token.symbol} />
          ) : (
            <div className="w-6 h-6 rounded-full bg-gray-300 mr-2" />
          )}
          <span className="font-bold text-white mr-1">{token?.symbol || '选择代币'}</span>
          <svg className="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" strokeWidth={2} viewBox="0 0 24 24">
            <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
          </svg>
        </button>
      </div>
    </div>
  )
} 