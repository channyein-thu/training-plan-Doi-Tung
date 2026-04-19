"use client";

import { useState } from "react";
import { Search } from "lucide-react";

function SearchBar() {
  const [searchTerm, setSearchTerm] = useState("");

  return (
    <div className="border border-[#006022] rounded-lg p-2 flex items-center gap-2 mt-6">
      <Search className="text-gray-400" />
      <input
        type="text"
        placeholder="Search..."
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
        className="w-full outline-none border-none focus:ring-0"
      />
    </div>
  );
}

export default SearchBar;
