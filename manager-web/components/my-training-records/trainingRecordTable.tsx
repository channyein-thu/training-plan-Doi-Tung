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
import { DownloadIcon, UploadIcon } from "lucide-react";

import { Button } from "../ui/button";
import { DropdownMenu, DropdownMenuTrigger } from "../ui/dropdown-menu";

import { useRouter } from "next/navigation";
import { OjtRecord } from "@/types/records";

// import { DeleteOjtDialog } from "./deleteOJTdialog";

function TrainingRecordTable({ records }: { records: OjtRecord[] }) {
  const router = useRouter();

  return (
    <>
      <div className="">
        <Table className="w-full table-fixed text-sm">
          <TableCaption>A list of training records</TableCaption>
          <TableHeader className={`bg-[#E8F7EC] rounded-t-md`}>
            <TableRow>
              <TableHead className="w-[250px] font-semibold">
                Training Plan
              </TableHead>
              <TableHead className="w-[200px] font-semibold">
                Location
              </TableHead>
              <TableHead className="w-[140px] font-semibold">
                Cost Per Person
              </TableHead>
              <TableHead className="w-[140px] font-semibold">
                Budget Code
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
              <TableHead className="w-[150px] font-bold">
                Upload Certificate
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {records.map((record) => (
              <TableRow key={record.id}>
                <TableCell>
                  <div className="max-w-[250px] line-clamp-2 break-words">
                    {record.trainingPlanName}
                  </div>
                </TableCell>

                <TableCell>
                  <div className="max-w-[250px] line-clamp-2 break-words">
                    {record.location}
                  </div>
                </TableCell>

                <TableCell className="text-center">
                  {record.costPerPerson}
                </TableCell>

                <TableCell>
                  <div className="max-w-[250px] line-clamp-2 break-words">
                    {record.budgetCode}
                  </div>
                </TableCell>

                <TableCell>
                  <div className="max-w-[250px] line-clamp-2 break-words">
                    {record.position}
                  </div>
                </TableCell>

                <TableCell className="text-center">
                  <div className="max-w-[250px] line-clamp-2 break-words">
                    {record.department}
                  </div>
                </TableCell>

                <TableCell>
                  <div className="max-w-[250px] line-clamp-2 break-words">
                    {record.division}
                  </div>
                </TableCell>

                <TableCell>{record.status}</TableCell>

                <TableCell className="text-center">
                  {record.preTestScore ?? "-"}
                </TableCell>
                <TableCell className="text-center">
                  {record.postTestScore ?? "-"}
                </TableCell>

                <TableCell>Excellent</TableCell>

                <TableCell className="">
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button
                        variant="ghost"
                        size="icon"
                        onClick={() =>
                          router.push(`/my-training-records/${record.id}`)
                        }
                      >
                        <UploadIcon className="w-4 h-4" />
                      </Button>
                    </DropdownMenuTrigger>

                    {/* <DropdownMenuContent align="end">
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
                    </DropdownMenuContent> */}
                  </DropdownMenu>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
        {/* <Dialog open={open} onOpenChange={setOpen}> */}
        {/* <DialogContent> */}
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
        {/* </DialogContent> */}
        {/* </Dialog> */}
      </div>
    </>
  );
}

export default TrainingRecordTable;
