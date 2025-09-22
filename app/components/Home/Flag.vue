<script setup lang="ts">
import {
	ChevronRightIcon,
	ShieldCheckIcon,
	ServerStackIcon,
} from "@heroicons/vue/24/outline";
</script>

<template>
	<div class="relative w-full h-[50vh] xl:h-[70vh] overflow-hidden">
		<div class="absolute inset-0 flex">
			<div class="flex-1 bg-[#0055A4]" />
			<div class="flex-1 bg-white" />
			<div class="flex-1 bg-[#EF4135]" />
		</div>

		<div class="absolute inset-0 bg-zinc-500/50 backdrop-blur-sm"/>

		<div class="absolute inset-0 flex flex-col items-center justify-center">
			<p class="text-lg md:text-2xl lg:text-3xl font-semibold text-zinc-800">
				{{ $t("home.flag.not_convinced") }}
			</p>
			<p class="text-lg md:text-2xl lg:text-3xl font-semibold text-zinc-800">
				{{ $t("home.flag.made_in_france") }}
			</p>
			<p class="text-zinc-800">
				{{ $t("home.flag.european_law") }}
			</p>
			<NuxtLink
				tabindex="0"
				to="https://account.gravitalia.com/signup"
				rel="noopener noreferrer"
				class="flex underline mt-6 xl:ml-1 text-md xl:text-lg text-zinc-800"
			>
				{{ $t("home.create_account") }}
				<ChevronRightIcon class="mt-1.5 size-4" />
			</NuxtLink>

			<hr class="h-px my-8 w-1/4 bg-zinc-400 border-0" >

			<div class="grid grid-cols-2 gap-6">
				<div class="flex text-zinc-800">
					<ShieldCheckIcon class="size-5 mr-2" />
					<p class="text-sm">
						{{ $t("home.flag.gdrp") }}
					</p>
				</div>
				<div class="flex text-zinc-800">
					<ServerStackIcon class="size-5 mr-2" />
					<p class="text-sm">{{ $t("home.flag.host") }}</p>
				</div>
			</div>
		</div>
	</div>
</template>
