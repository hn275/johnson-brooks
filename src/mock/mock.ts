import monorail from "./data.json";

export const mockFetch = () => {
  return new Promise((res, _) => {
    setTimeout(() => {
      res(monorail);
    }, 1000);
  });
};
