import { useState, useEffect } from "react";
import { BsCartFill } from "react-icons/bs/index";
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

  return (
    <>
      <button
        onClick={addToCart(id, item, onToast)}
        className={cx(
          "text-brand-100 text-xs uppercase font-bold bg-inherit",
          "flex items-center justify-center gap-1 rounded-full px-4 py-3",
          "hover:brightness-110 hover:bg-brand-200/5",
          "ml-auto transition-all duration-300",
        )}
      >
        Add to cart&nbsp;
        <span>
          <BsCartFill title="shopping-cart" size={16} />
        </span>
      </button>

      <Toast title={item.productTitle} toast={toast} />
    </>
  );
}

interface ToastProps {
  title: string;
  toast: boolean;
}

function Toast({ title, toast }: ToastProps) {
  const vars: Variants = {
    hidden: { opacity: 0, y: 30 },
    show: { opacity: 1, y: 0 },
  };

  return (
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
            <span className="font-semibold">{title}</span>
            &nbsp;added to cart!
          </p>

          <a className="underline italic" href={PAGES.checkout}>
            Checkout now.
          </a>
        </motion.div>
      )}
    </AnimatePresence>
  );
}

function useToast() {
  const [toast, setToast] = useState<boolean>(false);
  const onToast = () => setToast(() => true);

  useEffect(() => {
    const timeoutID = setTimeout(() => {
      if (!toast) return;
      setToast(() => false);
    }, 3000);

    return () => clearTimeout(timeoutID);
  }, [toast]);

  return { toast, onToast };
}
