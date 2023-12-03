import { Search } from "lucide-react";
import { Input } from "../ui/input";

const ChatHeader = () => {
  return (
    <div className="relative flex items-center m-2">
      <Search className="text-slate-400 absolute ml-2" />
      <Input className="m-1 pl-8 bg-chat" placeholder="検索" />
    </div>
  );
};

export default ChatHeader;
