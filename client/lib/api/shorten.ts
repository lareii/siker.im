import api from '@/lib/api';

export async function shortenUrl(targetUrl: string, slug?: string, turnstileToken?: string) {
  try {
    const payload = { target_url: targetUrl, slug };
    const headers: Record<string, string> = {};

    if (turnstileToken) {
      headers['Cf-Turnstile-Token'] = turnstileToken;
    }

    const response = await api.post('/urls/', payload, { headers });
    return response;
  } catch (error: any) {
    return error.response || error;
  }
}
