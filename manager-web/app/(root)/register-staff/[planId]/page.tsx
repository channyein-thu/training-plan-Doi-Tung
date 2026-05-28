import PlanDetails from "@/components/register-staff/planDetails";
import StaffList from "@/components/register-staff/staffList";
import { Button } from "@/components/ui/button";
import { getPlanById } from "@/lib/actions/register-staff/getPlanAction";
import { getStaff } from "@/lib/actions/register-staff/getStaffAction";
import { ArrowLeftIcon } from "lucide-react";
import Link from "next/link";

async function PlanDetailsRegister({
  params,
}: {
  params: Promise<{ planId: string }>;
}) {
  const resolvedParams = await params;
  const id = Number(resolvedParams.planId);

  const [planDetails, staff] = await Promise.all([getPlanById(id), getStaff()]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const plan = (planDetails?.data as any) ?? null;
  const staffItems = (staff?.data?.items as unknown[]) ?? [];

  if (!plan) {
    return (
      <div className="p-6 text-center text-gray-500">
        Training Plan not found or could not be loaded.
      </div>
    );
  }

  return (
    <div className="h-screen overflow-y-auto p-4 m-2 space-y-4">
      {/* Back Button */}
      <Button
        asChild
        variant="ghost"
        className="text-black hover:bg-transparent mr-4 border rounded-md px-2 py-1 border-[#006022]"
      >
        <Link href="/register-staff">
          <ArrowLeftIcon className="mr-2 h-4 w-4" />
          Back to Register Staff
        </Link>
      </Button>

      <h1 className="text-md font-semibold">Register Staff to Training Plan</h1>

      <div className="border rounded-md m-2 bg-[#E8F7EC] p-4 space-y-4">
        <p className="font-medium mb-2">Training Details</p>
        <h1>{plan.name}</h1>

        <div className="flex p-2 justify-between gap-3">
          <PlanDetails title="Date" subtitle={plan.date} />
          <PlanDetails title="Type" subtitle={plan.type} />
          <PlanDetails title="Category" subtitle={plan.category} />
          <PlanDetails title="Speaker" subtitle={plan.speakerInstitute} />
        </div>
      </div>
      {/* Staff List */}
      <div className="border rounded-md m-2 p-4 mt-4">
        <p className="font-medium mb-2">Staff List</p>
        {/* eslint-disable-next-line @typescript-eslint/no-explicit-any */}
        <StaffList staff={staffItems as any} planId={plan.id} />
      </div>
    </div>
  );
}

export default PlanDetailsRegister;
