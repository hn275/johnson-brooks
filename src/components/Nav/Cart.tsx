import icons from "react-icons/bs/index";

export function Cart() {
  const { BsCart2 } = icons;
  return (
    <BsCart2
      size={20}
      className="text-slate-500 md:text-slate-400 md:hover:text-slate-600 transition-colors cursor-pointer"
    />
  );
}
