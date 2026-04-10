<script setup lang="ts">
import { ShieldCheckIcon, ServerStackIcon } from "@heroicons/vue/24/outline";
</script>

<template>
	<div class="relative w-full py-32 overflow-hidden border-t border-white/5">
		<div
			class="absolute inset-0 flex opacity-20 dark:opacity-10 pointer-events-none -z-20"
		>
			<div class="flex-1 bg-[#0055A4]" />
			<div class="flex-1 bg-white" />
			<div class="flex-1 bg-[#EF4135]" />
		</div>

		<div
			class="absolute inset-0 bg-zinc-50/80 dark:bg-zinc-950/80 backdrop-blur-xl -z-10"
		/>

		<div
			class="relative z-10 max-w-4xl mx-auto px-6 text-center flex flex-col items-center"
		>
			<h2
				class="text-3xl md:text-5xl font-bold text-zinc-900 dark:text-white mb-6"
			>
				{{ $t("home.flag.made_in") }}<br >
				<span class="text-violet-600 dark:text-violet-400">{{
					$t("home.flag.for_world")
				}}</span>
			</h2>
			<p class="text-lg text-zinc-600 dark:text-zinc-400 max-w-2xl mb-10">
				{{ $t("home.flag.desc") }}
			</p>

			<div class="grid grid-cols-1 md:grid-cols-2 gap-4 w-full max-w-2xl mb-12">
				<div
					class="flex items-center justify-center p-4 rounded-2xl border border-white/10 bg-white/5 backdrop-blur-sm text-zinc-800 dark:text-zinc-200"
				>
					<ShieldCheckIcon class="size-6 mr-3 text-violet-500" />
					<span class="font-medium">{{ $t("home.flag.gdpr") }}</span>
				</div>
				<div
					class="flex items-center justify-center p-4 rounded-2xl border border-white/10 bg-white/5 backdrop-blur-sm text-zinc-800 dark:text-zinc-200"
				>
					<ServerStackIcon class="size-6 mr-3 text-violet-500" />
					<span class="font-medium">{{ $t("home.flag.hosted") }}</span>
				</div>
			</div>

			<NuxtLink to="https://account.gravitalia.com/signup">
				<ButtonDefault class="text-lg px-8 py-4">
					{{ $t("home.flag.btn") }}
				</ButtonDefault>
			</NuxtLink>
		</div>
	</div>
</template>
