export type StaffStatus = "ACTIVE" | "INACTIVE";

export type Staff = {
  id: number;
  employeeId: string;
  fullName: string;
  email: string;
  phone: string;
  departmentId: number;
  departmentName: string;
  jobRole: string;
  status: StaffStatus;
  isManager: boolean;
};
