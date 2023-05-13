import store from "nanostores";

export type CartItem = {
  productTitle: string;
  description: number;
  quantity: number;
  unitPrice: number;
};

export const cartItem = store.map<Record<string, CartItem>>({});

export function addToCart(id: string, item: CartItem) {
  return () => {
    const product = cartItem.get()[id];

    if (product) {
      const quantity = product.quantity + item.quantity;
      cartItem.setKey(id, { ...product, quantity });
      return;
    }

    cartItem.setKey(id, item);
  };
}
