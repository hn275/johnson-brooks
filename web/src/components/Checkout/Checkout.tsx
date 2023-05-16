import cx from "classnames";
import { CART_STORAGE, CartStorage, cartItem } from "../Cart/store"
import { useStore } from "@nanostores/react";
import { useState, useEffect } from "react";

export function Checkout() {
  const show = useCartIndicator();
  console.log(show)

  return (
    <div className={cx("grid sm:px-10 lg:grid-cols-2 lg:px-20 xl:px-32 p-20")}>
        <div className={cx("px-4 pt-8")}>
            <p className={cx("text-xl font-medium")}>Order Summary</p>
            <p className={cx("text-gray-400")}>Check your items. And select a suitable shipping method.</p>
            <div className={cx("mt-8 space-y-3 rounded-lg border bg-white px-2 py-4 sm:px-6")}>
                <div className={cx("flex flex-col rounded-lg bg-white sm:flex-row")}>
                    <img className={cx("m-2 h-24 w-28 rounded-md border object-cover object-center")} src="https://images.unsplash.com/flagged/photo-1556637640-2c80d3201be8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8c25lYWtlcnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60" alt="" />
                    <div className={cx("flex w-full flex-col px-4 py-4")}>
                        <span className={cx("font-semibold")}>Nike Air Max Pro 8888 - Super Light</span>
                        <span className={cx("float-right text-gray-400")}>42EU - 8.5US</span>
                        <p className={cx("text-lg font-bold")}>$138.99</p>
                    </div>
                </div>
                <div className={cx("flex flex-col rounded-lg bg-white sm:flex-row")}>
                    <img className={cx("m-2 h-24 w-28 rounded-md border object-cover object-center")} src="https://images.unsplash.com/photo-1600185365483-26d7a4cc7519?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8c25lYWtlcnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60" alt="" />
                    <div className={cx("flex w-full flex-col px-4 py-4")}>
                        <span className={cx("font-semibold")}>Nike Air Max Pro 8888 - Super Light</span>
                        <span className={cx("float-right text-gray-400")}>42EU - 8.5US</span>
                        <p className={cx("mt-auto text-lg font-bold")}>$238.99</p>
                    </div>
                </div>
            </div>
        </div>
           
        <div className={cx("mt-10 bg-gray-50 px-4 pt-8 lg:mt-0")}>
            <p className={cx("text-xl font-medium")}>Payment Details</p>
            <p className={cx("text-gray-400")}>Complete your order by providing your payment details.</p>
            <label htmlFor="email" className={cx("mt-4 mb-2 block text-sm font-medium")}>Email</label>
            <div className={cx("relative")}>
                <input type="text" id="email" name="email" className={cx("w-full rounded-md border border-gray-200 px-4 py-3 pl-11 text-sm shadow-sm outline-none focus:z-10 focus:border-blue-500 focus:ring-blue-500")} placeholder="your.email@gmail.com" />
                <div className={cx("pointer-events-none absolute inset-y-0 left-0 inline-flex items-center px-3")}>
                <svg xmlns="http://www.w3.org/2000/svg" className={cx("h-4 w-4 text-gray-400")} fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207" />
                </svg>
                </div>
            </div>

            <div className={cx("mt-6 flex items-center justify-between")}>
                <p className={cx("text-sm font-medium text-gray-900")}>Total</p>
                <p className={cx("text-2xl font-semibold text-gray-900")}>$408.00</p>
            </div>
            <button className={cx("mt-4 mb-8 w-full rounded-md bg-gray-900 px-6 py-3 font-medium text-white")}>Place Order</button>
        </div>
    </div>
  );
}

function useCartIndicator() {
  const [counter, setCounter] = useState<boolean>(false);
  const items = useStore(cartItem);

  useEffect(() => {
    const cachedCart = window.localStorage.getItem(CART_STORAGE);
    if (!cachedCart) return;

    const items = JSON.parse(cachedCart) as CartStorage;
    for (const [k, v] of Object.entries(items)) {
      cartItem.setKey(k, v);
    }
  }, []);

  useEffect(() => {
    const cartEntries = Object.keys(items);
    setCounter(() => cartEntries.length !== 0);
  }, [items]);

  return counter;
}