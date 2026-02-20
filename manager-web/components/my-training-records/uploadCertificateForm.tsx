import { UploadIcon } from "lucide-react";
import React from "react";

function UploadCertificateForm() {
  return (
    <div>
      <form>
        {/* Feedback */}
        <div className="border rounded-md m-2 p-4 space-y-6">
          <label className="font-medium">Description</label>
          <textarea
            className="w-full border rounded-md p-2"
            rows={4}
            placeholder="Brief description of what you learned..."
          />

          <div className="space-y-2">
            <label className="font-medium">
              Upload Certificate <span className="text-red-500">*</span>
            </label>

            <label className="flex flex-col items-center justify-center w-full h-52 border-2 border-dashed rounded-xl cursor-pointer bg-gray-50 hover:bg-gray-100 transition">
              <div className="flex flex-col items-center justify-center pt-5 pb-6">
                <UploadIcon className="mb-3 w-10 h-10 text-gray-400" />
                <p className="mb-2 text-sm text-gray-500">
                  Click to upload or drag and drop
                </p>
                <p className="text-xs text-gray-400">PDF, PNG, JPG (Max 5MB)</p>
              </div>

              <input
                type="file"
                className="hidden"
                accept=".pdf,.png,.jpg,.jpeg"
                //   onChange={handleFileChange}
              />
            </label>
          </div>
        </div>
      </form>
    </div>
  );
}

export default UploadCertificateForm;
