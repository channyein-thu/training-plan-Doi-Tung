import { ArrowLeftIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import Link from "next/link";

import { notFound } from "next/navigation";
import { getRecordById } from "@/lib/actions/training-records/getRecordById";
import UploadCertificateForm from "@/components/my-training-records/uploadCertificateForm";
type RecordDetailProps = {
  params: {
    recordId: string;
  };
};
async function UploadCertificate({ params }: RecordDetailProps) {
  const { recordId } = await params;
  const response = await getRecordById(recordId);
  const record = response?.data ?? response;

  if (!record) {
    notFound();
  }
  return (
    <div className="min-h-screen space-y-4 m-2">
      {/* Back Button */}
      <Button
        asChild
        variant="ghost"
        className="text-black hover:bg-transparent mr-4 border rounded-md px-2 py-1 border-[#006022]"
      >
        <Link href="/my-training-records">
          <ArrowLeftIcon className="mr-2 h-8 w-8" />
          Back to My Training Records
        </Link>
      </Button>

      <h1 className="text-md font-semibold">Upload Certificate</h1>
      <p className="text-gray-600">Upload Certificate for Training Record</p>

      {/* Training Details */}
      <div className="border rounded-md m-2  p-4 space-y-4">
        <p className="font-medium mb-2">Certificate Information</p>
        <div className="grid grid-cols-2 p-2 justify-between gap-4">
          <PlanDetails title="Employee" subtitle={record.staffName} />
          <PlanDetails title="Employee ID" subtitle={record.employeeID} />
          <PlanDetails title="Department" subtitle={record.departmentName} />
          <PlanDetails title="Division" subtitle={record.division} />
          <PlanDetails title="Position" subtitle={record.position} />

          <PlanDetails
            title="Training Name"
            subtitle={record.trainingPlanName}
          />
          <PlanDetails title="Category" subtitle={record.category} />
          <PlanDetails title="Type" subtitle={record.type} />
          <PlanDetails
            title="Speaker Institute"
            subtitle={record.speakerInstitute}
          />
          <PlanDetails title="Date Attended" subtitle={record.date} />
          <PlanDetails
            title="Duration"
            subtitle={`${record.numberOfDays} day(s)`}
          />
          <PlanDetails
            title="Number Of Hours"
            subtitle={String(record.numberOfHours)}
          />
          <PlanDetails title="Location" subtitle={record.location} />
        </div>
      </div>

      <UploadCertificateForm
        trainingId={record.trainingPlanId ?? record.trainingId ?? record.id}
      />
    </div>
  );
}

export default UploadCertificate;

function PlanDetails({ title, subtitle }: { title: string; subtitle: string }) {
  return (
    <div>
      {" "}
      <div>
        <p className="text-sm text-gray-600">{title}</p>
        <p className="text-md text-black">{subtitle}</p>
      </div>
    </div>
  );
}
