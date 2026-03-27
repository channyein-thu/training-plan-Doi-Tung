// import type { Metadata } from "next";
// import { Geist, Geist_Mono } from "next/font/google";
// import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
// import { AppSidebar } from "@/components/ui/app-sidebar";
// import { Separator } from "@/components/ui/separator";

// const geistSans = Geist({
//   variable: "--font-geist-sans",
//   subsets: ["latin"],
// });

// const geistMono = Geist_Mono({
//   variable: "--font-geist-mono",
//   subsets: ["latin"],
// });

// export const metadata: Metadata = {
//   title: "Manager Dashboard",
//   description: "Admin dashboard for training management",
// };

// export default function RootLayout({
//   children,
// }: Readonly<{
//   children: React.ReactNode;
// }>) {
//   return (
//     <html lang="en">
//       <body
//         className={`${geistSans.variable} ${geistMono.variable} antialiased`}
//       >
//         <SidebarProvider>
//           <div className="flex min-h-screen w-full">
//             {/* Sidebar — fixed on lg+, drawer on mobile */}
//             <AppSidebar />

//             {/* Main content area */}
//             <div className="flex flex-1 flex-col min-w-0">
//               {/* Mobile header with hamburger toggle */}
//               <header className="flex items-center gap-2 border-b px-4 py-3 lg:hidden">
//                 <SidebarTrigger className="-ml-1" />
//                 <Separator orientation="vertical" className="mr-2 h-4" />
//                 <span className="text-sm font-semibold">Manager Dashboard</span>
//               </header>

//               {/* Page content */}
//               <main className="flex-1 overflow-y-auto p-6">{children}</main>
//             </div>
//           </div>
//         </SidebarProvider>
//       </body>
//     </html>
//   );
// }
import type { Metadata } from "next";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/ui/app-sidebar";
import { Separator } from "@/components/ui/separator";

export const metadata: Metadata = {
  title: "Manager Dashboard",
  description: "Admin dashboard for training management",
};

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <SidebarProvider>
      <div className="flex min-h-screen w-full">
        {/* Sidebar */}
        <AppSidebar />

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
