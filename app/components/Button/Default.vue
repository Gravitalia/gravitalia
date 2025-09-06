<template>
	<button
		type="button"
		tabindex="0"
		class="flex justify-center text-sm py-2 cursor-pointer rounded border border-zinc-200 dark:border-zinc-800 bg-zinc-900 enabled:hover:bg-zinc-700 enabled:dark:hover:bg-zinc-900 font-medium dark:text-white text-white transition-all duration-300 disabled:cursor-not-allowed"
	>
		<slot />
	</button>
</template>
