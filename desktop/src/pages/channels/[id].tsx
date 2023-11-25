// import get_test_messages, { Message, useMessages } from "@/api/message"; aaa
import { Message, useMessages } from "@/api/message";
import Chat from "@/components/chat/Chat";
import ChatHeader from "@/components/chat/ChatHeader";
import ChatInput from "@/components/chat/ChatInput";
import Loading from "@/components/Loading";
import { Input } from "@/components/ui/input";
import { Laugh, PlusCircle, Search } from "lucide-react";
import React, { useEffect, useState } from "react";

const ChannelPage = () => {
  const { messages: pastMessages, isLoading, isError } = useMessages();
  const [messages, setMessages] = useState<Message[]>([]);

  useEffect(() => {
    if (pastMessages) {
      setMessages(pastMessages);
    }
  }, [pastMessages]);

  // const [messages, setMessages] = useState<Message[]>([]); // Message[] 型注釈を使用
  // useEffect(() => {
  //   // コンポーネントのマウント時にメッセージをロード
  //   async function loadMessages() {
  //     const loadedMessages = await get_test_messages();
  //     if (loadedMessages) {
  //       setMessages(loadedMessages);
  //     } else {
  //       console.log("エラー");
  //     }
  //   }

  //   loadMessages();
  // }, []);

  // 新しいメッセージを追加する関数
  const addMessage = (content: string) => {
    const newMessage: Message = {
      message_id: Date.now(), // 仮のID生成
      channel_id: 1, // 仮のチャンネルID
      user_id: 1, // 仮のユーザーID
      content: content,
      created_time: new Date(),
      updated_time: new Date(),
      mention: undefined, // 仮のメンション
      pinned: false,
      stamp_id: 0, // 仮のスタンプID
    };
    setMessages([...messages, newMessage]);
  };

  if (isLoading) return <Loading />;
  if (isError) return <div>Error loading messages.</div>;

  return (
    <div className="w-full ">
      <ChatHeader />
      <div className="flex-grow">
        <Chat messages={messages} />
      </div>
      <ChatInput sendNewMessage={addMessage} />
    </div>
  );
};

export default ChannelPage;
