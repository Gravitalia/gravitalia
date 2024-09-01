<script setup>
defineProps({
  text: {
    type: String,
    required: true,
  },
});
</script>

1
<template>
  <button
    class="group flex items-center justify-center overflow-hidden border border-white/10 px-6 py-3 text-white [background:var(--bg)] [border-radius:var(--radius)] dark:text-black transform-gpu transition-transform duration-300 ease-in-out active:translate-y-[1px] shadow-2xl"
    style="
      --spread: 90deg;
      --shimmer-color: #ffffff;
      --radius: 100px;
      --speed: 3s;
      --cut: 0.05em;
      --bg: rgba(0, 0, 0, 1);
    "
  >
    <div
      class="-z-30 blur-[2px] absolute inset-0 overflow-visible [container-type:size]"
    >
      <div
        class="absolute inset-0 h-[100cqh] animate-slide [aspect-ratio:1] [border-radius:0] [mask:none]"
      >
        <div
          class="animate-spin-around absolute inset-[-100%] w-auto rotate-0 [background:conic-gradient(from_calc(270deg-(var(--spread)*0.5)),transparent_0,var(--shimmer-color)_var(--spread),transparent_var(--spread))] [translate:0_0]"
        />
      </div>
    </div>
    <span
      class="text-center text-sm font-medium leading-none tracking-tight text-white dark:from-white dark:to-slate-900/10 lg:text-lg"
      >{{ $props.text }}</span
    >
    <div
      class="insert-0 absolute h-full w-full rounded-2xl px-4 py-1.5 text-sm font-medium shadow-[inset_0_-8px_10px_#ffffff1f] transform-gpu transition-all duration-300 ease-in-out group-hover:shadow-[inset_0_-6px_10px_#ffffff3f] group-active:shadow-[inset_0_-10px_10px_#ffffff3f]"
    />
    <div
      class="absolute -z-20 [background:var(--bg)] [border-radius:var(--radius)] [inset:var(--cut)]"
    />
  </button>
</template>
