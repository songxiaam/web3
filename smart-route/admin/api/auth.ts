import request from './request'

export async function adminLogin(username: string, password: string) {
  const res = await request.post(
    process.env.NEXT_PUBLIC_ADMIN_API_URL + '/login',
    { username, password }
  )
  return res.data
} 