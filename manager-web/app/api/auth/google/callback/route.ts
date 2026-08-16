import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { BACKEND_ORIGIN } from "@/lib/api/api";

const GOOGLE_EXCHANGE_URL = `${BACKEND_ORIGIN}/auth/google/exchange`;
const OAUTH_STATE_COOKIE = "oauth_state";

export async function GET(request: Request) {
  const url = new URL(request.url);

  try {
    const code = url.searchParams.get("code");
    const googleError = url.searchParams.get("error");
    const returnedState = url.searchParams.get("state");

    // Google returned an error (access_denied, etc.) or no code at all
    if (googleError || !code) {
      const reason = googleError ?? "missing_code";
      return NextResponse.redirect(
        new URL(`/login?oauth_error=${encodeURIComponent(reason)}`, url),
      );
    }

    // CSRF protection: the `state` echoed back by Google must match the
    // one-time value we committed to a cookie when starting the flow. A missing
    // or mismatched state means this callback was not initiated by this browser
    // — reject it before doing anything with the code.
    const cookieStore = await cookies();
    const expectedState = cookieStore.get(OAUTH_STATE_COOKIE)?.value;
    if (!expectedState || !returnedState || returnedState !== expectedState) {
      const res = NextResponse.redirect(
        new URL("/login?oauth_error=state_mismatch", url),
      );
      res.cookies.delete(OAUTH_STATE_COOKIE);
      return res;
    }

    // Exchange the code with our backend.
    // This is a public endpoint — no auth header needed.
    const exchangeRes = await fetch(GOOGLE_EXCHANGE_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ code }),
      cache: "no-store",
    });

    if (!exchangeRes.ok) {
      const text = await exchangeRes.text().catch(() => "");
      console.error("OAuth exchange failed:", exchangeRes.status, text);
      return NextResponse.redirect(
        new URL("/login?oauth_error=exchange_failed", url),
      );
    }

    const data = await exchangeRes.json();
    const accessToken = data?.accessToken;
    const isProfileComplete = data?.isProfileComplete;

    if (!accessToken) {
      return NextResponse.redirect(
        new URL("/login?oauth_error=no_token", url),
      );
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
      // Match JWT lifetime (7 days) so OAuth users stay logged in
      maxAge: 60 * 60 * 24 * 7,
    });

    // The one-time state has served its purpose; clear it.
    nextResponse.cookies.delete(OAUTH_STATE_COOKIE);

    return nextResponse;
  } catch (err) {
    console.error("OAuth callback unexpected error:", err);
    return NextResponse.redirect(
      new URL("/login?oauth_error=unexpected", url),
    );
  }
}
