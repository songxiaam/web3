import { useForm } from 'react-hook-form'
import axios from 'axios'
import { useRouter } from 'next/router'
import { useState } from 'react'
import { adminLogin } from '../api/auth'

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
    <div className="flex items-center justify-center min-h-screen bg-admin-bg">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-6 text-center">管理员登录</h2>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label className="block text-sm font-medium mb-1">账号</label>
            <input
              type="text"
              {...register('username', { required: '请输入账号' })}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
              placeholder="请输入账号"
            />
            {errors.username && <p className="text-red-500 text-xs mt-1">{errors.username.message as string}</p>}
          </div>
          <div>
            <label className="block text-sm font-medium mb-1">密码</label>
            <input
              type="password"
              {...register('password', { required: '请输入密码' })}
              className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-primary-500"
              placeholder="请输入密码"
            />
            {errors.password && <p className="text-red-500 text-xs mt-1">{errors.password.message as string}</p>}
          </div>
          {error && <p className="text-red-500 text-center text-sm">{error}</p>}
          <button
            type="submit"
            className="w-full py-2 bg-primary-600 text-white rounded-md hover:bg-primary-700 disabled:opacity-50"
            disabled={loading}
          >
            {loading ? '登录中...' : '登录'}
          </button>
        </form>
      </div>
    </div>
  )
} 