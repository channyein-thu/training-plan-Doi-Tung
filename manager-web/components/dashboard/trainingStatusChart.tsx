"use client";

import {
  PieChart,
  Pie,
  Cell,
  ResponsiveContainer,
  Tooltip,
} from "recharts";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

type Props = {
  registered: number;
  attended: number;
  absent: number;
};

const COLORS = {
  Attended: "#16a34a", // green
  Register: "#f59e0b", // amber
  Absent: "#dc2626", // red
};

export default function TrainingStatusChart({
  registered,
  attended,
  absent,
}: Props) {
  const total = registered + attended + absent;

  const data = [
    { name: "Attended", value: attended, color: COLORS.Attended },
    { name: "Register", value: registered, color: COLORS.Register },
    { name: "Absent", value: absent, color: COLORS.Absent },
  ];

  // Completion = attended out of all training records.
  const completion = total > 0 ? Math.round((attended / total) * 100) : 0;

  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="text-lg">Training Performance</CardTitle>
      </CardHeader>
      <CardContent>
        {total === 0 ? (
          <p className="py-12 text-center text-sm text-gray-500">
            No training records yet.
          </p>
        ) : (
          <div className="flex flex-col items-center gap-4 md:flex-row md:justify-around">
            {/* Donut with center completion % */}
            <div className="relative h-[220px] w-[220px]">
              <ResponsiveContainer width="100%" height="100%">
                <PieChart>
                  <Pie
                    data={data}
                    dataKey="value"
                    nameKey="name"
                    innerRadius={70}
                    outerRadius={100}
                    paddingAngle={2}
                    stroke="none"
                  >
                    {data.map((d) => (
                      <Cell key={d.name} fill={d.color} />
                    ))}
                  </Pie>
                  <Tooltip />
                </PieChart>
              </ResponsiveContainer>
              <div className="pointer-events-none absolute inset-0 flex flex-col items-center justify-center">
                <span className="text-3xl font-bold text-[#006022]">
                  {completion}%
                </span>
                <span className="text-xs text-gray-500">Completed</span>
              </div>
            </div>

            {/* Legend with counts */}
            <div className="space-y-2">
              {data.map((d) => (
                <div key={d.name} className="flex items-center gap-2 text-sm">
                  <span
                    className="inline-block h-3 w-3 rounded-sm"
                    style={{ backgroundColor: d.color }}
                  />
                  <span className="w-20 text-gray-700">
                    {d.name === "Register" ? "Left" : d.name}
                  </span>
                  <span className="font-semibold">{d.value}</span>
                </div>
              ))}
              <div className="mt-1 flex items-center gap-2 border-t pt-2 text-sm">
                <span className="w-[92px] text-gray-500">Total</span>
                <span className="font-semibold">{total}</span>
              </div>
            </div>
          </div>
        )}
      </CardContent>
    </Card>
  );
}
