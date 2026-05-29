import PageHeader from "@/components/dashboard/pageHeader";
import OjtTable from "@/components/ojt-records/ojtTable";
import { getOjtRecords } from "@/lib/actions/ojt-records/getOjtAction";

async function OjtRecordsPage() {
  const ojtRecords = await getOjtRecords();
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const items = (ojtRecords?.data?.items as any[]) ?? [];

  return (
    <div className="min-h-screen space-y-4 m-2">
      <PageHeader
        title="On-the-Job Training Records"
        subtitle="View and manage on-the-job training records"
      />

      <div>
        {items.length === 0 ? (
          <p className="text-center text-gray-500 py-12">
            No on-the-job training records found.
          </p>
        ) : (
          <OjtTable ojtRecords={items} />
        )}
      </div>
    </div>
  );
}

export default OjtRecordsPage;
