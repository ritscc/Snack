import React from "react";
import Image from "next/image";

interface PrivateChatProps {
  imgUrl: string;
  name: string;
}

const PrivateChat = ({ imgUrl, name }: PrivateChatProps) => {
  return (
    <div className="">
      <div className="flex items-center ">
        <Image src={imgUrl} alt={name} width={30} height={30} className="m-2 " />
        {name}
      </div>
    </div>
  );
};

export default PrivateChat;
