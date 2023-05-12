import { Dispatch, ReactNode, createContext, useReducer } from "react";

type ProductCart = {
  id: string;
  quantity: number;
};

interface CartRemove {
  type: "cart/add";
  payload: ProductCart;
}

interface CartAdd {
  type: "cart/rm";
  payload: ProductCart;
}

type CartDispatch = CartAdd | CartRemove;
type CartState = ProductCart[];
type CartContextType = {
  state: CartState;
  dispatch: Dispatch<CartDispatch>;
};

function reducer(state: CartState, action: CartDispatch) {
  switch (action.type) {
    case "cart/add":
      const ids = state.map((product) => product.id);
      if (ids.includes(action.payload.id)) {
        return state.map((prod) => {
          if (prod.id === action.payload.id) {
            return {
              id: prod.id,
              quantity: prod.quantity + action.payload.quantity,
            };
          }
          return prod;
        });
      }

      return [...state, action.payload];
    case "cart/rm":
      return state;
    default:
      return state;
  }
}

const initialState: CartState = [];
export const CartContext = createContext<CartContextType>(
  {} as CartContextType,
);

interface Props {
  children: ReactNode;
}

export function CartContextProvider({ children }: Props) {
  const [state, dispatch] = useReducer(reducer, initialState);
  return (
    <CartContext.Provider value={{ state, dispatch }}>
      {children}
    </CartContext.Provider>
  );
}
