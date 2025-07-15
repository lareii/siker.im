import { useState, useEffect } from 'react';

export function useTurnstile() {
  const [token, setToken] = useState<string | null>(null);
  const [showWidget, setShowWidget] = useState(false);
  const [widgetKey, setWidgetKey] = useState(0);

  const siteKey =
    process.env.NEXT_PUBLIC_TURNSTILE_SITE_KEY || '1x00000000000000000000AA';

  function onVerify(token: string) {
    setToken(token);
    setShowWidget(false);
  }

  useEffect(() => {
    if (showWidget && window.turnstile) {
      const container = document.getElementById(
        `turnstile-container-${widgetKey}`
      );
      if (container) {
        window.turnstile.render(container, {
          sitekey: siteKey,
          callback: onVerify
        });
      }
    }
  }, [showWidget, widgetKey, siteKey]);

  function showNewWidget() {
    setToken(null);
    setWidgetKey(prev => prev + 1);
    setShowWidget(true);
  }

  return {
    token,
    showWidget,
    showNewWidget,
    widgetKey
  };
}
