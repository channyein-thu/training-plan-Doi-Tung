import PageHeader from "@/components/dashboard/pageHeader";
import OjtTable from "@/components/ojt-records/ojtTable";
import React from "react";

function OjtRecordsPage() {
  return (
    <div className="h-screen overflow-y-auto p-4 m-2 space-y-4">
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
