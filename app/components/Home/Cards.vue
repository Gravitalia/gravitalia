<script setup lang="ts">
import {
	ShieldCheckIcon,
	FolderIcon,
	HomeIcon,
    UserCircleIcon,
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
    {
		title: computed(() => t("home.cards.sovereign.title")),
		icon: UserCircleIcon,
		description: computed(() => t("home.cards.sovereign.description")),
		image: "/svg/user.svg",
	},
];
</script>

<template>
	<div
		class="grid lg:grid-cols-2 2xl:grid-cols-4 gap-6 p-6 lg:-my-40"
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
							class="flex text-2xl text-zinc-800 dark:text-zinc-200 pb-4 font-serif"
						>
                            <component :is="card.icon" class="lg:hidden mt-1 mr-2 size-6" />
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
</template>
