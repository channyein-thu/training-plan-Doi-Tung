import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Search } from "lucide-react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import PageHeader from "@/components/dashboard/pageHeader";
import { getRecords } from "@/lib/actions/training-records/getRecords";
import TrainingRecordTable from "@/components/my-training-records/trainingRecordTable";

export async function MyTrainingRecords() {
  const records = await getRecords();
  return (
    <div className="min-h-screen space-y-4 m-2">
      <PageHeader
        title="My Training Records"
        subtitle="View your training history and progress"
      />

      <div className="flex items-center gap-4 my-6 justify-between">
        {/* Search */}
        {/* <div className="relative w-[70%]"> */}
        <div className="relative w-full">
          <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
          <Input
            placeholder="Search by name, email, or employee ID..."
            className="pl-9 border-[#006022]"
          />
        </div>

        {/* Filter */}
        {/* <Select>
          <SelectTrigger className="w-[180px] border-[#006022]">
            <SelectValue placeholder="Suspended" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="active">Active</SelectItem>
            <SelectItem value="inactive">Inactive</SelectItem>
            <SelectItem value="suspended">Suspended</SelectItem>
          </SelectContent>
        </Select> */}

        {/* Button */}
        {/* <Button className="bg-[#006022] hover:bg-[#005018] px-8">Search</Button> */}
      </div>
      <TrainingRecordTable records={records.data.items} />
    </div>
  );
}

export default MyTrainingRecords;
