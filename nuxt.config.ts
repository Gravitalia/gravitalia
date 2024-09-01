import { isDevelopment } from "std-env";

export default defineNuxtConfig({
  app: {
    keepalive: true,
    head: {
      charset: "utf-8",
      viewport: "width=device-width,initial-scale=1",
      title: "Gravitalia",
      htmlAttrs: {
        lang: "en",
      },
      meta: [
        { property: "og:type", content: "website" },
        { property: "og:site_name", content: "Gravitalia" },
        { property: "og:title", content: "Gravitalia" },
        { property: "og:image", content: "/favicon.webp" },
        {
          name: "og:description",
          content:
            "Gravitalia, all connected. Share your photos in complete security and privacy.",
        },
        { name: "theme-color", content: "#8b5cf6" },
        { name: "robots", content: "index, follow" },
        { name: "twitter:card", content: "summary" },
        { name: "twitter:site", content: "@gravitalianews" },
        {
          name: "description",
          content:
            "Gravitalia, all connected. Share your photos in complete security and privacy.",
        },
      ],
      link: [{ rel: "manifest", href: "/manifest.json" }],
      script: [
        {
          innerHTML: !isDevelopment
            ? '"serviceWorker"in navigator&&navigator.serviceWorker.register("/sw.js",{scope:"/"});'
            : "",
        },
      ],
      bodyAttrs: {
        class: "dark:bg-zinc-900 dark:text-white",
      },
    },
  },

  ssr: false,
  components: true,
  spaLoadingTemplate: "pages/loading.html",
  sourcemap: isDevelopment,

  modules: [
    "@pinia/nuxt",
    [
      "@nuxtjs/color-mode",
      {
        preference: "system",
        fallback: "light",
        hid: "color-script",
        globalName: "__NUXT_COLOR_MODE__",
        componentName: "ColorScheme",
        classPrefix: "",
        classSuffix: "",
        storageKey: "mode",
      },
    ],
    [
      "@nuxtjs/i18n",
      {
        defaultLocale: "en",
        strategy: "no_prefix",
        lazy: true,
        langDir: "locales",
        detectBrowserLanguage: {
          useCookie: true,
          cookieKey: "locale",
          redirectOn: "root",
          fallbackLocale: "en",
          alwaysRedirect: true,
        },
        locales: [
          {
            code: "en",
            iso: "en-US",
            file: "en-US.json",
            name: "English",
          },
          {
            code: "fr",
            iso: "fr-FR",
            file: "fr-FR.json",
            name: "Français",
          },
        ],
        baseUrl: "https://www.gravitalia.com",
      },
    ],
    "@nuxtjs/tailwindcss",
    "@nuxt/image",
    "@nuxt/eslint",
  ],

  routeRules: {
    // No JS.
    "/": { experimentalNoScripts: false },
  },

  devtools: { enabled: isDevelopment },

  nitro: {
    preset: "cloudflare_pages",
  },

  vite: {
    define: {
      // By default, Vite doesn't include shims for NodeJS/
      // necessary for segment analytics lib to work
      global: {},
    },
  },

  compatibilityDate: "2024-09-01",
});
