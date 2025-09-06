<script setup lang="ts">
import {
	ShieldCheckIcon,
	FolderIcon,
	HomeIcon,
} from "@heroicons/vue/24/outline";

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
				class="cursor-pointer shadow-2xl bg-zinc-50 p-4 lg:p-6"
				gradient-color="#8b5cf650"
			>
				<h3 class="text-2xl text-zinc-800 flex pb-4 font-serif">
					<div class="size-8 bg-white border border-zinc-200 rounded mx-2">
						<component :is="card.icon" class="w-full p-1.5" />
					</div>
					{{ card.title }}
				</h3>

				<p class="text-zinc-600">{{ card.description }}</p>

				<NuxtImg v-show="card.image" :src="card.image" class="size-56" alt="" />
			</CardSpotlight>
		</div>
	</div>

	<div class="h-screen" />

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

	<div class="flex justify-center items-center py-56 relative">
		<NuxtLink tabindex="0" to="https://account.gravitalia.com/signup">
			<ButtonDefault class="px-6 py-3">
				{{ $t("home.create_account") }}
			</ButtonDefault>
		</NuxtLink>
	</div>
</template>
