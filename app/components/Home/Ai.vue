<script setup lang="ts">
import { BoltIcon, ShieldCheckIcon } from "@heroicons/vue/24/outline";

const { t } = useI18n();
const features = computed(() => [
	{
		title: t("home.ai.features.cognitive.title"),
		description: t("home.ai.features.cognitive.desc"),
		icon: BoltIcon,
	},
	{
		title: t("home.ai.features.protection.title"),
		description: t("home.ai.features.protection.desc"),
		icon: ShieldCheckIcon,
	},
]);
</script>

<template>
	<div
		class="w-full max-w-5xl mx-auto px-6 py-24 flex flex-col items-center text-center"
	>
		<div
			class="absolute w-[600px] h-[600px] bg-violet-600/20 blur-[120px] rounded-full -z-10 pointer-events-none"
		/>

		<div
			class="p-4 rounded-3xl border border-white/10 bg-white/5 backdrop-blur-md mb-8 inline-block shadow-[0_0_30px_rgba(139,92,246,0.2)]"
		>
			<NuxtLink to="https://galadril.gravitalia.com" target="_blank">
				<NuxtImg
					src="/services/ai.webp"
					format="webp"
					quality="60"
					width="40"
					alt="Gravitalia AI logo"
					class="size-12"
				/>
			</NuxtLink>
		</div>

		<h2
			class="text-4xl md:text-6xl font-bold text-zinc-900 dark:text-white mb-6 tracking-tight"
		>
			{{ $t("home.ai.discover") }}
			<span
				class="text-transparent bg-clip-text bg-gradient-to-r from-violet-500 to-indigo-500"
				>{{ $t("home.ai.galadril") }}</span
			>
		</h2>
		<p class="text-xl text-zinc-600 dark:text-zinc-400 max-w-2xl mb-16">
			{{ $t("home.ai.desc") }}
		</p>

		<div class="grid md:grid-cols-2 gap-6 w-full">
			<CardSpotlight
				v-for="feat in features"
				:key="feat.title"
				class="p-8 text-left bg-zinc-50 dark:bg-zinc-900/40"
			>
				<component :is="feat.icon" class="size-10 text-violet-500 mb-6" />
				<h3 class="text-2xl font-semibold text-zinc-900 dark:text-white mb-3">
					{{ feat.title }}
				</h3>
				<p class="text-zinc-600 dark:text-zinc-400">{{ feat.description }}</p>
			</CardSpotlight>
		</div>
	</div>
</template>
