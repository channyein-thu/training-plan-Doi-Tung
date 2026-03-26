"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";

export async function getRecordById(recordId: string | number) {
  const cookieStore = await cookies();
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const response = await fetch(`${API_BASE_URL}/staff/records/${recordId}`, {
    method: "GET",
    credentials: "include",
    headers: { "Content-Type": "application/json", Cookie: cookieHeader },
    cache: "no-store",
  });

  if (response.status === 404) {
    return null;
  }

  if (!response.ok) {
    throw new Error("Failed to fetch record");
  }

  return response.json();
}
