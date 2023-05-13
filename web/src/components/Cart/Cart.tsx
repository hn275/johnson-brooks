import { BsCart2 } from "react-icons/bs/index";
import cx from "classnames";
import { CART_STORAGE, CartStorage, cartItem } from "./store";
import { useStore } from "@nanostores/react";
import { motion, AnimatePresence, Variants } from "framer-motion";
import { useState, useEffect } from "react";

export function Cart() {
  const show = useCartIndicator();

  const vars: Variants = {
    hidden: { scale: 0 },
    show: { scale: 1 },
  };

  return (
    <>
      <div className="relative">
        <BsCart2
          size={20}
          className={cx(
            "text-slate-500 md:text-slate-400",
            "md:hover:text-white",
            "transition-colors cursor-pointer",
          )}
        />

        <AnimatePresence>
          {show && (
            <motion.div
              variants={vars}
              initial="hidden"
              animate="show"
              exit="hidden"
              className={cx(
                "absolute -top-1 -right-1 z-50",
                "bg-brand-100 rounded-full w-2 h-2",
              )}
            ></motion.div>
          )}
        </AnimatePresence>
      </div>
    </>
  );
}

function useCartIndicator() {
  const [counter, setCounter] = useState<boolean>(false);
  const items = useStore(cartItem);

  useEffect(() => {
    const cachedCart = window.localStorage.getItem(CART_STORAGE);
    if (!cachedCart) return;

    const items = JSON.parse(cachedCart) as CartStorage;
    for (const [k, v] of Object.entries(items)) {
      cartItem.setKey(k, v);
    }
  }, []);

  useEffect(() => {
    const cartEntries = Object.keys(items);
    setCounter(() => cartEntries.length !== 0);
  }, [items]);

  return counter;
}
