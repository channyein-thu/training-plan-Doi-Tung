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
import { DeleteOjtDialog } from "./deleteOJTdialog";

function OjtTable() {
  const router = useRouter();

  const [open, setOpen] = useState(false);
  const [mode, setMode] = useState<"edit" | "delete">("edit");
  const [selectedRecord, setSelectedRecord] = useState<OjtRecord | null>(null);
  return (
    <>
      <div className="overflow-x-auto">
        <Table className="w-full table-fixed text-sm">
          <TableCaption>A list of training plans</TableCaption>

          <TableHeader style={{ backgroundColor: "#E8F7EC" }}>
            <TableRow>
              <TableHead className="w-[250px] font-semibold">
                Training Plan
              </TableHead>
              <TableHead className="w-[150px] font-semibold">
                Location
              </TableHead>
              <TableHead className="w-[140px] font-semibold">
                Cost Per Person
              </TableHead>
              <TableHead className="w-[140px] font-semibold">
                Budget Code
              </TableHead>
              <TableHead className="w-[120px] font-semibold">
                Employee ID
              </TableHead>
              <TableHead className="w-[160px] font-semibold">
                Name-Surname
              </TableHead>
              <TableHead className="w-[140px] font-semibold">
                Position
              </TableHead>
              <TableHead className="w-[140px] font-semibold">
                Department
              </TableHead>
              <TableHead className="w-[180px] font-semibold">
                Division
              </TableHead>
              <TableHead className="w-[120px] font-semibold">Status</TableHead>
              <TableHead className="w-[160px] font-semibold">
                Pre Test Score
              </TableHead>
              <TableHead className="w-[160px] font-semibold">
                Post Test Score
              </TableHead>
              <TableHead className="w-[120px] font-semibold">
                Evaluation
              </TableHead>
            </TableRow>
          </TableHeader>

          <TableBody>
            {ojtRecords.map((record) => (
              <TableRow key={record.id}>
                <TableCell>
                  <div className="max-w-[180px] line-clamp-2 break-words">
                    {record.course.name}
                  </div>
                </TableCell>

                <TableCell>
                  <div className="max-w-[140px] line-clamp-2 break-words">
                    {record.course.location}
                  </div>
                </TableCell>

                <TableCell className="text-center">
                  {record.course.costPerPerson}
                </TableCell>

                <TableCell>
                  <div className="max-w-[140px] line-clamp-2 break-words">
                    {record.course.budgetCode}
                  </div>
                </TableCell>

                <TableCell>{record.staff.id}</TableCell>

                <TableCell>
                  <div className="max-w-[160px] line-clamp-2 break-words">
                    {record.staff.name}
                  </div>
                </TableCell>

                <TableCell>
                  <div className="max-w-[140px] line-clamp-2 break-words">
                    {record.staff.position}
                  </div>
                </TableCell>

                <TableCell className="text-center">
                  <div className="max-w-[140px] line-clamp-2 break-words">
                    {record.staff.department.name}
                  </div>
                </TableCell>

                <TableCell>
                  <div className="max-w-[180px] line-clamp-2 break-words">
                    {record.staff.department.division}
                  </div>
                </TableCell>

                <TableCell>{record.status}</TableCell>

                <TableCell className="text-center">85</TableCell>
                <TableCell className="text-center">85</TableCell>

                <TableCell>Excellent</TableCell>
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
            {mode === "delete" && (
              <DeleteOjtDialog
                record={selectedRecord!}
                onConfirm={() => {
                  // deleteUser(selectedRecord!.staff.id);
                  setOpen(false);
                }}
                onCancel={() => setOpen(false)}
              />
            )}
          </DialogContent>
        </Dialog>
      </div>
    </>
  );
}

export default OjtTable;
