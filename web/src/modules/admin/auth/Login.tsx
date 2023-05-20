import { useState } from "react";
import cx from "classnames";
import Button from "../components/Button";
import { CgLogIn } from "react-icons/cg/index";
import { BiError } from "react-icons/bi/index";
import { useLogin } from "./useLogIn";
import { AnimatePresence, motion } from "framer-motion";

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
    <>
      <form
        className="flex flex-col gap-5 my-5 relative"
        onSubmit={handleSubmit}
      >
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

        <AnimatePresence>
          {error && (
            <motion.p
              className={cx(
                "fixed bottom-8 left-8",
                "bg-red-200 mx-auto",
                "p-3 rounded-md",
                "flex items-center gap-1",
              )}
              initial={{ opacity: 0, translateY: -20 }}
              animate={{ opacity: 1, translateY: 0 }}
              exit={{ opacity: 0, translateY: -20 }}
            >
              <span className="text-red-700">
                <BiError />
              </span>
              {error}
            </motion.p>
          )}
        </AnimatePresence>
      </form>
    </>
  );
}
