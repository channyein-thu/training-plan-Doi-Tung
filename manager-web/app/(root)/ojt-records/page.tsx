import PageHeader from "@/components/dashboard/pageHeader";
import OjtTable from "@/components/ojt-records/ojtTable";
import React from "react";

function OjtRecordsPage() {
  return (
    <div className="m-6 space-y-4">
      <PageHeader
        title="OJT Records"
        subtitle="View and manage on-the-job training records"
      />

      {/* OJT Records Table */}
      <div>
        <OjtTable />
      </div>
    </div>
  );
}

export default OjtRecordsPage;
