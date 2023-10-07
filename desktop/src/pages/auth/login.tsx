import Link from "next/link";
import React from "react";
import Head from "next/head";
import { useForm, SubmitHandler } from "react-hook-form";
import { useRouter } from "next/router";

type FormData = {
  userName: string;
  email: string;
  password: string;
};

const Login = () => {
  const {
    register,
    setValue,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>();

  const router = useRouter();

  const OnSubmit: SubmitHandler<FormData> = (data: FormData) => {
    console.log(data);
    router.push("/");
  };

  return (
    <div className="flex h-screen w-full flex-col justify-center bg-default py-12 sm:px-6 lg:px-8">
      <Head>
        <title>新規登録</title>
      </Head>

      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">ログイン</h2>
      </div>
      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-xl">
        <div className="bg-white px-4 py-8 shadow sm:rounded-lg sm:px-10">
          <form onSubmit={handleSubmit(OnSubmit)}>
            <div className="mt-6">
              <label className="block text-sm font-medium text-gray-700">メールアドレス</label>
              <input
                {...register("email", {
                  required: "パスワードを入力してください",
                  pattern: {
                    value: /^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/,
                    message: "不適切なメールアドレスです。",
                  },
                })}
                type="email"
                className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 text-base shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
              />
              {errors.email && <span className="text-xs text-red-500">{errors.email.message}</span>}
            </div>

            <div className="mt-6">
              <label className="block text-sm font-medium text-gray-700">パスワード</label>
              <input
                {...register("password", {
                  required: "パスワードを入力してください",
                  minLength: { value: 6, message: "パスワードは6文字以上である必要があります" },
                })}
                type="password"
                className="mt-1 block w-full rounded-md border border-gray-300 px-3 py-2 text-base shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-indigo-500"
              />
              {errors.password && <span className="text-xs text-red-500">{errors.password.message}</span>}
            </div>

            <div className="mt-12">
              <button
                type="submit"
                className=" w-full rounded-md border border-transparent bg-indigo-600  py-3 text-lg font-medium text-white shadow-lg hover:bg-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-400 focus:ring-offset-2"
              >
                ログイン
              </button>
            </div>
            <div className="mt-4 flex justify-center">
              新規登録は
              <Link href="/auth/register" className=" text-blue-300">
                こちら
              </Link>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
