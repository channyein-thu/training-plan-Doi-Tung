"use server";

import { authFetch } from "@/lib/api/authFetch";
import { getMe } from "@/lib/api/getMe";

async function getPrefix(): Promise<string> {
  const me = await getMe();
  return me?.user?.role === "DepartmentHead(manager)" ? "manager" : "staff";
}

export async function getVapidPublicKey(): Promise<string | null> {
  const prefix = await getPrefix();
  const { response } = await authFetch(
    `/${prefix}/notifications/vapid-public-key`,
    { method: "GET", cache: "no-store" },
  );
  if (!response.ok) return null;
  const payload = await response.json();
  return (payload?.data?.publicKey as string) ?? null;
}

type SubscriptionPayload = {
  endpoint: string;
  keys: { p256dh: string; auth: string };
};

export async function savePushSubscription(
  sub: SubscriptionPayload,
): Promise<{ success: boolean; message?: string }> {
  const prefix = await getPrefix();
  const { response } = await authFetch(
    `/${prefix}/notifications/push-subscribe`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(sub),
    },
  );
  if (!response.ok) {
    const data = await response.json().catch(() => ({}));
    return { success: false, message: data?.message ?? "Subscribe failed" };
  }
  return { success: true };
}

export async function removePushSubscription(
  endpoint: string,
): Promise<{ success: boolean }> {
  const prefix = await getPrefix();
  const { response } = await authFetch(
    `/${prefix}/notifications/push-unsubscribe`,
    {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ endpoint }),
    },
  );
  return { success: response.ok };
}

export type PushTestResult = {
  vapidConfigured: boolean;
  subscriptionsFound: number;
  results: Array<{
    endpoint: string;
    statusCode: number;
    error?: string;
  }>;
};

export async function sendTestPush(): Promise<PushTestResult | null> {
  const prefix = await getPrefix();
  const { response } = await authFetch(
    `/${prefix}/notifications/push-test`,
    { method: "POST" },
  );
  if (!response.ok) return null;
  const payload = await response.json();
  return (payload?.data as PushTestResult) ?? null;
}
