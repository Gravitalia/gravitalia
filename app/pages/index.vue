<script setup lang="ts">
import HomeBackground from "~/components/Home/Background.vue";
import HomeTitle from "~/components/Home/Title.vue";
import HomeContent from "~/components/Home/Content.vue";
import HomeMouse from "~/components/Home/Mouse.vue";
</script>

<template>
	<div
		class="min-h-screen bg-zinc-50 dark:bg-zinc-950 text-zinc-900 dark:text-zinc-50 font-sans overflow-hidden selection:bg-violet-500/30"
	>
		<HomeBackground class="hero-section">
			<div
				class="relative z-10 flex w-full flex-col items-center justify-center px-4"
			>
				<HomeTitle />
			</div>
			<div class="absolute bottom-10">
				<HomeMouse />
			</div>
		</HomeBackground>

		<main class="relative z-10 bg-zinc-50 dark:bg-zinc-950">
			<HomeContent />
		</main>
	</div>
</template>
