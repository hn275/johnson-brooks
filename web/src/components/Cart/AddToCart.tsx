import { useState, useEffect } from "react";
import { BsCart2 } from "react-icons/bs/index";
import cx from "classnames";
import { addToCart, CartItem } from "./store";
import { AnimatePresence, Variants, motion } from "framer-motion";
import { PAGES } from "@lib/routes";

interface Props {
  productID: string;
  item: CartItem;
}

export function AddToCart({ productID: id, item }: Props) {
  const { toast, onToast } = useToast();

  const vars: Variants = {
    hidden: { opacity: 0, y: 30 },
    show: { opacity: 1, y: 0 },
  };

  return (
    <>
      <button
        onClick={addToCart(id, item, onToast)}
        className={cx(
          "bg-brand-100 text-brand-200 text-sm uppercase",
          "flex items-center justify-center gap-1",
          "hover:brightness-110 rounded-md p-3",
        )}
      >
        Add to cart&nbsp;
        <span>
          <BsCart2 title="shopping-cart" size={18} />
        </span>
      </button>

      <AnimatePresence>
        {toast && (
          <motion.div
            key="toast"
            variants={vars}
            initial="hidden"
            animate="show"
            exit="hidden"
            className={cx(
              "fixed bottom-2 left-2 right-2",
              "md:bottom-8 md:left-8 md:w-max",
              "p-5 bg-green-300",
              "z-[999] rounded-md",
              "flex items-center gap-2",
            )}
          >
            <p className="text-center md:text-left">
              <span className="font-semibold">{item.productTitle}</span>
              &nbsp;added to cart!
            </p>

            <a className="underline italic" href={PAGES.checkout}>
              Checkout now.
            </a>
          </motion.div>
        )}
      </AnimatePresence>
    </>
  );
}

function useToast() {
  const [toast, setToast] = useState<boolean>(false);

  useEffect(() => {
    let timeoutID: number;
    if (!toast) return;

    timeoutID = setTimeout(() => {
      setToast(() => false);
    }, 3000);

    return () => clearTimeout(timeoutID);
  }, [toast]);

  return { toast, onToast: () => setToast(() => true) };
}
