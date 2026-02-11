import {
  Card,
  CardAction,
  CardContent,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Plus, Router } from "lucide-react";
// import ViewPlanForm from "./viewPlanDetails";
import InfoDetail from "./infoDetail";
import { Course } from "@/types/course";

import RouterButton from "./routerButton";

type RegisterCardProps = {
  plan: Course;
};
function RegisterCard({ plan }: RegisterCardProps) {
  return (
    <div>
      <Card className="w-full min-w-[360px]">
        <CardHeader className="flex flex-row items-center justify-between">
          <CardTitle className="">
            <p className="text-black text-md mb-2">{plan.name}</p>
            <p className="text-gray-500 text-sm">
              Speaker:{plan.speakerInstitute}
            </p>
          </CardTitle>

          <CardAction className="text-[#006022] font-medium">
            {/* <ButtonDialog
              icon={<Plus className="mr-2 h-4 w-4" />}
              name={"Register Staff"}
              className="bg-[#006022] border border-[#006022]  text-white px-4 py-2 rounded-md"
            >
              <RegisterStaffForm plan={plan} />
            </ButtonDialog> */}
            <RouterButton
              routerPath={`/register-staff/${plan.id}`}
              buttonText="Register Staff"
            />
          </CardAction>
        </CardHeader>

        <CardContent className="grid grid-cols-4 gap-4 mt-2">
          <InfoDetail title="Date" info={plan.date} />
          <InfoDetail title="Type" info={plan.type} />
          <InfoDetail title="Category" info={plan.category} />
          <InfoDetail
            title="Registered People"
            info={plan.numberOfPerson.toString()}
          />
        </CardContent>
      </Card>
    </div>
  );
}

export default RegisterCard;
