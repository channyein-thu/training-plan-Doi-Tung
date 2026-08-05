import { NextResponse } from "next/server";
import { randomBytes } from "crypto";
import { PUBLIC_BACKEND_ORIGIN } from "@/lib/api/api";

// Backend endpoint that builds the Google auth URL. The browser is redirected
// here, so it must be the public origin (not the internal Docker service name).
// It accepts the `state` we commit to here so the value Google echoes back to
// our callback matches the cookie we set below.
const BACKEND_GOOGLE_LOGIN_URL = `${PUBLIC_BACKEND_ORIGIN}/auth/google/login`;

export const OAUTH_STATE_COOKIE = "oauth_state";

// Initiates Google OAuth from the frontend origin so the CSRF `state` cookie is
// set on the SAME origin the callback lands on. Without this, the callback
// could never verify state (the old flow set the cookie on the backend origin,
// which the callback can't read — so state was silently ignored).
export async function GET() {
  const state = randomBytes(32).toString("base64url");

  const redirectUrl = new URL(BACKEND_GOOGLE_LOGIN_URL);
  redirectUrl.searchParams.set("state", state);

  const response = NextResponse.redirect(redirectUrl);
  response.cookies.set(OAUTH_STATE_COOKIE, state, {
    httpOnly: true,
    sameSite: "lax",
    secure: process.env.NODE_ENV === "production",
    path: "/",
    // Short-lived: the user should complete the round-trip in minutes.
    maxAge: 600,
  });

  return response;
}
