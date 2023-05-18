import { useState } from "react";
import type { ProductVariant } from "@schemas/index";
import cx from "classnames";
import { BsChevronLeft, BsChevronRight } from "react-icons/bs/index";

interface Props {
  variants: ProductVariant[];
}

export function Carousel({ variants }: Props) {
  const { variant, onNext, onPrev, onSel } = useActiveVariant(variants);
  const hasVariants = variants.length > 1;

  return (
    <div className="text-xs text-slate-600">
      <div className="relative isolate group">
        <img
          src={`data:image/png;base64,${variant?.thumbnail}`}
          className="w-full"
        />

        {hasVariants && (
          <>
            <BsChevronLeft
              className={cx("-left-2 carousel-btn")}
              role="button"
              onClick={onPrev}
              title="previous variant"
            />

            <BsChevronRight
              className={cx("-right-2 carousel-btn")}
              role="button"
              onClick={onNext}
            />

            <div className="flex items-center gap-2 my-1">
              <p className="">Color:</p>

              <ul className="flex items-center gap-2 h-full">
                {variants.map((v, i) => (
                  <li key={i}>
                    <button
                      style={{ backgroundColor: v.color }}
                      className="p-2 rounded-full border border-brand-200/50"
                      aria-label={`color: ${v.color}`}
                      onClick={() => onSel(i)}
                    />
                  </li>
                ))}
              </ul>
            </div>
          </>
        )}

        <p>Inventory: {variant.inventory}</p>
      </div>
    </div>
  );
}

function useActiveVariant(variants: ProductVariant[]) {
  const [index, setIndex] = useState<number>(0);
  const length = variants.length;

  const onNext = () => setIndex((i) => (i + 1) % length);
  const onPrev = () => setIndex((i) => Math.abs(i - 1) % length);
  const onSel = (i: number) => setIndex(() => i);

  return { variant: variants[index]!, onNext, onPrev, onSel };
}
