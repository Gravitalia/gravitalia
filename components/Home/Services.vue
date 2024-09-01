<template>
      <div
    class="relative flex h-screen w-full items-center justify-center overflow-hidden"
  >
    <div>
      <h1 class="font-bold text-4xl xl:text-6xl">
        <span class="block text-black dark:text-zinc-100">{{
          $t("home.discover")
        }}</span>
        <div class="h-20 overflow-hidden">
          <ul class="animate-change hover:[animation-play-state:paused]">
            <li class="leading-[5rem]">
              <span
                class="text-violet-600 dark:text-violet-500 text-4xl xl:text-6xl"
                >Gravitalia</span
              >
            </li>
            <li class="leading-[5rem]">
              <NuxtLink
                class="no-underline inline-block text-violet-600 dark:text-violet-500 text-4xl xl:text-6xl"
                to="https://account.gravitalia.com/"
              >
                Account</NuxtLink
              >
            </li>
            <li class="leading-[5rem]">
              <NuxtLink
                class="no-underline inline-block text-4xl xl:text-6xl"
                to="https://news.gravitalia.com/"
              >
                News</NuxtLink
              >
            </li>
            <li class="leading-[5rem]">
              <NuxtLink
                class="no-underline inline-block text-violet-600 dark:text-violet-500 text-4xl xl:text-6xl"
                to="https://turms.gravitalia.com/"
              >
                Turms</NuxtLink
              >
            </li>
            <!--
            <li class="leading-[5rem]">
              <span
                class="text-blue-600 dark:text-blue-500 text-4xl xl:text-6xl"
                >AirBlog</span
              >
            </li>
            -->
          </ul>
        </div>
      </h1>

      <p
        class="mt-6 xl:ml-1 text-lg xl:text-xl text-zinc-500 dark:text-zinc-400"
      >
        {{ $t("home.description") }}
      </p>

      <NuxtLink
        class="flex underline mt-6 xl:ml-1 text-md xl:text-lg text-zinc-500 dark:text-zinc-400"
        to="https://account.gravitalia.com/signup"
      >
        {{ $t("home.begin") }}

        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="mt-2 w-4 h-4"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="m8.25 4.5 7.5 7.5-7.5 7.5"
          />
        </svg>
      </NuxtLink>
    </div>
    <!-- Turms. -->
    <div
      class="hidden 2xl:block absolute flex transform-gpu animate-orbit items-center justify-center border [animation-delay:calc(var(--delay)*1000ms)] [animation-direction:reverse] border-none bg-transparent"
      style="--duration: 30; --radius: 400; --delay: -10"
    >
      <div
        class="size-8 bg-[length:100%]"
        style="background-image: url(/services/turms.webp)"
      />
    </div>
    <!-- Gravitalia. -->
    <div
      class="hidden xl:block absolute flex transform-gpu animate-orbit items-center justify-center border [animation-delay:calc(var(--delay)*1000ms)] [animation-direction:reverse] border-none bg-transparent"
      style="--duration: 20; --radius: 300; --delay: -10"
    >
      <div
        class="w-12 h-12 bg-[length:100%]"
        style="background-image: url(/favicon.webp)"
      />
    </div>
    <!-- Poios. -->
    <div
      class="hidden xl:block absolute flex transform-gpu animate-orbit items-center justify-center border [animation-delay:calc(var(--delay)*1000ms)] [animation-direction:reverse] border-none bg-transparent"
      style="--duration: 20; --radius: 300; --delay: -20"
    >
      <div
        class="w-24 h-24 bg-[length:100%] dark:hidden"
        style="background-image: url(/services/poios-black.webp)"
      />
      <!-- <div
        class="w-24 h-24 bg-[length:100%] dark:block"
        style="background-image: url(/services/poios-white.webp)"
      ></div> -->
    </div>
  </div>
  <HomeFont />
</template>