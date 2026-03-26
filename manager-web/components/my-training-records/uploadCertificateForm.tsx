"use client";

import { UploadIcon } from "lucide-react";
import React, { useEffect, useState } from "react";
import { Button } from "../ui/button";
import { useActionState } from "react";
import {
  uploadCertificate,
  UploadCertificateState,
} from "@/lib/actions/training-records/uploadCertificate";

type UploadCertificateFormProps = {
  trainingId: number | string;
};

const initialState: UploadCertificateState = { message: "", ok: undefined };

function UploadCertificateForm({ trainingId }: UploadCertificateFormProps) {
  const [state, formAction, pending] = useActionState(
    uploadCertificate,
    initialState,
  );
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);

  useEffect(() => {
    if (!selectedFile) {
      setPreviewUrl(null);
      return;
    }

    if (!selectedFile.type.startsWith("image/")) {
      setPreviewUrl(null);
      return;
    }

    const url = URL.createObjectURL(selectedFile);
    setPreviewUrl(url);
    return () => URL.revokeObjectURL(url);
  }, [selectedFile]);

  return (
    <div>
      <form action={formAction}>
        {/* Feedback */}
        <div className="m-2 rounded-2xl border bg-white/80 p-6 shadow-sm backdrop-blur">
          <input type="hidden" name="trainingId" value={String(trainingId)} />
          <div className="grid gap-6 lg:grid-cols-[1.1fr_1.4fr]">
            <div className="space-y-3">
              <label className="text-sm font-semibold text-gray-700">
                Description
              </label>
              <textarea
                className="min-h-[140px] w-full rounded-xl border border-gray-200 bg-white/70 p-3 text-sm shadow-inner focus:border-[#006022] focus:outline-none focus:ring-2 focus:ring-[#006022]/20"
                rows={5}
                placeholder="Brief description of what you learned..."
                name="description"
              />
              <div className="flex items-center justify-between rounded-xl bg-[#F4FAF6] px-3 py-2 text-xs text-gray-600">
                <span>Keep it short and specific.</span>
                <span>Optional</span>
              </div>
            </div>

            <div className="space-y-3">
              <label className="text-sm font-semibold text-gray-700">
                Upload Certificate <span className="text-red-500">*</span>
              </label>

              <label className="group flex flex-col items-center justify-center w-full rounded-2xl border border-dashed border-gray-300 bg-gradient-to-br from-white via-[#F6FBF8] to-[#EDF6F0] p-6 text-center shadow-sm transition hover:border-[#006022] hover:shadow-md cursor-pointer">
                <div className="flex flex-col items-center gap-3">
                  <div className="flex h-14 w-14 items-center justify-center rounded-full bg-white shadow">
                    {previewUrl ? (
                      <img
                        src={previewUrl}
                        alt="Certificate preview"
                        className="h-12 w-12 rounded-full object-cover"
                      />
                    ) : (
                      <UploadIcon className="h-6 w-6 text-[#006022]" />
                    )}
                  </div>
                  <div>
                    <p className="text-sm font-medium text-gray-700">
                      Click to upload or drag and drop
                    </p>
                    <p className="text-xs text-gray-500">
                      PDF, PNG, JPG up to 5MB
                    </p>
                  </div>
                  {selectedFile ? (
                    <div className="rounded-full bg-white px-3 py-1 text-xs text-gray-600 shadow">
                      Selected: {selectedFile.name}
                    </div>
                  ) : null}
                </div>

                <input
                  type="file"
                  className="hidden"
                  accept=".pdf,.png,.jpg,.jpeg"
                  name="image"
                  required
                  onChange={(event) =>
                    setSelectedFile(event.target.files?.[0] ?? null)
                  }
                />
              </label>
            </div>
          </div>

          {state.message ? (
            <p
              className={`mt-4 rounded-lg border px-3 py-2 text-sm ${
                state.ok
                  ? "border-green-200 bg-green-50 text-green-700"
                  : "border-red-200 bg-red-50 text-red-700"
              }`}
            >
              {state.message}
            </p>
          ) : null}

          {/* Action Buttons */}
          <div className="mt-6 flex flex-col gap-3 border-t border-gray-100 pt-4 sm:flex-row">
            <Button type="button" variant="outline" className="flex-1">
              Cancel
            </Button>
            <Button
              type="submit"
              disabled={pending}
              className="flex-1 bg-[#006022] hover:bg-[#004d1b] text-white"
            >
              {pending ? "Uploading..." : "Upload Certificate"}
            </Button>
          </div>
        </div>
      </form>
    </div>
  );
}

export default UploadCertificateForm;
