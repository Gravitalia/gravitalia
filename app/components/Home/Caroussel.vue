<template>
	<div>
		<h2 class="font-bold text-3xl lg:text-5xl flex dark:text-white">
			<span class="leading-[5rem] text-violet-600 dark:text-violet-500"
				>Gravitalia&nbsp;</span
			>
			<div class="h-20 overflow-hidden">
				<ul class="animate-change hover:[animation-play-state:paused]">
					<li class="leading-[5rem]">
						<NuxtLink
							class="no-underline inline-block"
							to="https://account.gravitalia.com/"
						>
							{{ $t("services.account") }}</NuxtLink
						>
					</li>
					<li class="leading-[5rem]">
						<NuxtLink
							class="no-underline inline-block"
							to="https://turms.gravitalia.com/"
						>
							{{ $t("services.turms") }}</NuxtLink
						>
					</li>
					<li class="leading-[5rem]">
						<NuxtLink
							class="no-underline inline-block"
							to="https://news.gravitalia.com/"
						>
							{{ $t("services.news") }}
						</NuxtLink>
					</li>
					<li class="leading-[5rem]">
						<NuxtLink
							class="no-underline inline-block"
							to="https://ai.gravitalia.com/"
						>
							{{ $t("services.ai") }}</NuxtLink
						>
					</li>
				</ul>
			</div>
		</h2>
	</div>
</template>
