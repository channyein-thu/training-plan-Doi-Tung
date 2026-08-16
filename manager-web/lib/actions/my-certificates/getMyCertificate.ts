"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { Certificate } from "@/types/certificate";

// Always returns an array. Never throws — on any error returns [].
export async function getCertificates(): Promise<Certificate[]> {
  try {
    const { response } = await authFetch(`${API_BASE_URL}/staff/certificates`, {
      method: "GET",
      next: { tags: ["certificates"] },
    });

    if (!response.ok) {
      console.warn(`getCertificates: backend returned ${response.status}`);
      return [];
    }

    const payload = await response.json().catch(() => null);
    return ((payload?.data ?? []) as Certificate[]) || [];
  } catch (err) {
    console.error("getCertificates: unexpected error", err);
    return [];
  }
}
