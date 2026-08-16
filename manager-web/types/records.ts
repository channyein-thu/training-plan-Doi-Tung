export type OjtRecord = {
  id: number;
  // Present on manager/admin record responses (AdminRecordResponse); absent on
  // a staff member's own records, hence optional.
  employeeId?: string;
  employeeName?: string;
  division: string;
  department: string;
  budgetCode: string;
  costPerPerson: number;
  position: string;
  trainingPlanId: number;
  trainingPlanName: string;
  status: "Register" | "Attended" | "Absent";
  location: string;
  trainingDate: string;
  numberOfHours: number;
  speakerInstitute: string;
  preTestScore?: number;
  postTestScore?: number;
  evaluation?: string;
  trainingType:
    | "In-house"
    | "Public"
    | "OJT"
    | "Self-learning"
    | "Online/Virtual";
  createdAt: string;
  updatedAt: string;
};
