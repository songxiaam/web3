import api from './index'

export const getUserInfo = () => api.get('/user/profile')
// 钱包登录，参数为 address, message, signature
export const login = (data: { address: string; message: string; signature: string }) => api.post<{ token: string }>('/user/login', data) 