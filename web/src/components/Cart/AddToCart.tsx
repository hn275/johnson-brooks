import { BsCart2 } from "react-icons/bs/index";
import cx from "classnames";

interface Props {
  productId: string;
}

export function AddToCart({ productId }: Props) {
  const handleClick = () => console.log(productId);
  return (
    <button
      onClick={handleClick}
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
  );
}
