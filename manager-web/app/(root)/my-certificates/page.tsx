import PageHeader from "@/components/dashboard/pageHeader";
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
      </div>
    </>
  );
}

export default MyCertificates;
