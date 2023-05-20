import { ChangeEvent, useEffect, useState } from "react";

type InputEvent = ChangeEvent<HTMLInputElement>;

export function useLogin() {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  useEffect(() => {
    const timeoutID = setTimeout(() => {
      setError(() => "");
    }, 5000);
    return () => clearTimeout(timeoutID);
  }, [error]);

  const onUserName = (e: InputEvent) => setUsername(() => e.target.value);
  const onPassword = (e: InputEvent) => setPassword(() => e.target.value);

  const api = import.meta.env["PUBLIC_API_URI"];
  async function handleSubmit(e: ChangeEvent<HTMLFormElement>) {
    e.preventDefault();
    try {
      setIsLoading(() => true);
      const res = await fetch(`${api}/auth/login`, {
        method: "POST",
        body: JSON.stringify({ username, password }),
      });

      const { status } = res;
      if (status === 200) {
        window.location.replace("/admin");
        return;
      }

      const payload = await res.json();
      setError(() => payload["error"]);
    } catch (e) {
      setError(() => "Something went wrong.");
      console.error(e);
    } finally {
      setIsLoading(() => false);
    }
  }

  return {
    username,
    onUserName,
    password,
    onPassword,
    isLoading,
    handleSubmit,
    error,
  };
}
