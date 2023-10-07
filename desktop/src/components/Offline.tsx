import Image from "next/image";
import React from "react";
import Overlay from "./Overlay";

interface UserProps {
  name: string;
  imageSrc: string;
}

const OfflineUser = ({ name, imageSrc }: UserProps) => {
  return (
    <div className="flex items-center">
      <Image src={imageSrc} alt={name} width={40} height={40} className="m-2 " />
      <p className="text-slate-400">{name}</p>
    </div>
  );
};

export default OfflineUser;
