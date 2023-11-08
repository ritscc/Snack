import { Skeleton } from "./ui/skeleton";

const Loading = () => {
  return (
    <div className="ml-3 space-y-4">
      <Skeleton className="h-8 w-3/5" />
      <div className="flex items-start space-x-4">
        <Skeleton className="h-12 w-12 rounded-full" />
        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px] bg-slate-100" />
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[200px]" />
            <Skeleton className="h-4 w-[120px]" />
            <Skeleton className="h-4 w-[70px]" />
          </div>
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[150px]" />
            <Skeleton className="h-4 w-[100px]" />
            <Skeleton className="h-4 w-[70px]" />
            <Skeleton className="h-4 w-[90px]" />
          </div>
        </div>
      </div>

      <div className="flex items-start space-x-4">
        <Skeleton className="h-12 w-12 rounded-full" />
        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px] bg-slate-100" />
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[200px]" />
            <Skeleton className="h-4 w-[70px]" />
          </div>
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[150px]" />
            <Skeleton className="h-4 w-[100px]" />
            <Skeleton className="h-4 w-[70px]" />
            <Skeleton className="h-4 w-[90px]" />
          </div>
          <Skeleton className="h-40 w-3/5" />
        </div>
      </div>
      <div className="flex items-start space-x-4">
        <Skeleton className="h-12 w-12 rounded-full" />
        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px] bg-slate-100" />
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[200px]" />
            <Skeleton className="h-4 w-[120px]" />
            <Skeleton className="h-4 w-[70px]" />
          </div>
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[150px]" />
            <Skeleton className="h-4 w-[100px]" />
            <Skeleton className="h-4 w-[70px]" />
            <Skeleton className="h-4 w-[90px]" />
          </div>
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[150px]" />
            <Skeleton className="h-4 w-[100px]" />
            <Skeleton className="h-4 w-[70px]" />
            <Skeleton className="h-4 w-[90px]" />
          </div>
        </div>
      </div>
      <div className="flex items-start space-x-4">
        <Skeleton className="h-12 w-12 rounded-full" />
        <div className="space-y-2">
          <Skeleton className="h-4 w-[100px] bg-slate-100" />
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[200px]" />
            <Skeleton className="h-4 w-[120px]" />
            <Skeleton className="h-4 w-[70px]" />
          </div>
          <div className="flex space-x-1">
            <Skeleton className="h-4 w-[150px]" />
            <Skeleton className="h-4 w-[100px]" />
            <Skeleton className="h-4 w-[70px]" />
            <Skeleton className="h-4 w-[90px]" />
            <Skeleton className="h-4 w-[120px]" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Loading;
