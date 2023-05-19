import { useState } from "react";
import type { ProductVariant } from "@schemas/index";
import cx from "classnames";

interface Props {
  variants: ProductVariant[];
}

export function Carousel({ variants }: Props) {
  const { variant, onSel } = useActiveVariant(variants);
  const hasVariants = variants.length > 1;
  const outOfStock = variant.inventory === 0;

  return (
    <>
      <div className="relative isolate">
        <div className="w-full aspect-[4/3] overflow-hidden">
          <img
            src={`data:image/png;base64,${variant?.thumbnail}`}
            className="w-full overflow-hidden object-cover"
          />
        </div>

        {hasVariants && (
          <>
            <div>
              <p className="">Select a color:</p>
              <ul className="flex items-center gap-2 h-full">
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
          </>
        )}
      </div>
    </>
  );
}

function useActiveVariant(variants: ProductVariant[]) {
  const [index, setIndex] = useState<number>(0);

  const onSel = (i: number) => setIndex(() => i);

  return { variant: variants[index]!, onSel };
}
