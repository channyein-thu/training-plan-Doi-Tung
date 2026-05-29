"use server";

import { authFetch } from "@/lib/api/authFetch";
import { getMe } from "@/lib/api/getMe";
import { NotificationsResponse } from "@/types/notification";

async function getPrefix(): Promise<string> {
  const me = await getMe();
  return me?.user?.role === "DepartmentHead(manager)" ? "manager" : "staff";
}

export async function getNotifications(
  page = 1,
  limit = 20,
): Promise<NotificationsResponse | null> {
  const prefix = await getPrefix();

  const { response } = await authFetch(
    `/${prefix}/notifications?page=${page}&limit=${limit}`,
    { method: "GET", cache: "no-store" },
  );

  if (!response.ok) return null;

  const payload = await response.json();
  return (payload.data as NotificationsResponse) ?? null;
}
