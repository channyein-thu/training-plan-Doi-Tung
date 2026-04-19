import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { cookies } from "next/headers";

export async function getOjtRecords() {
  const cookieStore = await cookies(); //  await

  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const { response } = await authFetch(`${API_BASE_URL}/manager/records`, {
    method: "GET",
    credentials: "include", //include credentials
    headers: {
      Cookie: cookieHeader, // correct now
    },
  });
  console.log("Response:", response.status); //log status for debugging

  if (!response.ok) {
    throw new Error("Failed to fetch OJT records");
  }

  return response.json();
}
