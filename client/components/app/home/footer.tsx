import Image from 'next/image';
import Link from 'next/link';

export default function Footer() {
  return (
    <footer className="mt-20 mb-5">
      <div className="max-w-screen-lg mx-auto flex justify-end gap-x-5 items-center px-5">
        <div className="text-muted-foreground text-xs">
          a{' '}
          <Link
            href="https://lareii.github.io"
            className="text-primary underline font-mono"
          >
            larei
          </Link>{' '}
          production <br /> created by human :)
        </div>
        <Link href="https://buymeacoffee.com/larei">
          <Image
            src="/buymeacoffee.png"
            alt="Buy Me a Coffee"
            width={128}
            height={0}
          />
        </Link>
      </div>
    </footer>
  );
}
