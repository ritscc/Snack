import Chat from "@/components/chat/Chat";
import ChatHeader from "@/components/chat/ChatHeader";
import ChatInput from "@/components/chat/ChatInput";
import Loading from "@/components/Loading";
import { Input } from "@/components/ui/input";
import { Laugh, PlusCircle, Search } from "lucide-react";
import React from "react";

const ChannelPage = () => {
  return (
    <div className="w-full ">
      <ChatHeader />

      <div className="flex-grow">
        <Chat />
      </div>

      <ChatInput />
    </div>
  );
};

export default ChannelPage;
