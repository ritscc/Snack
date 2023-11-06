import React, { useEffect, useState } from "react";
import Channnels from "../Channnels";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import PersonIcon from "@mui/icons-material/Person";
import PrivateChat from "./PrivateChat";
import AddIcon from "@mui/icons-material/Add";
import SearchIcon from "@mui/icons-material/Search";
import AddChannel from "./AddChannel";
import get_test_channels, { DefaultChannel } from "@/api/channel";

const SidebarChannels = () => {
  const [showAddChannel, setShowAddChannel] = useState(false);
  // useStateに型注釈<Channel[]>を追加
  const [channels, setChannels] = useState<DefaultChannel[]>([]);

  const toggleAddChannel = () => {
    setShowAddChannel(!showAddChannel);
  };
  useEffect(() => {
    const fetchChannels = async () => {
      const channels = await get_test_channels();
      if (channels) {
        setChannels(channels);
        console.log(channels);
      } else {
        console.log(channels);
      }
    };
    fetchChannels();
  }, []);

  return (
    //チャンネルデータを取得してきて表示する
    //useEffectまたはuseSWRを使う

    <div className="hide-scrollbar relative   h-screen overflow-y-auto ">
      <Channnels icon="#" name="general" />
      <Channnels icon="#" name="announce" />
      <div className="ml-10 ">
        <div className="flex">
          <span>
            <ExpandMoreIcon />
          </span>
          スター付きチャンネル
        </div>
        <div className="ml-2">
          {channels.map((channel) => (
            <Channnels key={channel.channel_id} icon="#" name={channel.name} />
          ))}
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
