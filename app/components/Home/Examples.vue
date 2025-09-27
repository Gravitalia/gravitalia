<template>
	<div
		class="h-full max-h-xl w-full py-12 xl:py-24 px-10 xl:px-56 2xl:px-80 grid grid-cols-1 xl:grid-cols-2 place-items-center gap-y-6"
	>
		<div class="lg:mr-48">
			<MockupMacbook
				class="h-[280px] w-[450px] lg:h-[425px] lg:w-[650px]"
				sclass="w-[500px] h-4 lg:w-[740px] lg:h-6"
				src="/mockup/turms.png"
			/>
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
