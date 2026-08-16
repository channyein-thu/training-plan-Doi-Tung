"use server";
import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { revalidatePath } from "next/cache";
import { redirect } from "next/navigation";
import { z } from "zod";

export type State = {
  errors?: {
    name?: string[];
    employeeID?: string[];
    email?: string[];
    phone?: string[];
    position?: string[];
    status?: string[];
  };
  message?: string | null;
};

const formSchema = z.object({
  name: z.string().trim().min(2).max(52),
  employeeID: z.string().trim().min(1).max(52),
  email: z.string().trim().email().max(52),
  phone: z.string().trim().min(1, "Phone is required").max(20),
  position: z.string().trim().min(1).max(100),
  status: z.enum(["Active", "Inactive", "Suspended"]),
});

// Bound with the staff id via EditStaffAction.bind(null, id) in the dialog.
async function EditStaffAction(
  staffId: number,
  prevState: State | void,
  formData: FormData,
): Promise<State | void> {
  const validatedFields = formSchema.safeParse({
    name: formData.get("name"),
    employeeID: formData.get("employeeID"),
    email: formData.get("email"),
    phone: formData.get("phone"),
    position: formData.get("position"),
    status: formData.get("status"),
  });

  if (!validatedFields.success) {
    const flattened = z.flattenError(validatedFields.error);
    return {
      errors: flattened.fieldErrors,
      message: "Failed to update staff. Please check your input.",
    };
  }

  let isUpdated = false;

  try {
    const { response } = await authFetch(
      `${API_BASE_URL}/manager/users/${staffId}`,
      {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(validatedFields.data),
        cache: "no-store",
      },
    );

    if (!response.ok) {
      const errorData = await response.json().catch(() => null);
      return {
        message:
          errorData?.message || "Failed to update staff. Please try again.",
      };
    }

    isUpdated = true;
  } catch {
    return { message: "Server error. Please try again later." };
  }

  if (isUpdated) {
    // Bust the cached staff list so the updated values show on the list page.
    revalidatePath("/department-staff");
    redirect("/department-staff");
  }
}

export default EditStaffAction;
