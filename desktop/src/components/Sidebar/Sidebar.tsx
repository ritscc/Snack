import React from "react";
import SidebarHeader from "./SidebarHeader";
import SidebarChannels from "./SidebarChannels";

const Sidebar = () => {
  return (
    <div className="w-[350px] relative  overflow-hidden bg-sidebar text-slate-200">
      <SidebarHeader />

      <SidebarChannels />
    </div>
  );
};

export default Sidebar;
