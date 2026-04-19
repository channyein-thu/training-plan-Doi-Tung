import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { cookies } from "next/headers";

export async function getPlans() {
  const cookieStore = await cookies(); //  await

  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const { response } = await authFetch(
    `${API_BASE_URL}/manager/training-plans`,
    {
      method: "GET",
      credentials: "include", //include credentials
      headers: {
        Cookie: cookieHeader, // correct now
      },
    },
  );
  console.log("Response:", response.status); //log status for debugging

  if (!response.ok) {
    throw new Error("Failed to fetch training plans");
  }

  return response.json();
}

export async function getPlanById(id: number) {
  const cookieStore = await cookies();

  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const { response } = await authFetch(
    `${API_BASE_URL}/manager/training-plans/${id}`,
    {
      method: "GET",
      credentials: "include", // ✅ send cookies
      headers: {
        Cookie: cookieHeader,
      },
      next: { tags: ["training-plans"] },
    },
  );

  if (!response.ok) {
    throw new Error("Failed to fetch training plan details");
  }

  return response.json();
}
