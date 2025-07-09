import {
  IconPointerFilled,
  IconLink,
  IconBrandGithubFilled,
  IconBoltFilled
} from '@tabler/icons-react';

export default function Features() {
  return (
    <div className="mt-54 grid md:grid-cols-3 gap-4">
      <div className="bg-linear-to-t from-background to-card rounded-md h-60 px-5 flex flex-col justify-center border">
        <IconLink className="w-16 h-16 text-primary mb-4" />
        <h2 className="text-primary text-5xl font-black">10+</h2>
        <p className="text-xl">Kısaltılmış URL</p>
        <p className="text-muted-foreground text-sm">
          Sizler tarafından kısaltılmış URL’ler
        </p>
      </div>
      <div className="bg-linear-to-b from-background to-card rounded-md h-60 md:col-span-2 px-5 flex flex-col justify-center border">
        <IconBoltFilled className="w-16 h-16 text-primary mb-4" />
        <h3 className="text-xl">Hızlı ve Kolay Kullanım</h3>
        <p className="text-muted-foreground text-sm">
          Hesap oluşturma, giriş yapma gibi işlemlerle uğraşmadan anında
          kısaltma yapabilirsin.
        </p>
      </div>
      <div className="bg-linear-to-t from-background to-card rounded-md h-60 md:col-span-2 px-5 flex flex-col justify-center border">
        <IconBrandGithubFilled className="w-16 h-16 text-primary mb-4" />
        <h3 className="text-xl">Açık Kaynak</h3>
        <p className="text-muted-foreground text-sm">
          Bu projeyi GitHub&apos;dan inceleyebilir ve katkıda bulunabilirsin.
        </p>
      </div>
      <div className="bg-linear-to-b from-background to-card rounded-md h-60 px-5 flex flex-col justify-center border">
        <IconPointerFilled className="w-16 h-16 text-primary mb-4" />
        <h2 className="text-primary text-5xl font-black">100+</h2>
        <p className="text-xl">Tıklama</p>
        <p className="text-muted-foreground text-sm">
          Kısaltılmış URL’lere yapılan tıklamalar
        </p>
      </div>
    </div>
  );
}
