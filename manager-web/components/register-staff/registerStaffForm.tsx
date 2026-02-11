import React from "react";
import { DialogDescription, DialogHeader, DialogTitle } from "../ui/dialog";
import { Course } from "@/types/course";
import PlanDetails from "./planDetails";
import { trainingPlanStaff } from "@/data/data";
import { Checkbox } from "../ui/checkbox";

function RegisterStaffForm({ plan }: { plan: Course }) {
  return (
    <div>
      <DialogHeader>
        <DialogTitle>Register Staff</DialogTitle>
        <DialogDescription>{plan.name}</DialogDescription>
      </DialogHeader>

      {/* Training Details */}
      <div className="border rounded-md m-2 bg-[#E8F7EC] p-2">
        <p className="font-medium mb-2">Training Details</p>
        <div className="grid grid-cols-2 gap-4 p-2">
          <PlanDetails title="Date" subtitle={plan.date} />
          <PlanDetails title="Type" subtitle={plan.type} />
          <PlanDetails title="Category" subtitle={plan.category} />
          <PlanDetails
            title="Speaker/Institute"
            subtitle={plan.speakerInstitute}
          />
        </div>
      </div>

      {/* Register Staff Form */}
      <div>
        {trainingPlanStaff.map((staff) => {
          return (
            <div
              key={staff.id}
              className="border-b last:border-0 p-2 flex items-center gap-3"
            >
              <Checkbox />
              <div>
                <p className="font-medium">{staff.fullName}</p>
                <p className="text-sm text-gray-600">{staff.position}</p>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}

export default RegisterStaffForm;
