"use server";
import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { z } from "zod";

export type State = {
  errors?: {
    name?: string[];
    empId?: string[];
    email?: string[];
    phone?: string[];
    position?: string[];
    password?: string[];
    status?: string[];
  };
  message?: string | null;
};

const formSchema = z.object({
  name: z.string().min(2).max(50),
  empId: z.string().min(2).max(50),
  email: z.string().email(),
  phone: z.string().min(10).max(15),
  position: z.string().min(2).max(50),
  password: z.string().min(8).max(100),
  status: z.enum(["active", "inactive", "suspended"]),
});

async function CreateStaffAction(
  prevState: State | void,
  formData: FormData,
): Promise<State | void> {
  const validatedFields = formSchema.safeParse({
    name: formData.get("name"),
    empId: formData.get("empId"),
    email: formData.get("email"),
    phone: formData.get("phone"),
    position: formData.get("position"),
    password: formData.get("password"),
    status: formData.get("status"),
  });

  const cookieStore = await cookies(); //  await

  // build cookie string manually
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  console.log("Received form data:", {
    validatedFields,
  });

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);

    return {
      errors: flattened.fieldErrors,
      message: "Failed to create department. Please check your input.",
    };
  }

  const { name, empId, email, phone, position, password, status } =
    validatedFields.data;
  let isCreated = false;

  try {
    const response = await fetch(`${API_BASE_URL}/admin/departments`, {
      method: "POST",
      credentials: "include",
      headers: { "Content-Type": "application/json", Cookie: cookieHeader },
      body: JSON.stringify({
        name,
        empId,
        email,
        phone,
        position,
        password,
        status,
      }),
      cache: "no-store",
    });

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
    redirect("/departments");
  }
}

export default CreateStaffAction;
