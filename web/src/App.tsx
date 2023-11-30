import { useEffect, useState } from "react";
import "./App.css";

import Loading from "react-loading";

import { api } from "./network/base.ts";

function App() {
  const [errMsg, setErrMsg] = useState("");

  const onLogin = async () => {
    const token = new URLSearchParams(window.location.search).get("token");
    if (!token) {
      location.assign("/");
      return;
    }
    try {
      await api.post("/", {
        token,
      });
      location.assign("/");
    } catch (err: any) {
      setErrMsg(err.message);
    }
  };

  useEffect(() => {
    onLogin();
  }, []);

  return (
    <div className={"container"}>
      <div className={"content"}>
        {errMsg ? (
          <>
            <h2>登录失败</h2>
            <h4
              style={{
                marginTop: "0",
                marginBottom: "3rem",
              }}
            >
              {errMsg}
            </h4>
            <button onClick={() => location.assign("/")}>重新登录</button>
          </>
        ) : (
          <>
            <Loading type={"cylon"} />
            <h3>正在登录</h3>
          </>
        )}
      </div>
    </div>
  );
}

export default App;
