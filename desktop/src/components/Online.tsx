import Image from "next/image";
import React from "react";
import OnlineUser from "./OnlineUser";
import OfflineUser from "./Offline";
import Overlay from "./Overlay";

const Online = () => {
  return (
    // オンライン
    <div className="hide-scrollbar h-screen  w-3/12  overflow-y-auto  bg-online">
      <div className="">
        <p className="mt-2 text-center text-slate-300">オンライン</p>
        <OnlineUser name="Taro" imageSrc="/images/rcc-favicon.png" />
        <OnlineUser name="Hanako" imageSrc="/images/rcc-favicon.png" />
        <OnlineUser name="RCCTarou" imageSrc="/images/rcc-favicon.png" />
      </div>

      {/* オフライン */}
      <div className="">
        <p className="mt-2 text-center  text-slate-300">オフライン</p>
        <div className="relative">
          <Overlay />

          <OfflineUser name="Taro" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Hanako" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="RCCTarou" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Taro" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Hanako" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="RCCTarou" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Taro" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Hanako" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Hanako" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="RCCTarou" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Taro" imageSrc="/images/rcc-favicon.png" />
          <OfflineUser name="Hanako" imageSrc="/images/rcc-favicon.png" />
        </div>
      </div>
    </div>
  );
};

export default Online;
