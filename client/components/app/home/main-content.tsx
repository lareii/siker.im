import AnimatedTextSwitcher from '@/components/app/home/text-switcher';
import Shortener from '@/components/app/home/shortener';
import AnnouncementBanner from '@/components/app/home/announcement-banner';
import Features from '@/components/app/home/features';

import Image from 'next/image';

export default function MainContent() {
  return (
    <main className="p-5 max-w-screen-lg mx-auto max-sm:mt-10 mt-20 relative">
      <div>
        <div className="relative z-20">
          <AnnouncementBanner />
          <div>
            <h1 className="text-4xl font-black">
              s**erim böyle linki diyorsan,
            </h1>
            <div className="flex w-full">
              <h1 className="text-4xl font-black text-primary">siker.im/</h1>
              <AnimatedTextSwitcher />
            </div>
            <h1 className="text-4xl font-black">tam sana göre!</h1>
            <div className="text-muted-foreground font-light mt-5">
              <p>bazen bir link paylaşırsın ama o link uzar da uzar.</p>
              <p>
                <span className="text-primary">siker.im</span>, URL&apos;lerini
                hızlıca kısaltmana yarayan basit bir araçtır.
              </p>
            </div>
          </div>
          <Shortener />
        </div>
        <div className="max-lg:opacity-10 max-md:hidden">
          <Image
            src="/undraw.svg"
            alt="vector image"
            width={400}
            height={400}
            className="absolute top-0 right-0 z-10"
          />
          <Image
            src="/blob.svg"
            alt="vector image"
            width={500}
            height={500}
            className="absolute top-0 right-0"
          />
        </div>
      </div>
      <Features />
    </main>
  );
}
