export type OjtRecord = {
  id: number;
  trainingPlanName: string;
  location: string;
  costPerPerson: number;
  budgetCode: string;
  employeeId: string;
  employeeName: string;
  position: string;
  department: string;
  division: string;
  status: "Register" | "Completed" | "Cancelled";
  createdAt: string;
  updatedAt: string;
};
