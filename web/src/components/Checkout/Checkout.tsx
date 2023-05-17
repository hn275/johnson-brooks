import cx from "classnames";
import { CART_STORAGE } from "../Cart/store"
import { useState, useEffect } from "react";

export function Checkout() {
  const [loading, setLoading] = useState(true);
  const [items, setItems] = useState<Object>({});
  const [price, setPrice] = useState<Number>(0);

  useEffect(()=> {
    const i = window.localStorage.getItem(CART_STORAGE);
    if (i) {
      const jsonItem = JSON.parse(i)
      setItems(jsonItem)
      Object.entries(jsonItem).map(([_, e])=> {
        setPrice((prev) => prev + (e as any).unitPrice)
      })
    }

    setLoading(false)
  }, [])

  if (loading) {
    return (
      <div>Loading...</div>
    )
  } else {
    return (
      <div className={cx("grid sm:px-10 lg:grid-cols-2 lg:px-20 xl:px-32 p-20")}>
        <div className={cx("px-4 pt-8")}>
            <p className={cx("text-xl font-medium")}>Order Summary</p>
            <p className={cx("text-gray-400")}>Check your items. And select a suitable shipping method.</p>
            <div className={cx("mt-8 space-y-3 rounded-lg border bg-white px-2 py-4 sm:px-6")}>
                {Object.entries(items).map(([_, item])=> (
                  <div className={cx("flex flex-col rounded-lg bg-white sm:flex-row")}>
                      <img className={cx("m-2 h-24 w-28 rounded-md border object-cover object-center")} src="https://images.unsplash.com/flagged/photo-1556637640-2c80d3201be8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8c25lYWtlcnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60" alt="" />
                      <div className={cx("flex w-full flex-col px-4 py-4")}>
                          <span className={cx("font-semibold")}>{item.productTitle}</span>
                          <span className={cx("float-right text-gray-400")}>Quantity:
                          <label htmlFor="years" className={cx("text-sm font-medium text-gray-900 dark:text-white")}>{item.quantity}</label>
                          <select id="years" className={cx("bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-10")}>
                            <option>1</option>
                            <option>2</option>
                            <option>3</option>
                            <option>4</option>
                          </select>
                          </span>
                          <p className={cx("text-lg font-bold")}>${item.unitPrice}</p>
                      </div>
                  </div>
                ))}
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
                <p className={cx("text-2xl font-semibold text-gray-900")}>{price.toString()}</p>
            </div>
            <button className={cx("mt-4 mb-8 w-full rounded-md bg-gray-900 px-6 py-3 font-medium text-white")}>Place Order</button>
        </div>
      </div>
    )
  }
}
