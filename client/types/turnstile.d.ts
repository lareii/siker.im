export {}

declare global {
  interface Window {
    turnstile: {
      render: (
        container: HTMLElement,
        options: {
          sitekey: string
          callback: (token: string) => void
        }
      ) => string
    }
  }
}
