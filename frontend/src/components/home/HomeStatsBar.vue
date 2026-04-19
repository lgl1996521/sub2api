<template>
  <section
    class="mb-16 rounded-3xl border border-gray-200/50 bg-gradient-to-br from-white/80 via-primary-50/30 to-brand-50/40 p-6 shadow-sm backdrop-blur-sm dark:border-dark-700/50 dark:from-dark-800/80 dark:via-primary-950/40 dark:to-dark-800/80 sm:p-8"
  >
    <div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-4">
      <div
        v-for="item in items"
        :key="item.key"
        class="flex items-center gap-4"
      >
        <span
          class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-gradient-to-br text-white shadow-sm"
          :class="item.gradient"
        >
          <Icon :name="item.icon" size="md" />
        </span>
        <div class="min-w-0">
          <div class="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
            {{ item.value }}
          </div>
          <div class="text-xs text-gray-500 dark:text-dark-400">
            {{ item.label }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'

type IconName = 'cube' | 'server' | 'shield' | 'bolt'

const props = defineProps<{
  modelCount?: number
  providerCount?: number
}>()

const { t } = useI18n()

type StatItem = {
  key: string
  value: string
  label: string
  icon: IconName
  gradient: string
}

const items = computed<StatItem[]>(() => {
  const models = props.modelCount && props.modelCount > 0 ? props.modelCount : 40
  const providers = props.providerCount && props.providerCount > 0 ? props.providerCount : 6
  return [
    {
      key: 'models',
      value: `${models}+`,
      label: t('home.stats.models'),
      icon: 'cube',
      gradient: 'from-primary-500 to-primary-600'
    },
    {
      key: 'providers',
      value: `${providers}+`,
      label: t('home.stats.providers'),
      icon: 'server',
      gradient: 'from-brand-500 to-brand-600'
    },
    {
      key: 'uptime',
      value: '99.9%',
      label: t('home.stats.uptime'),
      icon: 'shield',
      gradient: 'from-emerald-500 to-teal-600'
    },
    {
      key: 'latency',
      value: '< 200ms',
      label: t('home.stats.latency'),
      icon: 'bolt',
      gradient: 'from-amber-500 to-orange-500'
    }
  ]
})
</script>
