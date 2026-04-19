"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";
import { z } from "zod";
import { redirect } from "next/navigation";
import { authFetch } from "@/lib/api/authFetch";

export type State = {
  errors?: {
    status?: string[];
    evaluation?: string[];
    preTestScore?: string[];
    postTestScore?: string[];
  };
  message: string;
};

const FormSchema = z.object({
  status: z.string().trim().min(1, "Status is required!"),
  evaluation: z
    .string()
    .trim()
    .min(6, "Evaluation must be at least 6 characters.")
    .optional()
    .or(z.literal("")),

  preTestScore: z
    .number()
    .min(0, "Pre-test score must be at least 0!")
    .max(100, "Pre-test score cannot exceed 100!")
    .optional(),

  postTestScore: z
    .number()
    .min(0, "Post-test score must be at least 0!")
    .max(100, "Post-test score cannot exceed 100!")
    .optional(),
});

// export async function LoginAction(prevState: any, formData: FormData) {
//   const validatedFields = FormSchema.safeParse({
//     email: formData.get("email"),
//     password: formData.get("password"),
//   });

//   if (!validatedFields.success) {
//     return {
//       errors: validatedFields.error.flatten().fieldErrors,
//       message: "Invalid input",
//     };
//   }

//   const { email, password } = validatedFields.data;

export async function UpdateOjtRecordAction(
  id: number,
  prevState: State | void,
  formData: FormData,
): Promise<State | void> {
  const preScore = formData.get("preTestScore");
  const postScore = formData.get("postTestScore");
  const validatedFields = FormSchema.safeParse({
    status: formData.get("status"),
    evaluation: formData.get("evaluation"),
    preTestScore: preScore ? Number(preScore) : undefined,
    postTestScore: postScore ? Number(postScore) : undefined,
  });

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);

    return {
      errors: flattened.fieldErrors,
      message: "Failed to update OJT record. Please check your input.",
    };
  }

  const { status, evaluation, preTestScore, postTestScore } =
    validatedFields.data;
  const cookieStore = await cookies();

  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  let isCreated = false;

  try {
    const { response } = await authFetch(
      `${API_BASE_URL}/manager/records/${id}`,
      {
        method: "PUT",
        credentials: "include",
        headers: { "Content-Type": "application/json", Cookie: cookieHeader },
        body: JSON.stringify({
          status,
          evaluation,
          preTestScore,
          postTestScore,
        }),
        cache: "no-store",
      },
    );

    if (!response.ok) {
      const errorData = await response.json();

      return {
        message:
          errorData?.message ||
          "Failed to create department. Please try again.",
      };
    }

    isCreated = true;
  } catch {
    return {
      message: "Server error. Please try again later.",
    };
  }

  if (isCreated) {
    redirect("/ojt-records");
  }
}
