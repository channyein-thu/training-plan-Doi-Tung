"use server";

import { authFetch } from "@/lib/api/authFetch";

export async function getRecordById(recordId: string | number) {
  const { response, unauthorized } = await authFetch(
    `/staff/records/${recordId}`,
    {
      method: "GET",
      cache: "no-store",
    },
  );

  if (unauthorized) {
    return { unauthorized: true } as const;
  }

  if (response.status === 404) {
    return null;
  }

  //   if (!response.ok) {
  //     throw new Error("Failed to fetch record");
  //   }

  return response.json();
}
