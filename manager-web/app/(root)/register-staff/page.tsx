import PageHeader from "@/components/dashboard/pageHeader";
import RegisterCard from "@/components/register-staff/registerCard";
import { getPlans } from "@/lib/actions/register-staff/getPlanAction";
import { Course } from "@/types/course";

async function RegisterStaffPage() {
  const plans = await getPlans();
  const items = (plans?.data?.items as Course[] | undefined) ?? [];

  return (
    <div>
      <div className="min-h-screen space-y-4 m-2">
        <PageHeader
          title="Register Staff to Training Plans"
          subtitle="Register department staff to organization training programs"
        />

        {items.length === 0 ? (
          <p className="text-center text-gray-500 py-12">
            No training plans available.
          </p>
        ) : (
          <div className="space-y-6">
            {items.map((plan) => (
              <RegisterCard key={plan.id} plan={plan} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

export default RegisterStaffPage;
