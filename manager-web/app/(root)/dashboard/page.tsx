import {
  BookOpenIcon,
  UsersIcon,
  AwardIcon,
  CheckCircleIcon,
  GraduationCapIcon,
} from "lucide-react";
import PageHeader from "@/components/dashboard/pageHeader";
import DashboardCard from "@/components/dashboard/dashboardCard";
import TrainingStatusChart from "@/components/dashboard/trainingStatusChart";
import { getMe } from "@/lib/api/getMe";
import {
  getManagerDashboardStats,
  getStaffDashboardStats,
} from "@/lib/api/getDashboardStats";

export default async function DashboardPage() {
  const me = await getMe();
  const user = me?.user;
  let title = "Dashboard";
  let subtitle = "Overview";

  const isManager = user?.role === "DepartmentHead(manager)";
  const isStaff = user?.role === "Staff";

  if (isManager) {
    title = "Manager Dashboard";
    subtitle = "Department overview and personal training management";
  }

  if (isStaff) {
    title = "Staff Dashboard";
    subtitle = "Your training progress and assigned tasks";
  }

  const managerStats = isManager ? await getManagerDashboardStats() : null;
  const staffStats = isStaff ? await getStaffDashboardStats() : null;

  return (
    <div className="min-h-screen space-y-4 m-2">
      <PageHeader title={title} subtitle={subtitle} />
      <div className="grid grid-cols-2 lg:grid-cols-3 gap-4">
        {isManager ? (
          <>
            <DashboardCard
              icon={<UsersIcon className="w-6 h-6" />}
              count={managerStats?.departmentStaff ?? 0}
              description="Department Staff"
            />
            <DashboardCard
              icon={<BookOpenIcon className="w-6 h-6" />}
              count={managerStats?.activeTrainings ?? 0}
              description="Active Trainings"
            />
            <DashboardCard
              icon={<AwardIcon className="w-6 h-6" />}
              count={managerStats?.pendingCertificates ?? 0}
              description="My Pending Certificates"
            />
          </>
        ) : (
          <>
            <DashboardCard
              icon={<AwardIcon className="w-6 h-6" />}
              count={staffStats?.approvedCertificates ?? 0}
              description="Certificates"
            />
            <DashboardCard
              icon={<CheckCircleIcon className="w-6 h-6" />}
              count={staffStats?.attendedTrainings ?? 0}
              description="Finished Trainings"
            />
            <DashboardCard
              icon={<GraduationCapIcon className="w-6 h-6" />}
              count={staffStats?.registeredTrainings ?? 0}
              description="Left Trainings"
            />
          </>
        )}
      </div>

      {/* Staff training performance chart */}
      {isStaff && (
        <TrainingStatusChart
          registered={staffStats?.registeredTrainings ?? 0}
          attended={staffStats?.attendedTrainings ?? 0}
          absent={staffStats?.absentTrainings ?? 0}
        />
      )}

      {/* Training Calendar */}
      <div className="space-y-2">
        <h2 className="text-2xl font-bold tracking-tight">Training Calendar</h2>
        <div className="w-full rounded-2xl shadow-md border overflow-hidden">
          <iframe
            src="https://calendar.google.com/calendar/embed?src=145e678c35bb5fb0c74ca3359c85f0ab6e565f23297871e901550d2297b1f3cc%40group.calendar.google.com&ctz=Asia%2FBangkok"
            className="w-full h-[500px] md:h-[600px]"
            frameBorder="0"
            scrolling="no"
          />
        </div>
      </div>
    </div>
  );
}
