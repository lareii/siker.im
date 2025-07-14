import { NextResponse, type NextRequest } from 'next/server';

const isStaticAsset = (pathname: string) => {
  return !!pathname.match(/\.(.*)$/);
};

const isSystemPath = (pathname: string) => {
  return pathname.startsWith("/_next");
};

async function getUrl(slug: string): Promise<{ target_url?: string } | null> {
  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/urls/${slug}`, {
      cache: "no-store",
    });
    if (!res.ok) return null;
    const data = await res.json();
    return data;
  } catch {
    return null;
  }
}

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  if (pathname === "/" || isStaticAsset(pathname) || isSystemPath(pathname)) {
    return NextResponse.next();
  }

  const segments = pathname.split("/").filter(Boolean);

  if (segments.length === 1) {
    const slug = segments[0];
    const data = await getUrl(slug);

    if (data?.target_url) {
      return NextResponse.redirect(data.target_url);
    } else {
      return NextResponse.redirect(new URL("/", req.url));
    }
  }

  return NextResponse.next();
}