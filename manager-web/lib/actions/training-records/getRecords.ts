"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";

export async function getRecords() {
  const cookieStore = await cookies();
  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const response = await fetch(`${API_BASE_URL}/staff/records`, {
    method: "GET",
    credentials: "include",
    headers: { "Content-Type": "application/json", Cookie: cookieHeader },
    next: { tags: ["records"] },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch records");
  }

  return response.json();
}
