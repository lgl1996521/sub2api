<template>
  <section id="use-cases" class="mb-16 scroll-mt-24">
    <div class="mb-8 text-center">
      <h2 class="mb-2 text-2xl font-bold text-gray-900 dark:text-white md:text-3xl">
        {{ t('home.useCases.title') }}
      </h2>
      <p class="text-sm text-gray-600 dark:text-dark-400">
        {{ t('home.useCases.subtitle') }}
      </p>
    </div>

    <div class="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
      <article
        v-for="item in cases"
        :key="item.key"
        class="group relative overflow-hidden rounded-2xl border border-gray-200/60 bg-white/75 p-6 backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-primary-300/70 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/60 dark:bg-dark-800/75"
      >
        <span
          class="pointer-events-none absolute -right-12 -top-12 h-32 w-32 rounded-full opacity-30 blur-2xl transition-opacity group-hover:opacity-50"
          :class="item.blob"
        ></span>
        <span
          class="mb-4 inline-flex h-11 w-11 items-center justify-center rounded-xl bg-gradient-to-br text-white shadow-sm"
          :class="item.gradient"
        >
          <Icon :name="item.icon" size="md" />
        </span>
        <h3 class="mb-1.5 text-base font-semibold text-gray-900 dark:text-white">
          {{ item.title }}
        </h3>
        <p class="mb-4 text-sm leading-relaxed text-gray-600 dark:text-dark-300">
          {{ item.desc }}
        </p>
        <ul class="space-y-1.5 text-xs text-gray-500 dark:text-dark-400">
          <li
            v-for="(bullet, idx) in item.bullets"
            :key="idx"
            class="flex items-start gap-1.5"
          >
            <Icon name="check" size="xs" class="mt-0.5 shrink-0 text-primary-500" />
            <span>{{ bullet }}</span>
          </li>
        </ul>
      </article>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'

type IconName = 'cpu' | 'chat' | 'sparkles'

const { t } = useI18n()

type CaseItem = {
  key: string
  title: string
  desc: string
  icon: IconName
  gradient: string
  blob: string
  bullets: string[]
}

const cases = computed<CaseItem[]>(() => [
  {
    key: 'developer',
    title: t('home.useCases.items.developer.title'),
    desc: t('home.useCases.items.developer.desc'),
    icon: 'cpu',
    gradient: 'from-primary-500 to-primary-600',
    blob: 'bg-primary-400/40',
    bullets: [
      t('home.useCases.items.developer.bullets.0'),
      t('home.useCases.items.developer.bullets.1'),
      t('home.useCases.items.developer.bullets.2')
    ]
  },
  {
    key: 'team',
    title: t('home.useCases.items.team.title'),
    desc: t('home.useCases.items.team.desc'),
    icon: 'chat',
    gradient: 'from-brand-500 to-brand-600',
    blob: 'bg-brand-400/40',
    bullets: [
      t('home.useCases.items.team.bullets.0'),
      t('home.useCases.items.team.bullets.1'),
      t('home.useCases.items.team.bullets.2')
    ]
  },
  {
    key: 'product',
    title: t('home.useCases.items.product.title'),
    desc: t('home.useCases.items.product.desc'),
    icon: 'sparkles',
    gradient: 'from-emerald-500 to-teal-600',
    blob: 'bg-emerald-400/40',
    bullets: [
      t('home.useCases.items.product.bullets.0'),
      t('home.useCases.items.product.bullets.1'),
      t('home.useCases.items.product.bullets.2')
    ]
  }
])
</script>
