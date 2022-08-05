/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    //    "./index.html",
    "./src/**/*.{svelte,html,js,ts}",
  ],
  theme: {
    colors: {
      base: "#002B5B",
      primary: "#2B4865",
      secondary: "#256D85",
      accent: "#8FE3CF",
      white: '#FFFFFF',
      gray: '#8492a6',
    },
    extend: {},
  },
  plugins: [],
};
