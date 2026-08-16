import { DepartmentList } from "@/types/department";
import { API_BASE_URL } from "./api";

export async function getDepartmentList(): Promise<DepartmentList[]> {
  const response = await fetch(`${API_BASE_URL}/departments-list`, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
    cache: "no-store",
  });
  if (!response.ok) return [];

  const payload = await response.json();
  return (payload.data ?? []) as DepartmentList[];
}
