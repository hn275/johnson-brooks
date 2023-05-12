import { BsCart2 } from "react-icons/bs/index";
import cx from "classnames";

export function Cart() {
  return (
    <>
      <BsCart2
        size={20}
        className={cx(
          "text-slate-500 md:text-slate-400",
          "md:hover:text-slate-600",
          "transition-colors cursor-pointer",
        )}
      />
    </>
  );
}
