<template>
  <div
    class="relative min-h-screen overflow-hidden bg-gradient-to-br from-gray-50 via-primary-50/30 to-brand-50/30 dark:from-dark-950 dark:via-dark-900 dark:to-dark-950"
  >
    <!-- Ambient background blobs -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div class="absolute -top-40 right-0 h-96 w-96 rounded-full bg-primary-400/15 blur-3xl"></div>
      <div class="absolute bottom-0 -left-40 h-96 w-96 rounded-full bg-brand-500/10 blur-3xl"></div>
    </div>

    <PublicHeader />

    <main class="relative z-10 mx-auto max-w-7xl px-4 py-10 sm:px-6 lg:py-14">
      <!-- Page heading + refresh -->
      <div class="mb-8 flex flex-wrap items-start justify-between gap-4">
        <div>
          <h1 class="text-3xl font-extrabold tracking-tight text-gray-900 dark:text-white md:text-4xl">
            {{ t('models.plaza.title') }}
          </h1>
          <p class="mt-2 text-sm text-gray-500 dark:text-dark-400">
            {{ t('models.plaza.subtitle') }}
          </p>
        </div>
        <button
          type="button"
          class="inline-flex items-center gap-1.5 rounded-lg border border-gray-200 bg-white/80 px-3 py-1.5 text-sm text-gray-600 transition-all hover:border-primary-300 hover:text-primary-600 dark:border-dark-700 dark:bg-dark-800/80 dark:text-dark-300"
          :disabled="loading"
          @click="loadModels"
        >
          <Icon
            name="refresh"
            size="sm"
            :class="{ 'animate-spin': loading }"
          />
          <span>{{ t('models.plaza.refresh') }}</span>
        </button>
      </div>

      <!-- Stat cards -->
      <div class="mb-6 grid grid-cols-2 gap-3 md:grid-cols-4">
        <div
          class="group flex items-center gap-3 rounded-2xl border border-gray-200/70 bg-white/70 p-4 backdrop-blur-sm transition-all hover:border-primary-300/70 hover:shadow-lg hover:shadow-primary-500/5 dark:border-dark-700/70 dark:bg-dark-800/70"
        >
          <span class="flex h-11 w-11 items-center justify-center rounded-xl bg-gradient-to-br from-primary-500 to-brand-500 text-white shadow-sm">
            <Icon name="layers" size="md" />
          </span>
          <div>
            <div class="text-[11px] font-medium uppercase tracking-wider text-gray-400 dark:text-dark-500">
              {{ t('models.plaza.totalCount') }}
            </div>
            <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ models.length }}</div>
          </div>
        </div>
        <div
          v-for="p in topPlatforms"
          :key="p.key"
          class="flex items-center gap-3 rounded-2xl border border-gray-200/70 bg-white/70 p-4 backdrop-blur-sm transition-all hover:border-primary-300/70 hover:shadow-lg hover:shadow-primary-500/5 dark:border-dark-700/70 dark:bg-dark-800/70"
        >
          <ProviderIcon :provider="p.key" size="lg" />
          <div>
            <div class="text-[11px] font-medium uppercase tracking-wider text-gray-400 dark:text-dark-500">
              {{ p.label }}
            </div>
            <div class="text-2xl font-bold text-gray-900 dark:text-white">{{ p.count }}</div>
          </div>
        </div>
      </div>

      <!-- Toolbar row: platform tabs left, time window right -->
      <div class="mb-5 flex flex-wrap items-center justify-between gap-3">
        <div class="flex flex-wrap items-center gap-2">
          <button
            type="button"
            class="rounded-xl border px-3 py-1.5 text-sm transition-all"
            :class="activePlatform === 'all'
              ? 'border-transparent bg-gradient-to-r from-primary-500 to-brand-500 text-white shadow-sm'
              : 'border-gray-200 bg-white/60 text-gray-600 hover:border-primary-300 hover:text-primary-600 dark:border-dark-700 dark:bg-dark-800/60 dark:text-dark-300'"
            @click="activePlatform = 'all'"
          >
            {{ t('models.filter.all') }}
            <span class="ml-1 text-xs opacity-70">{{ models.length }}</span>
          </button>
          <button
            v-for="p in platforms"
            :key="p.key"
            type="button"
            class="inline-flex items-center gap-2 rounded-xl border px-3 py-1.5 text-sm transition-all"
            :class="activePlatform === p.key
              ? 'border-transparent bg-gradient-to-r from-primary-500 to-brand-500 text-white shadow-sm'
              : 'border-gray-200 bg-white/60 text-gray-600 hover:border-primary-300 hover:text-primary-600 dark:border-dark-700 dark:bg-dark-800/60 dark:text-dark-300'"
            @click="activePlatform = p.key"
          >
            <ProviderIcon :provider="p.key" size="xs" rounded="md" />
            {{ p.label }}
            <span class="ml-0.5 text-xs opacity-70">{{ p.count }}</span>
          </button>
        </div>

        <div class="inline-flex rounded-xl border border-gray-200 bg-white/60 p-0.5 text-xs dark:border-dark-700 dark:bg-dark-800/60">
          <button
            v-for="w in windows"
            :key="w"
            type="button"
            class="rounded-lg px-3 py-1 font-medium transition-all"
            :class="activeWindow === w
              ? 'bg-gradient-to-r from-primary-500 to-brand-500 text-white shadow-sm'
              : 'text-gray-500 hover:text-primary-600 dark:text-dark-400'"
            @click="activeWindow = w"
          >
            {{ w }}
          </button>
        </div>
      </div>

      <!-- Exchange rate hint -->
      <div
        class="mb-5 flex flex-wrap items-center justify-between gap-2 text-xs text-gray-500 dark:text-dark-400"
      >
        <span v-if="collectedAt" class="inline-flex items-center gap-1">
          <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-emerald-500"></span>
          {{ t('models.refreshed', { time: formattedCollectedAt }) }}
        </span>
        <span class="inline-flex items-center gap-1 rounded-full bg-white/60 px-3 py-1 dark:bg-dark-800/60">
          <Icon name="currency" size="xs" class="text-brand-500" />
          1 USD ≈ {{ usdRate.toFixed(2) }} CNY
        </span>
      </div>

      <!-- Empty state -->
      <div
        v-if="!loading && !filteredModels.length"
        class="mx-auto max-w-lg rounded-2xl border border-dashed border-gray-300 bg-white/60 p-10 text-center text-sm text-gray-500 backdrop-blur-sm dark:border-dark-700 dark:bg-dark-800/60 dark:text-dark-400"
      >
        {{ t('models.empty') }}
      </div>

      <!-- Loading skeleton -->
      <div
        v-else-if="loading && !models.length"
        class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3"
      >
        <div
          v-for="i in 6"
          :key="i"
          class="h-64 animate-pulse rounded-2xl border border-gray-200/60 bg-white/60 dark:border-dark-700/60 dark:bg-dark-800/60"
        ></div>
      </div>

      <!-- Model pricing cards -->
      <div v-else class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
        <article
          v-for="m in filteredModels"
          :key="m.model_name"
          class="group relative overflow-hidden rounded-2xl border border-gray-200/60 bg-white/85 p-5 backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:border-primary-300/70 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/60 dark:bg-dark-800/85 dark:hover:border-primary-500/50"
        >
          <!-- Decorative gradient -->
          <div
            class="pointer-events-none absolute -right-20 -top-20 h-44 w-44 rounded-full bg-gradient-to-br from-primary-400/15 to-brand-400/15 blur-3xl"
          ></div>

          <!-- Header: provider icon + name (copy) + release date -->
          <header class="mb-3 flex items-start gap-3">
            <ProviderIcon :provider="m.platform" size="md" rounded="xl" />
            <div class="min-w-0 flex-1">
              <div class="flex items-center gap-1.5">
                <h3 class="truncate text-sm font-semibold text-gray-900 dark:text-white">
                  {{ m.display_name || m.model_name }}
                </h3>
                <button
                  type="button"
                  class="shrink-0 rounded p-1 text-gray-400 transition-colors hover:bg-gray-100 hover:text-primary-600 dark:hover:bg-dark-700"
                  :title="copiedModel === m.model_name ? t('models.card.copied') : t('models.card.copyModelName')"
                  @click="copyModelName(m.model_name)"
                >
                  <Icon :name="copiedModel === m.model_name ? 'check' : 'copy'" size="xs" />
                </button>
              </div>
              <div class="mt-0.5 flex items-center gap-2 text-[11px] text-gray-400 dark:text-dark-500">
                <span class="font-semibold uppercase tracking-wider">
                  {{ platformLabel(m.platform) }}
                </span>
                <span v-if="m.release_date">{{ m.release_date }}</span>
              </div>
            </div>
            <span
              v-if="m.status && m.status !== 'unknown'"
              class="inline-flex shrink-0 items-center gap-1 rounded-full px-2 py-0.5 text-[10px] font-medium"
              :class="statusBadgeClass(m.status)"
              :title="statusTitle(m.status)"
            >
              <span class="h-1.5 w-1.5 rounded-full" :class="statusDotClass(m.status)"></span>
              {{ statusLabel(m.status) }}
            </span>
          </header>

          <!-- Tags -->
          <div v-if="m.tags && m.tags.length" class="mb-3 flex flex-wrap gap-1.5">
            <span
              v-for="tag in m.tags"
              :key="tag"
              class="inline-flex items-center gap-1 rounded-md bg-emerald-50 px-2 py-0.5 text-[11px] font-medium text-emerald-700 dark:bg-emerald-900/20 dark:text-emerald-300"
            >
              <Icon name="bolt" size="xs" />
              {{ tag }}
            </span>
          </div>

          <!-- Pricing: input / output -->
          <div class="mb-3 grid grid-cols-2 gap-3 border-t border-gray-100 pt-3 dark:border-dark-700/60">
            <div>
              <div class="mb-0.5 flex items-center gap-1 text-[11px] text-gray-400 dark:text-dark-500">
                <Icon name="arrowDownTray" size="xs" />
                {{ t('models.card.input') }}
              </div>
              <div class="text-2xl font-bold text-gray-900 dark:text-white">
                ¥{{ formatPrice(m.pricing.input_cny) }}
              </div>
              <div class="text-[11px] text-gray-400 dark:text-dark-500">
                ${{ formatPrice(m.pricing.input_usd) }}
              </div>
            </div>
            <div>
              <div class="mb-0.5 flex items-center gap-1 text-[11px] text-gray-400 dark:text-dark-500">
                <Icon name="arrowUpTray" size="xs" />
                {{ t('models.card.output') }}
              </div>
              <div class="text-2xl font-bold text-gray-900 dark:text-white">
                ¥{{ formatPrice(m.pricing.output_cny) }}
              </div>
              <div class="text-[11px] text-gray-400 dark:text-dark-500">
                ${{ formatPrice(m.pricing.output_usd) }}
              </div>
            </div>
          </div>

          <!-- Cache read -->
          <div v-if="m.pricing.cache_read_cny > 0" class="mb-3 border-t border-gray-100 pt-3 dark:border-dark-700/60">
            <div class="mb-0.5 flex items-center gap-1 text-[11px] text-gray-400 dark:text-dark-500">
              <Icon name="database" size="xs" />
              {{ t('models.card.cacheRead') }}
            </div>
            <div class="flex items-baseline gap-2">
              <span class="text-2xl font-bold text-gray-900 dark:text-white">
                ¥{{ formatPrice(m.pricing.cache_read_cny) }}
              </span>
              <span class="text-[11px] text-gray-400 dark:text-dark-500">
                ${{ formatPrice(m.pricing.cache_read_usd) }}
              </span>
            </div>
          </div>

          <footer class="flex items-center justify-end border-t border-gray-100 pt-2 text-[10px] uppercase tracking-wider text-gray-400 dark:border-dark-700/60 dark:text-dark-500">
            {{ t('models.card.unit') }}
          </footer>
        </article>
      </div>
    </main>

    <footer class="relative z-10 mt-10 border-t border-gray-200/50 px-6 py-6 text-center text-xs text-gray-500 dark:border-dark-800/50 dark:text-dark-400">
      &copy; {{ new Date().getFullYear() }} {{ siteName }}
    </footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onBeforeUnmount, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import PublicHeader from '@/components/layout/PublicHeader.vue'
