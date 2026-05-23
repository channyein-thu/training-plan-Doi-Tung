import type { Metadata } from "next";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/ui/app-sidebar";
import { Separator } from "@/components/ui/separator";
import { getMe } from "@/lib/api/getMe";
import { redirect } from "next/navigation";
export const metadata: Metadata = {
  title: "Manager Dashboard",
  description: "Admin dashboard for training management",
};

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const me = await getMe();
  const user = me?.user;
  console.log("DashboardLayout - User Info:", user); // Debug log to check user info

  if (!user) {
    redirect("/login");
  }
  return (
    <SidebarProvider>
      <div className="flex min-h-screen w-full">
        {/* Sidebar */}
        <AppSidebar
          user={{
            name: user.name,
            position: user.position,
            employeeID: user.employeeID,
            role: user.role,
          }}
        />

        {/* Main content */}
        <div className="flex flex-1 flex-col min-w-0">
          {/* Mobile header */}
          <header className="flex items-center gap-2 border-b px-4 py-3 lg:hidden">
            <SidebarTrigger className="-ml-1" />
            <Separator orientation="vertical" className="mr-2 h-4" />
            <span className="text-sm font-semibold">Manager Dashboard</span>
          </header>

          {/* Page content */}
          <main className="flex-1 overflow-y-auto p-6">{children}</main>
        </div>
      </div>
    </SidebarProvider>
  );
}
