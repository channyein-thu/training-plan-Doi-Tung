"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { cookies } from "next/headers";

export async function registerStaffAction(planId: string, staffIds: number[]) {
  const cookieStore = await cookies();

  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const { response } = await authFetch(
    `${API_BASE_URL}/manager/training-plans/${planId}/registrations`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
        Cookie: cookieHeader,
      },
      body: JSON.stringify({
        userIds: staffIds,
      }),
      cache: "no-store",
    },
  );

  if (!response.ok) {
    const errorBody = await response.text();
    throw new Error(
      `Failed to register staff: ${response.status} ${response.statusText} ${errorBody}`.trim(),
    );
  }

  return { success: true };
}