import ProviderIcon from '@/components/common/ProviderIcon.vue'
import Icon from '@/components/icons/Icon.vue'
import { useAppStore } from '@/stores'
import { publicAPI, type PlatformHealthStatus, type PublicModelItem } from '@/api/public'

const { t, locale } = useI18n()
const appStore = useAppStore()

const siteName = computed(
  () => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'Sub2API'
)

const models = ref<PublicModelItem[]>([])
const collectedAt = ref<string | undefined>(undefined)
const loading = ref(true)
const activePlatform = ref<string>('all')
const windows = ['90m', '24h', '7d'] as const
const activeWindow = ref<(typeof windows)[number]>('90m')
const copiedModel = ref<string>('')
const usdRate = ref(7.2)

let refreshTimer: ReturnType<typeof setInterval> | null = null
let copiedTimer: ReturnType<typeof setTimeout> | null = null

async function loadModels() {
  loading.value = true
  try {
    const resp = await publicAPI.getModelHealth()
    const data = (resp.data ?? resp) as unknown as {
      collected_at?: string
      models?: PublicModelItem[]
    }
    const list = Array.isArray(data.models) ? data.models : []
    models.value = list
    collectedAt.value = data.collected_at
    // Derive USD rate from first pricing row (pricing_usd > 0)
    const probe = list.find((m) => m.pricing?.input_usd > 0)
    if (probe && probe.pricing.input_usd > 0) {
      usdRate.value = probe.pricing.input_cny / probe.pricing.input_usd
    }
  } catch {
    models.value = []
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadModels()
  refreshTimer = setInterval(loadModels, 60 * 1000)
})

