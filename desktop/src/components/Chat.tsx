import React from "react";
import Loading from "./Loading";

const Chat = () => {
  return (
    <div className="hide-scrollbar h-screen w-full overflow-y-auto   bg-chat">
      <Loading />
    </div>
  );
};

export default Chat;
