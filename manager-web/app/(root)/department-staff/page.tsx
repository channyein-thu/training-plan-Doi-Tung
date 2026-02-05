import PageHeader from "@/components/dashboard/pageHeader";
import AddStaffForm from "@/components/department-staff/addStaffForm";
import ButtonDialog from "@/components/department-staff/buttonDialog";
import { Button } from "@/components/ui/button";
import { Search } from "lucide-react";
import React from "react";

function DepartmentStaffPage() {
  return (
    <>
      <div className="m-6 space-y-4">
        <PageHeader
          title="Department Staff"
          subtitle="View your department staff members"
          action={
            <ButtonDialog name="Add Staff Member">
              <AddStaffForm />
            </ButtonDialog>
          }
        />

        {/* Search bar */}
        <div className="border border-[#006022] rounded-lg p-2 flex items-center gap-2 mt-6">
          <Search className="text-gray-400" />
          <input
            type="text"
            placeholder="Search departments..."
            className="w-full outline-none border-none focus:ring-0"
            // value={searchTerm} // ✅ controlled value
            // onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
      </div>
    </>
  );
}

export default DepartmentStaffPage;
