/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}"],
  theme: {
    extend: {
      gridTemplateColumns: {
        nav: "1fr 70% 1fr",
      },
      colors: {
        brand: {
          100: "#f19a3e",
          200: "#403233",
          accent: "#b9ffb7",
        },
      },
    },
  },
  plugins: [],
};
