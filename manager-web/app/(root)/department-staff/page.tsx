import PageHeader from "@/components/dashboard/pageHeader";
import AddStaffForm from "@/components/department-staff/addStaffForm";
import ButtonDialog from "@/components/department-staff/buttonDialog";
import { Button } from "@/components/ui/button";
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
      </div>
    </>
  );
}

export default DepartmentStaffPage;
