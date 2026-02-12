import EditOJTdetails from "@/components/ojt-records/editOjtForm";
import TestFormExample from "@/components/ojt-records/pre/postProgress";
import PlanDetails from "@/components/register-staff/planDetails";
import { Button } from "@/components/ui/button";

import { courses, ojtRecords, trainingPlanStaff } from "@/data/data";
import { ArrowLeftIcon, UserPlus } from "lucide-react";
import Link from "next/link";

type PlanDetailsProps = {
  params: {
    ojtId: string;
  };
};
async function OjtDetails({ params }: PlanDetailsProps) {
  const { ojtId } = await params;
  const id = parseInt(ojtId);
  const record = ojtRecords.find((p) => p.id === id);

  if (!record) {
    return <div>Training Plan not found</div>;
  }
  return (
    <div className="m-6 space-y-4">
      {/* Back Button */}
      <Button
        asChild
        variant="ghost"
        className="text-black hover:bg-transparent mr-4 border rounded-md px-2 py-1 border-[#006022]"
      >
        <Link href="/ojt-records">
          <ArrowLeftIcon className="mr-2 h-8 w-8" />
          Back to OJT Records
        </Link>
      </Button>

      <h1 className="text-md font-semibold">Edit OJT Record</h1>
      <p className="text-gray-600">Update Training Record Status</p>

      {/* Training Details */}
      <div className="border rounded-md m-2  p-4 space-y-4">
        <p className="font-medium mb-2">Basic Information</p>
        <div className="grid grid-cols-2 p-2 justify-between gap-4">
          <PlanDetails title="Training Name" subtitle={record.course.name} />
          <PlanDetails title="Category" subtitle={record.course.category} />
          <PlanDetails title="Employee" subtitle={record.staff.fullName} />
          <PlanDetails
            title="Employee ID"
            subtitle={record.staff.id.toString()}
          />
          <PlanDetails title="Date" subtitle={record.course.date} />
          <PlanDetails
            title="Number Of Hours"
            subtitle={record.course.numberOfHours.toString()}
          />
          <PlanDetails title="Location" subtitle={record.course.location} />
          <PlanDetails
            title="Cost Per Person"
            subtitle={record.course.costPerPerson.toString()}
          />
          <PlanDetails
            title="Budget Code"
            subtitle={record.course.budgetCode}
          />
        </div>
      </div>

      <EditOJTdetails />
    </div>
  );
}

export default OjtDetails;
