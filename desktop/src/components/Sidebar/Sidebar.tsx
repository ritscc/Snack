import React from "react";
import SidebarHeader from "./SidebarHeader";
import SidebarChannels from "./SidebarChannels";

const Sidebar = () => {
  return (
    <div className="relative w-4/12 overflow-hidden bg-sidebar  text-sidebar-text">
      <SidebarHeader />

      <SidebarChannels />
    </div>
  );
};

export default Sidebar;
