import { Button } from "@/components/ui/button";
import { getStaff } from "@/lib/actions/department-staff/getStaff";
import { TrainingPlanStaff } from "@/types/staff";
import { ArrowLeftIcon } from "lucide-react";
import Link from "next/link";

type StaffPageProps = {
  params: {
    staffId: string;
  };
};
async function EditStaffPage({ params }: StaffPageProps) {
  const staff = await getStaff(); // Fetch OJT records on page load
  const { staffId } = await params;
  const id = parseInt(staffId);
  const staffMember = staff.find((p: TrainingPlanStaff) => p.id === id);
  if (!staffMember) {
    return <div>Staff member not found</div>;
  }
  return (
    <div className="min-h-screen space-y-4 m-2">
      {/* Back Button */}
      <Button
        asChild
        variant="ghost"
        className="text-black hover:bg-transparent mr-4 border rounded-md px-2 py-1 border-[#006022]"
      >
        <Link href="/department-staff">
          <ArrowLeftIcon className="mr-2 h-8 w-8" />
          Back to Department Staff
        </Link>
      </Button>

      <h1 className="text-md font-semibold">Edit Staff Details</h1>
      <p className="text-gray-600">Update Training Record Status</p>
    </div>
  );
}

export default EditStaffPage;
