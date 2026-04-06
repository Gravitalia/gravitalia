<template>
	<div
		class="flex items-center justify-center text-2xl lg:text-4xl font-semibold dark:text-zinc-300"
	>
		<span class="text-zinc-500 dark:text-zinc-500 mr-3">{{
			$t("home.explore")
		}}</span>
		<div
			class="h-[1.5em] overflow-hidden relative w-[260px] lg:w-[350px] text-left"
		>
			<ul
				class="animate-change hover:[animation-play-state:paused] flex flex-col absolute w-full text-violet-600 dark:text-violet-400"
			>
				<li class="h-[1.5em] leading-[1.5em] whitespace-nowrap">
					<NuxtLink
						class="hover:underline"
						to="https://turms.gravitalia.com/"
						>{{ $t("services.turms") }}</NuxtLink
					>
				</li>
				<li class="h-[1.5em] leading-[1.5em] whitespace-nowrap">
					<NuxtLink class="hover:underline" to="https://news.gravitalia.com/">{{
						$t("services.news")
					}}</NuxtLink>
				</li>
				<li class="h-[1.5em] leading-[1.5em] whitespace-nowrap">
					<NuxtLink
						class="hover:underline"
						to="https://galadril.gravitalia.com"
						>{{ $t("services.galadril") }}</NuxtLink
					>
				</li>
				<li class="h-[1.5em] leading-[1.5em] whitespace-nowrap">
					<NuxtLink
						class="hover:underline"
						to="https://account.gravitalia.com/"
						>{{ $t("services.account") }}</NuxtLink
					>
				</li>
			</ul>
		</div>
	</div>
</template>
