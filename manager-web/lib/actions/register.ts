"use server";

import { z } from "zod";

export type State = {
  errors?: {
    fullName?: string[];
    employeeId?: string[];
    email?: string[];
    phone?: string[];
    department?: string[];
    position?: string[];
    agency?: string[];
    cotton?: string[];
    line?: string[];
    password?: string[];
  };
  message?: string | null;
};

const FormSchema = z.object({
  fullName: z
    .string()
    .trim()
    .min(2, "Full name must be at least 2 characters!"),
  employeeId: z
    .string()
    .trim()
    .min(2, "Employee ID must be at least 2 characters!"),
  email: z.string().trim().email("Invalid email address!"),
  phone: z
    .string()
    .trim()
    .min(10, "Phone number must be at least 10 characters!"),
  department: z
    .string()
    .trim()
    .min(2, "Department must be at least 2 characters!"),
  position: z.string().trim().min(2, "Position must be at least 2 characters!"),
  agency: z.string().trim().min(2, "Agency must be at least 2 characters!"),
  cotton: z.string().trim().min(2, "Cotton must be at least 2 characters!"),
  line: z.string().trim().min(2, "Line must be at least 2 characters!"),
  password: z.string().trim().min(6, "Password must be at least 6 characters!"),
});

export async function RegisterAction(formData: FormData, prevState: State) {
  const validatedFields = FormSchema.safeParse({
    fullName: formData.get("fullName"),
    employeeId: formData.get("employeeId"),
    email: formData.get("email"),
    phone: formData.get("phone"),
    department: formData.get("department"),
    position: formData.get("position"),
    agency: formData.get("agency"),
    cotton: formData.get("cotton"),
    line: formData.get("line"),
    password: formData.get("password"),
  });

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);
    return {
      errors: flattened.fieldErrors,
      message: "Failed to register. Please check your input.",
    };
  }

  const { email, password } = validatedFields.data;

  try {
    console.log("Email:", email);
    console.log("Password:", password);

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
