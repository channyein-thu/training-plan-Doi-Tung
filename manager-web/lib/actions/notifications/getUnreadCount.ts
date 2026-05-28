"use server";

import { authFetch } from "@/lib/api/authFetch";
import { getMe } from "@/lib/api/getMe";

export async function getUnreadCount(): Promise<number> {
  const me = await getMe();
  const prefix =
    me?.user?.role === "DepartmentHead(manager)" ? "manager" : "staff";

  const { response } = await authFetch(
    `/${prefix}/notifications/unread-count`,
    { method: "GET", cache: "no-store" },
  );

  if (!response.ok) return 0;

  const payload = await response.json();
  return Number(payload?.data?.unread ?? 0);
}
