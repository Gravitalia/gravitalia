<script setup>
// import Footer from "~/components/Home/Footer.vue";

const route = useRoute();
const { t } = useI18n();
const head = useLocaleHead();
const title = computed(() =>
	route.meta.title ? t(route.meta.title) : t("layouts.title"),
);
</script>

<template>
	<div>
		<Html :lang="head.htmlAttrs.lang" :dir="head.htmlAttrs.dir">
			<Head>
				<Title>{{ title }}</Title>
				<template v-for="link in head.link" :key="link.key">
					<Link
						:id="link.key"
						:rel="link.rel"
						:href="link.href"
						:hreflang="link.hreflang"
					/>
				</template>
				<template v-for="meta in head.meta" :key="meta.key">
					<Meta
						:id="meta.key"
						:property="meta.property"
						:content="meta.content"
					/>
				</template>
			</Head>
			<Body>
				<slot />
				<!-- <Footer /> -->
			</Body>
		</Html>
	</div>
</template>
