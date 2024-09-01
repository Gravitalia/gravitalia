<template>
    <div class="max-w-7xl mx-auto p-12 space-y-12">
    <div class="flex gap-x-4">
      <div
        class="mt-5 w-16 h-16 p-[0.1875rem] rounded-full flex justify-center items-center ring-1 ring-zinc-900/10 shadow"
      >
        <div class="rounded-full bg-violet-600 dark:bg-violet-500 p-3">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            class="size-8 fill-white"
          >
            <path
              fill-rule="evenodd"
              d="M12 1.5a5.25 5.25 0 0 0-5.25 5.25v3a3 3 0 0 0-3 3v6.75a3 3 0 0 0 3 3h10.5a3 3 0 0 0 3-3v-6.75a3 3 0 0 0-3-3v-3c0-2.9-2.35-5.25-5.25-5.25Zm3.75 8.25v-3a3.75 3.75 0 1 0-7.5 0v3h7.5Z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
      </div>
      <div>
        <h2 class="mt-4 font-semibold text-violet-600 dark:text-violet-500">
          {{ $t("home.privacy.title") }}
        </h2>
        <p class="max-w-2xl text-zinc-500 dark:text-zinc-400">
          {{ $t("home.privacy.description") }}
        </p>
      </div>
    </div>

    <div class="flex gap-x-4">
      <div
        class="mt-5 w-16 h-16 rounded-full flex justify-center items-center ring-1 ring-zinc-900/10 shadow"
      >
        <div class="rounded-full bg-amber-400 dark:bg-amber-500 p-3">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="size-8 fill-white"
          >
            <defs>
              <filter id="thicken" x="-1" y="-1" width="200%" height="200%">
                <feMorphology operator="dilate" radius="1" in="SourceGraphic" />
              </filter>
            </defs>
            <path
              filter="url(#thicken)"
              fill-rule="evenodd"
              d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5 6.75 6.75 0 0 0 0-13.5ZM2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5Z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
      </div>
      <div>
        <h2 class="mt-4 font-semibold text-amber-600 dark:text-amber-500">
          {{ $t("home.transparent.title") }}
        </h2>
        <p class="max-w-2xl text-zinc-500 dark:text-zinc-400">
          {{ $t("home.transparent.description") }}
        </p>
      </div>
    </div>

    <div class="flex gap-x-4">
      <div
        class="mt-5 w-16 h-16 rounded-full flex justify-center items-center ring-1 ring-zinc-900/10 shadow"
      >
        <div class="rounded-full bg-zinc-400 dark:bg-zinc-500 p-3">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-8 fill-white dark:stroke-black"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="m20.893 13.393-1.135-1.135a2.252 2.252 0 0 1-.421-.585l-1.08-2.16a.414.414 0 0 0-.663-.107.827.827 0 0 1-.812.21l-1.273-.363a.89.89 0 0 0-.738 1.595l.587.39c.59.395.674 1.23.172 1.732l-.2.2c-.212.212-.33.498-.33.796v.41c0 .409-.11.809-.32 1.158l-1.315 2.191a2.11 2.11 0 0 1-1.81 1.025 1.055 1.055 0 0 1-1.055-1.055v-1.172c0-.92-.56-1.747-1.414-2.089l-.655-.261a2.25 2.25 0 0 1-1.383-2.46l.007-.042a2.25 2.25 0 0 1 .29-.787l.09-.15a2.25 2.25 0 0 1 2.37-1.048l1.178.236a1.125 1.125 0 0 0 1.302-.795l.208-.73a1.125 1.125 0 0 0-.578-1.315l-.665-.332-.091.091a2.25 2.25 0 0 1-1.591.659h-.18c-.249 0-.487.1-.662.274a.931.931 0 0 1-1.458-1.137l1.411-2.353a2.25 2.25 0 0 0 .286-.76m11.928 9.869A9 9 0 0 0 8.965 3.525m11.928 9.868A9 9 0 1 1 8.965 3.525"
            />
          </svg>
        </div>
      </div>
      <div>
        <h2 class="mt-4 font-semibold text-zinc-600 dark:text-zinc-200">
          {{ $t("home.communal.title") }}
        </h2>
        <p class="max-w-2xl text-zinc-500 dark:text-zinc-400">
          {{ $t("home.communal.description") }}
        </p>
      </div>
    </div>
  </div>
</template>