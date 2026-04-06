<template>
	<div
		:class="[
			'group relative flex size-full overflow-hidden rounded-2xl border border-zinc-200/50 dark:border-white/10 backdrop-blur-md bg-white/40 dark:bg-zinc-900/40 text-zinc-900 dark:text-zinc-100 transition-all duration-500 hover:border-violet-500/30 hover:shadow-[0_8px_30px_rgba(139,92,246,0.1)]',
			$props.class,
		]"
		@mousemove="handleMouseMove"
		@mouseleave="handleMouseLeave"
	>
		<div :class="cn('relative z-10 w-full', props.slotClass)">
			<slot />
		</div>
		<div
			class="pointer-events-none absolute inset-0 rounded-2xl opacity-0 transition-opacity duration-500 group-hover:opacity-100 mix-blend-screen"
			:style="{
				background: backgroundStyle,
				opacity: gradientOpacity,
			}"
		></div>
	</div>
</template>

<script setup lang="ts">
import { type HTMLAttributes } from "vue";
import { cn } from "~/utils/inspira";

const props = withDefaults(
	defineProps<{
		class?: HTMLAttributes["class"];
		slotClass?: HTMLAttributes["class"];
		gradientSize?: number;
		gradientColor?: string;
		gradientOpacity?: number;
	}>(),
	{
		class: "",
		slotClass: "",
		gradientSize: 250,
		gradientColor: "rgba(139, 92, 246, 0.15)",
		gradientOpacity: 1,
	},
);

const mouseX = ref(-props.gradientSize * 10);
const mouseY = ref(-props.gradientSize * 10);

function handleMouseMove(e: MouseEvent) {
	const target = e.currentTarget as HTMLElement;
	const rect = target.getBoundingClientRect();
	mouseX.value = e.clientX - rect.left;
	mouseY.value = e.clientY - rect.top;
}

function handleMouseLeave() {
	mouseX.value = -props.gradientSize * 10;
	mouseY.value = -props.gradientSize * 10;
}

onMounted(() => {
	mouseX.value = -props.gradientSize * 10;
	mouseY.value = -props.gradientSize * 10;
});

const backgroundStyle = computed(() => {
	return `radial-gradient(
      circle at ${mouseX.value}px ${mouseY.value}px,
      ${props.gradientColor} 0%,
      transparent 70%
    )`;
});
</script>
