import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Certificate } from "@/types/certificate";
import { Award } from "lucide-react";

function CertificateCard({ certificate }: { certificate: Certificate }) {
  // Load the image through the authenticated same-origin proxy instead of the
  // (now removed) open backend /uploads route.
  const imageUrl = `/api/certificate-file/${certificate.id}`;
  const statusClassName =
    certificate.status === "Approved"
      ? "bg-green-100 text-green-700"
      : certificate.status === "Pending"
        ? "bg-yellow-100 text-yellow-700"
        : "bg-red-100 text-red-700";

  return (
    <Card className="w-[350px] mt-4 relative">
      {/* Header Icon */}
      <CardHeader className="">
        <div className="border rounded-md bg-[#E8F7EC] w-12 h-12 flex items-center justify-center">
          <Award className="size-5 text-[#006022]" />
        </div>
        <span
          className={`absolute right-4 top-4 rounded-full px-3 py-1 text-xs font-medium ${statusClassName}`}
        >
          {certificate.status}
        </span>
      </CardHeader>

      {/* Content */}
      <CardContent className="space-y-2">
        <h3 className="text-md font-medium">{certificate.trainingName}</h3>

        <div className="w-full overflow-hidden rounded-md border bg-gray-50">
          <img
            src={imageUrl}
            alt={certificate.trainingName}
            className="h-40 w-full object-cover"
            loading="lazy"
          />
        </div>

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
      </CardContent>
    </Card>
  );
}

export default CertificateCard;
