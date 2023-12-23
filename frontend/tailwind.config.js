/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./dist/index.html", // build output location
    "../internal/templates/**/*.templ", // templ
    "./src/**/*.{html,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
