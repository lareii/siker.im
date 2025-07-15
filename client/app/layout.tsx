import { Toaster } from '@/components/ui/sonner';

// import type { Metadata } from 'next';
import Script from 'next/script';
import { Bricolage_Grotesque, JetBrains_Mono } from 'next/font/google';

import '@/styles/main.css';

const fontSans = Bricolage_Grotesque({
  variable: '--font-bricolage-grotesque-sans',
  subsets: ['latin']
});

const fontMono = JetBrains_Mono({
  variable: '--font-jetbrains-mono',
  subsets: ['latin']
});

// export const metadata: Metadata = {
//   metadataBase: new URL("https://siker.im"),
//   title: 'siker.im',
//   description: 'bazen bir link paylaşırsın ama o link uzar da uzar. siker.im, URL\'lerini hızlıca kısaltmana yarayan basit bir araçtır.',
//   keywords: [
//     'siker.im',
//     'url',
//     'shortener',
//     'short link',
//     'short url',
//     'link',
//     'shorten',
//     'short',
//     'url shortener',
//     'link shortener',
//   ],
//   openGraph: {
//     title: 'siker.im',
//     siteName: 'siker.im',
//     url: "https://siker.im",
//     description: 'bazen bir link paylaşırsın ama o link uzar da uzar. siker.im, URL\'lerini hızlıca kısaltmana yarayan basit bir araçtır.',
//   }
// };

export default function RootLayout({
  children
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="tr" translate="no">
      <body className={`${fontSans.variable} ${fontMono.variable} antialiased`}>
        {children}
        <Toaster />
      </body>
      <Script
        src="https://challenges.cloudflare.com/turnstile/v0/api.js"
        strategy="afterInteractive"
      />
    </html>
  );
}
