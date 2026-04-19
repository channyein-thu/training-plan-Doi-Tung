"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { Certificate } from "@/types/certificate";
import { cookies } from "next/headers";

export async function getCertificates(): Promise<Certificate[]> {
  const cookieStore = await cookies();
  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const { response } = await authFetch(`${API_BASE_URL}/staff/certificates`, {
    method: "GET",
    credentials: "include",
    headers: { "Content-Type": "application/json", Cookie: cookieHeader },
    next: { tags: ["certificates"] },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch certificates");
  }

  const payload = await response.json();
  return (payload.data ?? []) as Certificate[];
}
