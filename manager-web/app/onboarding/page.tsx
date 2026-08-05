"use client";

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { PUBLIC_API_BASE_URL } from "@/app/api/api";
import { completeProfileAction } from "@/lib/actions/completeProfile";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Field, FieldGroup, FieldLabel, FieldSet } from "@/components/ui/field";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

type FormState = {
  employeeId: string;
  departmentId: string;
  phone: string;
  position: string;
};

type Department = {
  id: number;
  name: string;
};

export default function OnboardingPage() {
  const router = useRouter();
  const [formState, setFormState] = useState<FormState>({
    employeeId: "",
    departmentId: "",
    phone: "",
    position: "",
  });
  const [departments, setDepartments] = useState<Department[]>([]);
  const [isDepartmentsLoading, setIsDepartmentsLoading] = useState(true);
  const [error, setError] = useState<string>("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    let isMounted = true;

    const loadDepartments = async () => {
      try {
        const response = await fetch(`${PUBLIC_API_BASE_URL}/departments-list`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
          cache: "no-store",
        });

        if (!response.ok) {
          throw new Error("Failed to load departments.");
        }

        const payload = await response.json();
        const data = Array.isArray(payload?.data) ? payload.data : [];

        if (isMounted) {
          setDepartments(data);
        }
      } catch (err) {
        if (isMounted) {
          const message =
            err instanceof Error ? err.message : "Failed to load departments.";
          setError(message);
        }
      } finally {
        if (isMounted) {
          setIsDepartmentsLoading(false);
        }
      }
    };

    loadDepartments();

    return () => {
      isMounted = false;
    };
  }, []);

  const handleChange = (key: keyof FormState, value: string) => {
    setFormState((prev) => ({ ...prev, [key]: value }));
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setError("");
    setIsSubmitting(true);

    try {
      const result = await completeProfileAction({
        employeeId: formState.employeeId,
        departmentId: Number(formState.departmentId),
        phone: formState.phone || undefined,
        position: formState.position || undefined,
      });

      if (!result.ok) {
        throw new Error(result.message || "Profile completion failed.");
      }

      router.replace("/dashboard");
    } catch (err) {
      const message =
        err instanceof Error ? err.message : "Something went wrong.";
      setError(message);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-[#F6FAF7] p-4">
      <Card className="w-full max-w-lg">
        <CardHeader>
          <CardTitle>Complete Your Profile</CardTitle>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="w-full">
            <FieldGroup>
              <FieldSet>
                <Field>
                  <FieldLabel htmlFor="employeeId">Employee ID</FieldLabel>
                  <Input
                    id="employeeId"
                    name="employeeId"
                    value={formState.employeeId}
                    onChange={(event) =>
                      handleChange("employeeId", event.target.value)
                    }
                    placeholder="Enter Employee ID"
                    required
                  />
                </Field>

                <Field>
                  <FieldLabel>Department</FieldLabel>
                  <Select
                    value={formState.departmentId}
                    onValueChange={(value) =>
                      handleChange("departmentId", value)
                    }
                  >
                    <SelectTrigger className="w-full">
                      <SelectValue
                        placeholder={
                          isDepartmentsLoading
                            ? "Loading departments..."
                            : "Select Department"
                        }
                      />
                    </SelectTrigger>
                    <SelectContent>
                      {departments.map((department) => (
                        <SelectItem
                          key={department.id}
                          value={String(department.id)}
                        >
                          {department.name}
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </Field>

                <Field>
                  <FieldLabel htmlFor="phone">Phone (optional)</FieldLabel>
                  <Input
                    id="phone"
                    name="phone"
                    value={formState.phone}
                    onChange={(event) =>
                      handleChange("phone", event.target.value)
                    }
                    placeholder="e.g. +66 90 000 0000"
                  />
                </Field>

                <Field>
                  <FieldLabel htmlFor="position">
                    Position (optional)
                  </FieldLabel>
                  <Input
                    id="position"
                    name="position"
                    value={formState.position}
                    onChange={(event) =>
                      handleChange("position", event.target.value)
                    }
                    placeholder="Enter Position"
                  />
                </Field>
              </FieldSet>

              {error && (
                <p className="text-red-500 text-sm text-center">{error}</p>
              )}

              <Button
                type="submit"
                className="bg-[#006022] text-white hover:bg-[#005018] w-full"
                disabled={isSubmitting}
              >
                {isSubmitting ? "Saving..." : "Complete Profile"}
              </Button>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
