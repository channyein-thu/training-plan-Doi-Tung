import React from "react";
import { Card, CardContent } from "../ui/card";

function UserInfo({
  name,
  position,
  employeeID,
  department,
  division,
}: {
  name: string;
  position: string;
  employeeID: string;
  department?: string;
  division?: string;
}) {
  return (
    <div>
      <Card className="bg-[#E8F7EC] border-none">
        <CardContent className="space-y-1 p-2">
          <p className="font-medium">Name - {name}</p>
          <p className="text-sm text-muted-foreground">Position - {position}</p>
          <p className="text-sm text-muted-foreground">
            Employee ID - {employeeID}
          </p>
          {department ? (
            <p className="text-sm text-muted-foreground">
              Department - {department}
            </p>
          ) : null}
          {division ? (
            <p className="text-sm text-muted-foreground">
              Division - {division}
            </p>
          ) : null}
        </CardContent>
      </Card>
    </div>
  );
}

export default UserInfo;
