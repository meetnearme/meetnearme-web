/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["templates/**/*.{html, go, templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["cupcake", "luxury"],
  },
};
