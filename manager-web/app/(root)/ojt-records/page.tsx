import PageHeader from "@/components/dashboard/pageHeader";
import OjtTable from "@/components/ojt-records/ojtTable";
import { getOjtRecords } from "@/lib/actions/ojt-records/getOjtAction";

async function OjtRecordsPage() {
  const ojtRecords = await getOjtRecords(); // Fetch OJT records on page load
  console.log("Fetched OJT Records:", ojtRecords.data.items); // Log fetched records for debugging
  return (
    <div className="min-h-screen space-y-4 m-2">
      <PageHeader
        title="On-the-Job Training Records"
        subtitle="View and manage on-the-job training records"
      />

      {/* OJT Records Table */}
      <div>
        <OjtTable ojtRecords={ojtRecords.data.items} />
      </div>
    </div>
  );
}

export default OjtRecordsPage;
