import { OjtRecord } from "@/types/records";
import { Button } from "../ui/button";
import { DialogTitle } from "../ui/dialog";
import { DeleteOjtAction, State } from "@/lib/actions/ojt-records/deleteOjt";
import { useActionState } from "react";

type Props = {
  record: OjtRecord;
  onConfirm: () => void;
  onCancel: () => void;
};

export function DeleteOjtDialog({ record, onCancel }: Props) {
  const initialState: State = {
    errors: {},
    message: null,
  };

  const [state, formAction, pending] = useActionState(
    DeleteOjtAction.bind(null, Number(record.id)),
    initialState,
  );
  return (
    <div className="space-y-4">
      <DialogTitle className="text-lg font-semibold text-red-600">
        Delete OJT Record
      </DialogTitle>

      <p className="text-sm">
        Are you sure you want to delete{" "}
        <span className="font-medium font-semibold">
          {record.trainingPlanName}
        </span>
        ? This action cannot be undone.
      </p>

      <div className="flex justify-end gap-2">
        <Button variant="outline" onClick={onCancel}>
          Cancel
        </Button>
        <form action={formAction}>
          <Button variant="destructive" type="submit" disabled={pending}>
            {pending ? "Deleting..." : "Delete"}
          </Button>
        </form>
      </div>
    </div>
  );
}
