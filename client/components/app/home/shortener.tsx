import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';

export default function Shortener() {
  return (
    <div className="flex flex-col gap-y-2 mt-10 rounded-lg md:w-1/2">
      <Input
        type="text"
        placeholder="Kısaltmak istediğin URL'yi buraya yapıştır"
        className="bg-card text-muted-foreground text-sm"
      />
      <div className="flex gap-x-2">
        <Input
          type="text"
          placeholder="Kısaltılmış URL (isteğe bağlı)"
          className="bg-card text-muted-foreground text-sm"
        />
        <Button className="cursor-pointer">kısalt</Button>
      </div>
      <div className="text-muted-foreground text-xs">
        Boş bırakırsan rastgele bir kısaltma oluşturulur.
      </div>
    </div>
  );
}
