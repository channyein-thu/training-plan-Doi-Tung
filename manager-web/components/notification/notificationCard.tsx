"use client";

import React from "react";
import { Notification } from "@/types/notification";
import {
  AwardIcon,
  BellIcon,
  BookOpenIcon,
  CheckIcon,
  Trash2Icon,
} from "lucide-react";
import { Button } from "@/components/ui/button";

function typeIcon(type: Notification["type"]) {
  switch (type) {
    case "training_registered":
      return <BookOpenIcon className="w-5 h-5 text-blue-600" />;
    case "certificate_approved":
      return <AwardIcon className="w-5 h-5 text-green-600" />;
    case "certificate_rejected":
      return <AwardIcon className="w-5 h-5 text-red-500" />;
    default:
      return <BellIcon className="w-5 h-5 text-gray-500" />;
  }
}

function formatDate(iso: string) {
  return new Date(iso).toLocaleString("en-GB", {
    day: "2-digit",
    month: "short",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

type Props = {
  notification: Notification;
  onMarkAsRead: () => void;
  onDelete: () => void;
  disabled?: boolean;
};

function NotificationCard({ notification, onMarkAsRead, onDelete, disabled }: Props) {
  const { title, message, isRead, type, createdAt } = notification;

  return (
    <div
      className={`border rounded-md p-4 flex items-start gap-4 transition-colors ${
        isRead ? "bg-white" : "bg-blue-50 border-blue-200"
      }`}
    >
      <div className="mt-0.5 shrink-0">{typeIcon(type)}</div>

      <div className="flex-1 min-w-0">
        <p className={`font-medium ${isRead ? "text-gray-800" : "text-gray-900"}`}>
          {title}
        </p>
        <p className="text-sm text-gray-600 mt-0.5">{message}</p>
        <p className="text-xs text-gray-400 mt-1">{formatDate(createdAt)}</p>
      </div>

      <div className="flex items-center gap-1 shrink-0">
        {!isRead && (
          <Button
            variant="ghost"
            size="icon"
            className="h-7 w-7 text-gray-400 hover:text-green-600"
            disabled={disabled}
            onClick={onMarkAsRead}
            title="Mark as read"
          >
            <CheckIcon className="w-4 h-4" />
          </Button>
        )}
        <Button
          variant="ghost"
          size="icon"
          className="h-7 w-7 text-gray-400 hover:text-red-500"
          disabled={disabled}
          onClick={onDelete}
          title="Delete"
        >
          <Trash2Icon className="w-4 h-4" />
        </Button>
      </div>
    </div>
  );
}

export default NotificationCard;
