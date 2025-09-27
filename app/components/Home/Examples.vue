<template>
	<div
		class="h-full max-h-xl w-full py-12 xl:py-24 px-10 xl:px-56 2xl:px-80 grid grid-cols-1 xl:grid-cols-2 place-items-center gap-y-6"
	>
		<div class="lg:mr-48 max-w-md xl:max-w-lg">
			<h3 class="font-semibold text-violet-500 dark:text-violet-400 uppercase">
				{{ $t("home.site.title") }}
			</h3>
			<p class="text-zinc-800 dark:text-zinc-200">
				{{ $t("home.site.description") }}
			</p>
		</div>
		<MockupSafari
			class="h-80 w-80 lg:h-[425px] lg:w-[650px]"
			url="news.gravitalia.com"
			src="/mockup/news.png"
		/>
	</div>

	<div
		class="h-full max-h-xl w-full py-12 xl:py-24 px-10 xl:px-56 2xl:px-80 grid grid-cols-1 xl:grid-cols-2 place-items-center gap-y-6"
	>
		<div class="lg:mr-48">
			<MockupMacbook
				class="h-44 w-64 lg:h-[425px] lg:w-[650px]"
				sclass="w-80 h-3 lg:w-[740px] lg:h-6"
			>
				<NuxtImg
					width="700"
					format="webp"
					quality="80"
					alt="Turms application"
					draggable="false"
					class="relative h-full w-full rounded-t-[10px] border-2 border-[rgb(18,18,18)] border-solid bg-black bg-cover object-cover"
					src="/mockup/turms.png"
			/></MockupMacbook>
		</div>
		<div class="lg:ml-24 max-w-md xl:max-w-lg">
			<h3 class="font-semibold text-violet-500 dark:text-violet-400 uppercase">
				{{ $t("home.computer.title") }}
			</h3>
			<p class="text-zinc-800 dark:text-zinc-200">
				{{ $t("home.computer.description") }}
			</p>
		</div>
	</div>
</template>
