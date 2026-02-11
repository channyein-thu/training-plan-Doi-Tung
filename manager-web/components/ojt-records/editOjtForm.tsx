import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";

function EditOJTdetails() {
  return (
    <div>
      <div className="border rounded-md m-2  p-4 space-y-4">
        <p className="font-medium mb-2">Please start fill from here</p>
        <div>
          {/* <FormField
            // control={form.control}
            name="status"
            render={({ field }) => (
              <FormItem> */}
          {/* <FormLabel>User Status</FormLabel> */}

          {/* <Select value={field.value} onValueChange={field.onChange}> */}
          <Select>
            {/* <FormControl> */}
            <SelectTrigger className="w-[220px]">
              <SelectValue placeholder="Select status" />
            </SelectTrigger>
            {/* </FormControl> */}

            <SelectContent>
              <SelectItem value="register">Register</SelectItem>
              <SelectItem value="attended">Attended</SelectItem>
              <SelectItem value="absent">Absent</SelectItem>
            </SelectContent>
          </Select>

          {/* <FormMessage /> */}
          {/* </FormItem> */}
        </div>

        {/* Feedback Section */}
        <div className="mt-4">
          <label className="block mb-2 font-medium">Feedback</label>
          <textarea
            className="w-full border border-gray-300 rounded-md p-2"
            rows={4}
            placeholder="Enter feedback here..."
          ></textarea>
        </div>

        {/* Evaluation */}
        <div className="mt-4">
          <label className="block mb-2 font-medium">
            Evaluate the Practical Application of Knowledge to Work
          </label>
          <textarea
            className="w-full border border-gray-300 rounded-md p-2"
            rows={4}
            placeholder="Describe how the training knowledge is applied to work..."
          ></textarea>
        </div>
      </div>
    </div>
  );
}

export default EditOJTdetails;
