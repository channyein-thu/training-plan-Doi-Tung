export type NotificationType =
  | "training_registered"
  | "certificate_approved"
  | "certificate_rejected";

export type Notification = {
  id: number;
  userId: number;
  type: NotificationType;
  title: string;
  message: string;
  isRead: boolean;
  createdAt: string;
};

export type NotificationMeta = {
  page: number;
  limit: number;
  totalItems: number;
  totalPages: number;
};

export type NotificationsResponse = {
  items: Notification[];
  meta: NotificationMeta;
};
