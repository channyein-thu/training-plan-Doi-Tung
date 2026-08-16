"use server";

import { authFetch } from "@/lib/api/authFetch";
import { BACKEND_ORIGIN } from "@/lib/api/api";

const COMPLETE_PROFILE_URL = `${BACKEND_ORIGIN}/user/complete-profile`;

export type CompleteProfilePayload = {
  employeeId: string;
  departmentId: number;
  phone?: string;
  position?: string;
};

export type CompleteProfileResult = {
  ok: boolean;
  message?: string;
};

export async function completeProfileAction(
  payload: CompleteProfilePayload,
): Promise<CompleteProfileResult> {
  try {
    const { response } = await authFetch(COMPLETE_PROFILE_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
      cache: "no-store",
    });

    if (!response.ok) {
      const data = await response.json().catch(() => null);
      return {
        ok: false,
        message: data?.message || "Profile completion failed.",
      };
    }

    return { ok: true };
  } catch {
    return { ok: false, message: "Profile completion failed." };
  }
}
