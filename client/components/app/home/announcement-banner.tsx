import Link from 'next/link';
import { IconConfettiFilled } from '@tabler/icons-react';

export default function AnnouncementBanner() {
  return (
    <Link
      href="https://github.com/lareii/siker.im"
      className="mb-2 text-sm w-fit px-3 py-1 rounded-full border bg-radial from-background to-neutral-900 flex items-center gap-x-2"
    >
      <IconConfettiFilled className="w-4 h-4" />
      ilk sürüm yayında.
    </Link>
  );
}
