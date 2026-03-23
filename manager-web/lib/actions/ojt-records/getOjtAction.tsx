import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";

export async function getOjtRecords() {
  const cookieStore = await cookies(); //  await

  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const res = await fetch(`${API_BASE_URL}/manager/records`, {
    method: "GET",
    credentials: "include", //include credentials
    headers: {
      Cookie: cookieHeader, // correct now
    },
  });
  console.log("Response:", res.status); //log status for debugging

  if (!res.ok) {
    throw new Error("Failed to fetch OJT records");
  }

  return res.json();
}
