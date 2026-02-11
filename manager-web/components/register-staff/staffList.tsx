import { trainingPlanStaff } from "@/data/data";
import { TrainingPlanStaff } from "@/types/staff";
import { Checkbox } from "../ui/checkbox";
type StaffSelectProps = {
  staff: TrainingPlanStaff[];
};

function StaffList({ staff }: StaffSelectProps) {
  return (
    <>
      <div className="grid grid-cols-2 gap-4">
        {staff.map((s) => (
          <div
            key={s.id}
            className="border rounded-md p-4 space-y-2 flex items-center gap-3"
          >
            <Checkbox />
            <div>
              <p className="font-medium">{s.fullName}</p>
              <p>{s.employeeId}</p>
              <p className="text-sm text-gray-600">{s.position}</p>
            </div>
          </div>
        ))}
      </div>
    </>
  );
}

export default StaffList;
