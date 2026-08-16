import * as React from "react";

import { cn } from "@/lib/utils";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { EyeClosedIcon } from "lucide-react";
import { EyeIcon } from "lucide-react";
const PasswordInput = React.forwardRef<
  HTMLInputElement,
  React.ComponentProps<"input">
>(({ className, type, ...props }, ref) => {
  const [showPass, setshowPass] = React.useState(false);
  return (
    <div className="relative">
      <Input
        type={showPass ? "text" : "password"}
        className={cn(className)}
        ref={ref}
        {...props}
      />
      <Button
        type="button"
        variant="ghost"
        className="absolute top-0 right-0 hover:bg-transparent"
        onClick={() => setshowPass((prev) => !prev)}
        disabled={props.value === "" || props.disabled}
      >
        {showPass ? (
          <EyeClosedIcon className="w-4 h-4" />
        ) : (
          <EyeIcon className="w-4 h-4" />
        )}
      </Button>
    </div>
  );
});
PasswordInput.displayName = "PasswordInput";

export { PasswordInput };