import React from "react";

function NotificationCard({
  icon,
  title,
  subtitle,
}: {
  icon: React.ReactNode;
  title: string;
  subtitle: string;
}) {
  return (
    <div className="border rounded-md p-4 flex items-center space-x-4 mt-4">
      {icon}
      <div>
        <p className="font-medium">{title}</p>
        <p className="text-sm text-gray-600">{subtitle}</p>
      </div>
    </div>
  );
}

export default NotificationCard;
