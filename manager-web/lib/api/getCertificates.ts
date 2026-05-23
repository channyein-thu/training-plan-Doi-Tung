"use server";

import { API_BASE_URL } from "@/lib/api/api";
import type { Certificate } from "@/types/certificate";
import { cookies } from "next/headers";
import { authFetch } from "./authFetch";

export async function getCertificates(): Promise<Certificate[]> {
  const { response, unauthorized } = await authFetch("/staff/certificates", {
    method: "GET",
    next: { tags: ["certificates"] },
  });

  if (unauthorized) return [];

  const payload = await response.json();
  return (payload.data ?? []) as Certificate[];
}
