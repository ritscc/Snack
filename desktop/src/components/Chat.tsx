import React from "react";
import Loading from "./Loading";

const Chat = () => {
  return (
    <div className="hide-scrollbar h-screen w-full overflow-y-auto   bg-chat text-chat-text">
      <Loading />
    </div>
  );
};

export default Chat;
