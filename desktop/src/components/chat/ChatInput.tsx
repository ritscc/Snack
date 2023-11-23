import { Laugh, PlusCircle } from "lucide-react";
import { Input } from "../ui/input";
import { useState } from "react";

interface ChatInputProps {
  sendNewMessage: (message: string) => void;
}

const ChatInput = ({ sendNewMessage }: ChatInputProps) => {
  const [inputText, setInputText] = useState("");
  const handleSendMessage = () => {
    if (inputText.trim()) {
      sendNewMessage(inputText);
      setInputText("");
    }
  };

  return (
    <div className="flex justify-between sticky bottom-1">
      <div className="relative flex items-center m-2  w-full ">
        <PlusCircle className="text-slate-400 absolute ml-2" />
        <Input
          className="m-1 pl-8 bg-slate-100 "
          placeholder=""
          value={inputText}
          onChange={(e) => setInputText(e.target.value)}
          onKeyDown={(e) => e.key === "Enter" && handleSendMessage()}
        />
        <Laugh className="right-2 absolute text-slate-400 " />
      </div>
    </div>
  );
};

export default ChatInput;
