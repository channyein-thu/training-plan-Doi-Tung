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
