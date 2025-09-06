<script setup lang="ts">
import { ChevronRightIcon } from "@heroicons/vue/24/outline";
</script>

<template>
	<div
		class="relative w-full h-[30vh] xl:h-[50vh] shadow-xl overflow-hidden rounded-md"
	>
		<div class="absolute inset-0 flex">
			<div class="flex-1 bg-[#0055A4]" />
			<div class="flex-1 bg-white" />
			<div class="flex-1 bg-[#EF4135]" />
		</div>

		<div class="absolute inset-0 flex flex-col items-center justify-center">
			<div class="bg-white px-6 py-3 rounded">
				<p
					class="text-lg md:text-2xl lg:text-3xl font-semibold dark:text-zinc-800"
				>
					{{ $t("home.flag.not_convinced") }}
				</p>
				<p
					class="text-lg md:text-2xl lg:text-3xl font-semibold dark:text-zinc-800"
				>
					{{ $t("home.flag.made_in_france") }}
				</p>
				<p
					class="text-lg md:text-xl lg:text-2xl font-semibold dark:text-zinc-800"
				>
					{{ $t("home.flag.european_law") }}
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
