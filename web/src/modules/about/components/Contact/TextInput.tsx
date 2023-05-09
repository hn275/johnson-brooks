import type { InputHTMLAttributes } from "react";
import cx from "classnames";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label: string;
}

export function TextInput(props: InputProps) {
  const { label, id } = props;

  return (
    <div className="w-full bg-transparent form-control">
      <label className="label" htmlFor={id}>
        {label}
      </label>
      <input
        {...props}
        className={cx("p-2 rounded-md bg-slate-100", "w-full")}
      />
    </div>
  );
}
