import dynamic from "next/dynamic";

const CSR = dynamic(() => import("./CSRInner"), { ssr: false });

export default CSR;
