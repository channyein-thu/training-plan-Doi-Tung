"use client";

import {
  AwardIcon,
  BellIcon,
  BookOpen,
  Calendar,
  DownloadIcon,
  Home,
  Inbox,
  LogOut,
  Notebook,
  NotebookIcon,
  Search,
  User,
  UsersIcon,
} from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
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
import { Button } from "./button";
import { logoutAction } from "@/lib/actions/logout";

// Menu items.
const managerFunctions = [
  {
    title: "Dashboard",
    url: "/dashboard",
    icon: Home,
  },
  {
    title: "Department Staff",
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
];

const myTraining = [
  {
    title: "My Certificates",
    url: "/my-certificates",
    icon: AwardIcon,
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

type AppSidebarProps = {
  user: {
    name: string;
    position: string;
    employeeID: string;
    role: string;
  };
};

export function AppSidebar({ user }: AppSidebarProps) {
  const pathname = usePathname();
  const isManager = user.role === "DepartmentHead(manager)";

  return (
    <Sidebar collapsible="offcanvas">
      <SidebarContent>
        {/* Application / Profile */}
        <SidebarGroup>
          <SidebarGroupLabel>Application</SidebarGroupLabel>
          <SidebarGroupContent>
            <div className="m-4 space-y-4">
              <LogoCard />
              <UserInfo
                name={user.name}
                position={user.position}
                employeeID={user.employeeID}
              />
            </div>
          </SidebarGroupContent>
        </SidebarGroup>

        {/* Manager Functions */}
        {isManager && (
          <SidebarGroup>
            <SidebarGroupLabel>Manager Functions</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {managerFunctions.map((item) => {
                  const isActive = pathname === item.url;

                  return (
                    <SidebarMenuItem key={item.title}>
                      <SidebarMenuButton
                        asChild
                        isActive={isActive}
                        className="
                        data-[active=true]:bg-[#006022]
                        data-[active=true]:text-white
                        px-3 py-4 rounded-md
                      "
                      >
                        <Link
                          href={item.url}
                          className="flex items-center gap-2"
                        >
                          <item.icon className="h-4 w-4" />
                          <span>{item.title}</span>
                        </Link>
                      </SidebarMenuButton>
                    </SidebarMenuItem>
                  );
                })}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        )}

        {/* My Training */}
        <SidebarGroup>
          <SidebarGroupLabel>My Training</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {myTraining.map((item) => {
                const isActive = pathname === item.url;

                return (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton
                      asChild
                      isActive={isActive}
                      className="
                        data-[active=true]:bg-[#006022]
                        data-[active=true]:text-white
                        px-3 py-4 rounded-md
                      "
                    >
                      <Link href={item.url} className="flex items-center gap-2">
                        <item.icon className="h-4 w-4" />
                        <span>{item.title}</span>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <form action={logoutAction} className="p-2">
          <Button
            type="submit"
            variant="ghost"
            className="w-full justify-start gap-2 text-red-600 hover:text-red-700"
          >
            <LogOut className="h-4 w-4" />
            <span>Log out</span>
          </Button>
        </form>
      </SidebarFooter>
    </Sidebar>
  );
}
