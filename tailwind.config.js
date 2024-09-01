/** @type {import("tailwindcss").Config} */
module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./app.vue",
    "./error.vue",
    "./nuxt.config.{js,ts}",
  ],
  theme: {
    extend: {
      keyframes: {
        change: {
          "0%, 12.66%, 100%": { transform: "translate3d(0,0,0)" },
          "16.66%, 29.32%": { transform: "translate3d(0,-25%,0)" },
          "33.32%, 45.98%": { transform: "translate3d(0,-50%,0)" },
          "49.98%, 62.64%": { transform: "translate3d(0,-75%,0)" },
          "66.64%, 79.3%": { transform: "translate3d(0,-50%,0)" },
          "83.3%, 95.96%": { transform: "translate3d(0,-25%,0)" },
        },
        shimmer: {
          "0%, 90%, 100%": {
            "background-position": "calc(-100% - var(--shimmer-width)) 0",
          },
          "30%, 60%": {
            "background-position": "calc(100% + var(--shimmer-width)) 0",
          },
        },
        orbit: {
          "0%": {
            transform:
              "rotate(0deg) translateY(calc(var(--radius) * 1px)) rotate(0deg)",
          },
          "100%": {
            transform:
              "rotate(360deg) translateY(calc(var(--radius) * 1px)) rotate(-360deg)",
          },
        },
        gradient: {
          to: {
            backgroundPosition: "var(--bg-size) 0",
          },
        },
        "spin-around": {
          "0%": {
            transform: "translateZ(0) rotate(0)",
          },
          "15%, 35%": {
            transform: "translateZ(0) rotate(90deg)",
          },
          "65%, 85%": {
            transform: "translateZ(0) rotate(270deg)",
          },
          "100%": {
            transform: "translateZ(0) rotate(360deg)",
          },
        },
        slide: {
          to: {
            transform: "translate(calc(100cqw - 100%), 0)",
          },
        },
        slideDown: {
          "0%": { opacity: "0", height: "0", padding: "0" },
          "100%": { opacity: "1", height: "auto" },
        },
        showContent: {
          "0%": { opacity: "0", height: "0" },
          "100%": { opacity: "1", height: "auto" },
        },
      },
      animation: {
        change: "change 16s infinite",
        shimmer: "shimmer 8s infinite",
        orbit: "orbit calc(var(--duration)*1s) linear infinite",
        gradient: "gradient 8s linear infinite",
        "spin-around": "spin-around calc(var(--speed) * 2) infinite linear",
        slide: "slide var(--speed) ease-in-out infinite alternate",
        "slide-down": "slideDown 0.3s ease-in-out",
        "show-content": "showContent 0.6s 0.2s forwards",
      },
    },
  },
  plugins: [],
};
