import EditOJTdetails from "@/components/ojt-records/editOjtForm";
import PlanDetails from "@/components/register-staff/planDetails";
import { Button } from "@/components/ui/button";
import { getOjtRecords } from "@/lib/actions/ojt-records/getOjtAction";
import { ArrowLeftIcon } from "lucide-react";
import Link from "next/link";

type PlanDetailsProps = {
  params: Promise<{ ojtId: string }>;
};

async function OjtDetails({ params }: PlanDetailsProps) {
  const ojtRecords = await getOjtRecords();
  // Loose typing here — the backend's record shape doesn't match the local
  // OjtRecord type, but the JSX already assumes the looser server shape.
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const items = (ojtRecords?.data?.items as any[]) ?? [];
  const { ojtId } = await params;
  const id = parseInt(ojtId);
  const record = items.find((p) => p.id === id);

  if (!record) {
    return (
      <div className="p-6 text-center text-gray-500">
        Training record not found.
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
          <PlanDetails
            title="Training Name"
            subtitle={record.trainingPlanName}
          />
          <PlanDetails title="Employee" subtitle={record.employeeName} />
          <PlanDetails title="Employee ID" subtitle={record.employeeId} />
          <PlanDetails title="Date" subtitle={record.createdAt} />
          <PlanDetails title="Location" subtitle={record.location} />
          <PlanDetails
            title="Cost Per Person"
            subtitle={record.costPerPerson?.toString() ?? ""}
          />
          <PlanDetails title="Budget Code" subtitle={record.budgetCode} />
        </div>
      </div>

      <EditOJTdetails record={record} />
    </div>
  );
}

export default OjtDetails;
