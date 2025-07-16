import { ReactNode } from 'react'
import { ConnectButton } from '@rainbow-me/rainbowkit'
import { 
  HomeIcon, 
  ChartBarIcon, 
  CogIcon, 
  UserGroupIcon,
  CurrencyDollarIcon,
  ShieldCheckIcon
} from '@heroicons/react/24/outline'
import Link from 'next/link'
import { useRouter } from 'next/router'

interface LayoutProps {
  children: ReactNode
}

const navigation = [
  { name: '仪表板', href: '/', icon: HomeIcon },
  { name: '路由管理', href: '/routes', icon: ChartBarIcon },
  { name: '用户管理', href: '/users', icon: UserGroupIcon },
  { name: '交易记录', href: '/transactions', icon: CurrencyDollarIcon },
  { name: '合约管理', href: '/contracts', icon: ShieldCheckIcon },
  { name: '系统设置', href: '/settings', icon: CogIcon },
]

export default function Layout({ children }: LayoutProps) {
  const router = useRouter()

  return (
    <div className="admin-layout">
      {/* Sidebar */}
      <div className="admin-sidebar">
        <div className="p-6">
          <h1 className="text-xl font-bold">Smart Route Admin</h1>
        </div>
        
        <nav className="mt-6">
          <div className="px-3">
            {navigation.map((item) => {
              const isActive = router.pathname === item.href
              return (
                <Link
                  key={item.name}
                  href={item.href}
                  className={`
                    flex items-center px-3 py-2 text-sm font-medium rounded-md mb-1
                    ${isActive 
                      ? 'bg-primary-600 text-white' 
                      : 'text-gray-300 hover:bg-gray-700 hover:text-white'
                    }
                  `}
                >
                  <item.icon className="mr-3 h-5 w-5" />
                  {item.name}
                </Link>
              )
            })}
          </div>
        </nav>
      </div>

      {/* Main content */}
      <div className="admin-main">
        {/* Header */}
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-gray-900">管理后台</h2>
          <ConnectButton />
        </div>

        {/* Page content */}
        {children}
      </div>
    </div>
  )
} 