<script setup lang="ts">
import Cards from "./Cards.vue";
import Examples from "./Examples.vue";
import Ai from "./Ai.vue";
import Flag from "./Flag.vue";
</script>

<template>
	<div class="relative w-full flex flex-col items-center">
		<div class="reveal-section w-full">
			<Cards />
		</div>

		<div
			class="w-px h-32 bg-gradient-to-b from-transparent via-violet-500/30 to-transparent my-12"
		/>

		<div class="reveal-section w-full">
			<Examples />
		</div>

		<div class="reveal-section w-full py-24 relative">
			<div
				class="absolute top-0 left-1/2 -translate-x-1/2 w-1/3 h-px bg-gradient-to-r from-transparent via-violet-500/20 to-transparent"
			/>
			<Ai />
		</div>

		<div class="reveal-section w-full">
			<Flag />
		</div>
	</div>
</template>
