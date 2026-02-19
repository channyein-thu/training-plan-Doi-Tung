import { BookOpenIcon, LineChartIcon, UsersIcon } from "lucide-react";
import PageHeader from "@/components/dashboard/pageHeader";
import DashboardCard from "@/components/dashboard/dashboardCard";

export default function DashboardPage() {
  return (
    <div className="h-screen overflow-y-auto p-4 m-2 space-y-4">
      <PageHeader
        title="Manager Dashboard"
        subtitle="Department overview and personal training management"
      />
      <div className="flex flex-wrap gap-10">
        <DashboardCard
          icon={<UsersIcon className="w-6 h-6" />}
          percentChange="+12%"
          count={245}
          description="Department Staff"
        />

        <DashboardCard
          icon={<BookOpenIcon className="w-6 h-6" />}
          percentChange="+8%"
          count={120}
          description="Active Training"
        />

        <DashboardCard
          icon={<LineChartIcon className="w-6 h-6" />}
          percentChange="+5%"
          count={32}
          description="Avg Training Hours"
        />
      </div>
    </div>
  );
}
