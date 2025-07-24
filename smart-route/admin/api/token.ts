import request from './request'

export interface Token {
  id: string
  chain: string
  chainId: number
  symbol: string
  name: string
  address: string
  decimals: number
  logo: string
  isNative: boolean
  isStable: boolean
  createdAt: string
  updatedAt: string
}

export interface TokenListParams {
  page?: number
  pageSize?: number
  symbol?: string
  name?: string
  chain?: string
  chainId?: number
}

export interface TokenListResponse {
  tokens: Token[]
  total: number
  page: number
  pageSize: number
}

export async function fetchTokenList(params: TokenListParams = {}): Promise<TokenListResponse> {
  const res = await request.get(
    process.env.NEXT_PUBLIC_ADMIN_API_URL + '/token/list',
    { params }
  )
  const { tokens, total, page, pageSize } = res.data
  return {
    tokens: tokens.map((t: any) => ({
      id: t.id,
      chain: t.chain,
      chainId: t.chainId,
      symbol: t.symbol,
      name: t.name,
      address: t.address,
      decimals: t.decimals,
      logo: t.logo,
      isNative: t.isNative,
      isStable: t.isStable,
      createdAt: t.createdAt,
      updatedAt: t.updatedAt,
    })),
    total,
    page,
    pageSize: pageSize,
  }
} 