"use server";
import { API_BASE_URL } from "@/lib/api/api";
import { cookies } from "next/headers";
import { redirect } from "next/navigation";
import { z } from "zod";

export type State = {
  errors?: {
    name?: string[];
    empId?: string[];
    email?: string[];
    phone?: string[];
    position?: string[];
    password?: string[];
    status?: string[];
  };
  message?: string | null;
};
