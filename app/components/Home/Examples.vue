<template>
	<div
		class="w-full max-w-7xl mx-auto py-24 px-6 lg:px-12 flex flex-col gap-40"
	>
		<div
			class="grid grid-cols-1 xl:grid-cols-2 place-items-center gap-16 xl:gap-12"
		>
			<div class="flex flex-col gap-6 max-w-xl z-20 relative">
				<div
					class="inline-flex w-fit items-center px-3 py-1 rounded-full border border-violet-500/20 bg-violet-500/10 text-violet-600 dark:text-violet-400 text-sm font-semibold tracking-wider uppercase"
				>
					{{ $t("home.turms.badge") }}
				</div>
				<h3
					class="text-3xl md:text-5xl font-bold text-zinc-900 dark:text-white tracking-tight"
				>
					{{ $t("home.turms.title") }}
				</h3>
				<p class="text-lg text-zinc-600 dark:text-zinc-400 leading-relaxed">
					{{ $t("home.turms.desc") }}
				</p>
				<NuxtLink to="https://turms.gravitalia.com/">
					<ButtonDefault class="w-fit mt-4">{{
						$t("home.turms.btn")
					}}</ButtonDefault>
				</NuxtLink>
			</div>

			<div
				class="relative w-full flex items-center justify-center min-h-[300px] lg:min-h-[450px]"
			>
				<div
					class="absolute w-[80%] h-[80%] bg-violet-500/20 dark:bg-violet-500/30 blur-[100px] rounded-full pointer-events-none z-0"
				/>

				<div
					class="relative z-10 transition-transform duration-700 origin-center"
				>
					<MockupMacbook
						class="h-[176px] w-[256px] lg:h-[310px] lg:w-[480px]"
						sclass="w-[320px] h-[12px] lg:w-[560px] lg:h-[16px]"
					>
						<NuxtImg
							format="webp"
							alt="Turms application"
							draggable="false"
							class="relative h-full w-full rounded-t-[10px] border-2 border-[#121212] border-solid bg-black bg-cover object-cover"
							src="/mockup/turms.webp"
						/>
					</MockupMacbook>
				</div>
			</div>
		</div>

		<div
			class="grid grid-cols-1 xl:grid-cols-2 place-items-center gap-16 xl:gap-12 mt-12"
		>
			<div
				class="relative w-full flex items-center justify-center min-h-[300px] lg:min-h-[450px] xl:order-1 order-2"
			>
				<div
					class="absolute w-[80%] h-[80%] bg-indigo-500/20 dark:bg-indigo-500/30 blur-[100px] rounded-full pointer-events-none z-0"
				/>

				<div
					class="relative z-10 transition-transform duration-700 origin-center"
				>
					<MockupSafari
						class="h-[220px] w-[340px] lg:h-[360px] lg:w-[550px] shadow-2xl"
						url="news.gravitalia.com"
						src="/mockup/news.webp"
					/>
				</div>
			</div>

			<div
				class="flex flex-col gap-6 max-w-xl z-20 relative xl:order-2 order-1"
			>
				<div
					class="inline-flex w-fit items-center px-3 py-1 rounded-full border border-indigo-500/20 bg-indigo-500/10 text-indigo-600 dark:text-indigo-400 text-sm font-semibold tracking-wider uppercase"
				>
					{{ $t("home.news.badge") }}
				</div>
				<h3
					class="text-3xl md:text-5xl font-bold text-zinc-900 dark:text-white tracking-tight"
				>
					{{ $t("home.news.title") }}
				</h3>
				<p class="text-lg text-zinc-600 dark:text-zinc-400 leading-relaxed">
					{{ $t("home.news.desc") }}
				</p>
				<NuxtLink to="https://news.gravitalia.com/">
					<ButtonDefault class="w-fit mt-4">{{
						$t("home.news.btn")
					}}</ButtonDefault>
				</NuxtLink>
			</div>
		</div>
	</div>
</template>
