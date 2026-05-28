import PageHeader from "@/components/dashboard/pageHeader";
import NotificationsClient from "@/components/notification/NotificationsClient";
import PushToggle from "@/components/notification/PushToggle";
import { getNotifications } from "@/lib/actions/notifications/getNotifications";

export default async function NotificationPage() {
  const data = await getNotifications();

  return (
    <div className="min-h-screen space-y-4 m-2">
      <div className="flex flex-col md:flex-row md:items-start md:justify-between gap-3">
        <PageHeader title="Notifications" subtitle="View your notifications" />
        <PushToggle />
      </div>
      <NotificationsClient initialData={data} />
    </div>
  );
}
