"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export type State = {
  errors?: {
    name?: string[];
  };
  message?: string | null;
};

export async function DeleteOjtAction(
  id: number,
  prevState: State | void,
  formData: FormData,
): Promise<State | void> {
  const cookieStore = await cookies(); //  await

  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  let isDeleted = false;

  try {
    const { response } = await authFetch(
      `${API_BASE_URL}/manager/records/${id}`,
      {
        method: "DELETE",
        credentials: "include",
        headers: { "Content-Type": "application/json", Cookie: cookieHeader },
        cache: "no-store",
      },
    );

    if (!response.ok) {
      const errorData = await response.json();

      return {
        message:
          errorData?.message ||
          "Failed to delete OJT record. Please try again.",
      };
    }

    isDeleted = true;
  } catch {
    return {
      message: "Server error. Please try again later.",
    };
  }

  if (isDeleted) {
    redirect("/ojt-records");
  }
}
