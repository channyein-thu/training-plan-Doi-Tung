"use client";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import {
  State,
  UpdateOjtRecordAction,
} from "@/lib/actions/ojt-records/updateOjtRecord";
import { useActionState, useState } from "react";
import { OjtRecord } from "@/types/records";

function EditOJTdetails({ record }: { record: OjtRecord }) {
  const initialState: State = { errors: {}, message: "" };
  const [status, setStatus] = useState(record.status);

  const [state, formAction, pending] = useActionState(
    UpdateOjtRecordAction.bind(null, record.id),
    initialState,
  );

  return (
    <form action={formAction} className="border rounded-md m-2 p-4 space-y-6">
      <input type="hidden" name="status" value={status} />
      <p className="font-medium">Please start fill from here</p>

      {/* Status */}
      <div className="space-y-2">
        <label className="font-medium">User Status</label>

        <Select value={status} onValueChange={setStatus as any}>
          <SelectTrigger className="w-[220px]">
            <SelectValue placeholder="Select status" />
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="Register">Register</SelectItem>
            <SelectItem value="Attended">Attended</SelectItem>
            <SelectItem value="Absent">Absent</SelectItem>
          </SelectContent>
        </Select>

        {state?.errors?.status && (
          <p className="text-sm text-red-500">{state.errors.status[0]}</p>
        )}
      </div>

      {/* Evaluation */}
      <div className="space-y-2">
        <label className="font-medium">
          Evaluate the Practical Application of Knowledge to Work
        </label>

        <textarea
          name="evaluation"
          className="w-full border rounded-md p-2"
          rows={4}
          placeholder="Describe how the training knowledge is applied..."
        />

        {state?.errors?.evaluation && (
          <p className="text-sm text-red-500">{state.errors.evaluation[0]}</p>
        )}
      </div>

      {/* Scores */}
      <div className="grid grid-cols-2 gap-8">
        {/* Pretest */}
        <div className="space-y-2">
          <label className="text-lg font-medium">Pretest Score</label>

          <div className="relative">
            <Input
              name="preTestScore"
              type="number"
              min={0}
              max={100}
              className="h-14 text-xl rounded-xl pr-10"
            />

            <span className="absolute right-4 top-1/2 -translate-y-1/2 text-lg text-muted-foreground">
              %
            </span>
          </div>

          {state?.errors?.preTestScore && (
            <p className="text-sm text-red-500">
              {state.errors.preTestScore[0]}
            </p>
          )}
        </div>

        {/* Posttest */}
        <div className="space-y-2">
          <label className="text-lg font-medium">Posttest Score</label>

          <div className="relative">
            <Input
              name="postTestScore"
              type="number"
              min={0}
              max={100}
              className="h-14 text-xl rounded-xl pr-10"
            />

            <span className="absolute right-4 top-1/2 -translate-y-1/2 text-lg text-muted-foreground">
              %
            </span>
          </div>

          {state?.errors?.postTestScore && (
            <p className="text-sm text-red-500">
              {state.errors.postTestScore[0]}
            </p>
          )}
        </div>
      </div>

      {/* Action Buttons */}
      <div className="px-6 py-4 flex gap-4">
        <Button type="button" variant="outline" className="flex-1">
          Cancel
        </Button>

        <Button
          type="submit"
          disabled={pending}
          className="flex-1 bg-[#006022] hover:bg-[#004d1b] text-white"
        >
          {pending ? "Saving..." : "Save Changes"}
        </Button>
      </div>

      {state?.message && (
        <p className="text-sm text-red-500">{state.message}</p>
      )}
    </form>
  );
}

export default EditOJTdetails;
