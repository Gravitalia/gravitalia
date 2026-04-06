<script setup lang="ts">
import {
	ShieldCheckIcon,
	FolderIcon,
	HomeIcon,
	UserCircleIcon,
} from "@heroicons/vue/24/outline";

const { t } = useI18n();
const cards = computed(() => [
	{
		title: t("home.cards.privacy.title"),
		icon: ShieldCheckIcon,
		description: t("home.cards.privacy.description"),
	},
	{
		title: t("home.cards.oss.title"),
		icon: FolderIcon,
		description: t("home.cards.oss.description"),
	},
	{
		title: t("home.cards.selfhost.title"),
		icon: HomeIcon,
		description: t("home.cards.selfhost.description"),
	},
	{
		title: t("home.cards.sovereign.title"),
		icon: UserCircleIcon,
		description: t("home.cards.sovereign.description"),
	},
]);
</script>

<template>
	<div
		class="grid grid-cols-1 md:grid-cols-2 gap-6 lg:gap-8 px-6 lg:px-12 max-w-7xl mx-auto py-24 relative z-20"
	>
		<CardSpotlight
			v-for="card in cards"
			:key="card.title"
			class="group p-8 md:p-10 flex flex-col gap-4 backdrop-blur-xl bg-white/40 dark:bg-zinc-900/40 border border-zinc-200/50 dark:border-white/10 transition-all duration-500 hover:-translate-y-1 hover:shadow-[0_20px_40px_-15px_rgba(139,92,246,0.2)]"
		>
			<div class="flex items-center gap-4 relative z-10 mb-2">
				<div
					class="p-3 shrink-0 rounded-xl bg-violet-100 dark:bg-violet-500/20 text-violet-600 dark:text-violet-400 border border-violet-200 dark:border-violet-500/30"
				>
					<component :is="card.icon" class="size-7" />
				</div>
				<h3
					class="text-2xl font-bold font-sans text-zinc-900 dark:text-white tracking-tight"
				>
					{{ card.title }}
				</h3>
			</div>

			<div class="relative z-10">
				<p class="text-zinc-600 dark:text-zinc-400 leading-relaxed">
					{{ card.description }}
				</p>
			</div>
		</CardSpotlight>
	</div>
</template>
