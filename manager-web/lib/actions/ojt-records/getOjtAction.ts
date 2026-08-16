import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";

type OjtRecordsResponse = {
  status?: string;
  data?: {
    items?: unknown[];
    meta?: unknown;
  };
};

// Returns a SAFE default when anything goes wrong (rate limit, auth issue,
// network blip, unexpected response shape). The page that consumes this never
// has to worry about crashing the render — it always gets a structured object.
export async function getOjtRecords(): Promise<OjtRecordsResponse> {
  const empty: OjtRecordsResponse = { data: { items: [], meta: {} } };

  try {
    const { response } = await authFetch(`${API_BASE_URL}/manager/records`, {
      method: "GET",
      cache: "no-store",
    });

    if (!response.ok) {
      console.warn(
        `getOjtRecords: backend returned ${response.status}`,
      );
      return empty;
    }

    const payload = (await response.json().catch(() => null)) as
      | OjtRecordsResponse
      | null;

    if (!payload) return empty;
    return {
      ...payload,
      data: {
        items: payload.data?.items ?? [],
        meta: payload.data?.meta ?? {},
      },
    };
  } catch (err) {
    console.error("getOjtRecords: unexpected error", err);
    return empty;
  }
}
