import { BsCart2 } from "react-icons/bs/index";
import cx from "classnames";
import { cartItem } from "./store";
import { useStore } from "@nanostores/react";
import { motion, AnimatePresence, Variants } from "framer-motion";

export function Cart() {
  const items = useStore(cartItem);
  const itemCount = Object.keys(items).length;
  const count = itemCount <= 10 ? String(itemCount) : "9+";

  const vars: Variants = {
    hidden: { scale: 0 },
    show: { scale: 1 },
  };

  return (
    <div className="relative">
      <BsCart2
        size={20}
        className={cx(
          "text-slate-500 md:text-slate-400",
          "md:hover:text-slate-600",
          "transition-colors cursor-pointer",
        )}
      />

      <AnimatePresence>
        {itemCount !== 0 && (
          <motion.div
            variants={vars}
            initial="hidden"
            animate="show"
            exit="hidden"
            className={cx(
              "absolute -top-2 -right-2",
              "bg-brand-100 text-brand-200",
              "rounded-full w-4 h-4",
              "grid place-items-center",
              "origin-center",
            )}
          >
            <p className="text-[0.5rem]">{count}</p>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
