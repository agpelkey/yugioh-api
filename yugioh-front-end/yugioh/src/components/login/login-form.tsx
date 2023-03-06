import Head from "next/head";

function LoginForm() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-3 bg-gray-100">
      <Head>
        <title>Login</title>
      </Head>

      <h1 className="flex flex-col mt-20">Login</h1>
      <main className="flex-flex-col items-center justify-center w-full flex-1 px-20 text-center">
        <div className="w-3/5 mb-10"></div>
        <form
          className="space-x-3"
          action="http://localhost:8080/login"
          method="post"
        >
          <label htmlFor="username">Username</label>

          <input
            className="outline outline-offset-2 outline-1 shadow-xl"
            type="text"
            id="username"
            name="username"
            required
          />
          <form className="space-x-3">
            <label htmlFor="password">Password</label>
            <input
              className="mt-6 outline outline-offset-2 outline-1 shadow-xl"
              type="text"
              id="paswword"
              name="password"
              required
            />
            <div className="list-none">
              <button
                className="w-1/3 bg-orange-200 hover:bg-orange-400 rounded-lg my-4"
                type="submit"
              >
                Submit
              </button>
            </div>
          </form>
        </form>
      </main>
    </div>
  );
}

export default LoginForm;
