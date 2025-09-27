<script setup lang="ts">
import Ai from "./Ai.vue";
import Cards from "./Cards.vue";
import Examples from "./Examples.vue";
import Flag from "./Flag.vue";
</script>

<template>
	<Cards />
	<div class="h-[35svh]" />
	<Examples />
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
	<Ai />
	<div class="h-[35svh]" />
	<Flag />
</template>
