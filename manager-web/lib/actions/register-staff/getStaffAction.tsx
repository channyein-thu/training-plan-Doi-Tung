import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";

export async function getStaff() {
  const cookieStore = await cookies(); //  await

  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const res = await fetch(`${API_BASE_URL}/manager/users`, {
    method: "GET",
    credentials: "include", //include credentials
    headers: {
      Cookie: cookieHeader, // correct now
    },
  });
  console.log("Response:", res.status); //log status for debugging

  if (!res.ok) {
    throw new Error("Failed to fetch staff members");
  }

  return res.json();
}

// export async function getPlanById(id: number) {
//   const cookieStore = await cookies();

//   const cookieHeader = cookieStore
//     .getAll()
//     .map((c) => `${c.name}=${c.value}`)
//     .join("; ");

//   const res = await fetch(`${API_BASE_URL}/manager/training-plans/${id}`, {
//     method: "GET",
//     credentials: "include", // ✅ send cookies
//     headers: {
//       Cookie: cookieHeader,
//     },
//     next: { tags: ["training-plans"] },
//   });

//   if (!res.ok) {
//     throw new Error("Failed to fetch training plan details");
//   }

//   return res.json();
// }
