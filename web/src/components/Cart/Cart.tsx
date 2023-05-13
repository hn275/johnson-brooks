import { BsCart2 } from "react-icons/bs/index";
import cx from "classnames";
import { CART_STORAGE, CartStorage, cartItem } from "./store";
import { useStore } from "@nanostores/react";
import { motion, AnimatePresence, Variants } from "framer-motion";
import { useState, useEffect } from "react";

export function Cart() {
  const counter = useCartCount();

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
        {counter !== 0 && (
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
            <p className="text-[0.5rem]">{counter <= 9 ? counter : "9+"}</p>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}

function useCartCount() {
  const [counter, setCounter] = useState<number>(0);
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
    let total = 0;
    Object.values(items).forEach(({ quantity }) => (total += quantity));
    setCounter(() => total);
  }, [items]);

  return counter;
}
