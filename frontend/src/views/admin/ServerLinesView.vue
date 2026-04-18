<template>
  <AppLayout>
    <div class="mx-auto w-full max-w-5xl space-y-6 p-4">
      <!-- Header -->
      <div class="flex flex-wrap items-center justify-between gap-3">
        <div>
          <h1 class="text-xl font-semibold text-gray-900 dark:text-white">
            {{ t('admin.serverLines.title') }}
          </h1>
          <p class="mt-1 text-sm text-gray-500 dark:text-dark-400">
            {{ t('admin.serverLines.description') }}
          </p>
        </div>
        <div class="flex flex-wrap items-center gap-2">
          <button
            type="button"
            class="btn btn-secondary"
            :disabled="loading"
            @click="loadLines"
          >
            <Icon name="refresh" size="md" :class="loading ? 'animate-spin' : ''" />
          </button>
          <button
            type="button"
            class="btn btn-secondary"
            :disabled="loading"
            @click="checkStatus"
          >
            {{ t('admin.serverLines.probeNow') }}
          </button>
          <button type="button" class="btn btn-primary" @click="addLine">
            <Icon name="plus" size="md" class="mr-1" />
            {{ t('admin.serverLines.addLine') }}
          </button>
        </div>
      </div>

      <!-- Empty state -->
      <div
        v-if="!loading && lines.length === 0"
        class="rounded-2xl border border-dashed border-gray-300 p-10 text-center text-sm text-gray-500 dark:border-dark-700 dark:text-dark-400"
      >
        {{ t('admin.serverLines.empty') }}
      </div>

      <!-- Line cards -->
      <div v-else class="space-y-4">
        <div
          v-for="(line, index) in lines"
          :key="line._key"
          class="rounded-2xl border border-gray-200 bg-white p-5 shadow-sm dark:border-dark-700 dark:bg-dark-800"
        >
          <div class="mb-4 flex flex-wrap items-center justify-between gap-3">
            <div class="flex items-center gap-3">
              <span
                class="inline-flex h-8 w-8 items-center justify-center rounded-full bg-gray-100 text-xs font-medium text-gray-600 dark:bg-dark-700 dark:text-dark-300"
              >
                #{{ index + 1 }}
              </span>
              <span v-if="statusById[line.id]" class="flex items-center gap-2 text-xs">
                <span
                  class="inline-flex h-1.5 w-1.5 rounded-full"
                  :class="statusDot(statusById[line.id].status)"
                ></span>
                <span class="text-gray-600 dark:text-dark-300">
                  {{ t('admin.serverLines.status.' + statusById[line.id].status) }}
                  <span class="ml-1 font-mono">
                    {{
                      statusById[line.id].status === 'unknown'
                        ? '—'
                        : statusById[line.id].latency_ms + ' ms'
                    }}
                  </span>
                </span>
              </span>
            </div>
            <div class="flex items-center gap-2">
              <label class="flex items-center gap-2 text-sm text-gray-600 dark:text-dark-300">
                <input
                  type="checkbox"
                  class="h-4 w-4 rounded border-gray-300 text-primary-600 focus:ring-primary-500"
                  v-model="line.enabled"
                />
                {{ t('admin.serverLines.fields.enabled') }}
              </label>
              <button
                type="button"
                class="rounded-lg border border-gray-200 p-1.5 text-gray-500 hover:bg-gray-50 dark:border-dark-700 dark:hover:bg-dark-700/50"
                :disabled="index === 0"
                @click="move(index, -1)"
                :title="t('admin.serverLines.moveUp')"
              >
                <Icon name="chevronUp" size="sm" />
              </button>
              <button
                type="button"
                class="rounded-lg border border-gray-200 p-1.5 text-gray-500 hover:bg-gray-50 dark:border-dark-700 dark:hover:bg-dark-700/50"
                :disabled="index === lines.length - 1"
                @click="move(index, 1)"
                :title="t('admin.serverLines.moveDown')"
              >
                <Icon name="chevronDown" size="sm" />
              </button>
              <button
                type="button"
                class="rounded-lg border border-rose-200 p-1.5 text-rose-500 hover:bg-rose-50 dark:border-rose-900/50 dark:hover:bg-rose-900/20"
                @click="remove(index)"
                :title="t('admin.serverLines.remove')"
              >
                <Icon name="trash" size="sm" />
              </button>
            </div>
          </div>

          <div class="grid gap-4 md:grid-cols-2">
            <div>
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.id') }}
              </label>
              <input
                v-model="line.id"
                type="text"
                class="input"
                :placeholder="t('admin.serverLines.placeholders.id')"
              />
            </div>
            <div>
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.name') }}
              </label>
              <input
                v-model="line.name"
                type="text"
                class="input"
                :placeholder="t('admin.serverLines.placeholders.name')"
              />
            </div>
            <div>
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.region') }}
              </label>
              <select v-model="line.region" class="input">
                <option v-for="opt in regionOptions" :key="opt.value" :value="opt.value">
                  {{ opt.label }}
                </option>
              </select>
            </div>
            <div>
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.sortOrder') }}
              </label>
              <input
                v-model.number="line.sort_order"
                type="number"
                class="input"
                min="0"
              />
            </div>
            <div class="md:col-span-2">
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.url') }}
              </label>
              <input
                v-model="line.url"
                type="url"
                class="input"
                placeholder="https://api.example.com"
              />
            </div>
            <div>
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.probePath') }}
              </label>
              <input
                v-model="line.probe_path"
                type="text"
                class="input"
                placeholder="/health"
              />
            </div>
            <div>
              <label class="mb-1 block text-xs font-medium text-gray-500 dark:text-dark-400">
                {{ t('admin.serverLines.fields.description') }}
              </label>
              <input
                v-model="line.description"
                type="text"
                class="input"
                :placeholder="t('admin.serverLines.placeholders.description')"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Save bar -->
      <div
        class="sticky bottom-4 flex flex-wrap items-center justify-between gap-3 rounded-2xl border border-gray-200 bg-white/90 p-3 shadow-lg backdrop-blur dark:border-dark-700 dark:bg-dark-800/90"
      >
        <p class="text-xs text-gray-500 dark:text-dark-400">
          {{ t('admin.serverLines.saveHint') }}
        </p>
        <div class="flex items-center gap-2">
          <button
            type="button"
            class="btn btn-secondary"
            :disabled="saving || loading"
            @click="loadLines"
          >
            {{ t('common.reset') }}
          </button>
          <button
            type="button"
            class="btn btn-primary"
            :disabled="saving || loading"
            @click="save"
          >
            {{ saving ? t('common.saving') : t('common.save') }}
          </button>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { adminAPI } from '@/api/admin'
