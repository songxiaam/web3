import { useForm } from 'react-hook-form'
import { adminLogin } from '../api/auth'
import { useRouter } from 'next/router'
import { useState } from 'react'

export default function Login() {
  const { register, handleSubmit, formState: { errors } } = useForm()
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const onSubmit = async (data: any) => {
    setLoading(true)
    setError('')
    try {
      const res = await adminLogin(data.username, data.password)
      if (res && res.token) {
        localStorage.setItem('admin_token', res.token)
        router.push('/')
      } else {
        setError('登录失败，请检查账号和密码')
      }
    } catch (e: any) {
      setError(e.response?.data?.message || '登录失败')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-[#0f2027] via-[#2c5364] to-[#24243e] relative overflow-hidden">
      {/* web3 背景装饰 */}
      <div className="absolute inset-0 pointer-events-none select-none">
        <div className="absolute -top-32 -left-32 w-96 h-96 bg-gradient-to-tr from-blue-500 via-purple-500 to-pink-500 opacity-30 rounded-full blur-3xl animate-pulse" />
        <div className="absolute bottom-0 right-0 w-80 h-80 bg-gradient-to-br from-green-400 via-cyan-400 to-blue-500 opacity-20 rounded-full blur-2xl animate-blob" />
      </div>
      <div className="relative z-10 w-full max-w-md">
        <div className="flex flex-col items-center mb-8">
          {/* 可替换为项目logo */}
          <div className="bg-gradient-to-tr from-blue-500 via-purple-500 to-pink-500 p-1 rounded-full mb-3 shadow-lg">
            <svg width="48" height="48" viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
              <circle cx="24" cy="24" r="24" fill="url(#paint0_linear_1_2)" />
              <path d="M24 12L32 36H16L24 12Z" fill="white" fillOpacity="0.9" />
              <defs>
                <linearGradient id="paint0_linear_1_2" x1="0" y1="0" x2="48" y2="48" gradientUnits="userSpaceOnUse">
                  <stop stopColor="#00F2FE" />
                  <stop offset="1" stopColor="#4A00E0" />
                </linearGradient>
              </defs>
            </svg>
          </div>
          <h2 className="text-3xl font-extrabold text-white drop-shadow mb-2 tracking-wide">Web3 管理员登录</h2>
          <p className="text-gray-300 text-sm">Smart Route Admin Dashboard</p>
        </div>
        <form onSubmit={handleSubmit(onSubmit)} className="bg-white/90 rounded-2xl shadow-2xl p-8 space-y-6 backdrop-blur-md border border-white/30">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">账号</label>
            <input
              type="text"
              {...register('username', { required: '请输入账号' })}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400 bg-white/80"
              placeholder="请输入账号"
              autoComplete="username"
            />
            {errors.username && <p className="text-red-500 text-xs mt-1">{errors.username.message as string}</p>}
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">密码</label>
            <input
              type="password"
              {...register('password', { required: '请输入密码' })}
              className="w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-400 bg-white/80"
              placeholder="请输入密码"
              autoComplete="current-password"
            />
            {errors.password && <p className="text-red-500 text-xs mt-1">{errors.password.message as string}</p>}
          </div>
          {error && <p className="text-red-500 text-center text-sm font-medium">{error}</p>}
          <button
            type="submit"
            className="w-full py-2 bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 text-white font-bold rounded-lg shadow-md hover:scale-105 hover:shadow-xl transition-all duration-200 disabled:opacity-60"
            disabled={loading}
          >
            {loading ? '登录中...' : '登录'}
          </button>
        </form>
        <div className="mt-6 text-center text-xs text-gray-400">© {new Date().getFullYear()} Smart Route. All rights reserved.</div>
      </div>
    </div>
  )
} 