import {
  FormControl,
  FormField,
  FormItem,
  FormMessage
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';

import { Control } from 'react-hook-form';

interface Props {
  control: Control<any>;
  isLoading: boolean;
}

export function TargetUrlInput({ control, isLoading }: Props) {
  return (
    <FormField
      control={control}
      name="targetUrl"
      render={({ field }) => (
        <FormItem>
          <FormControl>
            <Input
              type="text"
              placeholder="kısaltmak istediğin URL'yi buraya yapıştır"
              className="bg-card text-muted-foreground text-sm"
              {...field}
              disabled={isLoading}
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}

export function SlugInput({ control, isLoading }: Props) {
  return (
    <FormField
      control={control}
      name="slug"
      render={({ field }) => (
        <FormItem className="w-full">
          <FormControl>
            <Input
              type="text"
              placeholder="etiket (isteğe bağlı)"
              className="bg-card text-muted-foreground text-sm"
              {...field}
              disabled={isLoading}
            />
          </FormControl>
          <FormMessage />
        </FormItem>
      )}
    />
  );
}