import AppLayout from '@/components/layout/AppLayout.vue'
import Icon from '@/components/icons/Icon.vue'
import type { ServerLine, ServerLineWithStatus } from '@/api/admin/serverLines'

type EditableLine = ServerLine & { _key: number }

const { t } = useI18n()

const lines = ref<EditableLine[]>([])
const statusById = reactive<Record<string, ServerLineWithStatus>>({})
const loading = ref(false)
const saving = ref(false)
let keyCounter = 0

const regionOptions = computed(() => [
  { value: 'cn', label: t('home.lines.region.cn') },
  { value: 'hk', label: t('home.lines.region.hk') },
  { value: 'tw', label: t('home.lines.region.tw') },
  { value: 'jp', label: t('home.lines.region.jp') },
  { value: 'sg', label: t('home.lines.region.sg') },
  { value: 'us', label: t('home.lines.region.us') },
  { value: 'eu', label: t('home.lines.region.eu') },
  { value: 'intl', label: t('home.lines.region.intl') },
  { value: 'global', label: t('home.lines.region.global') }
])

function toEditable(list: ServerLine[]): EditableLine[] {
  return list.map((l) => ({ ...l, _key: ++keyCounter }))
}

async function loadLines() {
  loading.value = true
  try {
    const resp = await adminAPI.serverLines.list()
    const data = (resp.data ?? resp) as unknown as ServerLine[]
    lines.value = toEditable(data || [])
    await checkStatus()
  } catch (err) {
    console.error(err)
  } finally {
    loading.value = false
  }
}

async function checkStatus() {
  try {
    const resp = await adminAPI.serverLines.listWithStatus()
    const data = (resp.data ?? resp) as unknown as ServerLineWithStatus[]
    for (const key of Object.keys(statusById)) {
      delete statusById[key]
    }
    for (const item of data || []) {
      statusById[item.id] = item
    }
  } catch (err) {
    console.error(err)
  }
}

function addLine() {
  const id = 'line-' + Math.random().toString(36).slice(2, 8)
  lines.value.push({
    _key: ++keyCounter,
    id,
    name: '',
    region: 'intl',
    url: '',
    probe_path: '/health',
    description: '',
    sort_order: (lines.value[lines.value.length - 1]?.sort_order ?? 0) + 1,
    enabled: true
  })
}

function remove(index: number) {
  lines.value.splice(index, 1)
}

function move(index: number, delta: number) {
  const target = index + delta
  if (target < 0 || target >= lines.value.length) return
  const arr = lines.value.slice()
  const [moved] = arr.splice(index, 1)
  arr.splice(target, 0, moved)
  lines.value = arr.map((l, i) => ({ ...l, sort_order: i + 1 }))
}

async function save() {
  saving.value = true
  try {
    const payload: ServerLine[] = lines.value.map((l, idx) => ({
      id: l.id.trim(),
      name: l.name.trim(),
      region: l.region,
      url: l.url.trim(),
      probe_path: (l.probe_path || '/health').trim(),
      description: l.description.trim(),
      sort_order: Number.isFinite(l.sort_order) ? l.sort_order : idx + 1,
      enabled: !!l.enabled
    }))
    const resp = await adminAPI.serverLines.save(payload)
    const data = (resp.data ?? resp) as unknown as ServerLine[]
    lines.value = toEditable(data || [])
    await checkStatus()
  } catch (err) {
    console.error(err)
    const message = (err as { message?: string })?.message || ''
    if (message) {
      window.alert(message)
    }
  } finally {
    saving.value = false
  }
}

function statusDot(status: string): string {
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

onMounted(() => {
  loadLines()
})
</script>
