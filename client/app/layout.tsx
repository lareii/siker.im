import { Toaster } from "@/components/ui/sonner"

import type { Metadata } from 'next';
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

export const metadata: Metadata = {
  title: 'siker.im',
  description: 'siker.im desc'
};

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
    </html>
  );
}
