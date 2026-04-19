<template>
  <section id="faq" class="mb-16 scroll-mt-24">
    <div class="mb-8 text-center">
      <h2 class="mb-2 text-2xl font-bold text-gray-900 dark:text-white md:text-3xl">
        {{ t('home.faq.title') }}
      </h2>
      <p class="text-sm text-gray-600 dark:text-dark-400">
        {{ t('home.faq.subtitle') }}
      </p>
    </div>

    <div class="mx-auto max-w-3xl space-y-3">
      <details
        v-for="(item, idx) in items"
        :key="item.key"
        class="group rounded-2xl border border-gray-200/60 bg-white/75 backdrop-blur-sm transition-all hover:border-primary-300/60 open:border-primary-400/70 open:shadow-lg open:shadow-primary-500/5 dark:border-dark-700/60 dark:bg-dark-800/75 dark:open:border-primary-500/50"
        :open="idx === 0"
      >
        <summary
          class="flex cursor-pointer list-none items-center justify-between gap-4 px-5 py-4 text-sm font-medium text-gray-900 marker:hidden dark:text-white"
        >
          <span class="flex items-center gap-3">
            <span
              class="flex h-6 w-6 shrink-0 items-center justify-center rounded-full bg-primary-100 text-[11px] font-semibold text-primary-600 dark:bg-primary-900/40 dark:text-primary-300"
            >
              {{ idx + 1 }}
            </span>
            <span>{{ item.q }}</span>
          </span>
          <Icon
            name="chevronDown"
            size="sm"
            class="shrink-0 text-gray-400 transition-transform group-open:rotate-180 dark:text-dark-400"
          />
        </summary>
        <div
          class="border-t border-gray-200/60 px-5 py-4 text-sm leading-relaxed text-gray-600 dark:border-dark-700/60 dark:text-dark-300"
        >
          {{ item.a }}
        </div>
      </details>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()

const keys = ['what', 'billing', 'models', 'security', 'migration', 'refund'] as const

const items = computed(() =>
  keys.map((k) => ({
    key: k,
    q: t(`home.faq.items.${k}.q`),
    a: t(`home.faq.items.${k}.a`)
  }))
)
</script>
