import type { AppProps } from 'next/app'
import { RainbowKitProvider, getDefaultWallets } from '@rainbow-me/rainbowkit'
import { configureChains, createConfig, WagmiConfig } from 'wagmi'
import { mainnet, polygon, sepolia } from 'wagmi/chains'
import { publicProvider } from 'wagmi/providers/public'
import '@rainbow-me/rainbowkit/styles.css'
import '../styles/globals.css'

const { chains, publicClient, webSocketPublicClient } = configureChains(
  [mainnet, polygon, sepolia],
  [publicProvider()]
)

const { connectors } = getDefaultWallets({
  appName: 'Smart Route',
  projectId: 'YOUR_PROJECT_ID', // 需要替换为实际的 WalletConnect Project ID
  chains,
})

const wagmiConfig = createConfig({
  autoConnect: true,
  connectors,
  publicClient,
  webSocketPublicClient,
})

export default function App({ Component, pageProps }: AppProps) {
  return (
    <WagmiConfig config={wagmiConfig}>
      <RainbowKitProvider chains={chains}>
        <Component {...pageProps} />
      </RainbowKitProvider>
    </WagmiConfig>
  )
} 