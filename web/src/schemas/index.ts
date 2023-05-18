export enum Depth {
  sm = "12mm",
  md = "15mm",
  lg = "20mm",
  custom = "Custom",
}

export type Product = {
  id: string;
  title: string;
  description: string;
  material: string;
  price: number;
  variants: ProductVariant[];
};

export type ProductVariant = {
  variant: string;
  thumbnail: string;
  color: string;
  inventory: number;
};
