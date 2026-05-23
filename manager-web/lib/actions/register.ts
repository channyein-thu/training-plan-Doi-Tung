"use server";

import { z } from "zod";
import { API_BASE_URL } from "../api/api";

export type State = {
  errors?: {
    name?: string[];
    employeeID?: string[];
    email?: string[];
    phone?: string[];
    departmentId?: string[];
    position?: string[];
    confirmPassword?: string[];
    password?: string[];
  };
  message?: string | null;
};

const FormSchema = z
  .object({
    name: z.string().trim().min(2, "Name must be at least 2 characters!"),
    employeeID: z
      .string()
      .trim()
      .min(2, "Employee ID must be at least 2 characters!"),
    email: z.string().trim().email("Invalid email address!"),
    phone: z
      .string()
      .trim()
      .min(10, "Phone number must be at least 10 characters!"),
    departmentId: z.coerce.number().int().positive("Department is required!"),
    position: z
      .string()
      .trim()
      .min(2, "Position must be at least 2 characters!"),
    password: z
      .string()
      .trim()
      .min(6, "Password must be at least 6 characters!"),
    confirmPassword: z.string().trim().min(6, "Confirm your password!"),
  })
  .refine((values) => values.password === values.confirmPassword, {
    message: "Passwords do not match.",
    path: ["confirmPassword"],
  });

export async function RegisterAction(prevState: State, formData: FormData) {
  const validatedFields = FormSchema.safeParse({
    name: formData.get("name"),
    employeeID: formData.get("employeeID"),
    email: formData.get("email"),
    phone: formData.get("phone"),
    departmentId: formData.get("departmentId"),
    position: formData.get("position"),
    password: formData.get("password"),
    confirmPassword: formData.get("confirmPassword"),
  });

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);
    return {
      errors: flattened.fieldErrors,
      message: "Failed to register. Please check your input.",
    };
  }

  const requestBody = validatedFields.data;

  try {
    const response = await fetch(`${API_BASE_URL}/auth/staff/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(requestBody),
      cache: "no-store",
    });
    console.log("Register payload:", requestBody);

    return {
      message: "Registration successful!",
    };
  } catch (error) {
    return {
      message: "Failed to register. Please try again later.",
    };
  }
}

// revalidatePath - revalidating specific pages or layout

// revalidateTag - revalidating data in server action and route handler
// updateTag - revalidating data in server action only

// cacheTag
