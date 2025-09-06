<script setup lang="ts">
import { ChevronRightIcon } from "@heroicons/vue/24/outline";
import Caroussel from "./Caroussel.vue";
</script>

<template>
	<div class="absolute">
		<h1 class="text-2xl xl:text-4xl font-bold">
			{{ $t("home.slogan") }}
		</h1>
		<Caroussel />
		<NuxtLink
			tabindex="0"
			to="https://account.gravitalia.com/signup"
			rel="noopener noreferrer"
			class="flex underline mt-6 xl:ml-1 text-md xl:text-lg text-zinc-500 dark:text-zinc-400"
		>
			{{ $t("home.start_by_create_account") }}
			<ChevronRightIcon class="mt-1.5 size-4"
		/></NuxtLink>
	</div>
</template>
