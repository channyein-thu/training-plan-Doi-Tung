import { getStaff } from "@/lib/actions/department-staff/getStaff";
import DepartmentStaff from "./departmentStaff";

export default async function DepartmentStaffPage() {
  const trainingPlanStaff = await getStaff();

  return <DepartmentStaff initialStaff={trainingPlanStaff} />;
}
