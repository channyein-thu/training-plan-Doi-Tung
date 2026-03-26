import PageHeader from "@/components/dashboard/pageHeader";
import CertificateCard from "@/components/my-certificates/certificateCard";
import { getCertificates } from "@/lib/actions/my-certificates/getMyCertificate";

async function MyCertificates() {
  const certificates = await getCertificates(); // Fetch certificates on page load
  return (
    <>
      <div className="min-h-screen space-y-4 m-2">
        <PageHeader
          title="My Certificates"
          subtitle="View and download all your training certificates"
        />

        {/* Certificates List*/}
        <div className="grid lg:grid-cols-3 md:grid-cols-2 mt-2">
          {certificates.map((certificate) => (
            <CertificateCard key={certificate.id} certificate={certificate} />
          ))}
        </div>
      </div>
    </>
  );
}

export default MyCertificates;
