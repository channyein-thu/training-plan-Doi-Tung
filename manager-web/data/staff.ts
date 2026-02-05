import { Staff } from "@/types/data";

export const staffItems: Staff[] = [
  {
    id: 1,
    employeeId: "ADMIN001",
    fullName: "System Admin",
    email: "admin@company.com",
    phone: "0999999999",
    departmentId: 1,
    departmentName: "IT",
    jobRole: "HR Administrator",
    status: "ACTIVE",
    isManager: false,
  },
  {
    id: 2,
    employeeId: "EMP00123",
    fullName: "Aung Min Oo",
    email: "aung.min.oo@company.com",
    phone: "0912345678",
    departmentId: 1,
    departmentName: "IT",
    jobRole: "Software Engineer",
    status: "ACTIVE",
    isManager: false,
  },
];
