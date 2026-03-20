"use client";
import { Button } from "../ui/button";
import { Plus } from "lucide-react";
import { useRouter } from "next/navigation";

function RouterButton({
  routerPath,
  buttonText,
}: {
  routerPath: string;
  buttonText: string;
}) {
  const router = useRouter();
  console.log("RouterButton Props:", { routerPath, buttonText }); // Log props for debugging
  return (
    <div>
      <Button
        onClick={() => {
          console.log("Navigating to:", routerPath); // Log navigation path for debugging }
          router.push(routerPath);
        }}
        className="bg-[#006022] border border-[#006022] hover:bg-[#004d17] text-white px-4 py-2 rounded-md"
      >
        <Plus className="mr-2 h-4 w-4" />
        {buttonText}
      </Button>
    </div>
  );
}

export default RouterButton;
