import { useState } from "react";
import cx from "classnames";
import Button from "../components/Button";
import { CgLogIn } from "react-icons/cg/index";
import { useLogin } from "./useLogIn";

export default function Login() {
  const [showPass, setShowPass] = useState<boolean>(false);
  const {
    username,
    onUserName,
    password,
    onPassword,
    error,
    handleSubmit,
    isLoading,
  } = useLogin();

  return (
    <form className="flex flex-col gap-5 my-5" onSubmit={handleSubmit}>
      <input
        id="username"
        type="text"
        className="text-input"
        placeholder="User name"
        value={username}
        onChange={onUserName}
      />

      <input
        id="password"
        type={showPass ? "text" : "password"}
        placeholder="Password"
        className="text-input"
        value={password}
        onChange={onPassword}
      />

      <div className="flex justify-end items-center gap-1">
        <input
          type="checkbox"
          onChange={() => setShowPass((s) => !s)}
          checked={showPass}
        />
        <label className={cx("text-slate-500")} htmlFor="show-pass">
          show password
        </label>
      </div>

      <Button
        className="btn btn-primary"
        loading={isLoading ? true : undefined}
        icon={<CgLogIn />}
        type="submit"
      >
        Sign in
      </Button>

      {error && <p>{error}</p>}
    </form>
  );
}
