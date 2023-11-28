import Image from "next/image";
import React from "react";
import MicIcon from "@mui/icons-material/Mic";
import HeadphonesIcon from "@mui/icons-material/Headphones";
import SettingsIcon from "@mui/icons-material/Settings";
import { useState } from "react";
import USerData from "./USerData";
import { useUser, useUsers } from "@/api/user";

const SidebarHeader = () => {
  // const { user, isLoading, isError } = useUser(1);
  const { users } = useUsers();
  const [showUserData, setShowUserData] = useState(false);
  const toggleUserData = () => {
    setShowUserData(!showUserData);
  };

  return (
    <div className=" border-b-2 border-slate-800 bg-sidebar ">
      <div className="flex items-center justify-between ">
        <div className="flex items-center justify-center">
          <Image
            src="/images/rcc-favicon.png"
            alt="RCCのfaviconです"
            width={40}
            height={40}
            className="m-2 "
            onClick={toggleUserData}
          />
          <div className="text-sm font-bold md:text-xl">{users?.[1]?.username}</div>
        </div>

        <div className="mr-3 flex space-x-2 ">
          <MicIcon />
          <HeadphonesIcon />
          <SettingsIcon />
        </div>
      </div>
      {showUserData && (
        <div className="absolute z-10 -mt-4 ml-10 ">
          <USerData />
        </div>
      )}
    </div>
  );
};

export default SidebarHeader;
