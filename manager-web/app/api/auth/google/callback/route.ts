import { NextResponse } from "next/server";
import { authFetch } from "@/lib/api/authFetch";

const GOOGLE_EXCHANGE_URL = "http://localhost:8080/auth/google/exchange";

export async function GET(request: Request) {
  try {
    const url = new URL(request.url);
    const code = url.searchParams.get("code");

    if (!code) {
      return NextResponse.redirect(new URL("/login", url));
    }

    const { response } = await authFetch(GOOGLE_EXCHANGE_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ code }),
      cache: "no-store",
    });

    if (!response.ok) {
      return NextResponse.redirect(new URL("/login", url));
    }

    const data = await response.json();
    const accessToken = data?.accessToken;
    const isProfileComplete = data?.isProfileComplete;

    if (!accessToken) {
      return NextResponse.redirect(new URL("/login", url));
    }

    const redirectPath =
      isProfileComplete === false ? "/onboarding" : "/dashboard";
    const nextResponse = NextResponse.redirect(new URL(redirectPath, url));

    nextResponse.cookies.set({
      name: "token",
      value: accessToken,
      httpOnly: true,
      sameSite: "lax",
      secure: process.env.NODE_ENV === "production",
      path: "/",
    });

    return nextResponse;
  } catch {
    return NextResponse.redirect(new URL("/login", request.url));
  }
}

// import { NextResponse } from "next/server";
// import type { NextRequest } from "next/server";

// export function middleware(request: NextRequest) {
//   const token = request.cookies.get("token")?.value;

//   if (!token) {
//     return NextResponse.redirect(new URL("/login", request.url));
//   }

//   return NextResponse.next();
// }

// export const config = {
//   matcher: [
//     "/dashboard/:path*",
//     "/my-certificates/:path*",
//     "/my-training-records/:path*",
//   ],
// };
