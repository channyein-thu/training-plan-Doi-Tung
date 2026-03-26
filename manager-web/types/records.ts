export type OjtRecord = {
  id: number;
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
  trainingType:
    | "In-house"
    | "Public"
    | "OJT"
    | "Self-learning"
    | "Online/Virtual";
  createdAt: string;
  updatedAt: string;
};
