"use client";

import {
  AwardIcon,
  BellIcon,
  BookOpen,
  Calendar,
  DownloadIcon,
  Home,
  Inbox,
  Notebook,
  NotebookIcon,
  Search,
  User,
  UsersIcon,
} from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import Link from "next/link";
import LogoCard from "../dashboard/logoCard";
import UserInfo from "../dashboard/userInfo";
import { usePathname } from "next/navigation";

// Menu items.
const items = [
  {
    title: "Dashboard",
    url: "/dashboard",
    icon: Home,
  },
  {
    title: "Department-Staff",
    url: "/department-staff",
    icon: UsersIcon,
  },
  {
    title: "Register Staff",
    url: "/register-staff",
    icon: User,
  },
  {
    title: "OJT Records",
    url: "/ojt-records",
    icon: NotebookIcon,
  },
  {
    title: "My Certificates",
    url: "/my-certificates",
    icon: AwardIcon,
  },
  {
    title: "Upload Certificate",
    url: "/upload-certificate",
    icon: DownloadIcon,
  },
  {
    title: "My Training Records",
    url: "/my-training-records",
    icon: Notebook,
  },
  {
    title: "Notifications",
    url: "/notifications",
    icon: BellIcon,
  },
];

export function AppSidebar() {
  const pathname = usePathname();

  return (
    <>
      <Sidebar collapsible="none">
        <SidebarContent>
          <SidebarGroup>
            <SidebarGroupLabel>Application</SidebarGroupLabel>
            <SidebarGroupContent>
              <div className="m-4 space-y-4">
                <LogoCard />
                <UserInfo />
              </div>
              <h1 className="m-2 text-sm font-semibold">Manager Functions</h1>
              <SidebarMenu>
                {items.map((item) => {
                  const isActive = pathname === item.url;

                  return (
                    <SidebarMenuItem key={item.title}>
                      <SidebarMenuButton
                        asChild
                        isActive={isActive}
                        className="
                        data-[active=true]:bg-[#006022]
                        data-[active=true]:text-white
                        text-md px-3 py-4 rounded-md 
                      "
                      >
                        <Link
                          href={item.url}
                          className="flex items-center gap-2"
                        >
                          <item.icon className="h-4 w-4" />
                          <span className="">{item.title}</span>
                        </Link>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  );
                })}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </SidebarContent>
      </Sidebar>
    </>
  );
}
