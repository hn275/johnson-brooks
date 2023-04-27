import cx from "classnames";

export type ServiceCardProps = {
  href: string;
  text: string;
  src: string;
};

export function ServiceCard({ href, text, src }: ServiceCardProps) {
  return (
    <div
      className={cx(
        "flex flex-col items-stretch",
        "overflow-hidden lg:relative",
        "group"
      )}
    >
      <img loading="lazy" src={src} alt={text} className="w-full" />

      <div
        className={cx(
          "btn text-center",
          "lg:absolute top-0 left-0 w-full h-full",
          "lg:grid place-items-center",
          "lg:bg-slate-900/0 lg:group-hover:bg-slate-900/20",
          "lg:transition-all"
        )}
      >
        <a
          className={cx(
            "lg:translate-y-5 lg:group-hover:translate-y-0",
            "lg:opacity-0 lg:group-hover:opacity-100",
            "lg:font-semibold lg:btn lg:bg-slate-900",
            "lg:shadow",
            "transition-all cursor-pointer"
          )}
          href={href}
        >
          {text}
        </a>
      </div>
    </div>
  );
}
