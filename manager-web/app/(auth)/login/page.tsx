"use client";

import React, { Suspense, useActionState, useState } from "react";
import { useSearchParams } from "next/navigation";
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
import { Users } from "lucide-react";
import { Eye, EyeOff } from "lucide-react";

import Link from "next/link";
import { LoginAction, State } from "@/lib/actions/login";

// Human-readable messages for the ?oauth_error= query param set by the callback
const OAUTH_ERRORS: Record<string, string> = {
  access_denied: "You cancelled the Google sign-in.",
  exchange_failed:
    "Google sign-in failed at the server. Check the backend logs.",
  no_token: "Google sign-in returned no token.",
  missing_code: "Google did not return an authorization code.",
  unexpected: "Unexpected error during Google sign-in.",
};

function LoginForm() {
  const [showPassword, setShowPassword] = useState(false);
  const searchParams = useSearchParams();
  const oauthErrorKey = searchParams.get("oauth_error");
  const oauthErrorMsg = oauthErrorKey
    ? (OAUTH_ERRORS[oauthErrorKey] ??
      `Google sign-in failed (${oauthErrorKey}).`)
    : null;

  const initialState: State = {
    message: "",
  };

  const [state, formAction, pending] = useActionState(
    LoginAction,
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
              <h2 className="text-2xl font-semibold mt-2">Login</h2>
              <p className="text-gray-500 text-sm">
                Please enter your credentials to access the system.
              </p>
            </div>
          </div>

          {/* Form */}
          <div className="mb-6">
            <a
              href="/api/auth/google/login"
              className="w-full inline-flex items-center justify-center rounded-md border border-[#006022] px-4 py-2 text-sm font-medium text-[#006022] transition hover:bg-[#E8F7EC]"
            >
              Continue with Google
            </a>
            {oauthErrorMsg && (
              <p className="mt-3 text-sm text-red-600 text-center">
                {oauthErrorMsg}
              </p>
            )}
          </div>
          <form action={formAction} className="w-full">
            <FieldGroup>
              <FieldSet>
                {/* <FieldLegend>Login Information</FieldLegend>
                <FieldDescription>
                  Please enter your credentials to access the system.
                </FieldDescription> */}

                <FieldGroup className="mt-4">
                  {/* Email or phone */}
                  <Field>
                    <FieldLabel htmlFor="email">Email or Phone</FieldLabel>
                    <Input
                      id="email"
                      type="text"
                      placeholder="Enter email or phone"
                      name="email"
                      required
                    />

                    {state.errors?.email && (
                      <p className="text-red-500 text-sm mt-1">
                        {state.errors.email.join(", ")}
                      </p>
                    )}
                  </Field>

                  {/* Password */}
                  <Field>
                    <FieldLabel htmlFor="password">Password</FieldLabel>
                    <div className="relative">
                      <Input
                        id="password"
                        type={showPassword ? "text" : "password"}
                        placeholder="Enter Password"
                        name="password"
                        className="pr-10"
                        required
                      />
                      <button
                        type="button"
                        onClick={() => setShowPassword(!showPassword)}
                        className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500"
                        aria-label={
                          showPassword ? "Hide password" : "Show password"
                        }
                      >
                        {showPassword ? (
                          <EyeOff size={18} />
                        ) : (
                          <Eye size={18} />
                        )}
                      </button>
                    </div>
                    {state.errors?.password && (
                      <p className="text-red-500 text-sm mt-1">
                        {state.errors.password.join(", ")}
                      </p>
                    )}
                  </Field>
                </FieldGroup>
              </FieldSet>
              {state.message && (
                <p className="text-red-500 text-sm text-center">
                  {state.message}
                </p>
              )}
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

// useSearchParams() requires a Suspense boundary for static prerendering.
export default function LoginPage() {
  return (
    <Suspense fallback={null}>
      <LoginForm />
    </Suspense>
  );
}
