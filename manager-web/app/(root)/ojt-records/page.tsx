import PageHeader from "@/components/dashboard/pageHeader";
import OjtTable from "@/components/ojt-records/ojtTable";
import React from "react";

function OjtRecordsPage() {
  return (
    <div className="min-h-screen space-y-4 m-2">
      <PageHeader
        title="On-the-Job Training Records"
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
