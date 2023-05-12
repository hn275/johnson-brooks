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
