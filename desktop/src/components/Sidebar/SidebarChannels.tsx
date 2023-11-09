import React, { useEffect, useState } from "react";
import Channnels from "../Channnels";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import PersonIcon from "@mui/icons-material/Person";
import PrivateChat from "./PrivateChat";
import AddIcon from "@mui/icons-material/Add";
import SearchIcon from "@mui/icons-material/Search";
import AddChannel from "./AddChannel";
import get_test_channels from "@/api/channel";
import Link from "next/link";

const SidebarChannels = () => {
  const [showAddChannel, setShowAddChannel] = useState(false);
  const [channels, setChannels] = useState([]);

  const toggleAddChannel = () => {
    setShowAddChannel(!showAddChannel);
  };

  useEffect(() => {
    const getChannels = async () => {
      const data = await get_test_channels();
      console.log(data);
    };
    getChannels();
  }, []);

  return (
    //チャンネルデータを取得してきて表示する
    //useEffectまたはuseSWRを使う

    <div className="hide-scrollbar relative   h-screen overflow-y-auto ">
      <Link href="/channels/1">
        <Channnels icon="#" name="general" />
      </Link>
      <Channnels icon="#" name="announce" />
      <div className="ml-10 ">
        <div className="flex">
          <span>
            <ExpandMoreIcon />
          </span>
          スター付きチャンネル
        </div>
        {/* <div className="ml-2">
          {channels.map((channel) => (
            <Channnels key={channel.id}  icon="#" name={channel.name} />
          ))}
        </div> */}

        <div className="ml-2">
          {/* {channels.map((channel) => (
            <Channnels key={channel.id}  icon="#" name={channel.name} />
          ))} */}
          {/* <Channnels icon="#" name="general" />
          <Channnels icon="#" name="announce" />
          <Channnels icon="#" name="general" />
          <Channnels icon="#" name="announce" />
          <Channnels icon="#" name="general" /> */}
        </div>
      </div>

      {/* チャンネル */}
      <div className="ml-10 mt-4 ">
        <div className="my-1 flex cursor-pointer">
          <AddIcon onClick={toggleAddChannel} />
          <div>チャンネルを追加</div>
          {showAddChannel && (
            <div className="absolute z-10 -mt-40 ml-0 flex items-center justify-center ">
              <AddChannel />
            </div>
          )}
        </div>
        <div className="my-1 flex">
          <SearchIcon />
          <div>チャンネルを探す</div>
        </div>
        <div className="flex">
          <ExpandMoreIcon />
          チャンネル
        </div>
        <div className="ml-2">
          <Channnels icon="#" name="general" />
          <Channnels icon="#" name="announce" />
          <Channnels icon="#" name="general" />
          <Channnels icon="#" name="announce" />
          <Channnels icon="#" name="general" />
        </div>
      </div>

      {/* プライベートメッセージ */}
      <div className="ml-10 mt-4">
        <div className="flex">
          <span>
            <PersonIcon />
          </span>
          プライベートメッセージ
        </div>
        <div className="mb-6 ml-2">
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
          <PrivateChat imgUrl="/images/rcc-favicon.png" name="Taro" />
        </div>
      </div>
    </div>
  );
};

export default SidebarChannels;
