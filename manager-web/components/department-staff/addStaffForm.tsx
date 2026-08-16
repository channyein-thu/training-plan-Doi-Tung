"use client";
import React, { useActionState } from "react";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { email, z } from "zod";
import {
  DialogClose,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import CreateStaffAction, {
  State,
} from "@/lib/actions/department-staff/createStaffAction";
import { Field, FieldGroup, FieldLabel } from "../ui/field";

const formSchema = z.object({
  name: z.string().min(2).max(50),
  employeeID: z.string().min(1).max(52),
  email: z.string().email(),
  phone: z.string().max(20),
  position: z.string().min(1).max(100),
  password: z.string().min(6).max(100),
  status: z.enum(["Active", "Inactive", "Suspended"]),
});

function DialogForm() {
  const initialState: State = { errors: {}, message: null };

  const [state, formAction, pending] = useActionState(
    CreateStaffAction,
    initialState,
  );
  // const form = useForm<z.infer<typeof formSchema>>({
  //   resolver: zodResolver(formSchema),
  //   defaultValues: {
  //     name: "",
  //     empId: "",
  //     email: "",
  //     phone: "",
  //     position: "",
  //     password: "",
  //     status: "active",
  //   },
  // });

  return (
    <form action={formAction}>
      <DialogHeader>
        <DialogTitle>Add New Staff</DialogTitle>
      </DialogHeader>

      <div className="grid grid-cols-2 gap-4 mt-4">
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="name">Add Staff Name</FieldLabel>
            <Input
              id="name"
              type="text"
              placeholder="Enter Staff Name"
              name="name"
              required
            />
          </Field>
        </FieldGroup>
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="employeeID">Add Employee ID</FieldLabel>
            <Input
              id="employeeID"
              type="text"
              placeholder="Enter Employee ID"
              name="employeeID"
              required
            />
          </Field>
        </FieldGroup>
      </div>
      <div className="grid grid-cols-2 gap-4 mt-4">
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="email">Add Email</FieldLabel>
            <Input
              id="email"
              type="text"
              placeholder="Enter Email"
              name="email"
              required
            />
          </Field>
        </FieldGroup>
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="phone">Add Phone</FieldLabel>
            <Input
              id="phone"
              type="text"
              placeholder="Enter Phone"
              name="phone"
              required
            />
          </Field>
        </FieldGroup>
      </div>

      <div className="grid grid-cols-2 gap-4 mt-4">
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="position">Add Position</FieldLabel>
            <Input
              id="position"
              type="text"
              placeholder="Enter Position"
              name="position"
              required
            />
          </Field>
        </FieldGroup>
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="password">Add Password</FieldLabel>
            <Input
              id="password"
              type="text"
              placeholder="Enter Password"
              name="password"
              required
            />
          </Field>
        </FieldGroup>
      </div>

      <div className="grid grid-cols-2 gap-4 mt-4">
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="status">Select Status</FieldLabel>

            <select
              id="status"
              name="status"
              required
              className="w-full border border-[#006022] rounded-md px-3 py-2"
            >
              <option value="Active">Active</option>

              <option value="Inactive">Inactive</option>

              <option value="Suspended">Suspended</option>
            </select>
          </Field>
        </FieldGroup>
        <FieldGroup>
          <Field>
            <FieldLabel htmlFor="workStartDate">Work Start Date</FieldLabel>
            <Input
              id="workStartDate"
              name="workStartDate"
              type="date"
              required
            />
          </Field>
        </FieldGroup>
      </div>

      <DialogFooter className="mt-4">
        <DialogClose asChild>
          <Button type="button" variant="outline">
            Cancel
          </Button>
        </DialogClose>
        <Button
          type="submit"
          className="bg-[#006022] text-white hover:bg-[#005018]"
        >
          Add Department Staff
        </Button>
      </DialogFooter>
    </form>
  );
}

export default DialogForm;
