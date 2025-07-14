'use client';

import { useRef } from 'react';
import Script from 'next/script';

interface TurnstileProps {
  onVerify: (token: string) => void;
}

export default function Turnstile({ onVerify }: TurnstileProps) {
  const widgetRef = useRef<HTMLDivElement>(null);

  if (!process.env.NEXT_PUBLIC_TURNSTILE_SITE_KEY) return;

  function handleScriptLoad() {
    if (window.turnstile && widgetRef.current) {
      window.turnstile.render(widgetRef.current, {
        sitekey: process.env.NEXT_PUBLIC_TURNSTILE_SITE_KEY || '',
        callback: (token: string) => {
          onVerify(token);
        }
      });
    }
  }

  return (
    <>
      <Script
        src="https://challenges.cloudflare.com/turnstile/v0/api.js"
        strategy="afterInteractive"
        onLoad={handleScriptLoad}
      />
      <div id="turnstile-widget" ref={widgetRef}></div>
    </>
  );
}
