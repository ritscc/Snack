import Image from "next/image";
import React, { useState } from "react";

interface AddChannelProps {
  channelId: number;
  channelType: string;
  channelPosition: number;
  channelName: string;
  //プライベートorパブリック
  channelPrivate: boolean;
  channelUserId: number;
}

const AddChannel = () => {
  const [isVisible, setIsVisible] = useState(true);

  if (!isVisible) {
    return null;
  }

  return (
    <div className="bg-orange-50 p-5 text-black ">
      <div className="flex justify-between">
        <h1 className="mb-5  text-xl font-extrabold">Create a channel</h1>
        <p className="cursor-pointer" onClick={() => setIsVisible(false)}>
          X
        </p>
      </div>
      <div>
        <h4 className="text-base font-medium">チャンネル名</h4>
        <input
          type="text"
          className="my-4 block w-full rounded-md border-0 px-3.5 py-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:outline-none focus:ring-4 focus:ring-inset focus:ring-blue-400 sm:text-sm sm:leading-6"
        />
        <div className="mb-8 mt-3 text-gray-400">
          チャンネルとは、特定のトピックに関する会話が行われる場所です。
          <br />
          見つけやすく、わかりやすい名前を使用してください。
        </div>
      </div>
      <div className="flex justify-end">
        <button className="rounded  bg-green-600 px-4 py-2 text-center text-base font-medium text-white transition duration-500 ease-in-out hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 lg:px-7">
          作成
        </button>
      </div>
    </div>
  );
};

export default AddChannel;
