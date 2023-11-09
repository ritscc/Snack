import React from "react";
import Loading from "../Loading";
import { Skeleton } from "../ui/skeleton";
import Image from "next/image";
import { ThumbsUp } from "lucide-react";

const Chat = () => {
  return (
    <div className="hide-scrollbar h-screen w-full overflow-y-auto   bg-chat">
      {/* <Loading /> */}
      <div className="ml-3 space-y-4">
        <div className="flex items-start space-x-4">
          <Image
            src="/images/rcc-favicon.png"
            alt="RCCのfaviconです"
            width={40}
            height={40}
            className="h-12 w-12 rounded-full "
          />

          <div className="space-y-2">
            {/* ここにユーザー名を表示 */}
            <div className="flex">
              <p className="h-4 w-[100px]">RCC太郎</p>
              <p className="text-xs">2023/10/29</p>
            </div>

            <div className="flex space-x-1">
              {/* ここにメッセージを表示 */}
              本日の例会は18:00からです.場所はF106です
            </div>
            <div className="flex space-x-1">
              {/* ここにえもじを表表示 */}
              <ThumbsUp className="text-yellow-500" />
              {/* 言い値数を表示する */}
              <span>4</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Chat;
