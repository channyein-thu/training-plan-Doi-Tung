import { Course } from "./course";
import { TrainingPlanStaff } from "./staff";

export type OjtRecord = {
  id: number;
  staff: TrainingPlanStaff;
  course: Course;
  status: "COMPLETED" | "IN_PROGRESS" | "NOT_STARTED";
};
