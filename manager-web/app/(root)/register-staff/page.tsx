import PageHeader from "@/components/dashboard/pageHeader";
import { Input } from "@/components/ui/input";
import { Search } from "lucide-react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "@/components/ui/button";
import { courses } from "@/data/data";
import RegisterCard from "@/components/register-staff/registerCard";

function RegisterStaffPage() {
  return (
    <div>
      <div className="h-screen overflow-y-auto p-4 m-2 space-y-4">
        <PageHeader
          title="Register Staff to Training Plans"
          subtitle="Register department staff to organization training programs"
        />

        <div className="flex items-center gap-4 my-6 justify-between">
          {/* Search */}
          <div className="relative w-[70%]">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 w-4 h-4" />
            <Input
              placeholder="Search by name, email, or employee ID..."
              className="pl-9 border-[#006022]"
            />
          </div>

          {/* Filter */}
          <Select>
            <SelectTrigger className="w-[180px] border-[#006022]">
              <SelectValue placeholder="Suspended" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="active">Active</SelectItem>
              <SelectItem value="inactive">Inactive</SelectItem>
              <SelectItem value="suspended">Suspended</SelectItem>
            </SelectContent>
          </Select>

          {/* Button */}
          <Button className="bg-[#006022] hover:bg-[#005018] px-8">
            Search
          </Button>
        </div>

        <div className="space-y-6">
          {courses.map((plan) => (
            <RegisterCard key={plan.id} plan={plan} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default RegisterStaffPage;
