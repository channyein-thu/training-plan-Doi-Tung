export type TrainingPlanStaff = {
  id: number;
  employeeId: string;
  fullName: string;
  email: string;
  phone: string;
  departmentId: number;
  departmentName: string;
  position: string;
  status: "ACTIVE" | "INACTIVE";
};
