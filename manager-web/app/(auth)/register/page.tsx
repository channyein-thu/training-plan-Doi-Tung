import Logo from "@/components/login/logo";
import { Card, CardContent } from "@/components/ui/card";
import { Users } from "lucide-react";
import { getDepartmentList } from "@/lib/api/getDepartment";
import RegisterForm from "@/components/register/register-form";

export default async function RegisterPage() {
  const departments = await getDepartmentList();

  return (
    <div className="min-h-screen bg-[#dbe7dd] flex flex-col items-center py-10 px-4 w-[100%]">
      {/* Logo */}
      <div className="mb-6">
        <Logo />
      </div>

      <Card className="w-full max-w-4xl rounded-2xl shadow-lg">
        <CardContent className="p-8">
          <div className="flex items-center gap-4 mb-8 justify-center">
            <div className="bg-[#006022] w-14 h-14 flex items-center justify-center text-white rounded-xl">
              <Users size={28} />
            </div>
            <div>
              <h2 className="text-2xl font-semibold">Manager Registration</h2>
              <p className="text-gray-500 text-sm">
                Fill in your details below
              </p>
            </div>
          </div>

          <RegisterForm departments={departments} />
        </CardContent>
      </Card>
    </div>
  );
}
