import type { ButtonHTMLAttributes, ReactNode } from "react";
import { CgSpinner } from "react-icons/cg/index";

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  icon?: ReactNode;
  loading?: boolean | undefined;
  children: ReactNode;
}

export default function Button(props: Props) {
  const { loading, disabled, children, icon, ...rest } = props;
  return (
    <button disabled={loading || disabled} {...rest}>
      {children}
      <span>
        {loading ? <CgSpinner className="animate-spin text-inherit" /> : icon}
      </span>
    </button>
  );
}
