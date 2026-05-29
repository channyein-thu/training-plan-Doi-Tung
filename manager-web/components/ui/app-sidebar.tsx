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
import { useEffect, useState } from "react";
import { getUnreadCount } from "@/lib/actions/notifications/getUnreadCount";

// Menu items.
const generalNav = [
  {
    title: "Dashboard",
    url: "/dashboard",
    icon: Home,
  },
];

const managerFunctions = [
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
  const [unreadCount, setUnreadCount] = useState(0);

  // Poll the backend for the unread notification count.
  // - Fetches once on mount
  // - Re-fetches whenever the user navigates (so leaving /notifications drops the badge)
  // - Polls every 30s for new notifications while the tab is open
  useEffect(() => {
    let cancelled = false;
    const refresh = async () => {
      const count = await getUnreadCount();
      if (!cancelled) setUnreadCount(count);
    };
    refresh();
    const id = setInterval(refresh, 30000);
    return () => {
      cancelled = true;
      clearInterval(id);
    };
  }, [pathname]);

  return (
    <Sidebar collapsible="offcanvas">
      <SidebarContent>
        {/* Logo + Profile */}
        <SidebarGroup>
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

        {/* General (visible to all roles) */}
        <SidebarGroup>
          <SidebarGroupLabel>General</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {generalNav.map((item) => {
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
                const isNotifications = item.url === "/notifications";
                const showBadge = isNotifications && unreadCount > 0;

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
                        className="flex items-center gap-2 w-full"
                      >
                        <item.icon className="h-4 w-4" />
                        <span className="flex-1">{item.title}</span>
                        {showBadge && (
                          <span className="ml-auto inline-flex items-center justify-center min-w-[20px] h-5 px-1.5 rounded-full bg-red-500 text-white text-[11px] font-semibold leading-none">
                            {unreadCount > 99 ? "99+" : unreadCount}
                          </span>
                        )}
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
