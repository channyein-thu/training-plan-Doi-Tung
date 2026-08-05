"use server";
import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { z } from "zod";

export type State = {
  errors?: {
    name?: string[];
    employeeID?: string[];
    email?: string[];
    phone?: string[];
    position?: string[];
    password?: string[];
    status?: string[];
  };
  message?: string | null;
};

const formSchema = z.object({
  name: z.string().trim().min(2).max(52),
  employeeID: z.string().trim().min(1).max(52),
  email: z.string().trim().email().max(52),
  phone: z.string().trim().max(20),
  position: z.string().trim().min(1).max(100),
  password: z.string().min(6).max(100),
  status: z.enum(["Active", "Inactive", "Suspended"]),
});

async function CreateStaffAction(
  prevState: State | void,
  formData: FormData,
): Promise<State | void> {
  const validatedFields = formSchema.safeParse({
    name: formData.get("name"),
    employeeID: formData.get("employeeID"),
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

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);

    return {
      errors: flattened.fieldErrors,
      message: "Failed to create department. Please check your input.",
    };
  }

  const { name, employeeID, email, phone, position, password, status } =
    validatedFields.data;
  let isCreated = false;

  try {
    const { response } = await authFetch(`${API_BASE_URL}/manager/users`, {
      method: "POST",
      credentials: "include",
      headers: { "Content-Type": "application/json", Cookie: cookieHeader },
      body: JSON.stringify({
        name,
        employeeID,
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
    redirect("/department-staff");
  }
}

export default CreateStaffAction;
