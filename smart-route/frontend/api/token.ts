import api from './index'

export const getTokenList = (params?: any) => api.get('/token/list', { params })
export const getTokenDetail = (address: string) => api.get(`/token/${address}`) 