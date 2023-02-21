import Head from "next/head";

function LoginForm() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-100">
      <Head>
        <title>Login</title>
      </Head>

      <main className="flex-flex-col items-center justify-center w-full flex-1 px-20 text-center">
        <div className="w-3/5 p-10"></div>
        <div className="mb-10"></div>
        <form className="space-x-3" action="http://localhost:8080/login" method="post">
          <label htmlFor="username">Username</label>

          <input className="mt-6 outline outline-offset-2 outline-1 shadow-xl" type="text" id="username" name="username" required />
          <form className="space-x-3">
            <label htmlFor="password">Password</label>
            <input className="mt-6 outline outline-offset-2 outline-1 shadow-xl" type="text" id="paswword" name="password" required />

            <button
              className="w-3/5 py-2 mt-6 text-medium text-white bg-blue-500 rounded-md"
              type="submit"
            >
              Submit
            </button>
          </form>
        </form>
      </main>
    </div>
  );
}

export default LoginForm;
