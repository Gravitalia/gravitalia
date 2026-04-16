<script setup lang="ts">
import { ChevronRightIcon } from "@heroicons/vue/24/outline";
import Caroussel from "./Caroussel.vue";
</script>

<template>
	<div
		class="relative z-10 flex flex-col items-center w-full max-w-5xl mx-auto px-6 mt-20"
	>
		<h1
			class="text-5xl md:text-6xl lg:text-7xl tracking-tight text-center leading-[1.1] text-zinc-900 dark:text-white break-words w-full"
		>
			<span class="font-extrabold font-sans block mb-2">{{
				$t("home.slogan_main")
			}}</span>
			<span
				class="font-serif italic font-medium text-transparent bg-clip-text bg-gradient-to-r from-violet-600 to-indigo-500 dark:from-violet-400 dark:to-indigo-300"
			>
				{{ $t("home.slogan_italic") }}
			</span>
		</h1>

		<div class="w-full overflow-hidden">
			<Caroussel class="mt-10" />
		</div>

		<p
			class="mt-8 text-lg md:text-xl text-zinc-600 dark:text-zinc-400 max-w-2xl text-center font-light leading-relaxed"
		>
			{{ $t("home.title_desc") }}
		</p>

		<div
			class="mt-12 flex flex-col sm:flex-row gap-4 items-center w-full justify-center"
		>
			<NuxtLink
				tabindex="0"
				to="https://account.gravitalia.com/signup"
				class="group relative inline-flex items-center justify-center gap-2 px-8 py-4 rounded-full bg-zinc-900 dark:bg-white text-white dark:text-zinc-950 font-semibold transition-all duration-300 hover:scale-105 hover:shadow-[0_0_40px_rgba(139,92,246,0.4)] w-full sm:w-auto"
			>
				{{ $t("home.start_by_create_account") }}
				<ChevronRightIcon
					class="size-4 transition-transform group-hover:translate-x-1"
				/>
			</NuxtLink>
			<!-- <ButtonDefault class="px-8 py-4 rounded-full">
				{{ $t("home.discover_projects") }}
			</ButtonDefault> -->
		</div>
	</div>
</template>
