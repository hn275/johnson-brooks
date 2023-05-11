import { ChangeEvent, useState } from "react";

type InputEvent = ChangeEvent<HTMLInputElement>;

export function useLogin() {
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  const onUserName = (e: InputEvent) => setUsername(() => e.target.value);
  const onPassword = (e: InputEvent) => setPassword(() => e.target.value);

  async function handleSubmit(e: ChangeEvent<HTMLFormElement>) {
    e.preventDefault();
    try {
      setIsLoading(() => true);
      const res = await mockFetch();
      console.log(res);
      // window.location.replace(res as string);
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

function mockFetch() {
  return new Promise((res) => {
    setTimeout(() => {
      res("localhost:3000/admin/products");
    }, 2000);
  });
}
