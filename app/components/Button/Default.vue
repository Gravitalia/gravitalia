<template>
	<button
		type="button"
		tabindex="0"
		class="relative inline-flex items-center justify-center px-6 py-2.5 text-sm font-medium tracking-wide transition-all duration-300 rounded-full cursor-pointer bg-white/5 dark:bg-zinc-900/50 backdrop-blur-md border border-zinc-200 dark:border-white/10 text-zinc-800 dark:text-zinc-200 enabled:hover:bg-white/10 enabled:dark:hover:bg-white/10 enabled:hover:border-violet-500/50 enabled:hover:shadow-[0_0_20px_rgba(139,92,246,0.3)] enabled:hover:text-violet-600 enabled:dark:hover:text-violet-400 focus:outline-none focus:ring-2 focus:ring-violet-500/50 disabled:opacity-50 disabled:cursor-not-allowed"
	>
		<slot />
	</button>
</template>
