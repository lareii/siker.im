'use client';

import { shortenUrl } from '@/lib/api/shorten';
import { TargetUrlInput, SlugInput } from '@/components/app/shortener/Inputs';
import { useTurnstile } from '@/components/app/shortener/useTurnstile';
import { Form } from '@/components/ui/form';
import { Button } from '@/components/ui/button';

import { IconInnerShadowBottomLeft } from '@tabler/icons-react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { toast } from 'sonner';
import Script from 'next/script';
import { useEffect, useState } from 'react';
import { z } from 'zod';

const formSchema = z.object({
  targetUrl: z.url({ message: 'Geçerli bir URL girin.' }),
  slug: z
    .string()
    .regex(/^[a-zA-Z0-9-_]*$/, {
      message: 'kısaltma sadece harf, rakam, tire ve alt çizgi içerebilir.'
    })
    .max(50, {
      message: 'kısaltma en fazla 50 karakter olmalıdır.'
    })
    .optional()
});

export function Shortener() {
  const [isLoading, setIsLoading] = useState(false);
  const { token, showWidget, showNewWidget, widgetKey } = useTurnstile();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: { targetUrl: '', slug: '' }
  });

  async function onSubmit() {
    showNewWidget();
  }

  useEffect(() => {
    async function submitWithToken() {
      if (!token) return;

      setIsLoading(true);

      const data = form.getValues();
      const response = await shortenUrl(data.targetUrl, data.slug, token);
      const shortUrl = `${process.env.NEXT_PUBLIC_BASE_URL}/${data.slug}`;

      switch (response.status) {
        case 201:
          navigator.clipboard.writeText(shortUrl);
          toast.success('URL başarıyla kısaltıldı', {
            description: `kısaltılmış URL: ${shortUrl}`,
            action: {
              label: 'kopyala',
              onClick: () => {
                navigator.clipboard.writeText(shortUrl);
                toast.success('kısaltılmış URL panoya kopyalandı.');
              }
            },
            duration: 30000
          });
          form.reset();
          break;
        case 400:
          toast.error('geçersiz url veya kısaltma etiketi.');
          break;
        case 409:
          toast.error('bu kısaltma etiketi zaten mevcut.');
          break;
        default:
          toast.error('bir hata oluştu, lütfen daha sonra tekrar deneyin.');
          break;
      }

      setIsLoading(false);
    }

    submitWithToken();
  }, [token, form]);

  return (
    <>
      <Script
        src="https://challenges.cloudflare.com/turnstile/v0/api.js"
        strategy="afterInteractive"
      />
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="flex flex-col gap-y-2 mt-10 rounded-lg md:w-1/2"
        >
          <TargetUrlInput control={form.control} isLoading={isLoading} />
          <div className="flex gap-x-2">
            <SlugInput control={form.control} isLoading={isLoading} />
            <Button
              type="submit"
              variant="ghost"
              className="w-24 cursor-pointer animated-glow"
              disabled={isLoading}
            >
              {isLoading ? (
                <>
                  <IconInnerShadowBottomLeft className="inline w-3.5 h-3.5 animate-spin" />
                  kısalt
                </>
              ) : (
                'kısalt'
              )}
            </Button>
          </div>
          <div className="text-xs text-muted-foreground">
            eğer etiket kısmını boş bırakırsan biz senin için bir tane
            oluşturacağız.
          </div>
          {showWidget && (
            <div id={`turnstile-container-${widgetKey}`} className="mt-4" />
          )}
        </form>
      </Form>
    </>
  );
}
