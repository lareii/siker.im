import { NextResponse, type NextRequest } from 'next/server';

const isStaticAsset = (pathname: string) => {
  return !!pathname.match(/\.(.*)$/);
};

const isSystemPath = (pathname: string) => {
  return pathname.startsWith("/_next");
};

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  if (pathname === "/" || isStaticAsset(pathname) || isSystemPath(pathname)) {
    return NextResponse.next();
  }

  const segments = pathname.split("/").filter(Boolean);
  if (segments.length === 1) {
    const slug = segments[0];
    const redirectUrl = new URL(`/redirect/${slug}`, process.env.NEXT_PUBLIC_API_URL);
    return NextResponse.redirect(redirectUrl);
  }

  return NextResponse.next();
}