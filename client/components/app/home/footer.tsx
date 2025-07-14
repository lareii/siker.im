import Link from 'next/link';
import {IconBrandGithubFilled, IconMugFilled} from '@tabler/icons-react';

export default function Footer() {
  return (
    <footer className="mt-20 mb-5">
      <div className="max-w-screen-lg mx-auto flex max-md:flex-col justify-between md:items-center px-5 text-xs text-muted-foreground">
        <div className='flex justify-between items-center'>
          <Link href="/" className="text-lg font-bold text-primary">
            siker.im/
          </Link>
          <div className='flex md:hidden gap-x-5'>
            <Link href="https://buymeacoffee.com/larei" className='text-xs text-muted-foreground items-center flex gap-x-1'>
              <IconMugFilled className='inline w-3.5 h-3.5' /> Buy Me a Coffee
            </Link>
            <Link href="https://github.com/lareii/siker.im" className='text-xs text-muted-foreground items-center flex gap-x-1'>
              <IconBrandGithubFilled className='inline w-3.5 h-3.5' /> GitHub
            </Link>
          </div>
        </div>
        <div className='flex gap-x-10 max-md:hidden'>
          <Link href="https://buymeacoffee.com/larei" className='text-xs text-muted-foreground items-center flex gap-x-1'>
            <IconMugFilled className='inline w-3.5 h-3.5' /> Buy Me a Coffee
          </Link>
          <Link href="https://github.com/lareii/siker.im" className='text-xs text-muted-foreground items-center flex gap-x-1'>
            <IconBrandGithubFilled className='inline w-3.5 h-3.5' /> GitHub
          </Link>
        </div>
        <hr className='my-3 md:hidden' />
        <div className='flex flex-col md:items-end'>
          <div>
            a <Link href="https://lareii.github.io" className='underline text-primary'>larei</Link> production
          </div>
          <div>
            siker.im, <Link href="https://github.com/lareii/siker.im/blob/master/LICENSE" className='underline text-primary'>AGPL-3.0</Link> altında lisanslanmıştır.
          </div>
        </div>
      </div>
    </footer>
  );
}
