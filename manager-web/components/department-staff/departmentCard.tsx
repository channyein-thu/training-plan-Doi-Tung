import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Edit2Icon, Trash2Icon, UsersIcon } from "lucide-react";
import { Button } from "../ui/button";
import { Separator } from "@/components/ui/separator";

import { TrainingPlanStaff } from "../../types/staff";

type DepartmentCardProps = {
  staff: TrainingPlanStaff;
  onEdit: (staff: TrainingPlanStaff) => void;
  onDelete: (staff: TrainingPlanStaff) => void;
};

function DepartmentCard({ staff, onEdit, onDelete }: DepartmentCardProps) {
  return (
    <Card className="w-full sm:max-w-[360px]">
      <CardHeader className="flex flex-row items-center justify-between">
        <CardTitle className="bg-[#E8F7EC] w-16 h-16 flex items-center justify-center text-white rounded-md">
          <UsersIcon className="w-6 h-6 text-[#006022]" />
        </CardTitle>

        <CardAction className="text-[#006022] font-medium">
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

      <CardContent>
        <p className="text-black text-md">{staff.departmentName}</p>
        <p className="text-gray-600 text-sm ">
          {""}Division: {staff.departmentName}
        </p>
        <p className="text-gray-600 text-sm "> Manager: {staff.fullName}</p>
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
