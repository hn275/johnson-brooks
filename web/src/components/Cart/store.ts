import store from "nanostores";

export type CartItem = {
  productTitle: string;
  description: number;
  quantity: number;
  unitPrice: number;
};

export const cartItem = store.map<Record<string, CartItem>>({});
