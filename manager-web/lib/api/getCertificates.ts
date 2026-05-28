"use server";

import type { Certificate } from "@/types/certificate";
import { authFetch } from "./authFetch";

// Always returns a fresh array. Never throws.
//
// IMPORTANT: certificates get approved/rejected by the admin app on the
// backend directly, which means the manager app has NO way to know to bust
// a cache. So this MUST be `cache: "no-store"` — otherwise staff sees
// stale "Pending" status forever even after admin approval.
export async function getCertificates(): Promise<Certificate[]> {
  try {
    const { response, unauthorized } = await authFetch("/staff/certificates", {
      method: "GET",
      cache: "no-store",
    });

    if (unauthorized) return [];

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
