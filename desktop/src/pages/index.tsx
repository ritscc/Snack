import Image from "next/image";
import { Inter } from "next/font/google";
import Chat from "@/components/chat/Chat";
import Online from "@/components/Online";

const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  return <div className="flex w-full  ">{/* <Chat /> */}</div>;
}
