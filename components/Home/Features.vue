<template>
  <div
    class="relative flex flex-col w-full h-[115vh] xl:h-screen items-center justify-center"
  >
    <!-- Title. -->
    <p
      class="p-12 text-center text-4xl xl:text-5xl font-bold text-black dark:text-white"
    >
      {{ $t("home.is_also") }}
    </p>

    <!-- List of features. -->
    <div class="grid gap-12 xl:gap-24 grid-cols-1 xl:grid-cols-3">
      <div class="w-72 space-y-4">
        <div class="flex space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6 stroke-zinc-500 dark:stroke-zinc-300"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 0 1 3 19.875v-6.75ZM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 0 1-1.125-1.125V8.625ZM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 0 1-1.125-1.125V4.125Z"
            />
          </svg>

          <p class="text-zinc-600 dark:text-zinc-300">
            {{ $t("home.analysis_tools.title") }}
          </p>
        </div>

        <p class="text-lg font-semibold max-w-64">
          {{ $t("home.analysis_tools.description") }}
        </p>
      </div>

      <div class="w-72 space-y-4">
        <div class="flex space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6 stroke-zinc-500 dark:stroke-zinc-300"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48.432.447.74 1.04.586 1.641a4.483 4.483 0 0 1-.923 1.785A5.969 5.969 0 0 0 6 21c1.282 0 2.47-.402 3.445-1.087.81.22 1.668.337 2.555.337Z"
            />
          </svg>

          <p class="text-zinc-600 dark:text-zinc-300">
            {{ $t("home.private_conversations.title") }}
          </p>
        </div>

        <p class="text-lg font-semibold max-w-64">
          {{ $t("home.private_conversations.description") }}
        </p>
      </div>

      <div class="w-72 space-y-4">
        <div class="flex space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6 stroke-zinc-500 dark:stroke-zinc-300"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M2.25 7.125C2.25 6.504 2.754 6 3.375 6h6c.621 0 1.125.504 1.125 1.125v3.75c0 .621-.504 1.125-1.125 1.125h-6a1.125 1.125 0 0 1-1.125-1.125v-3.75ZM14.25 8.625c0-.621.504-1.125 1.125-1.125h5.25c.621 0 1.125.504 1.125 1.125v8.25c0 .621-.504 1.125-1.125 1.125h-5.25a1.125 1.125 0 0 1-1.125-1.125v-8.25ZM3.75 16.125c0-.621.504-1.125 1.125-1.125h5.25c.621 0 1.125.504 1.125 1.125v2.25c0 .621-.504 1.125-1.125 1.125h-5.25a1.125 1.125 0 0 1-1.125-1.125v-2.25Z"
            />
          </svg>

          <p class="text-zinc-600 dark:text-zinc-300">
            {{ $t("home.api.title") }}
          </p>
        </div>

        <p class="text-lg font-semibold max-w-64">
          {{ $t("home.api.description") }}
        </p>
      </div>

      <div class="hidden xl:block w-72 space-y-4">
        <div class="flex space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6 stroke-zinc-500 dark:stroke-zinc-300"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M15 19.128a9.38 9.38 0 0 0 2.625.372 9.337 9.337 0 0 0 4.121-.952 4.125 4.125 0 0 0-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.106A12.318 12.318 0 0 1 8.624 21c-2.331 0-4.512-.645-6.374-1.766l-.001-.109a6.375 6.375 0 0 1 11.964-3.07M12 6.375a3.375 3.375 0 1 1-6.75 0 3.375 3.375 0 0 1 6.75 0Zm8.25 2.25a2.625 2.625 0 1 1-5.25 0 2.625 2.625 0 0 1 5.25 0Z"
            />
          </svg>

          <p class="text-zinc-600 dark:text-zinc-300">
            {{ $t("home.user_security.title") }}
          </p>
        </div>

        <p class="text-lg font-semibold max-w-64">
          {{ $t("home.user_security.description") }}
        </p>
      </div>
      <div class="hidden xl:block w-72 space-y-4">
        <div class="flex space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6 stroke-zinc-500 dark:stroke-zinc-300"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M9.75 3.104v5.714a2.25 2.25 0 0 1-.659 1.591L5 14.5M9.75 3.104c-.251.023-.501.05-.75.082m.75-.082a24.301 24.301 0 0 1 4.5 0m0 0v5.714c0 .597.237 1.17.659 1.591L19.8 15.3M14.25 3.104c.251.023.501.05.75.082M19.8 15.3l-1.57.393A9.065 9.065 0 0 1 12 15a9.065 9.065 0 0 0-6.23-.693L5 14.5m14.8.8 1.402 1.402c1.232 1.232.65 3.318-1.067 3.611A48.309 48.309 0 0 1 12 21c-2.773 0-5.491-.235-8.135-.687-1.718-.293-2.3-2.379-1.067-3.61L5 14.5"
            />
          </svg>

          <p class="text-zinc-600 dark:text-zinc-300">
            {{ $t("home.perpetual_development.title") }}
          </p>
        </div>

        <p class="text-lg font-semibold max-w-64">
          {{ $t("home.perpetual_development.description") }}
        </p>
      </div>

      <div class="hidden xl:block space-y-4">
        <div class="flex space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6 stroke-zinc-500 dark:stroke-zinc-300"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M12 18v-5.25m0 0a6.01 6.01 0 0 0 1.5-.189m-1.5.189a6.01 6.01 0 0 1-1.5-.189m3.75 7.478a12.06 12.06 0 0 1-4.5 0m3.75 2.383a14.406 14.406 0 0 1-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 1 0-7.517 0c.85.493 1.509 1.333 1.509 2.316V18"
            />
          </svg>

          <p class="text-zinc-600 dark:text-zinc-300">
            {{ $t("home.philosophy.title") }}
          </p>
        </div>

        <p class="text-lg font-semibold max-w-64">
          {{ $t("home.philosophy.description") }}
        </p>
      </div>
    </div>

    <!-- Grid background. -->
    <svg
      aria-hidden="true"
      class="pointer-events-none absolute inset-0 w-full h-full fill-gray-400/30 stroke-gray-400/30"
      xmlns="http://www.w3.org/2000/svg"
    >
      <defs>
        <pattern
          id="pattern"
          width="80"
          height="80"
          patternUnits="userSpaceOnUse"
          x="-1"
          y="-1"
        >
          <path d="M.5 80V.5H80" fill="none" stroke-dasharray="0"/>
        </pattern>

        <mask
          id="mask"
          maskUnits="userSpaceOnUse"
          x="0"
          y="0"
          width="100%"
          height="100%"
        >
          <rect width="100%" height="100%" fill="url(#gradient)" />
        </mask>

        <linearGradient id="gradient" x1="0" y1="0" x2="0" y2="1">
          <stop offset="0%" style="stop-color: white; stop-opacity: 0" />
          <stop offset="25%" style="stop-color: white; stop-opacity: 0.9" />
          <stop offset="75%" style="stop-color: white; stop-opacity: 0.9" />
          <stop offset="100%" style="stop-color: white; stop-opacity: 0" />
        </linearGradient>
      </defs>

      <rect
        width="100%"
        height="100%"
        stroke-width="0"
        fill="url(#pattern)"
        mask="url(#mask)"
      />
    </svg>
  </div>
</template>
