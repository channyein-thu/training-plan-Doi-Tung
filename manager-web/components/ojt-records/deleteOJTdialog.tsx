import { OjtRecord } from "@/types/records";
import { Button } from "../ui/button";
import { DialogTitle } from "../ui/dialog";

type Props = {
  record: OjtRecord;
  onConfirm: () => void;
  onCancel: () => void;
};

export function DeleteOjtDialog({ record, onConfirm, onCancel }: Props) {
  return (
    <div className="space-y-4">
      <DialogTitle className="text-lg font-semibold text-red-600">
        Delete OJT Record
      </DialogTitle>

      <p className="text-sm">
        Are you sure you want to delete{" "}
        <span className="font-medium font-semibold">{record.course.name}</span>?
        This action cannot be undone.
      </p>

      <div className="flex justify-end gap-2">
        <Button variant="outline" onClick={onCancel}>
          Cancel
        </Button>

        <Button variant="destructive" onClick={onConfirm}>
          Delete
        </Button>
      </div>
    </div>
  );
}
