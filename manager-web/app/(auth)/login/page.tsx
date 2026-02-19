"use client";

import React, { useActionState, useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Field,
  FieldGroup,
  FieldLabel,
  FieldDescription,
  FieldSet,
  FieldLegend,
} from "@/components/ui/field";
import Logo from "@/components/login/logo";
import { Card, CardContent } from "@/components/ui/card";
import { PersonStanding, Shield, Users } from "lucide-react";
import Link from "next/link";
import { LoginAction, State } from "@/lib/actions/login";

export default function LoginForm() {
  const initialState: State = { errors: {}, message: null };

  const [state, formAction, pending] = useActionState(
    LoginAction as any,
    initialState,
  );

  return (
    <div className="flex flex-col items-center gap-6 w-full max-w-sm">
      {/* Logo + System Title */}
      <div className="flex flex-col items-center gap-2">
        <Logo />
        <h1 className="text-[#4A5565] text-center">
          Training and Record Management System
        </h1>
      </div>

      <Card className="w-full">
        <CardContent>
          {/* Header Section */}
          <div className="flex items-start gap-4 mb-2">
            <div className="bg-[#006022] w-16 h-16 flex items-center justify-center text-white rounded-lg">
              <Users size={32} />
            </div>
            <div>
              <h2 className="text-2xl font-semibold mt-2">Manager Login</h2>
              <p className="text-gray-500 text-sm">Department manager login</p>
            </div>
          </div>

          {/* Form */}
          <form onSubmit={formAction} className="w-full">
            <FieldGroup>
              <FieldSet>
                {/* <FieldLegend>Login Information</FieldLegend>
                <FieldDescription>
                  Please enter your credentials to access the system.
                </FieldDescription> */}

                <FieldGroup className="mt-4">
                  {/* Email */}
                  <Field>
                    <FieldLabel htmlFor="email">Email Address</FieldLabel>
                    <Input
                      id="email"
                      type="email"
                      placeholder="Enter Email"
                      name="email"
                      required
                    />
                  </Field>

                  {/* Password */}
                  <Field>
                    <FieldLabel htmlFor="password">Password</FieldLabel>
                    <Input
                      id="password"
                      type="password"
                      placeholder="Enter Password"
                      name="password"
                      required
                    />
                  </Field>
                </FieldGroup>
              </FieldSet>

              {/* Button */}
              <Field orientation="horizontal" className="">
                <Button
                  type="submit"
                  className="bg-[#006022] text-white hover:bg-[#005018] w-full"
                >
                  Login
                </Button>
              </Field>
            </FieldGroup>
          </form>
          <p className="text-sm mt-2 items-center text-center">
            Don't have an account?{" "}
            <Link href="/register" className="text-[#006022] underline">
              Register here
            </Link>
          </p>
        </CardContent>
      </Card>
    </div>
  );
}
