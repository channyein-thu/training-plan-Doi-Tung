"use server";

import { redirect } from "next/navigation";
import { cookies } from "next/headers";
import { z } from "zod";
import { API_BASE_URL } from "../api/api";

export type State = {
  errors?: {
    email?: string[];
    password?: string[];
  };
  message: string;
};

const FormSchema = z.object({
  email: z.string().trim().email("Invalid email address!"),
  password: z.string().trim().min(6, "Password must be at least 6 characters!"),
});

export async function LoginAction(
  prevState: State,
  formData: FormData,
): Promise<State> {
  const validatedFields = FormSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });

  // Validation error
  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
      message: "Invalid input",
    };
  }

  const { email, password } = validatedFields.data;

  const response = await fetch(`${API_BASE_URL}/auth/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
    cache: "no-store",
  });

  if (!response.ok) {
    // Use .text() + try/catch so non-JSON bodies (rate-limit plain text, HTML
    // error pages, etc.) don't crash JSON.parse.
    const text = await response.text();
    let parsedMessage: string | undefined;
    try {
      parsedMessage = JSON.parse(text)?.message;
    } catch {
      // Fall through — body was not JSON.
    }
    return {
      message:
        parsedMessage ??
        (response.status === 429
          ? "Too many requests. Please wait a moment and try again."
          : "Login failed"),
    };
  }
  const data = await response.json().catch(() => ({}));
  const accessToken = (data as { accessToken?: string })?.accessToken;

  if (!accessToken) {
    return {
      message: "Login failed",
    };
  }

  const cookieStore = await cookies();

  cookieStore.set("token", accessToken, {
    httpOnly: true,
    sameSite: "lax",
    secure: process.env.NODE_ENV === "production",
    path: "/",
    // Keep the session alive for the full JWT lifetime (7 days).
    // Without maxAge the cookie is a session cookie and dies on browser close,
    // which is why "the token feels too short" — it's the cookie, not the JWT.
    maxAge: 60 * 60 * 24 * 7,
  });

  redirect("/dashboard");
}
