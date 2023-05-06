import type { TextareaHTMLAttributes } from "react";
import cx from "classnames";

interface TextAreaProps extends TextareaHTMLAttributes<HTMLTextAreaElement> {
  label: string;
}
export function TextArea(props: TextAreaProps) {
  const { label, id, className } = props;
  return (
    <div className="w-full form-control bg-transparent">
      <label htmlFor={id} className="label">
          {label}
      </label>
      <textarea {...props} className={cx("bg-slate-100 p-3 rounded-md", className)}/>
    </div>
  )
}
