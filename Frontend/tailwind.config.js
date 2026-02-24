/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}", // This ensures Vue files are scanned
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

