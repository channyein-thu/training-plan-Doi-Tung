"use client";

import { useActionState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Field, FieldGroup, FieldLabel } from "@/components/ui/field";
import { TrainingPlanStaff } from "../../types/staff";
import EditStaffAction, {
  State,
} from "@/lib/actions/department-staff/editStaffAction";

type Props = {
  staff: TrainingPlanStaff | null;
  onClose: () => void;
};

function EditStaffDialog({ staff, onClose }: Props) {
  const initialState: State = { errors: {}, message: null };
  const [state, formAction, pending] = useActionState(
    // Bind the staff id so the action can PUT /manager/users/:id.
    EditStaffAction.bind(null, staff?.id ?? 0),
    initialState,
  );

  if (!staff) return null;

  return (
    <Dialog open={true} onOpenChange={(open) => !open && onClose()}>
      <DialogContent className="max-w-lg">
        <DialogHeader>
          <DialogTitle>Edit Staff Member</DialogTitle>
        </DialogHeader>

        <form action={formAction}>
          <div className="grid grid-cols-2 gap-4">
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="name">Full Name</FieldLabel>
                <Input id="name" name="name" defaultValue={staff.name} required />
              </Field>
            </FieldGroup>
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="employeeID">Employee ID</FieldLabel>
                <Input
                  id="employeeID"
                  name="employeeID"
                  defaultValue={staff.employeeID}
                  required
                />
              </Field>
            </FieldGroup>
          </div>

          <div className="grid grid-cols-2 gap-4 mt-4">
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="email">Email</FieldLabel>
                <Input
                  id="email"
                  name="email"
                  type="email"
                  defaultValue={staff.email}
                  required
                />
              </Field>
            </FieldGroup>
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="phone">Phone</FieldLabel>
                <Input
                  id="phone"
                  name="phone"
                  defaultValue={staff.phone}
                  required
                />
              </Field>
            </FieldGroup>
          </div>

          <div className="grid grid-cols-2 gap-4 mt-4">
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="position">Position</FieldLabel>
                <Input
                  id="position"
                  name="position"
                  defaultValue={staff.position}
                  required
                />
              </Field>
            </FieldGroup>
            <FieldGroup>
              <Field>
                <FieldLabel htmlFor="status">Status</FieldLabel>
                <select
                  id="status"
                  name="status"
                  required
                  defaultValue={staff.status}
                  className="w-full border border-[#006022] rounded-md px-3 py-2"
                >
                  <option value="Active">Active</option>
                  <option value="Inactive">Inactive</option>
                  <option value="Suspended">Suspended</option>
                </select>
              </Field>
            </FieldGroup>
          </div>

          {state?.message ? (
            <p className="mt-3 text-sm text-red-600">{state.message}</p>
          ) : null}

          <DialogFooter className="mt-4">
            <Button type="button" variant="outline" onClick={onClose}>
              Cancel
            </Button>
            <Button
              type="submit"
              disabled={pending}
              className="bg-[#006022] text-white hover:bg-[#005018]"
            >
              {pending ? "Saving..." : "Save"}
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}

export default EditStaffDialog;
