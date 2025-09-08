<script setup lang="ts">
import {
	ShieldCheckIcon,
	FolderIcon,
	HomeIcon,
} from "@heroicons/vue/24/outline";
import Flag from "./Flag.vue";

const { t } = useI18n();

const cards = [
	{
		title: computed(() => t("home.cards.privacy.title")),
		icon: ShieldCheckIcon,
		description: computed(() => t("home.cards.privacy.description")),
		image: "/svg/secure-files.svg",
	},
	{
		title: computed(() => t("home.cards.oss.title")),
		icon: FolderIcon,
		description: computed(() => t("home.cards.oss.description")),
		image: "/svg/oss.svg",
	},
	{
		title: computed(() => t("home.cards.selfhost.title")),
		icon: HomeIcon,
		description: computed(() => t("home.cards.selfhost.description")),
		image: "/svg/selfhostable.svg",
	},
];
</script>

<template>
	<div
		class="grid lg:grid-cols-2 2xl:grid-cols-3 gap-6 p-6 px-42 2xl:px-80 lg:-my-40"
	>
		<div v-for="card in cards" :key="card.title.value" class="flex h-82 w-full">
			<CardSpotlight
				class="cursor-pointer shadow-2xl bg-zinc-50 dark:border-zinc-800 p-4"
				gradient-color="#8b5cf650"
			>
				<div class="flex justify-between items-center gap-x-4">
					<NuxtImg
						:src="card.image"
						class="mt-6 size-48 hidden lg:block"
						alt=""
						draggable="false"
					/>

					<div>
						<h3
							class="text-2xl text-zinc-800 dark:text-zinc-200 pb-4 font-serif"
						>
							{{ card.title }}
						</h3>

						<p class="text-zinc-500 dark:text-zinc-400">
							{{ card.description }}
						</p>
					</div>
				</div>
			</CardSpotlight>
		</div>
	</div>

	<div class="h-[50vh]" />

	<div
		class="h-screen z-10 flex min-h-64 items-center justify-center rounded-lg"
	>
		<AnimationTextScrollReveal :text="$t('home.scroll.data_master')" />
	</div>

	<div
		class="h-screen z-10 flex min-h-64 items-center justify-center rounded-lg"
	>
		<AnimationTextScrollReveal :text="$t('home.scroll.keep_secret')" />
	</div>

	<div class="h-[50vh]" />

	<Flag />
</template>
