// import { TrainingPlanStaff } from "@/types/staff";
// import { Checkbox } from "../ui/checkbox";
// type StaffSelectProps = {
//   staff: TrainingPlanStaff[];
// };

// function StaffList({ staff }: StaffSelectProps) {
//   return (
//     <>
//       <div className="grid grid-cols-2 gap-4">
//         {staff.map((s) => (
//           <div
//             key={s.id}
//             className="border rounded-md p-4 space-y-2 flex items-center gap-3"
//           >
//             <Checkbox />
//             <div>
//               <p className="font-medium">{s.name}</p>

//               <p className="text-sm text-gray-600">{s.position}</p>
//             </div>
//           </div>
//         ))}
//       </div>
//     </>
//   );
// }

// export default StaffList;

"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { TrainingPlanStaff } from "@/types/staff";
import { Checkbox } from "../ui/checkbox";
import { Button } from "../ui/button";
import { UserPlus } from "lucide-react";
import { registerStaffAction } from "@/lib/actions/register-staff/registerStaffAction";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "../ui/dialog";
type StaffSelectProps = {
  staff: TrainingPlanStaff[];
  planId?: string;
};

function StaffList({ staff, planId }: StaffSelectProps) {
  const router = useRouter();
  const [selectedStaff, setSelectedStaff] = useState<number[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isSuccessOpen, setIsSuccessOpen] = useState(false);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const toggleStaff = (id: number) => {
    setSelectedStaff((prev) =>
      prev.includes(id)
        ? prev.filter((staffId) => staffId !== id)
        : [...prev, id],
    );
  };

  const handleSubmit = async () => {
    if (!planId) {
      setErrorMessage("Plan ID is required.");
      return;
    }
    if (selectedStaff.length === 0) {
      setErrorMessage("Please select at least one staff member.");
      return;
    }
    try {
      setIsSubmitting(true);
      await registerStaffAction(planId, selectedStaff);
      setSelectedStaff([]);
      setIsSuccessOpen(true);
    } catch (error) {
      console.error(error);
      const message =
        error instanceof Error ? error.message : "Registration failed.";
      setErrorMessage(message);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <>
      <div className="grid grid-cols-2 gap-4">
        {staff.map((s) => (
          <div
            key={s.id}
            className="border rounded-md p-4 space-y-2 flex items-center gap-3"
          >
            <Checkbox
              checked={selectedStaff.includes(Number(s.id))}
              onCheckedChange={() => toggleStaff(Number(s.id))}
            />

            <div>
              <p className="font-medium">{s.name}</p>
              <p className="text-sm text-gray-600">{s.position}</p>
            </div>
          </div>
        ))}
        {/* Action Buttons */}
      </div>
      <div className="border-gray-200 gap-4 grid grid-cols-2 mt-4">
        <Button
          variant="outline"
          className="w-full"
          onClick={() => setSelectedStaff([])}
        >
          Cancel
        </Button>
        <Button
          className="w-full bg-[#006022] hover:bg-[#004d1b] text-white"
          onClick={handleSubmit}
          disabled={isSubmitting}
        >
          <UserPlus className="w-4 h-4 mr-2" />
          {isSubmitting ? "Registering..." : "Register Staff"}
        </Button>
      </div>

      <Dialog
        open={isSuccessOpen}
        onOpenChange={(open) => {
          setIsSuccessOpen(open);
          if (!open && planId) {
            router.push(`/training-plans/${planId}`);
          }
        }}
      >
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Staff registered</DialogTitle>
            <DialogDescription>
              The selected staff members have been added to this training plan.
            </DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <Button
              onClick={() => {
                setIsSuccessOpen(false);
              }}
            >
              Close
            </Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>

      <Dialog
        open={Boolean(errorMessage)}
        onOpenChange={() => setErrorMessage(null)}
      >
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Registration failed</DialogTitle>
            <DialogDescription>{errorMessage}</DialogDescription>
          </DialogHeader>
          <DialogFooter>
            <DialogClose asChild>
              <Button className="bg-[#006022]">OK</Button>
            </DialogClose>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  );
}

export default StaffList;
