"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { TrainingPlanStaff } from "@/types/staff";

// Always returns an array. Never throws — on any error returns [].
export async function getStaff(): Promise<TrainingPlanStaff[]> {
  try {
    const { response } = await authFetch(`${API_BASE_URL}/manager/users`, {
      method: "GET",
      next: { tags: ["users"] },
    });

    if (!response.ok) {
      console.warn(`getStaff (department-staff): backend ${response.status}`);
      return [];
    }

    const payload = await response.json().catch(() => null);
    return ((payload?.data?.items ?? []) as TrainingPlanStaff[]) || [];
  } catch (err) {
    console.error("getStaff: unexpected error", err);
    return [];
  }
}
