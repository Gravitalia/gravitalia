<script setup lang="ts">
const currentYear = new Date().getFullYear();
</script>

<template>
	<footer
		class="relative w-full border-t border-zinc-200 dark:border-white/10 bg-white/50 dark:bg-zinc-950/50 backdrop-blur-xl"
	>
		<div class="max-w-7xl mx-auto px-6 py-12 md:py-16 lg:py-24">
			<div
				class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-x-6 gap-y-12 lg:gap-8"
			>
				<div
					class="col-span-2 md:col-span-4 lg:col-span-2 flex flex-col gap-4 lg:gap-6"
				>
					<div
						class="text-2xl font-bold font-sans tracking-tight text-zinc-900 dark:text-white flex items-center gap-2"
					>
						Gravitalia
					</div>
					<p class="text-zinc-500 dark:text-zinc-400 text-sm max-w-sm">
						© {{ currentYear }} Gravitalia.<br >
						{{ $t("footer.rights") }}
					</p>
				</div>

				<div class="col-span-1 flex flex-col gap-4">
					<h4 class="font-semibold text-zinc-900 dark:text-white">
						{{ $t("footer.products") }}
					</h4>
					<NuxtLink
						to="https://news.gravitalia.com"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors"
					>
						Gravitalia News
					</NuxtLink>
					<NuxtLink
						to="https://turms.gravitalia.com"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors"
					>
						Gravitalia Turms
					</NuxtLink>
					<NuxtLink
						to="https://galadril.gravitalia.com"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors"
					>
						Galadril Engine
					</NuxtLink>
					<NuxtLink
						to="https://account.gravitalia.com"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors"
					>
						Gravitalia Account
					</NuxtLink>
				</div>

				<div class="col-span-1 flex flex-col gap-4">
					<h4 class="font-semibold text-zinc-900 dark:text-white">
						{{ $t("footer.socials") }}
					</h4>
					<NuxtLink
						to="https://github.com/Gravitalia"
						target="_blank"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors flex items-center gap-2"
					>
						GitHub
					</NuxtLink>
					<NuxtLink
						to="https://discord.gg/4dcEwKj2KM"
						target="_blank"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors flex items-center gap-2"
					>
						Discord
					</NuxtLink>
					<NuxtLink
						to="https://www.youtube.com/@GravitaliaNews"
						target="_blank"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors flex items-center gap-2"
					>
						YouTube
					</NuxtLink>
					<a
						href="mailto:contact@gravitalia.com"
						class="mt-2 text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors underline decoration-zinc-300 dark:decoration-zinc-700 underline-offset-4 break-all"
					>
						contact@gravitalia.com
					</a>
				</div>

				<div class="col-span-2 md:col-span-2 lg:col-span-1 flex flex-col gap-4">
					<h4 class="font-semibold text-zinc-900 dark:text-white">
						{{ $t("footer.legal") }}
					</h4>
					<NuxtLink
						to="https://account.gravitalia.com/privacy.pdf"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors"
					>
						{{ $t("footer.privacy") }}
					</NuxtLink>
					<NuxtLink
						to="https://account.gravitalia.com/terms.pdf"
						class="text-sm text-zinc-500 hover:text-violet-600 dark:text-zinc-400 dark:hover:text-violet-400 transition-colors"
					>
						{{ $t("footer.terms") }}
					</NuxtLink>
				</div>
			</div>
		</div>
	</footer>
</template>
