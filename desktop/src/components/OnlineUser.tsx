import Image from "next/image";
import React from "react";

interface UserProps {
  name: string;
  imageSrc: string;
}

const OnlineUser = ({ name, imageSrc }: UserProps) => {
  return (
    <div className="flex items-center ">
      <Image src={imageSrc} alt={name} width={40} height={40} className="m-2" />
      <p className="text-online-text">{name}</p>
    </div>
  );
};

export default OnlineUser;
