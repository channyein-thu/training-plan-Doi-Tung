import { Card, CardContent } from "@/components/ui/card";
import { ArrowLeftIcon, Users } from "lucide-react";

import { Button } from "@/components/ui/button";
import Link from "next/link";
import UploadCertificateForm from "@/components/my-training-records/uploadCertificateForm";
import { ojtRecords } from "@/data/data";
type RecordDetailProps = {
  params: {
    recordId: string;
  };
};
async function UploadCertificate({ params }: RecordDetailProps) {
  const { recordId } = await params;
  const id = parseInt(recordId);
  console.log("recordId", recordId);
  console.log("id", id);
  const record = ojtRecords.find((p) => p.id === id);

  if (!record) {
    return <div>Record not found</div>;
  }
  return (
    <div className="h-screen overflow-y-auto p-4 m-2 space-y-4">
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
          <PlanDetails title="Employee" subtitle={record.staff.fullName} />

          <PlanDetails title="Training Name" subtitle={record.course.name} />
          <PlanDetails title="Category" subtitle={record.course.category} />
          <PlanDetails title="Date Attended" subtitle={record.course.date} />
          <PlanDetails
            title="Duration"
            subtitle={record.course.numberOfDays.toString() + "day"}
          />

          <PlanDetails title="Date" subtitle={record.course.date} />
          <PlanDetails
            title="Number Of Hours"
            subtitle={record.course.numberOfHours.toString()}
          />
          <PlanDetails title="Location" subtitle={record.course.location} />
        </div>
      </div>

      <UploadCertificateForm />
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
