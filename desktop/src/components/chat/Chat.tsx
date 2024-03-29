import React from "react";
import Loading from "../Loading";
import { Skeleton } from "../ui/skeleton";
import Image from "next/image";
import { ThumbsUp } from "lucide-react";
import { Message, useMessages } from "@/api/message";
import { useUser, useUsers } from "@/api/user";
import { useChannels } from "@/api/channel";

interface ChatProps {
  messages: Message[];
}

const Chat = ({ messages }: ChatProps) => {
  // const { user, isLoading, isError } = useUser(1);
  const { users } = useUsers();
  //TODO:デスクトップAPIデータを任意のidを取ってくるよう変更し ユーザーごとの情報を取得する
  // const { user, isLoading, isError } = useUser(messages.user_id);
  const { channels } = useChannels();
  console.log(channels);
  console.log(users);

  return (
    <div className="hide-scrollbar h-screen w-full overflow-y-auto   bg-chat">
      {/*TODO: データ取得時にローディングコンポーネントの表示　 <Loading /> */}
      <div className="ml-3 space-y-4">
        {messages.map((message) => (
          <div key={message.message_id} className="flex items-start space-x-4">
            {/*TODO: メッセージごとのユーザーの画像を表示 */}
            <Image
              src="/images/rcc-favicon.png"
              alt="RCCのfaviconです"
              width={40}
              height={40}
              className="h-12 w-12 rounded-full "
            />

            <div className="space-y-2">
              <div className="flex">
                {/* TODO: ここにメッセージごとのユーザー名を表示  */}
                <p className="h-4 w-[100px]">{users?.[0]?.username}</p>
                {/* 送信日時の表示 */}
                <p className="text-xs">{new Date(message.created_time).toLocaleDateString()}</p>
              </div>

              <div className="flex space-x-1">
                {/* メッセージを表示  */}
                {message.content}
              </div>
              <div className="flex space-x-1">
                {/* TODO: ここに絵文字を表表示 */}
                {/* //{message.stamp} */}
                <ThumbsUp className="text-yellow-500" />
                {/*TODO: いいね数を表示する */}
                <span>4</span>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Chat;
