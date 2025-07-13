import type { AppProps } from 'next/app'
import { RainbowKitProvider, getDefaultWallets } from '@rainbow-me/rainbowkit'
import { configureChains, createConfig, WagmiConfig } from 'wagmi'
import { mainnet, polygon, sepolia } from 'wagmi/chains'
import { publicProvider } from 'wagmi/providers/public'
import { QueryClient, QueryClientProvider } from 'react-query'
import { Toaster } from 'react-hot-toast'
import '@rainbow-me/rainbowkit/styles.css'
import '../styles/globals.css'

const { chains, publicClient, webSocketPublicClient } = configureChains(
  [mainnet, polygon, sepolia],
  [publicProvider()]
)

const { connectors } = getDefaultWallets({
  appName: 'Smart Route Admin',
  projectId: 'YOUR_PROJECT_ID',
  chains,
})

const wagmiConfig = createConfig({
  autoConnect: true,
  connectors,
  publicClient,
  webSocketPublicClient,
})

const queryClient = new QueryClient()

export default function App({ Component, pageProps }: AppProps) {
  return (
    <WagmiConfig config={wagmiConfig}>
      <RainbowKitProvider chains={chains}>
        <QueryClientProvider client={queryClient}>
          <Component {...pageProps} />
          <Toaster position="top-right" />
        </QueryClientProvider>
      </RainbowKitProvider>
    </WagmiConfig>
  )
} 