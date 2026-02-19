"use client";

import PageHeader from "@/components/dashboard/pageHeader";
import AddStaffForm from "@/components/department-staff/addStaffForm";
import ButtonDialog from "@/components/department-staff/buttonDialog";
import DepartmentCard from "@/components/department-staff/departmentCard";
import { trainingPlanStaff } from "@/data/data";
import { TrainingPlanStaff } from "@/types/staff";
import { Search } from "lucide-react";
import React, { useState } from "react";
import EditStaffDialog from "@/components/department-staff/editStaffDialog";
import DeleteStaffDialog from "@/components/department-staff/deleteStaffDialog";

function DepartmentStaffPage() {
  const [selectedStaff, setSelectedStaff] = useState<TrainingPlanStaff | null>(
    null,
  );

  const [activeDialog, setActiveDialog] = useState<
    "staff" | "edit" | "delete" | null
  >(null);

  const [searchTerm, setSearchTerm] = useState("");

  // ✅ Filter staff by name, employeeId, or department
  const filteredStaff = trainingPlanStaff.filter((staff) =>
    `${staff.fullName} ${staff.employeeId} ${staff.departmentName}`
      .toLowerCase()
      .includes(searchTerm.toLowerCase()),
  );

  return (
    <div className="h-screen overflow-y-auto p-4 m-2 space-y-4">
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
          placeholder="Search staff..."
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          className="w-full outline-none border-none focus:ring-0"
        />
      </div>

      {/* Staff Cards Grid */}
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6 mt-4">
        {filteredStaff.map((staff) => (
          <DepartmentCard
            key={staff.id}
            staff={staff}
            onEdit={(staff) => {
              setSelectedStaff(staff);
              setActiveDialog("edit");
            }}
            onDelete={(staff) => {
              setSelectedStaff(staff);
              setActiveDialog("delete");
            }}
          />
        ))}
      </div>

      {/* EDIT */}
      {activeDialog === "edit" && selectedStaff && (
        <EditStaffDialog
          staff={selectedStaff}
          onClose={() => {
            setActiveDialog(null);
            setSelectedStaff(null);
          }}
        />
      )}

      {/* DELETE */}
      {activeDialog === "delete" && selectedStaff && (
        <DeleteStaffDialog
          staff={selectedStaff}
          onCancel={() => {
            setActiveDialog(null);
            setSelectedStaff(null);
          }}
          onConfirm={() => {
            // deleteDepartment(selectedDepartment.id);
            setActiveDialog(null);
            setSelectedStaff(null);
          }}
        />
      )}
    </div>
  );
}

export default DepartmentStaffPage;
