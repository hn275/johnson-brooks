import { useState } from "react";
import type { ProductVariant } from "@schemas/index";
import cx from "classnames";
import { AiOutlineArrowLeft, AiOutlineArrowRight } from "react-icons/ai/index";

interface Props {
  variants: ProductVariant[];
}

export function Carousel({ variants }: Props) {
  const { variant, onNext, onPrev, onSel } = useActiveVariant(variants);
  const hasVariants = variants.length > 1;

  return (
    <>
      <div className="relative isolate">
        <img
          src={`data:image/png;base64,${variant?.thumbnail}`}
          className="w-full"
        />

        {hasVariants && (
          <>
            <div className="flex justify-center items-center gap-3">
              <AiOutlineArrowLeft
                className={cx("-left-2 carousel-btn")}
                role="button"
                onClick={onPrev}
                title="previous variant"
              />

              <AiOutlineArrowRight
                className={cx("-right-2 carousel-btn")}
                role="button"
                onClick={onNext}
              />
            </div>

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

        <p>Inventory: {variant.inventory}</p>
      </div>
    </>
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
