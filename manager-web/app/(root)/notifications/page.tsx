import PageHeader from "@/components/dashboard/pageHeader";
import NotificationCard from "@/components/notification/notificationCard";
import { BellIcon } from "lucide-react";
import React from "react";

function NotificationPage() {
  return (
    <div className="m-6 space-y-4">
      <PageHeader title="Notifications" subtitle="View your notifications" />

      <NotificationCard
        icon={<BellIcon />}
        title="New Message"
        subtitle="You have received a new message."
      />
    </div>
  );
}

export default NotificationPage;
