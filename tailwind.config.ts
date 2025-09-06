import type { Config } from "tailwindcss";

export default <Partial<Config>>{
	content: [
		"./app/components/**/*.{js,vue,ts}",
		"./app/layouts/**/*.vue",
		"./app/pages/**/*.vue",
		"./app/plugins/**/*.{js,ts}",
		"./app/app.vue",
		"./nuxt.config.{js,ts}",
	],
	theme: {
		extend: {
			animation: {
				"animate-aurora": "aurora 60s linear infinite",
				change: "change 16s infinite",
				scroll: "scroll 2.2s cubic-bezier(.15,.41,.69,.94) infinite",
			},
			keyframes: {
				aurora: {
					from: {
						"background-position": "50% 50%, 50% 50%",
					},
					to: {
						"background-position": "350% 50%, 350% 50%",
					},
				},
				change: {
					"0%, 12.66%, 100%": { transform: "translate3d(0,0,0)" },
					"16.66%, 29.32%": { transform: "translate3d(0,-25%,0)" },
					"33.32%, 45.98%": { transform: "translate3d(0,-50%,0)" },
					"49.98%, 62.64%": { transform: "translate3d(0,-75%,0)" },
					"66.64%, 79.3%": { transform: "translate3d(0,-50%,0)" },
					"83.3%, 95.96%": { transform: "translate3d(0,-25%,0)" },
				},
				scroll: {
					"0%": { opacity: "0" },
					"10%": { transform: "translateY(0)", opacity: "1" },
					"100%": { transform: "translateY(15px)", opacity: "0" },
				},
			},
		},
	},
};
