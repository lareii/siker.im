import {
  Tooltip,
  TooltipContent,
  TooltipTrigger
} from '@/components/ui/tooltip';

import Link from 'next/link';

export default function Navbar() {
  return (
    <div className="border-b backdrop-blur-xl sticky top-0 z-30">
      <div className="p-5 max-w-screen-lg mx-auto flex items-center justify-between">
        <Link href="/" className="text-2xl font-bold text-primary">
          siker.im/
        </Link>
        <Tooltip>
          <TooltipTrigger asChild>
            <div className="select-none bg-primary/50 rounded-lg px-4 py-2 text-sm font-medium text-accent">
              giriş yap
            </div>
          </TooltipTrigger>
          <TooltipContent>
            <p>henüz aktif değil</p>
          </TooltipContent>
        </Tooltip>
      </div>
    </div>
  );
}
