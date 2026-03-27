"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { TrainingPlanStaff } from "@/types/staff";
import { cookies } from "next/headers";

export async function getStaff(): Promise<TrainingPlanStaff[]> {
  const cookieStore = await cookies();
  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const response = await fetch(`${API_BASE_URL}/manager/users`, {
    method: "GET",
    credentials: "include",
    headers: { "Content-Type": "application/json", Cookie: cookieHeader },
    next: { tags: ["users"] },
  });

  if (!response.ok) {
    throw new Error("Failed to fetch staff members");
  }

  const payload = await response.json();
  return (payload.data.items ?? []) as TrainingPlanStaff[];
}
