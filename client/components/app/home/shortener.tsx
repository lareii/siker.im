'use client'

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Button } from '@/components/ui/button';
import Turnstile from "@/components/app/turnstile";
import { shortenUrl } from "@/lib/api/shorten";

import { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { toast } from "sonner";

const formSchema = z.object({
  targetUrl: z.url({ message: "Geçerli bir URL girin." }),
  slug: z.string().regex(/^[a-zA-Z0-9-_]*$/, {
    message: "Kısaltma sadece harf, rakam, tire ve alt çizgi içerebilir."
  }).max(50, {
    message: "Kısaltma en fazla 50 karakter olmalıdır."
  }).optional(),
})

export default function Shortener() {
  const [turnstileToken, setTurnstileToken] = useState<string | null>(null);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      targetUrl: "",
      slug: "",
    },
  })

  async function onSubmit(data: z.infer<typeof formSchema>) {
    if (!turnstileToken && process.env.NEXT_PUBLIC_TURNSTILE_SITE_KEY) {
      toast.error("Lütfen doğrulamayı tamamlayın.");
      return;
    }

    const response = await shortenUrl(data.targetUrl, data.slug, turnstileToken);

    if (response.status === 201) {
      toast("URL başarıyla kısaltıldı!", {
        description: `Kısaltılmış URL: ${process.env.NEXT_PUBLIC_APP_URL}/${response.data.slug}`,
        action: {
          label: "Kopyala",
          onClick: () => {
            navigator.clipboard.writeText(process.env.NEXT_PUBLIC_APP_URL + '/' + response.data.slug);
            toast("Kısaltılmış URL panoya kopyalandı!");
          },
        },
      });
      setTurnstileToken(null);
      form.reset();
    } else {
      toast("Hay aksi, bir şeyler ters gitti.", {
        description: "Bir hata oluştu. Lütfen daha sonra tekrar deneyin.",
      });
    }
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-col gap-y-2 mt-10 rounded-lg md:w-1/2">
        <FormField
          control={form.control}
          name="targetUrl"
          render={({ field }) => (
            <FormItem>
              <FormControl>
                <Input
                  type="text"
                  placeholder="Kısaltmak istediğin URL'yi buraya yapıştır"
                  className="bg-card text-muted-foreground text-sm"
                  {...field}
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <div className="flex gap-x-2">
          <FormField
            control={form.control}
            name="slug"
            render={({ field }) => (
              <FormItem className="w-full">
                <FormControl>
                  <Input
                    type="text"
                    placeholder="Kısaltılmış URL (isteğe bağlı)"
                    className="bg-card text-muted-foreground text-sm"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit" variant="ghost" className="w-24 cursor-pointer animated-glow">
            Kısalt
          </Button>
        </div>
        <Turnstile onVerify={setTurnstileToken} />
      </form>
    </Form>
  );
}
