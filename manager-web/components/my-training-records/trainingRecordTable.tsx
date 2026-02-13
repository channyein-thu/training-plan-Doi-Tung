"use client";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { MoreVertical } from "lucide-react";
import React, { useState } from "react";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Dialog, DialogContent } from "../ui/dialog";

import { useRouter } from "next/navigation";

import { ojtRecords, trainingPlanStaff } from "@/data/data";
import { OjtRecord } from "@/types/records";
// import { DeleteOjtDialog } from "./deleteOJTdialog";

function OjtTable() {
  const router = useRouter();

  const [open, setOpen] = useState(false);
  const [mode, setMode] = useState<"edit" | "delete">("edit");
  const [selectedRecord, setSelectedRecord] = useState<OjtRecord | null>(null);
  return (
    <>
      <div className="">
        <Table className="table-fixed w-full rounded-t-md ">
          <TableCaption>A list of OJT records</TableCaption>
          <TableHeader className={`bg-[#E8F7EC] rounded-t-md`}>
            <TableRow>
              <TableHead className="w-[250px] font-bold">Course Name</TableHead>
              <TableHead className="w-[200px] font-bold">Category</TableHead>
              <TableHead className="w-[150px] font-bold">Name</TableHead>
              <TableHead className=" w-[150px] font-bold">
                Employee ID
              </TableHead>
              <TableHead className=" w-[150px] font-bold">Date</TableHead>
              <TableHead className=" w-[100px] font-bold">Hours</TableHead>
              <TableHead className="w-[100px] font-bold">Location</TableHead>
              <TableHead className="w-[100px] font-bold">Status</TableHead>
              <TableHead className="w-[100px] font-bold">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {ojtRecords.map((record) => (
              <TableRow key={record.id}>
                <TableCell className="font-medium">
                  {record.course.name}
                </TableCell>
                <TableCell>{record.course.category}</TableCell>
                <TableCell>{record.staff.fullName}</TableCell>
                <TableCell>{record.staff.employeeId}</TableCell>
                <TableCell>{record.course.date}</TableCell>
                <TableCell>{record.course.numberOfHours}</TableCell>
                <TableCell>{record.course.location}</TableCell>
                <TableCell>{record.status.replace("_", " ")}</TableCell>

                <TableCell className="">
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="icon">
                        <MoreVertical />
                      </Button>
                    </DropdownMenuTrigger>

                    <DropdownMenuContent align="end">
                      <DropdownMenuItem
                        onClick={() =>
                          router.push(`/ojt-records/${record.staff.id}`)
                        }
                      >
                        Edit
                      </DropdownMenuItem>
                      <DropdownMenuItem
                        onClick={() =>
                          router.push(`/ojt-records/${record.staff.id}`)
                        }
                      >
                        View Details
                      </DropdownMenuItem>
                      <DropdownMenuItem
                        className="text-red-600"
                        onClick={() => {
                          setMode("delete");
                          setSelectedRecord(record);
                          setOpen(true);
                        }}
                      >
                        Delete
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogContent>
            {/* {mode === "edit" && (
              <EditUserForm
                user={selectedRecord?.staff ?? undefined}
                onClose={() => setOpen(false)}
              />
            )} */}
            {/* {mode === "delete" && (
              <DeleteOjtDialog
                record={selectedRecord!}
                onConfirm={() => {
                  // deleteUser(selectedRecord!.staff.id);
                  setOpen(false);
                }}
                onCancel={() => setOpen(false)}
              />
            )} */}
          </DialogContent>
        </Dialog>
      </div>
    </>
  );
}

export default OjtTable;