onBeforeUnmount(() => {
  if (refreshTimer) clearInterval(refreshTimer)
  if (copiedTimer) clearTimeout(copiedTimer)
})

const platforms = computed(() => {
  const map = new Map<string, { key: string; label: string; count: number }>()
  for (const m of models.value) {
    const key = (m.platform || 'unknown').toLowerCase()
    const existing = map.get(key)
    if (existing) existing.count += 1
    else map.set(key, { key, label: platformLabel(key), count: 1 })
  }
  // Stable sort: anthropic / openai / gemini first, then alpha
  const rank: Record<string, number> = {
    openai: 0,
    anthropic: 1,
    gemini: 2,
    google: 2,
    antigravity: 3,
    codex: 4,
    bedrock: 5
  }
  return Array.from(map.values()).sort((a, b) => {
    const ra = rank[a.key] ?? 99
    const rb = rank[b.key] ?? 99
    if (ra !== rb) return ra - rb
    return a.key.localeCompare(b.key)
  })
})

const topPlatforms = computed(() => platforms.value.slice(0, 3))

const filteredModels = computed(() => {
  if (activePlatform.value === 'all') return models.value
  return models.value.filter((m) => (m.platform || 'unknown').toLowerCase() === activePlatform.value)
})

const formattedCollectedAt = computed(() => {
  if (!collectedAt.value) return ''
  try {
    return new Intl.DateTimeFormat(locale.value, {
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    }).format(new Date(collectedAt.value))
  } catch {
    return collectedAt.value
  }
})

