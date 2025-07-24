import axios, { type AxiosRequestHeaders } from 'axios'

const request = axios.create()

// 请求拦截器：自动附加JWT
request.interceptors.request.use((config) => {
  if (!config.url?.includes('/login')) {
    if (typeof window !== 'undefined') {
      const token = localStorage.getItem('admin_token')
      if (token) {
        (config.headers as AxiosRequestHeaders)["Authorization"] = `Bearer ${token}`
      }
    }
  }
  return config
})

// 响应拦截器：401自动登出
request.interceptors.response.use(
  response => response,
  error => {
    if (typeof window !== 'undefined' && error.response && error.response.status === 401) {
      localStorage.removeItem('admin_token')
      window.location.replace('/login')
    }
    return Promise.reject(error)
  }
)

export default request 