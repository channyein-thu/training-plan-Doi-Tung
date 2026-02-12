"use client"

import { useForm } from "react-hook-form"
import { Input } from "@/components/ui/input"

type TestForm = {
  prePercent: number
  postPercent: number
}

export default function TestForm() {
  const form = useForm<TestForm>({
    defaultValues: {
      prePercent: 0,
      postPercent: 0,
    },
  })

  const pre = form.watch("prePercent")
  const post = form.watch("postPercent")

  return (
    <form className="grid grid-cols-2 gap-6">
      
      {/* Pretest Card */}
      <div className="border rounded-xl p-6 text-center space-y-4">
        <p className="text-sm text-muted-foreground">Pretest</p>

        <p className="text-4xl font-bold">
          {pre || 0}%
        </p>

        <Input
          type="number"
          placeholder="Enter %"
          {...form.register("prePercent", {
            valueAsNumber: true,
            min: 0,
            max: 100,
          })}
        />
      </div>

      {/* Posttest Card */}
      <div className="border rounded-xl p-6 text-center space-y-4">
        <p className="text-sm text-muted-foreground">Posttest</p>

        <p className="text-4xl font-bold">
          {post || 0}%
        </p>

        <Input
          type="number"
          placeholder="Enter %"
          {...form.register("postPercent", {
            valueAsNumber: true,
            min: 0,
            max: 100,
          })}
        />
      </div>

    </form>
  )
}
