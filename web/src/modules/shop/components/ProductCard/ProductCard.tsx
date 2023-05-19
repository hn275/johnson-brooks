import type { Product, ProductVariant } from "@schemas/index";
import cx from "classnames";
import { AddToCart } from "@components/Cart/AddToCart";
import type { CartItem } from "@components/Cart/store";
import { useState } from "react";
import { getImageSource } from "@lib/getImageSource";

export function ProductCard(props: Product) {
  const { title, price, description, id, variants, material } = props;
  const { variant, onSel } = useActiveVariant(variants);

  const cartItem: CartItem = {
    description,
    quantity: 1,
    unitPrice: price,
    productTitle: title,
  };

  return (
    <div
      className={cx(
        "bg-white p-3 shadow-lg shadow-slate-400/20 rounded-md w-full",
      )}
    >
      <div
        className={cx(
          "text-xs text-slate-600",
          "w-full aspect-[4/3] overflow-hidden",
        )}
      >
        <img
          src={getImageSource(variant.thumbnail)}
          className="h-full w-auto object-cover rounded-sm"
        />
      </div>

      <div className="my-3 flex flex-col gap-4 flex-grow">
        <div>
          <h3 className="text-lg font-semibold text-brand-200 uppercase">
            {title}{" "}
            <span className="text-xs font-normal normal-case opacity-50">
              &nbsp;{material}
            </span>
          </h3>
          <p className="text-lg">${price.toFixed(2)}</p>

          <div className="mt-2">
            <p className="text-sm">Select a color</p>

            <ul className="flex gap-2">
              {variants.map((v, i) => (
                <li key={i}>
                  <button
                    style={{ backgroundColor: v.color }}
                    className={cx(
                      "p-2 rounded-full border",
                      v.color === variant.color
                        ? "border-2 border-brand-100"
                        : "border-slate-300/50",
                    )}
                    aria-label={`color: ${v.color}`}
                    onClick={() => onSel(i)}
                  />
                </li>
              ))}
            </ul>
          </div>
        </div>

        <p>{description}</p>
      </div>

      <AddToCart productID={id} item={cartItem} />
    </div>
  );
}

function useActiveVariant(variants: ProductVariant[]) {
  const [index, setIndex] = useState<number>(0);

  const onSel = (i: number) => setIndex(() => i);

  return { variant: variants[index]!, onSel };
}
