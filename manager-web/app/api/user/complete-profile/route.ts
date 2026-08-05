import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { authFetch } from "@/lib/api/authFetch";
import { BACKEND_ORIGIN } from "@/lib/api/api";

const COMPLETE_PROFILE_URL = `${BACKEND_ORIGIN}/user/complete-profile`;

export async function POST(request: Request) {
  const cookieStore = await cookies();
  const token = cookieStore.get("token")?.value;

  if (!token) {
    return NextResponse.json({ message: "Unauthorized" }, { status: 401 });
  }

  try {
    const body = await request.json();
    const { response } = await authFetch(COMPLETE_PROFILE_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
      cache: "no-store",
    });

    const contentType = response.headers.get("content-type") || "";
    if (contentType.includes("application/json")) {
      const data = await response.json();
      return NextResponse.json(data, { status: response.status });
    }

    const text = await response.text();
    return new NextResponse(text, { status: response.status });
  } catch {
    return NextResponse.json(
      { message: "Failed to complete profile" },
      { status: 500 },
    );
  }
}
