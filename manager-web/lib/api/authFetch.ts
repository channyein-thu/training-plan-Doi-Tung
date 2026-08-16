"use server";

import { cookies } from "next/headers";
import { API_BASE_URL } from "./api";

export type AuthFetchResult = {
  response: Response;
  unauthorized: boolean;
};

export async function authFetch(
  path: string,
  init: RequestInit = {},
): Promise<AuthFetchResult> {
  const cookieStore = await cookies();
  const token = cookieStore.get("token")?.value;

  const headers = new Headers(init.headers);

  // ONLY send Authorization header
  if (token && !headers.has("Authorization")) {
    headers.set("Authorization", `Bearer ${token}`);
  }

  // Auto JSON
  if (!headers.has("Content-Type") && !(init.body instanceof FormData)) {
    headers.set("Content-Type", "application/json");
  }

  const url = path.startsWith("http") ? path : `${API_BASE_URL}${path}`;

  const response = await fetch(url, {
    ...init,
    headers,
  });

  const unauthorized = response.status === 401 || response.status === 403;

  return { response, unauthorized };
}

/**
 * Safely parse a Response as JSON.
 *
 * Returns the parsed body, or null if the body is missing or not JSON
 * (e.g. plain-text "Too Many Requests" from a rate limiter, gateway HTML
 * error pages, empty 204 responses, etc.).
 *
 * Use this instead of `await response.json()` anywhere a non-OK response
 * is being inspected for an error message.
 */
export async function safeJson<T = unknown>(
  response: Response,
): Promise<T | null> {
  try {
    const text = await response.text();
    if (!text) return null;
    return JSON.parse(text) as T;
  } catch {
    return null;
  }
}
