import {
  IconPointerFilled,
  IconLink,
  IconBrandGithubFilled,
  IconBoltFilled
} from '@tabler/icons-react';
import Link from 'next/link';

export default function Features() {
  return (
    <div className="mt-54 grid md:grid-cols-3 gap-4">
      <div className="bg-linear-to-t from-background to-card rounded-md h-60 px-5 flex flex-col justify-center">
        <IconLink className="w-16 h-16 text-primary mb-4" />
        <h2 className="text-primary text-5xl font-black">10+</h2>
        <p className="text-xl">kısaltılmış URL</p>
        <p className="text-muted-foreground text-sm">
          sizler tarafından kısaltılmış URL’ler
        </p>
      </div>
      <div className="bg-linear-to-b from-background to-card rounded-md h-60 md:col-span-2 px-5 flex flex-col justify-center">
        <IconBoltFilled className="w-16 h-16 text-primary mb-4" />
        <h3 className="text-xl">hızlı ve kolay kullanım</h3>
        <p className="text-muted-foreground text-sm">
          hesap oluşturma, giriş yapma gibi işlemlerle uğraşmadan anında
          kısaltma yapabilirsin.
        </p>
      </div>
      <Link
        href="https://github.com/lareii/siker.im"
        className="bg-linear-to-t from-background to-card rounded-md h-60 md:col-span-2 px-5 flex flex-col justify-center"
      >
        <IconBrandGithubFilled className="w-16 h-16 text-primary mb-4" />
        <h3 className="text-xl">özgür yazılım</h3>
        <p className="text-muted-foreground text-sm">
          bu projeyi GitHub&apos;dan inceleyebilir ve katkıda bulunabilirsin.
        </p>
      </Link>
      <div className="bg-linear-to-b from-background to-card rounded-md h-60 px-5 flex flex-col justify-center">
        <IconPointerFilled className="w-16 h-16 text-primary mb-4" />
        <h2 className="text-primary text-5xl font-black">100+</h2>
        <p className="text-xl">tıklama</p>
        <p className="text-muted-foreground text-sm">
          kısaltılmış URL’lere yapılan tıklamalar
        </p>
      </div>
    </div>
  );
}
