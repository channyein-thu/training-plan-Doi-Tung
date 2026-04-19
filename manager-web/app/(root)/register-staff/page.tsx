import PageHeader from "@/components/dashboard/pageHeader";
import RegisterCard from "@/components/register-staff/registerCard";
import { getPlans } from "@/lib/actions/register-staff/getPlanAction";
import { Course } from "@/types/course";

async function RegisterStaffPage() {
  const plans = await getPlans(); // Fetch training plans on page load

  return (
    <div>
      <div className="min-h-screen space-y-4 m-2">
        <PageHeader
          title="Register Staff to Training Plans"
          subtitle="Register department staff to organization training programs"
        />

        <div className="space-y-6">
          {plans.data.items.map((plan: Course) => (
            // Log each plan being rendered for debugging
            <RegisterCard key={plan.id} plan={plan} />
          ))}
        </div>
      </div>
    </div>
  );
}

export default RegisterStaffPage;
