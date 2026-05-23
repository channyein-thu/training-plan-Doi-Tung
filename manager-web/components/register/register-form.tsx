"use client";

import React, { useActionState, useEffect } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Field,
  FieldGroup,
  FieldLabel,
  FieldSet,
  FieldLegend,
  FieldSeparator,
} from "@/components/ui/field";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { RegisterAction, State } from "@/lib/actions/register";
import { DepartmentList } from "@/types/department";

type RegisterFormProps = {
  departments: DepartmentList[];
};

export default function RegisterForm({ departments }: RegisterFormProps) {
  const initialState: State = { errors: {}, message: null };
  const [state, formAction, pending] = useActionState(
    RegisterAction,
    initialState,
  );
  const router = useRouter();

  useEffect(() => {
    if (state.message === "Registration successful!") {
      router.replace("/login");
    }
  }, [router, state.message]);

  return (
    <form action={formAction}>
      <FieldGroup>
        <FieldSet>
          <FieldLegend>Personal Information</FieldLegend>
          <FieldSeparator />
          <FieldGroup className="grid gap-3 md:grid-cols-2">
            <Field>
              <FieldLabel>Full Name *</FieldLabel>
              <Input name="name" placeholder="Enter your full name" required />
            </Field>

            <Field>
              <FieldLabel>Employee ID *</FieldLabel>
              <Input name="employeeID" placeholder="e.g., EMP001" required />
            </Field>
          </FieldGroup>
        </FieldSet>

        <FieldSet>
          <FieldLegend>Contact Information</FieldLegend>
          <FieldSeparator />
          <FieldGroup className="grid gap-3 md:grid-cols-2">
            <Field>
              <FieldLabel>Email *</FieldLabel>
              <Input
                type="email"
                name="email"
                placeholder="email@company.com"
                required
              />
            </Field>

            <Field>
              <FieldLabel>Phone *</FieldLabel>
              <Input name="phone" placeholder="+66 xxx-xxx-xxxx" required />
            </Field>
          </FieldGroup>
        </FieldSet>

        <FieldSet>
          <FieldLegend>Work Information</FieldLegend>
          <FieldSeparator />
          <FieldGroup className="grid gap-3 md:grid-cols-2">
            <Field>
              <FieldLabel>Department *</FieldLabel>
              <Select name="departmentId" required>
                <SelectTrigger>
                  <SelectValue placeholder="Select department" />
                </SelectTrigger>
                <SelectContent>
                  {departments.length > 0 ? (
                    departments.map((department) => (
                      <SelectItem key={department.id} value={department.id}>
                        {department.name}
                      </SelectItem>
                    ))
                  ) : (
                    <SelectItem value="no-departments" disabled>
                      No departments available
                    </SelectItem>
                  )}
                </SelectContent>
              </Select>
            </Field>

            <Field>
              <FieldLabel>Position *</FieldLabel>
              <Input
                name="position"
                placeholder="e.g., Senior Developer"
                required
              />
            </Field>
          </FieldGroup>
        </FieldSet>

        <FieldSeparator />

        <FieldSet>
          <FieldGroup className="grid gap-3 md:grid-cols-2">
            <Field>
              <FieldLabel>Create Password *</FieldLabel>
              <Input
                type="password"
                name="password"
                placeholder="Enter your password"
                required
              />
            </Field>

            <Field>
              <FieldLabel>Confirm Password *</FieldLabel>
              <Input
                type="password"
                name="confirmPassword"
                placeholder="Re-enter your password"
                required
              />
            </Field>
          </FieldGroup>
        </FieldSet>

        {state.message ? (
          <p className="text-sm text-center text-gray-600">{state.message}</p>
        ) : null}

        <div className="mt-8">
          <Button
            type="submit"
            disabled={pending}
            className="h-12 w-full rounded-lg bg-[#006022] text-base text-white hover:bg-[#004e1b]"
          >
            {pending ? "Submitting..." : "Complete Registration"}
          </Button>
        </div>

        <p className="text-center text-sm">
          Already have an account?{" "}
          <Link href="/login">
            <span className="cursor-pointer font-medium text-[#006022]">
              Login here
            </span>
          </Link>
        </p>
      </FieldGroup>
    </form>
  );
}
