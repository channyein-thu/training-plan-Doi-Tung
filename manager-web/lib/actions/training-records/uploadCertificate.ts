"use server";

import { API_BASE_URL } from "@/lib/api/api";
import { authFetch } from "@/lib/api/authFetch";
import { cookies } from "next/headers";

export type UploadCertificateState = {
  message?: string;
  ok?: boolean;
};

export async function uploadCertificate(
  _prevState: UploadCertificateState,
  formData: FormData,
): Promise<UploadCertificateState> {
  const trainingId = formData.get("trainingId");
  const description = formData.get("description");
  const image = formData.get("image");

  if (!trainingId) {
    return { message: "Training ID is required.", ok: false };
  }

  if (!image || !(image instanceof File) || image.size === 0) {
    return { message: "Certificate file is required.", ok: false };
  }

  const cookieStore = await cookies();
  const cookieHeader = cookieStore
    .getAll()
    .map((c) => `${c.name}=${c.value}`)
    .join("; ");

  const body = new FormData();
  body.append("trainingId", String(trainingId));
  if (description) {
    body.append("description", String(description));
  }
  body.append("image", image);

  const { response } = await authFetch(`${API_BASE_URL}/staff/certificates`, {
    method: "POST",
    credentials: "include",
    headers: { Cookie: cookieHeader },
    body,
    cache: "no-store",
  });

  if (!response.ok) {
    return { message: "Failed to upload certificate.", ok: false };
  }

  return { message: "Upload successful.", ok: true };
}
