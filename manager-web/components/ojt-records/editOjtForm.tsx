"use client";

import { useForm } from "react-hook-form";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Input } from "../ui/input";
import { Button } from "../ui/button";

type TestForm = {
  status: string;
  feedback: string;
  evaluation: string;
  prePercent: number;
  postPercent: number;
};

function EditOJTdetails() {
  const form = useForm<TestForm>({
    defaultValues: {
      status: "",
      feedback: "",
      evaluation: "",
      prePercent: 0,
      postPercent: 0,
    },
  });

  const onSubmit = (data: TestForm) => {
    console.log("Form Data:", data);
  };

  return (
    <form
      onSubmit={form.handleSubmit(onSubmit)}
      className="border rounded-md m-2 p-4 space-y-6"
    >
      <p className="font-medium">Please start fill from here</p>

      {/* Status */}
      <div className="space-y-2">
        <label className="font-medium">User Status</label>
        <Select onValueChange={(value) => form.setValue("status", value)}>
          <SelectTrigger className="w-[220px]">
            <SelectValue placeholder="Select status" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="register">Register</SelectItem>
            <SelectItem value="attended">Attended</SelectItem>
            <SelectItem value="absent">Absent</SelectItem>
          </SelectContent>
        </Select>
      </div>

      {/* Feedback */}
      <div className="space-y-2">
        <label className="font-medium">Feedback</label>
        <textarea
          className="w-full border rounded-md p-2"
          rows={4}
          {...form.register("feedback")}
          placeholder="Enter feedback here..."
        />
      </div>

      {/* Evaluation */}
      <div className="space-y-2">
        <label className="font-medium">
          Evaluate the Practical Application of Knowledge to Work
        </label>
        <textarea
          className="w-full border rounded-md p-2"
          rows={4}
          {...form.register("evaluation")}
          placeholder="Describe how the training knowledge is applied..."
        />
      </div>

      {/* Test Scores */}
      <div className="grid grid-cols-2 gap-8">
        {/* Pretest */}
        <div className="space-y-2">
          <label className="text-lg font-medium">Pretest Score</label>

          <div className="relative">
            <Input
              type="number"
              min={0}
              max={100}
              {...form.register("prePercent", {
                valueAsNumber: true,
              })}
              className="h-14 text-xl rounded-xl pr-10"
            />
            <span className="absolute right-4 top-1/2 -translate-y-1/2 text-lg text-muted-foreground">
              %
            </span>
          </div>
        </div>

        {/* Posttest */}
        <div className="space-y-2">
          <label className="text-lg font-medium">Posttest Score</label>

          <div className="relative">
            <Input
              type="number"
              min={0}
              max={100}
              {...form.register("postPercent", {
                valueAsNumber: true,
              })}
              className="h-14 text-xl rounded-xl pr-10"
            />
            <span className="absolute right-4 top-1/2 -translate-y-1/2 text-lg text-muted-foreground">
              %
            </span>
          </div>
        </div>
      </div>

      {/* Action Buttons */}
      <div className="px-6 py-4  border-gray-200 flex gap-4">
        <Button variant="outline" className="flex-1">
          Cancel
        </Button>
        <Button className="flex-1 bg-[#006022] hover:bg-[#004d1b] text-white">
          Save Changes
        </Button>
      </div>
    </form>
  );
}

export default EditOJTdetails;
