import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Edit2Icon, Trash2Icon } from "lucide-react";
import { Button } from "../ui/button";

import { TrainingPlanStaff } from "../../types/staff";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

type DepartmentCardProps = {
  staff: TrainingPlanStaff;
  onEdit: (staff: TrainingPlanStaff) => void;
  onDelete: (staff: TrainingPlanStaff) => void;
};

function DepartmentCard({ staff, onEdit, onDelete }: DepartmentCardProps) {
  return (
    <Card className="w-full sm:max-w-[360px]">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle className="flex items-center gap-3">
          <Avatar className="bg-green h-12 w-12 border-2 border-[#E8F7EC]">
            <AvatarFallback className="bg-[#4E7F57] text-white text-md">
              {staff.name.charAt(0)}
            </AvatarFallback>
          </Avatar>
          <div>
            <h3 className="text-md font-medium text-black">{staff.name}</h3>
            <p className="text-sm text-gray-600 space-y-1">
              <span className="font-semibold">ID:</span> {staff.employeeID}
            </p>
            <p className="text-sm text-gray-600">{staff.position}</p>
          </div>
        </CardTitle>

        <CardAction className="flex flex-row items-center gap-1 text-[#006022] font-medium">
          <Button variant="ghost" size="icon-sm" onClick={() => onEdit(staff)}>
            <Edit2Icon />
          </Button>
          <Button
            variant="ghost"
            size="icon-sm"
            onClick={() => onDelete(staff)}
          >
            <Trash2Icon className="text-red-500" />
          </Button>
        </CardAction>
      </CardHeader>

      <CardContent className="space-y-1">
        <p className="text-black text-md">{staff.role}</p>
        <p className="text-gray-600 text-sm ">
          {""}Department: {staff.department}
        </p>
        <p className="text-gray-600 text-sm "> Phone: {staff.phone}</p>
        {/* <Separator className="w-24 self-center mt-4" /> */}
      </CardContent>

      {/* <CardFooter className="flex flex-col items-start gap-2"> */}
      {/* <p className="text-sm text-muted-foreground">Staff</p> */}
      {/* <h4 className="">{staff.length}</h4> */}
      {/* </CardFooter> */}
    </Card>
  );
}

export default DepartmentCard;
