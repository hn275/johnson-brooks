export interface Hangboard {
  id?: string;
  img?: string;
  material: string;
  title: string;
  description: string;
  inventory: number;
  price: number;
}

export interface MonoRail extends Hangboard {
  variants: Variant[];
}

export interface Variant {
  color: string;
  img: string;
  depth: Depth;
}

export enum Depth {
  sm = "12mm",
  md = "15mm",
  lg = "20mm",
  custom = "Custom",
}

export type Product = {
  id: string;
  title: string;
  thumbnail: string;
  description: string;
  material: string;
  price: number;
  inventory: number;
};

export type ProductVariant = {
  variant: string;
  thumbnail: string;
  color: string;
  inventory: number;
};
