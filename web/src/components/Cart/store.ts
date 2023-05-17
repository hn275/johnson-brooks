import { map } from "nanostores";

export type CartItem = {
  productTitle: string;
  description: string;
  quantity: number;
  unitPrice: number;
};

export const CART_STORAGE = "cart_items";

export interface CartStorage {
  [key: string]: CartItem;
}

export const cartItem = map<Record<string, CartItem>>({});

export function addToCart(id: string, item: CartItem, cb: () => void) {
  return () => {
    const product = cartItem.get()[id];

    if (product) {
      const quantity = product.quantity + item.quantity;
      cartItem.setKey(id, { ...product, quantity });
    } else {
      cartItem.setKey(id, item);
    }

    window.localStorage.setItem(CART_STORAGE, JSON.stringify(cartItem.get()));
    cb();
  };
}