"use client";

import {
  getVapidPublicKey,
  savePushSubscription,
  removePushSubscription,
} from "@/lib/actions/notifications/push";

// VAPID keys are URL-safe base64; pushManager.subscribe needs a Uint8Array.
function urlBase64ToUint8Array(base64String: string): Uint8Array {
  const padding = "=".repeat((4 - (base64String.length % 4)) % 4);
  const base64 = (base64String + padding).replace(/-/g, "+").replace(/_/g, "/");
  const raw = window.atob(base64);
  const arr = new Uint8Array(raw.length);
  for (let i = 0; i < raw.length; i++) arr[i] = raw.charCodeAt(i);
  return arr;
}

export function isPushSupported(): boolean {
  if (typeof window === "undefined") return false;
  return (
    "serviceWorker" in navigator &&
    "PushManager" in window &&
    "Notification" in window
  );
}

export async function getCurrentSubscription(): Promise<PushSubscription | null> {
  if (!isPushSupported()) return null;
  const reg = await navigator.serviceWorker.getRegistration("/sw.js");
  if (!reg) return null;
  return reg.pushManager.getSubscription();
}

export async function enablePush(): Promise<{ ok: boolean; error?: string }> {
  if (!isPushSupported()) {
    return { ok: false, error: "This browser does not support push notifications." };
  }

  // 1. Ask the user for permission
  const permission = await Notification.requestPermission();
  if (permission !== "granted") {
    return { ok: false, error: "Notification permission was not granted." };
  }

  // 2. Register the service worker (idempotent)
  const reg = await navigator.serviceWorker.register("/sw.js");
  await navigator.serviceWorker.ready;

  // 3. Fetch the VAPID public key from the backend
  const publicKey = await getVapidPublicKey();
  if (!publicKey) {
    return { ok: false, error: "Server is missing VAPID configuration." };
  }

  // 4. Subscribe
  const sub = await reg.pushManager.subscribe({
    userVisibleOnly: true,
    applicationServerKey: urlBase64ToUint8Array(publicKey),
  });

  // 5. Send the subscription to the backend
  const json = sub.toJSON();
  const result = await savePushSubscription({
    endpoint: sub.endpoint,
    keys: {
      p256dh: json.keys?.p256dh ?? "",
      auth: json.keys?.auth ?? "",
    },
  });

  if (!result.success) {
    // Roll back the local subscription if the backend rejected it
    await sub.unsubscribe();
    return { ok: false, error: result.message ?? "Failed to save subscription." };
  }

  return { ok: true };
}

export async function disablePush(): Promise<{ ok: boolean }> {
  const sub = await getCurrentSubscription();
  if (!sub) return { ok: true };

  const endpoint = sub.endpoint;
  await sub.unsubscribe();
  await removePushSubscription(endpoint);
  return { ok: true };
}
