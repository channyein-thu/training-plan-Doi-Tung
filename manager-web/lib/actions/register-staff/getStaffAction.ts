import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";

type StaffListResponse = {
  data?: {
    items?: unknown[];
    meta?: unknown;
  };
};

// Safe paginated shape. Never throws — returns { data: { items: [], meta: {} } } on any error.
export async function getStaff(): Promise<StaffListResponse> {
  const empty: StaffListResponse = { data: { items: [], meta: {} } };

  try {
    const { response } = await authFetch(`${API_BASE_URL}/manager/users`, {
      method: "GET",
      cache: "no-store",
    });

    if (!response.ok) {
      console.warn(`getStaff (register-staff): backend ${response.status}`);
      return empty;
    }

    const payload = (await response.json().catch(() => null)) as
      | StaffListResponse
      | null;
    if (!payload) return empty;
    return {
      data: {
        items: payload.data?.items ?? [],
        meta: payload.data?.meta ?? {},
      },
    };
  } catch (err) {
    console.error("getStaff: unexpected error", err);
    return empty;
  }
}
