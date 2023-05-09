import monorail from "./monorail.json";

export const mockFetch = () => {
  return new Promise((res, _) => {
    setTimeout(() => {
      res(JSON.stringify(monorail));
    }, 1000);
  });
};
