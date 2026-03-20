import PlanDetails from "@/components/register-staff/planDetails";
import StaffList from "@/components/register-staff/staffList";
import { Button } from "@/components/ui/button";
import { courses, trainingPlanStaff } from "@/data/data";
import { getPlanById } from "@/lib/actions/register-staff/getPlanAction";
import { ArrowLeftIcon } from "lucide-react";
import Link from "next/link";

async function PlanDetailsRegister({
  params,
}: {
  params: Promise<{ planId: string }>;
}) {
  const resolvedParams = await params;
  const id = Number(resolvedParams.planId);

  console.log("Received ID:", id);
  const planDetails = await getPlanById(id);

  if (!planDetails) {
    return <div>Training Plan not found</div>;
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
        <h1>{planDetails.data.name}</h1>

        <div className="flex p-2 justify-between gap-3">
          <PlanDetails title="Date" subtitle={planDetails.data.date} />
          <PlanDetails title="Type" subtitle={planDetails.data.type} />
          <PlanDetails title="Category" subtitle={planDetails.data.category} />
          <PlanDetails
            title="Speaker"
            subtitle={planDetails.data.speakerInstitute}
          />
        </div>
      </div>
    </div>
  );
}

export default PlanDetailsRegister;
