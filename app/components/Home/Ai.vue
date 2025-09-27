<script setup lang="ts">
const { t } = useI18n();

const cards = [
	{
		description: computed(() => t("home.ai.on_device")),
		image: "/svg/ai.svg",
	},
	{
		description: computed(() => t("home.ai.protected")),
		image: "/svg/email.svg",
	},
];
</script>

<template>
	<div class="flex flex-col justify-center items-center">
		<div class="border border-zinc-200 dark:border-zinc-800 rounded">
			<NuxtLink to="https://ai.gravitalia.com" target="_blank">
				<NuxtImg
					src="/services/ai.png"
					format="webp"
					quality="70"
					width="100"
					alt="Gravitalia AI logo"
					class="size-12"
				/>
			</NuxtLink>
		</div>

		<AnimationBlurReveal class="mt-6 text-xl font-bold xl:text-3xl md:text-2xl">
			{{ $t("home.ai.that_why") }}
		</AnimationBlurReveal>
		<AnimationBlurReveal
			class="text-center text-2xl font-bold xl:text-5xl p-4 lg:text-4xl md:p-8 md:text-3xl"
		>
			{{ $t("home.ai.app_ai") }}
		</AnimationBlurReveal>

		<AnimationBlurReveal>
			<div class="mt-12 grid lg:grid-cols-2 gap-6 p-6 md:px-32 2xl:px-80">
				<div v-for="card in cards" key="" class="flex h-82 w-full">
					<CardSpotlight
						class="cursor-pointer shadow-2xl bg-zinc-50 dark:border-zinc-800 p-4"
						gradient-color="#8b5cf650"
					>
						<div class="flex justify-between items-center gap-x-4">
							<NuxtImg
								:src="card.image"
								class="mt-6 size-36 hidden lg:block"
								alt=""
								draggable="false"
							/>

							<p class="text-zinc-500 dark:text-zinc-400">
								{{ card.description }}
							</p>
						</div>
					</CardSpotlight>
				</div>
			</div>
		</AnimationBlurReveal>
	</div>
</template>
