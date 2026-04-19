<template>
  <section id="quickstart" class="mb-16 scroll-mt-24">
    <div class="mb-6 text-center">
      <h2 class="mb-2 text-2xl font-bold text-gray-900 dark:text-white md:text-3xl">
        {{ t('home.quickstart.title') }}
      </h2>
      <p class="text-sm text-gray-600 dark:text-dark-400">
        {{ t('home.quickstart.subtitle') }}
      </p>
    </div>

    <div class="grid gap-6 lg:grid-cols-5">
      <!-- Steps -->
      <ol class="space-y-4 lg:col-span-2">
        <li
          v-for="(step, idx) in steps"
          :key="step.key"
          class="group flex items-start gap-4 rounded-2xl border border-gray-200/60 bg-white/70 p-4 backdrop-blur-sm transition-all hover:border-primary-300/70 hover:shadow-lg hover:shadow-primary-500/5 dark:border-dark-700/60 dark:bg-dark-800/70"
        >
          <span
            class="flex h-9 w-9 shrink-0 items-center justify-center rounded-full bg-gradient-to-br from-primary-500 to-brand-500 text-sm font-semibold text-white shadow-sm"
          >
            {{ idx + 1 }}
          </span>
          <div class="min-w-0 flex-1">
            <h3 class="text-sm font-semibold text-gray-900 dark:text-white">
              {{ step.title }}
            </h3>
            <p class="mt-1 text-xs leading-relaxed text-gray-600 dark:text-dark-400">
              {{ step.desc }}
            </p>
          </div>
        </li>
      </ol>

      <!-- Code panel -->
      <div
        class="overflow-hidden rounded-2xl border border-gray-900/70 bg-gradient-to-br from-accent-900 to-accent-950 shadow-xl lg:col-span-3"
      >
        <!-- Tab bar -->
        <div class="flex items-center justify-between border-b border-white/5 px-4 py-2">
          <div class="flex items-center gap-1">
            <button
              v-for="tab in tabs"
              :key="tab.key"
              type="button"
              class="rounded-md px-3 py-1.5 text-xs font-medium transition-colors"
              :class="activeTab === tab.key
                ? 'bg-white/10 text-white'
                : 'text-dark-400 hover:text-white'"
              @click="activeTab = tab.key"
            >
              {{ tab.label }}
            </button>
          </div>
          <button
            type="button"
            class="inline-flex items-center gap-1 rounded-md px-2 py-1 text-xs text-dark-400 transition-colors hover:bg-white/10 hover:text-white"
            @click="copyCurrent"
          >
            <Icon :name="copied ? 'check' : 'copy'" size="xs" />
            <span>{{ copied ? t('home.quickstart.copied') : t('home.quickstart.copy') }}</span>
          </button>
        </div>
        <!-- Fake window chrome -->
        <div class="flex items-center gap-1.5 px-4 pt-3">
          <span class="h-2.5 w-2.5 rounded-full bg-rose-400/80"></span>
          <span class="h-2.5 w-2.5 rounded-full bg-amber-400/80"></span>
          <span class="h-2.5 w-2.5 rounded-full bg-emerald-400/80"></span>
          <span class="ml-3 font-mono text-[11px] text-dark-500">{{ currentTab.filename }}</span>
        </div>
        <!-- Code -->
        <pre
          class="max-h-80 overflow-auto px-5 pb-5 pt-3 font-mono text-[12.5px] leading-relaxed text-dark-100 sm:text-[13px]"
        ><code>{{ currentTab.code }}</code></pre>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()

type TabKey = 'curl' | 'python' | 'node'

const tabs = computed(() => [
  { key: 'curl' as TabKey, label: 'cURL', filename: 'request.sh' },
  { key: 'python' as TabKey, label: 'Python', filename: 'client.py' },
  { key: 'node' as TabKey, label: 'Node.js', filename: 'client.mjs' }
])

const activeTab = ref<TabKey>('curl')

const curlCode = `curl https://api.example.com/v1/chat/completions \\
  -H "Content-Type: application/json" \\
  -H "Authorization: Bearer sk-your-api-key" \\
  -d '{
    "model": "gpt-5",
    "messages": [
      { "role": "user", "content": "Hello from Sub2API!" }
    ]
  }'`

const pythonCode = `from openai import OpenAI

client = OpenAI(
    base_url="https://api.example.com/v1",
    api_key="sk-your-api-key",
)

resp = client.chat.completions.create(
    model="gpt-5",
    messages=[{"role": "user", "content": "Hello from Sub2API!"}],
)
print(resp.choices[0].message.content)`

const nodeCode = `import OpenAI from 'openai'

const client = new OpenAI({
  baseURL: 'https://api.example.com/v1',
  apiKey: 'sk-your-api-key'
})

const resp = await client.chat.completions.create({
  model: 'gpt-5',
  messages: [{ role: 'user', content: 'Hello from Sub2API!' }]
})

console.log(resp.choices[0].message.content)`

const codeMap: Record<TabKey, string> = {
  curl: curlCode,
  python: pythonCode,
  node: nodeCode
}

const currentTab = computed(() => {
  const tab = tabs.value.find((t) => t.key === activeTab.value) ?? tabs.value[0]
  return { ...tab, code: codeMap[tab.key] }
})

const steps = computed(() => [
  {
    key: 'register',
    title: t('home.quickstart.steps.register.title'),
    desc: t('home.quickstart.steps.register.desc')
  },
  {
    key: 'key',
    title: t('home.quickstart.steps.key.title'),
    desc: t('home.quickstart.steps.key.desc')
  },
  {
    key: 'call',
    title: t('home.quickstart.steps.call.title'),
    desc: t('home.quickstart.steps.call.desc')
  }
])

const copied = ref(false)
let copiedTimer: ReturnType<typeof setTimeout> | null = null

async function copyCurrent() {
  try {
    await navigator.clipboard.writeText(currentTab.value.code)
    copied.value = true
    if (copiedTimer) clearTimeout(copiedTimer)
    copiedTimer = setTimeout(() => {
      copied.value = false
    }, 1500)
  } catch {
    // ignore clipboard failures
  }
}
</script>
