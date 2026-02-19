"use server";

import { z } from "zod";

export type State = {
  errors?: {
    email?: string[];
    password?: string[];
  };
  message?: string | null;
};

const FormSchema = z.object({
  email: z.string().trim().email("Invalid email address!"),
  password: z.string().trim().min(6, "Password must be at least 6 characters!"),
});

export async function LoginAction(formData: FormData, prevState: State) {
  const validatedFields = FormSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);
    return {
      errors: flattened.fieldErrors,
      message: "Failed to login. Please check your input.",
    };
  }

  const { email, password } = validatedFields.data;

  try {
    console.log("Email:", email);
    console.log("Password:", password);

    return {
      message: "Login successful! (This is a mock response.)",
    };
  } catch (error) {
    return {
      message: "Failed to login. Please try again later.",
    };
  }
}

// revalidatePath - revalidating specific pages or layout

// revalidateTag - revalidating data in server action and route handler
// updateTag - revalidating data in server action only

// cacheTag
