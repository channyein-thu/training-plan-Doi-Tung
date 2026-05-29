import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";

type PlansListResponse = {
  data?: { items?: unknown[]; meta?: unknown };
};

type PlanDetailResponse = {
  data?: Record<string, unknown>;
} | null;

// Safe paginated shape for the list view. Never throws.
export async function getPlans(): Promise<PlansListResponse> {
  const empty: PlansListResponse = { data: { items: [], meta: {} } };

  try {
    const { response } = await authFetch(
      `${API_BASE_URL}/manager/training-plans`,
      { method: "GET", cache: "no-store" },
    );

    if (!response.ok) {
      console.warn(`getPlans: backend ${response.status}`);
      return empty;
    }

    const payload = (await response.json().catch(() => null)) as
      | PlansListResponse
      | null;
    if (!payload) return empty;
    return {
      data: {
        items: payload.data?.items ?? [],
        meta: payload.data?.meta ?? {},
      },
    };
  } catch (err) {
    console.error("getPlans: unexpected error", err);
    return empty;
  }
}

// Returns null on any error (including 404). Page should check and render "not found".
export async function getPlanById(id: number): Promise<PlanDetailResponse> {
  try {
    const { response } = await authFetch(
      `${API_BASE_URL}/manager/training-plans/${id}`,
      {
        method: "GET",
        cache: "no-store",
        next: { tags: ["training-plans"] },
      },
    );

    if (!response.ok) {
      console.warn(`getPlanById(${id}): backend ${response.status}`);
      return null;
    }

    const payload = await response.json().catch(() => null);
    return payload as PlanDetailResponse;
  } catch (err) {
    console.error("getPlanById: unexpected error", err);
    return null;
  }
}
