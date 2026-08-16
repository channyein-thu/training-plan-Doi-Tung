// Mirrors the backend training-plan response (data/response/training_plan_response.go).
// Used for the register-staff plan list and cards.
export type Course = {
  id: number;
  name: string;
  speakerInstitute?: string;
  type: string;
  category: string;
  date: string;
  content: string;
  numberOfDays: number;
  numberOfHours?: number;
  location?: string;
  totalCost?: number;
  budgetCode?: string;
  numberOfPerson: number;
  costPerPerson?: number;
  createdAt: string;
  updatedAt: string;
};