function platformLabel(p: string): string {
  switch (p.toLowerCase()) {
    case 'anthropic':
    case 'claude':
      return 'Anthropic'
    case 'openai':
      return 'OpenAI'
    case 'gemini':
    case 'google':
      return 'Gemini'
    case 'antigravity':
      return 'Antigravity'
    case 'codex':
      return 'Codex'
    case 'bedrock':
      return 'Bedrock'
    default:
      return p ? p.charAt(0).toUpperCase() + p.slice(1) : 'Other'
  }
}

function statusBadgeClass(status: PlatformHealthStatus): string {
  switch (status) {
    case 'healthy':
      return 'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/30 dark:text-emerald-300'
    case 'degraded':
      return 'bg-amber-100 text-amber-700 dark:bg-amber-900/30 dark:text-amber-300'
    case 'down':
      return 'bg-rose-100 text-rose-700 dark:bg-rose-900/30 dark:text-rose-300'
    default:
      return 'bg-gray-100 text-gray-600 dark:bg-dark-700 dark:text-dark-300'
  }
}

function statusDotClass(status: PlatformHealthStatus): string {
  switch (status) {
    case 'healthy':
      return 'bg-emerald-500'
    case 'degraded':
      return 'bg-amber-500'
    case 'down':
      return 'bg-rose-500'
    default:
      return 'bg-gray-400'
  }
}

function statusLabel(status: PlatformHealthStatus): string {
  switch (status) {
    case 'healthy':
      return t('models.card.statusNormal')
    case 'degraded':
      return t('models.card.statusDegraded')
    case 'down':
      return t('models.card.statusDown')
    default:
      return status
  }
}

function statusTitle(status: PlatformHealthStatus): string {
  return statusLabel(status)
}

function formatPrice(value: number): string {
  if (!value) return '0'
  if (value >= 100) return value.toFixed(0)
  if (value >= 10) return value.toFixed(1)
  if (value >= 1) return value.toFixed(2)
  return value.toFixed(3).replace(/0+$/, '').replace(/\.$/, '')
}

async function copyModelName(name: string) {
  try {
    await navigator.clipboard.writeText(name)
    copiedModel.value = name
    if (copiedTimer) clearTimeout(copiedTimer)
    copiedTimer = setTimeout(() => (copiedModel.value = ''), 1500)
  } catch {
    // ignore
  }
}
</script>
