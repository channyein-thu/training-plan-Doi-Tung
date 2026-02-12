import PageHeader from "@/components/dashboard/pageHeader";
import CertificateCard from "@/components/my-certificates/certificateCard";
import { certificates } from "@/data/data";
import { Certificate } from "crypto";
import React from "react";

function MyCertificates() {
  return (
    <>
      <div className="m-6 space-y-4">
        <PageHeader
          title="My Certificates"
          subtitle="View and download all your training certificates"
        />

        {/* Certificates List*/}
        {certificates.map((certificate) => (
          <CertificateCard key={certificate.id} certificate={certificate} />
        ))}
      </div>
    </>
  );
}

export default MyCertificates;
