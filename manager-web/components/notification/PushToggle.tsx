"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { BellIcon, BellOffIcon } from "lucide-react";
import {
  enablePush,
  disablePush,
  getCurrentSubscription,
  isPushSupported,
} from "@/lib/push/client";

type Status = "loading" | "unsupported" | "denied" | "off" | "on";

export default function PushToggle() {
  const [status, setStatus] = useState<Status>("loading");
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let cancelled = false;
    const check = async () => {
      if (!isPushSupported()) {
        if (!cancelled) setStatus("unsupported");
        return;
      }
      if (Notification.permission === "denied") {
        if (!cancelled) setStatus("denied");
        return;
      }
      const sub = await getCurrentSubscription();
      if (!cancelled) setStatus(sub ? "on" : "off");
    };
    check();
    return () => {
      cancelled = true;
    };
  }, []);

  const handleEnable = async () => {
    setError(null);
    setBusy(true);
    const result = await enablePush();
    setBusy(false);
    if (result.ok) {
      setStatus("on");
    } else {
      setError(result.error ?? "Failed to enable push.");
      if (Notification.permission === "denied") setStatus("denied");
    }
  };

  const handleDisable = async () => {
    setError(null);
    setBusy(true);
    await disablePush();
    setBusy(false);
    setStatus("off");
  };

  if (status === "loading") return null;

  if (status === "unsupported") {
    return (
      <div className="text-xs text-gray-500 italic">
        Push notifications are not supported in this browser.
      </div>
    );
  }

  if (status === "denied") {
    return (
      <div className="text-xs text-red-600">
        You have blocked notifications for this site. Open your browser&apos;s site
        settings to re-allow them.
      </div>
    );
  }

  return (
    <div className="flex items-center gap-2">
      {status === "on" ? (
        <Button
          variant="outline"
          size="sm"
          onClick={handleDisable}
          disabled={busy}
          className="border-[#006022] text-[#006022] hover:bg-[#E8F7EC]"
        >
          <BellOffIcon className="h-4 w-4 mr-1" />
          {busy ? "Disabling..." : "Disable push"}
        </Button>
      ) : (
        <Button
          size="sm"
          onClick={handleEnable}
          disabled={busy}
          className="bg-[#006022] text-white hover:bg-[#005018]"
        >
          <BellIcon className="h-4 w-4 mr-1" />
          {busy ? "Enabling..." : "Enable push notifications"}
        </Button>
      )}
      {error && <span className="text-xs text-red-600">{error}</span>}
    </div>
  );
}
