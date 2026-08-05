"use server";

import { authFetch } from "@/lib/api/authFetch";

export type ManagerDashboardStats = {
  departmentStaff: number;
  activeTrainings: number;
  pendingCertificates: number;
};

export type StaffDashboardStats = {
  myTrainings: number;
  pendingCertificates: number;
  approvedCertificates: number;
};

// Manager dashboard aggregate counts. Returns null on auth failure/error so the
// page can render a safe zero-state.
export async function getManagerDashboardStats(): Promise<ManagerDashboardStats | null> {
  const { response, unauthorized } = await authFetch(
    "/manager/dashboard-stats",
    { method: "GET", cache: "no-store" },
  );

  if (unauthorized || !response.ok) {
    return null;
  }

  const json = await response.json().catch(() => null);
  return (json?.data as ManagerDashboardStats) ?? null;
}

// Staff dashboard aggregate counts (the manager app is also used by staff).
export async function getStaffDashboardStats(): Promise<StaffDashboardStats | null> {
  const { response, unauthorized } = await authFetch("/staff/dashboard-stats", {
    method: "GET",
    cache: "no-store",
  });

  if (unauthorized || !response.ok) {
    return null;
  }

  const json = await response.json().catch(() => null);
  return (json?.data as StaffDashboardStats) ?? null;
}
