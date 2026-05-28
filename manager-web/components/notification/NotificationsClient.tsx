"use client";

import React, { useTransition } from "react";
import { useRouter } from "next/navigation";
import { NotificationsResponse } from "@/types/notification";
import NotificationCard from "./notificationCard";
import { markAsRead } from "@/lib/actions/notifications/markAsRead";
import { markAllRead } from "@/lib/actions/notifications/markAllRead";
import { deleteNotification } from "@/lib/actions/notifications/deleteNotification";
import { Button } from "@/components/ui/button";

type Props = {
  initialData: NotificationsResponse | null;
};

export default function NotificationsClient({ initialData }: Props) {
  const router = useRouter();
  const [isPending, startTransition] = useTransition();

  const notifications = initialData?.items ?? [];
  const hasUnread = notifications.some((n) => !n.isRead);

  const handleMarkAsRead = (id: number) => {
    startTransition(async () => {
      await markAsRead(id);
      router.refresh();
    });
  };

  const handleMarkAllRead = () => {
    startTransition(async () => {
      await markAllRead();
      router.refresh();
    });
  };

  const handleDelete = (id: number) => {
    startTransition(async () => {
      await deleteNotification(id);
      router.refresh();
    });
  };

  if (notifications.length === 0) {
    return (
      <p className="text-center text-gray-500 py-12">No notifications yet.</p>
    );
  }

  return (
    <div>
      {hasUnread && (
        <div className="flex justify-end mb-4">
          <Button
            variant="outline"
            size="sm"
            disabled={isPending}
            onClick={handleMarkAllRead}
            className="border-[#006022] text-[#006022] hover:bg-[#006022] hover:text-white"
          >
            Mark all as read
          </Button>
        </div>
      )}

      <div className="space-y-2">
        {notifications.map((n) => (
          <NotificationCard
            key={n.id}
            notification={n}
            onMarkAsRead={() => handleMarkAsRead(n.id)}
            onDelete={() => handleDelete(n.id)}
            disabled={isPending}
          />
        ))}
      </div>
    </div>
  );
}
