"use client";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { MoreVertical } from "lucide-react";
import { useState } from "react";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Dialog, DialogContent } from "../ui/dialog";

import { useRouter } from "next/navigation";

import { OjtRecord } from "@/types/records";
import { DeleteOjtDialog } from "./deleteOJTdialog";

function OjtTable({ ojtRecords }: { ojtRecords: OjtRecord[] }) {
  const router = useRouter();

  const [open, setOpen] = useState(false);
  const [mode, setMode] = useState<"edit" | "delete">("edit");
  const [selectedRecord, setSelectedRecord] = useState<OjtRecord | null>(null);
  return (
    <>
      <div className="overflow-x-auto border rounded-md">
        <Table className="table-fixed w-full rounded-t-md ">
          <TableHeader className={`bg-[#E8F7EC] rounded-t-md`}>
            <TableRow>
              <TableHead className="w-[250px] font-bold">Plan Name</TableHead>
              {/* <TableHead className="w-[200px] font-bold">Category</TableHead> */}
              <TableHead className="w-[150px] font-bold">Employee ID</TableHead>
              <TableHead className=" w-[150px] font-bold">
                Employee Name
              </TableHead>
              <TableHead className=" w-[150px] font-bold">Position</TableHead>
              <TableHead className=" w-[150px] font-bold">Department</TableHead>
              <TableHead className="w-[120px] font-bold">Division</TableHead>
              <TableHead className="w-[100px] font-bold">Status</TableHead>
              <TableHead className="w-[110px] font-bold">Pre-Test</TableHead>
              <TableHead className="w-[110px] font-bold">Post-Test</TableHead>
              <TableHead className="w-[150px] font-bold">Location</TableHead>

              <TableHead className="w-[150px] font-bold">
                Cost Per Person
              </TableHead>

              <TableHead className="w-[100px] font-bold">Budget Code</TableHead>

              <TableHead className="w-[220px] font-bold">Evaluation</TableHead>

              <TableHead className="w-[100px] font-bold">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {ojtRecords.map((record) => (
              <TableRow key={record.id}>
                <TableCell className="font-medium break-words whitespace-normal">
                  {record.trainingPlanName}
                </TableCell>
                <TableCell>{record.employeeId}</TableCell>
                <TableCell className="break-words whitespace-normal">
                  {record.employeeName}
                </TableCell>
                <TableCell className="break-words whitespace-normal">
                  {record.position}
                </TableCell>
                <TableCell className="break-words whitespace-normal">
                  {record.department}
                </TableCell>
                <TableCell className="break-words whitespace-normal">
                  {record.division}
                </TableCell>
                <TableCell>{record.status.replace("_", " ")}</TableCell>
                <TableCell>{record.preTestScore ?? "—"}</TableCell>
                <TableCell>{record.postTestScore ?? "—"}</TableCell>
                <TableCell className="break-words whitespace-normal">
                  {record.location}
                </TableCell>
                <TableCell>{record.costPerPerson}</TableCell>
                <TableCell>{record.budgetCode}</TableCell>
                <TableCell className="break-words whitespace-normal max-w-[240px]">
                  {record.evaluation || "—"}
                </TableCell>

                <TableCell className="">
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="icon">
                        <MoreVertical />
                      </Button>
                    </DropdownMenuTrigger>

                    <DropdownMenuContent align="end">
                      <DropdownMenuItem
                        onClick={() => router.push(`/ojt-records/${record.id}`)}
                      >
                        Edit
                      </DropdownMenuItem>
                      {/* <DropdownMenuItem
                        onClick={() =>
                          router.push(`/ojt-records/${record.id}`)
                        }
                      >
                        View Details
                      </DropdownMenuItem> */}
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
