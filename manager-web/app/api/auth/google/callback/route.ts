import { NextResponse } from "next/server";

const GOOGLE_EXCHANGE_URL = "http://localhost:8080/auth/google/exchange";

export async function GET(request: Request) {
  const url = new URL(request.url);

  try {
    const code = url.searchParams.get("code");
    const googleError = url.searchParams.get("error");

    // Google returned an error (access_denied, etc.) or no code at all
    if (googleError || !code) {
      const reason = googleError ?? "missing_code";
      return NextResponse.redirect(
        new URL(`/login?oauth_error=${encodeURIComponent(reason)}`, url),
      );
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

    return nextResponse;
  } catch (err) {
    console.error("OAuth callback unexpected error:", err);
    return NextResponse.redirect(
      new URL("/login?oauth_error=unexpected", url),
    );
  }
}
