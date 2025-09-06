import { isDevelopment } from "std-env";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: "2025-07-15",
	devtools: { enabled: true },

	modules: [
		"@nuxt/eslint",
		"@nuxt/image",
		"@nuxtjs/color-mode",
		"@nuxtjs/i18n",
		"@nuxtjs/tailwindcss",
		"motion-v/nuxt",
		...(isDevelopment ? [] : ["nuxt-security"]),
	],

	ssr: true,
	components: true,
	sourcemap: isDevelopment,

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
				// { name: "theme-color", content: "#8b5cf6" },
				{ name: "robots", content: "index, follow" },
				{ name: "twitter:card", content: "summary" },
				{ name: "twitter:site", content: "@gravitalianews" },
				{
					name: "description",
					content:
						"Gravitalia, all connected. Share your photos in complete security and privacy.",
				},
			],
			bodyAttrs: {
				class: "bg-zinc-50 dark:bg-zinc-900 dark:text-white",
			},
		},
	},

	i18n: {
		defaultLocale: "en",
		strategy: "prefix_except_default",
		langDir: ".",
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
				language: "en-US",
				file: "en-US.json",
				name: "English",
			},
			{
				code: "fr",
				language: "fr-FR",
				file: "fr-FR.json",
				name: "Français",
			},
		],
		baseUrl: "https://www.gravitalia.com",
	},

	colorMode: {
		preference: "system",
		classPrefix: "",
		classSuffix: "",
	},

	sri: true,
	security: {
		headers: {
			crossOriginEmbedderPolicy: "credentialless",
			crossOriginOpenerPolicy: "same-origin",
			crossOriginResourcePolicy: "same-site",
			originAgentCluster: "?1",
			referrerPolicy: "no-referrer",
			strictTransportSecurity: {
				maxAge: 63072000, // 2 years
				includeSubdomains: true,
				preload: true,
			},
			xFrameOptions: "DENY", // also managed by CSP.
			contentSecurityPolicy: {
				"default-src": ["'self'"],
				"font-src": ["'none'"],
				"form-action": ["'none'"],
				"frame-ancestors": ["'none'"],
				"frame-src": ["'none'"],
				"script-src-attr": ["'none'"],
				"object-src": ["'none'"],
				"connect-src": ["'self'"],
				"img-src": ["'self'"],
				"media-src": ["'self'"],
				"script-src": ["'self'", "'strict-dynamic'", "'nonce-{{nonce}}'"],
				"style-src": ["'self'", "'unsafe-inline'"],
				"upgrade-insecure-requests": false,
			},
			permissionsPolicy: {
				camera: [],
				geolocation: [],
				microphone: [],
			},
		},
		corsHandler: {
			origin: undefined,
			methods: "OPTIONS, GET",
			allowHeaders: "Content-Type, Accept",
			credentials: false,
			maxAge: "86400",
			preflight: {
				statusCode: 200,
			},
		},
		allowedMethodsRestricter: {
			methods: ["OPTIONS", "GET"],
		},
		rateLimiter: false,
	},

	nitro: {
		preset: isDevelopment ? "node-server" : "cloudflare_module",
		prerender: {
			autoSubfolderIndex: !isDevelopment,
		},
	},
});
