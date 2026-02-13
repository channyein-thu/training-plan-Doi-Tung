import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";
import { Award, DownloadIcon, EyeIcon } from "lucide-react";
import { Separator } from "@/components/ui/separator";
import { Button } from "../ui/button";
import { Course } from "@/types/course";
import { Certificate } from "@/types/certificate";

function CertificateCard({ certificate }: { certificate: Certificate }) {
  return (
    <Card className="w-[350px] mt-4">
      {/* Header Icon */}
      <CardHeader className="">
        <div className="border rounded-md bg-[#E8F7EC] w-12 h-12 flex items-center justify-center">
          <Award className="size-5 text-[#006022]" />
        </div>
      </CardHeader>

      {/* Content */}
      <CardContent className="space-y-2">
        <h3 className="text-md font-medium">{certificate.trainingName}</h3>

        <p className="text-sm text-gray-600">{certificate.userName}</p>
        <p className="text-sm text-gray-600">
          Department - {certificate.department}
        </p>

        <div className="border rounded-md px-2 py-1 w-fit bg-[#E8F7EC] mt-1">
          <p className="text-sm text-[#006022]">
            Division: {certificate.division}
          </p>
        </div>

        <p className="text-sm text-gray-600">
          Category - {certificate.category}
        </p>

        <Separator className="" />
      </CardContent>

      {/* Footer Buttons */}
      <CardFooter className="flex gap-2">
        <Button
          variant="outline"
          className="flex-1 border-[#006022] text-[#006022]"
        >
          <DownloadIcon className="mr-2 h-4 w-4" />
          Download
        </Button>
      </CardFooter>
    </Card>
  );
}

export default CertificateCard;
