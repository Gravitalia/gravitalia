<script setup>
const fill = ref([0, 200]);
const opacity = ref(0);

onMounted(() => {
  const fillInterval = setInterval(() => {
    fill.value = [fill.value[0] + 4, fill.value[1] - 4];

    if (fill.value[0] >= 200) {
      clearInterval(fillInterval);

      const opacityInterval = setInterval(() => {
        opacity.value += 0.025;

        if (opacity.value >= 1) clearInterval(opacityInterval);
      }, 25);
    }
  }, 25);
});
</script>

<template>
  <div
    class="relative flex flex-col h-screen w-full items-center justify-center overflow-hidden"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlns:xlink="http://www.w3.org/1999/xlink"
      version="1.1"
      baseProfile="full"
      width="899"
      height="1011"
      class="select-none"
    >
      <rect width="899" height="1011" x="0" y="0" fill="none"/>
      <g>
        <text
          dominant-baseline="central"
          text-anchor="start"
          y="40"
          transform="translate(211.6875 425.5)"
          paint-order="stroke"
          :stroke-dasharray="fill.join(',')"
          stroke-miterlimit="2"
          class="fill-black stroke-black text-7xl font-bold font-[sans-serif]"
          stroke-dashoffset="146"
          :fill-opacity="opacity"
        >
          Plongez dans
        </text>
        <text
          dominant-baseline="central"
          text-anchor="start"
          y="120"
          transform="translate(211.6875 425.5)"
          paint-order="stroke"
          :stroke-dasharray="fill.join(',')"
          stroke-miterlimit="2"
          class="fill-black stroke-black text-7xl font-bold font-[sans-serif]"
          stroke-dashoffset="146"
          :fill-opacity="opacity"
        >
          Gravitalia !
        </text>
      </g>
    </svg>

    <div v-if="fill[0] === 200" class="absolute pt-96">
      <NuxtLink to="/">
        <ButtonShimmer text="Commencer mon aventure" />
      </NuxtLink>
    </div>
  </div>

  <div
    class="absolute pointer-events-none inset-0"
    style="
      background: linear-gradient(
        rgba(0, 0, 0, 0),
        50%,
        rgba(254, 243, 199, 0.25)
      );
    "
  />
</template>
