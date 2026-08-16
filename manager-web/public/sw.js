// Service worker for Web Push notifications.
// This file is served at https://<origin>/sw.js — it must live at the origin root
// so it can receive push events for the whole site, even when no tab is open.

self.addEventListener("install", (event) => {
  // Activate immediately on first install so we don't have to wait
  self.skipWaiting();
});

self.addEventListener("activate", (event) => {
  event.waitUntil(self.clients.claim());
});

// Browser → "you have a push" → this handler runs even when no tab is open.
self.addEventListener("push", (event) => {
  let data = {};
  try {
    data = event.data ? event.data.json() : {};
  } catch (e) {
    data = { title: "Notification", message: event.data ? event.data.text() : "" };
  }

  const title = data.title || "Notification";
  const options = {
    body: data.message || "",
    icon: "/logo.png",
    badge: "/logo.png",
    tag: "training-notification", // collapses duplicate notifications
    renotify: true,
    data: { url: data.url || "/notifications" },
  };

  event.waitUntil(self.registration.showNotification(title, options));
});

// User clicks the notification → focus an existing tab or open a new one.
self.addEventListener("notificationclick", (event) => {
  event.notification.close();
  const targetUrl = (event.notification.data && event.notification.data.url) || "/notifications";

  event.waitUntil(
    self.clients.matchAll({ type: "window", includeUncontrolled: true }).then((clientsArr) => {
      // Try to focus an existing tab on the same origin
      for (const client of clientsArr) {
        if ("focus" in client) {
          client.navigate(targetUrl);
          return client.focus();
        }
      }
      // Otherwise open a new tab
      if (self.clients.openWindow) {
        return self.clients.openWindow(targetUrl);
      }
    }),
  );
});
