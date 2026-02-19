"use client";

import React, { useActionState, useState } from "react";
import Logo from "@/components/login/logo";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Field,
  FieldGroup,
  FieldLabel,
  FieldSet,
  FieldLegend,
  FieldDescription,
  FieldSeparator,
} from "@/components/ui/field";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Card, CardContent } from "@/components/ui/card";
import { Users } from "lucide-react";
import { init } from "next/dist/compiled/webpack/webpack";
import { LoginAction } from "@/lib/actions/login";
import { RegisterAction } from "@/lib/actions/register";

export default function RegisterPage() {
  const initialState = { errors: {}, message: null };
  const [state, formAction, pending] = useActionState(
    RegisterAction as any,
    initialState,
  );

  return (
    <div className="min-h-screen bg-[#dbe7dd] flex flex-col items-center py-10 px-4 w-[100%]">
      {/* Logo */}
      <div className="mb-6">
        <Logo />
      </div>

      <Card className="w-full max-w-4xl rounded-2xl shadow-lg">
        <CardContent className="p-8">
          {/* Header */}
          <div className="flex items-center gap-4 mb-8 justify-center">
            <div className="bg-[#006022] w-14 h-14 flex items-center justify-center text-white rounded-xl">
              <Users size={28} />
            </div>
            <div>
              <h2 className="text-2xl font-semibold">Manager Registration</h2>
              <p className="text-gray-500 text-sm">
                Fill in your details below
              </p>
            </div>
          </div>

          <form onSubmit={formAction}>
            <FieldGroup>
              {/* ================= PERSONAL INFO ================= */}
              <FieldSet>
                <FieldLegend>Personal Information</FieldLegend>
                <FieldSeparator />
                <FieldGroup className="grid md:grid-cols-2 gap-3">
                  <Field>
                    <FieldLabel>Full Name *</FieldLabel>
                    <Input
                      name="fullName"
                      placeholder="Enter your full name"
                      required
                    />
                  </Field>

                  <Field>
                    <FieldLabel>Employee ID *</FieldLabel>
                    <Input
                      name="employeeId"
                      placeholder="e.g., EMP001"
                      required
                    />
                  </Field>
                </FieldGroup>
              </FieldSet>

              {/* ================= CONTACT INFO ================= */}
              <FieldSet>
                <FieldLegend>Contact Information</FieldLegend>
                <FieldSeparator />
                <FieldGroup className="grid md:grid-cols-2 gap-3">
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
                    <Input
                      name="phone"
                      placeholder="+66 xxx-xxx-xxxx"
                      required
                    />
                  </Field>
                </FieldGroup>
              </FieldSet>

              {/* ================= WORK INFO ================= */}
              <FieldSet>
                <FieldLegend>Work Information</FieldLegend>
                <FieldSeparator />

                <FieldGroup className="grid md:grid-cols-2 gap-3">
                  <Field>
                    <FieldLabel>Department *</FieldLabel>
                    <Select>
                      <SelectTrigger>
                        <SelectValue placeholder="All Departments" />
                      </SelectTrigger>
                      <SelectContent>
                        <SelectItem value="IT">IT</SelectItem>
                        <SelectItem value="HR">HR</SelectItem>
                        <SelectItem value="Finance">Finance</SelectItem>
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

                  <Field>
                    <FieldLabel>Agency</FieldLabel>
                    <Input name="agency" placeholder="e.g., HQ, Branch A" />
                  </Field>

                  <Field>
                    <FieldLabel>Cotton</FieldLabel>
                    <Input name="cotton" placeholder="e.g., A, B, C" />
                  </Field>

                  <Field>
                    <FieldLabel>Line</FieldLabel>
                    <Input name="line" placeholder="e.g., 1, 2, 3" />
                  </Field>
                </FieldGroup>
              </FieldSet>

              <FieldSeparator />

              {/* ================= PASSWORD ================= */}
              <FieldSet>
                <FieldGroup className="grid md:grid-cols-2 gap-3">
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

              {/* Submit */}
              <div className="mt-8">
                <Button
                  type="submit"
                  className="w-full bg-[#006022] hover:bg-[#004e1b] text-white h-12 text-base rounded-lg"
                >
                  Complete Registration
                </Button>
              </div>

              <p className="text-center text-sm">
                Already have an account?{" "}
                <span className="text-[#006022] font-medium cursor-pointer">
                  Login here
                </span>
              </p>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
