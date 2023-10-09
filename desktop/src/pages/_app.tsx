import CSR from "@/components/CSR";
import Header from "@/components/Header";
import Online from "@/components/Online";
import Sidebar from "@/components/Sidebar/Sidebar";
import "@/styles/globals.css";
import type { AppProps } from "next/app";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

export default function App({ Component, pageProps }: AppProps) {
  const router = useRouter();
  const isAuthPage = router.pathname.includes("/auth");

  return (
    <CSR>
      <div className="w-full">
        {!isAuthPage && (
          <div className="sticky top-0 ">
            <Header />
          </div>
        )}

        <div className="flex h-screen ">
          {!isAuthPage && <Sidebar />}
          <Component {...pageProps} />
          {!isAuthPage && <Online />}
        </div>
      </div>
    </CSR>
  );
}
