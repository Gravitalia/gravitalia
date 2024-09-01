<template>
      <div
    class="p-12 flex flex-col xl:flex-row justify-center items-center xl:space-x-24"
  >
    <h2
      class="text-4xl xl:text-5xl font-bold text-black dark:text-white xl:max-w-md"
    >
      {{ $t("faq.title") }}
    </h2>
    <div>
      <Accordion
        :question="$t('faq.what_is.question')"
        :answer="$t('faq.what_is.answer')"
      />
      <Accordion
        :question="$t('faq.when_out.question')"
        :answer="$t('faq.when_out.answer')"
      />
      <Accordion
        :question="$t('faq.what_missions.question')"
        :answer="$t('faq.what_missions.answer')"
      />
      <Accordion
        :question="$t('faq.privacy.question')"
        :answer="$t('faq.privacy.answer')"
      />
      <Accordion
        :question="$t('faq.transparency.question')"
        :answer="$t('faq.transparency.answer')"
      />
      <Accordion
        :question="$t('faq.funding.question')"
        :answer="$t('faq.funding.answer')"
      />
    </div>
</div>
</template>