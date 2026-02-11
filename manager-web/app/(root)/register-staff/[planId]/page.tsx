import PlanDetails from "@/components/register-staff/planDetails";
import { Button } from "@/components/ui/button";
import { courses } from "@/data/data";
import { ArrowLeftIcon } from "lucide-react";
import Link from "next/link";
import React from "react";

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
    <div className="m-6 space-y-4">
      {/* Back Button */}
      <Button
        asChild
        variant="ghost"
        className="text-black hover:bg-transparent mr-4"
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
    </div>
  );
}

export default PlanDetailsRegister;
