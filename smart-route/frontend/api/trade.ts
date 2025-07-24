import api from './index'

export const getBestRoute = (params: { tokenIn: string; tokenOut: string; amount: string }) =>
  api.get('/trade/best-route', { params })

export const executeTrade = (data: { tokenIn: string; tokenOut: string; amount: string }) =>
  api.post('/trade/execute', data) 