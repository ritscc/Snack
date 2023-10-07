import React from "react";

interface ChannnelProps {
  icon: string;
  name: string;
}

const Channnels = ({ icon, name }: ChannnelProps) => {
  return (
    <div className="">
      <div className="flex cursor-pointer rounded-sm p-2 duration-200 hover:bg-slate-400">
        <span className="mx-2">{icon}</span>
        {name}
      </div>
    </div>
  );
};

export default Channnels;
