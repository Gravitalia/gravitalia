<script lang="ts" setup>
import { cn } from "~/utils/inspira";
type SafariMockupProps = {
	class?: string;
	sclass?: string;
};

const props = withDefaults(defineProps<SafariMockupProps>(), {
	class: "h-[618px] w-[418px]",
	sclass: "w-[740px] h-6",
});
</script>

<template>
	<div class="relative my-4">
		<div
			:class="
				cn(
					'relative z-1 mx-auto my-0 overflow-hidden rounded-[20px] border-2 border-[rgb(200,202,203)] px-[9px] pt-[9px] pb-[23px] [background:rgb(13,13,13)]',
					props.class,
				)
			"
		>
			<slot />
			<div
				class="absolute right-0 bottom-0 left-0 h-6 bg-linear-to-b from-[#272727] to-[#0d0d0d]"
			/>
		</div>

		<div
			class="-ml-8 absolute top-[11px] left-2/4 z-2 h-3 w-16 rounded-br rounded-bl bg-[rgb(13,13,13)]"
		/>
		<div
			:class="
				cn(
					'-mt-2.5 relative z-9 rounded-[2px_2px_12px_12px] border-[1px_2px_0px] border-[rgb(160,163,167)] border-solid shadow-[rgb(108,112,116)_0px_-2px_8px_0px_inset] [background:radial-gradient(circle,rgb(226,227,228)_85%,rgb(200,202,203)_100%)] [border-image:initial]',
					props.sclass,
				)
			"
		>
			<div
				class="absolute top-0 left-1/2 ml-[-60px] h-2.5 rounded-b-[10px] shadow-[inset_0_0_4px_2px_#babdbf]"
			/>
		</div>
		<div
			class="-bottom-0.5 absolute left-12 h-0.5 w-10 rounded-b-full bg-neutral-600"
		/>
		<div
			class="-bottom-0.5 absolute right-12 h-0.5 w-10 rounded-b-full bg-neutral-600"
		/>
	</div>
</template>
