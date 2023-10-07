import Image from "next/image";
import React, { useState } from "react";
// interface onlineUserProps{
//   online:0;
//   offline:false
// }

const USerData = () => {
  const [isVisible, setIsVisible] = useState(true);

  if (!isVisible) {
    return null;
  }
  return (
    <div className="bg-orange-50 p-1 px-4 text-black">
      <div className="mb-2 flex space-x-2 border-b-4 border-slate-600">
        <div className="flex items-center justify-center">
          <Image src="/images/rcc-favicon.png" alt="RCCのfaviconです" width={40} height={40} className="m-2" />

          <div className="my-1">
            <p className="text-lg">RCC太郎</p>
            <p className="text-sm">オンライン</p>
          </div>
        </div>
        <div>🖋</div>
        <div className="cursor-pointer" onClick={() => setIsVisible(false)}>
          ×
        </div>
      </div>
      <div className=" border-b-2 border-slate-800">
        <div className="flex">
          <div>役職</div>
          <span className="mx-2">:</span>
          <div className="">システム管理</div>
        </div>
      </div>

      <div className=" mb-10 border-b-2 border-slate-800">
        <div className="flex">
          <div>所属日</div>
          <span className="mx-2">:</span>
          <div className="">2023年1月1日</div>
        </div>
      </div>
    </div>
  );
};

export default USerData;
