"use server";

import { authFetch } from "@/lib/api/authFetch";
import { getMe } from "@/lib/api/getMe";

export async function deleteNotification(
  id: number,
): Promise<{ success: boolean }> {
  const me = await getMe();
  const prefix =
    me?.user?.role === "DepartmentHead(manager)" ? "manager" : "staff";

  const { response } = await authFetch(`/${prefix}/notifications/${id}`, {
    method: "DELETE",
  });

  return { success: response.ok };
}
