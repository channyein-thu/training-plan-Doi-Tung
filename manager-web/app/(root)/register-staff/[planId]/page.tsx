import PlanDetails from "@/components/register-staff/planDetails";
import StaffList from "@/components/register-staff/staffList";
import { Button } from "@/components/ui/button";
import { courses, trainingPlanStaff } from "@/data/data";
import { ArrowLeftIcon, UserPlus } from "lucide-react";
import Link from "next/link";

type PlanDetailsProps = {
  params: {
    planId: string;
  };
};
async function PlanDetailsRegister({ params }: PlanDetailsProps) {
  const { planId } = await params;
  const id = parseInt(planId);
  const plan = courses.find((p) => p.id === id);

  if (!plan) {
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
          <ArrowLeftIcon className="mr-2 h-8 w-8" />
          Back to Register Staff
        </Link>
      </Button>

      <h1 className="text-md font-semibold">Register Staff to Training Plan</h1>
      <p className="text-gray-600">
        Select department staff to register to this training
      </p>

      {/* Training Details */}
      <div className="border rounded-md m-2 bg-[#E8F7EC] p-4 space-y-4">
        <p className="font-medium mb-2">Training Details</p>
        <h1>{plan.name}</h1>
        <div className="flex  p-2 justify-between">
          <PlanDetails title="Date" subtitle={plan.date} />
          <PlanDetails title="Type" subtitle={plan.type} />
          <PlanDetails title="Category" subtitle={plan.category} />
          <PlanDetails
            title="Speaker/Institute"
            subtitle={plan.speakerInstitute}
          />
        </div>
      </div>

      {/* Staff List */}
      <div className="border rounded-md m-2 p-4">
        <p className="font-medium mb-2">Staff List</p>
        <StaffList staff={trainingPlanStaff} />
      </div>

      {/* Action Buttons */}
      <div className="px-6 py-4  border-gray-200 flex gap-4">
        <Button variant="outline" className="flex-1">
          Cancel
        </Button>
        <Button className="flex-1 bg-[#006022] hover:bg-[#004d1b] text-white">
          <UserPlus className="w-4 h-4 mr-2" />
          Register Staff
        </Button>
      </div>
    </div>
  );
}

export default PlanDetailsRegister;
