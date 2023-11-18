/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pkg/**/*.{html,js,go,ts}", "./client/**/*.{html,js,go,ts}"],
  theme: {
    extend: {},
    colors: {
      "navy": "#1f2335",
      "sky": "#7aa2f7",
      "cyan": "#2bc3df",
      "cyanLight": "#70e9ffff",
      "orange": "#e06b43ff",
      "navyLight": "#282d40",
      "white": "#ffffff",
      "green": "#8fba65",
      "lightGray": "#a6afd4",
      "red": "#eb7289",
      "yellow": "#e0af67",
    },
  },
  plugins: [],
}