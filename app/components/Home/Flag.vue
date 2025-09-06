<script setup lang="ts">
import { ChevronRightIcon } from "@heroicons/vue/24/outline";
</script>

<template>
	<div
		class="relative w-full h-[20vh] xl:h-[50vh] shadow-xl overflow-hidden rounded-md"
	>
		<!-- Trois bandes verticales : bleu, blanc, rouge -->
		<div class="absolute inset-0 flex">
			<div class="flex-1 bg-[#0055A4]"/>
			<div class="flex-1 bg-white"/>
			<div class="flex-1 bg-[#EF4135]"/>
		</div>

		<div class="absolute inset-0 flex flex-col items-center justify-center">
			<div class="bg-white px-6 py-3 rounded">
				<p
					class="text-lg md:text-2xl lg:text-3xl font-semibold dark:text-zinc-800"
				>
					Pas encore convaincu ?
				</p>
				<p
					class="text-lg md:text-2xl lg:text-3xl font-semibold dark:text-zinc-800"
				>
					Gravitalia est développé en France 🇫🇷
				</p>
				<p
					class="text-lg md:text-xl lg:text-2xl font-semibold dark:text-zinc-800"
				>
					et protégé par le droit européen.
				</p>
			</div>
			<NuxtLink
				tabindex="0"
				to="https://account.gravitalia.com/signup"
				rel="noopener noreferrer"
				class="flex underline mt-6 xl:ml-1 text-md xl:text-lg dark:text-zinc-800"
			>
				{{ $t("home.create_account") }}
				<ChevronRightIcon class="mt-1.5 size-4"
			/></NuxtLink>
		</div>
	</div>
</template>
