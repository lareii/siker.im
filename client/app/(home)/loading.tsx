import { IconInnerShadowBottomLeftFilled } from '@tabler/icons-react';

export default function TestMotion() {
  return (
    <div className="flex flex-col gap-y-2 items-center justify-center h-screen">
      <div className="text-4xl font-bold text-primary">siker.im/</div>
      <div className="flex items-center gap-x-2 text-sm">
        <IconInnerShadowBottomLeftFilled className="animate-spin w-5 h-5 text-primary" />
        yükleniyor...
      </div>
    </div>
  );
}
