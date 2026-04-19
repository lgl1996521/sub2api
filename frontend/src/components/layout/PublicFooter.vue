<template>
  <footer
    class="relative z-10 border-t border-gray-200/60 bg-white/70 backdrop-blur-xl dark:border-dark-800/60 dark:bg-dark-900/70"
  >
    <div class="mx-auto max-w-7xl px-6 py-12 lg:py-14">
      <div class="grid gap-10 lg:grid-cols-12">
        <!-- Brand column -->
        <div class="lg:col-span-4">
          <router-link to="/home" class="inline-flex items-center">
            <BrandLogo :site-name="siteName" :custom-logo="siteLogo" />
          </router-link>
          <p class="mt-4 max-w-sm text-sm leading-relaxed text-gray-600 dark:text-dark-300">
            {{ t('publicFooter.tagline') }}
          </p>
          <div class="mt-5 flex items-center gap-2">
            <a
              v-if="githubUrl"
              :href="githubUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 bg-white/70 text-gray-500 transition-all hover:border-primary-300 hover:text-primary-600 dark:border-dark-700 dark:bg-dark-800/70 dark:text-dark-300 dark:hover:border-primary-500/50 dark:hover:text-primary-300"
              :title="t('publicFooter.social.github')"
            >
              <svg viewBox="0 0 24 24" fill="currentColor" class="h-4 w-4" aria-hidden="true">
                <path
                  fill-rule="evenodd"
                  clip-rule="evenodd"
                  d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.919.678 1.852 0 1.336-.012 2.415-.012 2.743 0 .268.18.58.688.482A10.02 10.02 0 0022 12.017C22 6.484 17.522 2 12 2z"
                />
              </svg>
            </a>
            <a
              v-if="docUrl"
              :href="docUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 bg-white/70 text-gray-500 transition-all hover:border-primary-300 hover:text-primary-600 dark:border-dark-700 dark:bg-dark-800/70 dark:text-dark-300 dark:hover:border-primary-500/50 dark:hover:text-primary-300"
              :title="t('publicFooter.social.docs')"
            >
              <Icon name="book" size="sm" />
            </a>
            <a
              v-if="contactEmail"
              :href="'mailto:' + contactEmail"
              class="inline-flex h-9 w-9 items-center justify-center rounded-lg border border-gray-200 bg-white/70 text-gray-500 transition-all hover:border-primary-300 hover:text-primary-600 dark:border-dark-700 dark:bg-dark-800/70 dark:text-dark-300 dark:hover:border-primary-500/50 dark:hover:text-primary-300"
              :title="t('publicFooter.social.email')"
            >
              <Icon name="mail" size="sm" />
            </a>
          </div>
        </div>

        <!-- Link columns -->
        <div class="grid grid-cols-2 gap-8 sm:grid-cols-3 lg:col-span-8">
          <div>
            <h4 class="mb-3 text-[11px] font-semibold uppercase tracking-wider text-gray-500 dark:text-dark-400">
              {{ t('publicFooter.columns.product') }}
            </h4>
            <ul class="space-y-2 text-sm">
              <li>
                <router-link to="/home" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.home') }}
                </router-link>
              </li>
              <li>
                <router-link to="/models" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.models') }}
                </router-link>
              </li>
              <li>
                <router-link to="/home#plans" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.plans') }}
                </router-link>
              </li>
              <li>
                <router-link to="/home#lines" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.lines') }}
                </router-link>
              </li>
            </ul>
          </div>
          <div>
            <h4 class="mb-3 text-[11px] font-semibold uppercase tracking-wider text-gray-500 dark:text-dark-400">
              {{ t('publicFooter.columns.resources') }}
            </h4>
            <ul class="space-y-2 text-sm">
              <li>
                <router-link to="/home#quickstart" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.quickstart') }}
                </router-link>
              </li>
              <li v-if="docUrl">
                <a :href="docUrl" target="_blank" rel="noopener noreferrer" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.docs') }}
                </a>
              </li>
              <li>
                <router-link to="/home#faq" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.faq') }}
                </router-link>
              </li>
              <li>
                <router-link to="/key-usage" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.keyUsage') }}
                </router-link>
              </li>
            </ul>
          </div>
          <div>
            <h4 class="mb-3 text-[11px] font-semibold uppercase tracking-wider text-gray-500 dark:text-dark-400">
              {{ t('publicFooter.columns.account') }}
            </h4>
            <ul class="space-y-2 text-sm">
              <li v-if="isAuthenticated">
                <router-link :to="dashboardPath" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.dashboard') }}
                </router-link>
              </li>
              <template v-else>
                <li>
                  <router-link to="/login" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                    {{ t('publicFooter.links.login') }}
                  </router-link>
                </li>
                <li>
                  <router-link to="/register" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                    {{ t('publicFooter.links.register') }}
                  </router-link>
                </li>
              </template>
              <li>
                <router-link to="/forgot-password" class="text-gray-600 transition-colors hover:text-primary-600 dark:text-dark-300 dark:hover:text-primary-300">
                  {{ t('publicFooter.links.forgotPassword') }}
                </router-link>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <!-- Bottom row -->
      <div
        class="mt-10 flex flex-col items-start justify-between gap-3 border-t border-gray-200/60 pt-6 text-xs text-gray-500 dark:border-dark-800/60 dark:text-dark-400 sm:flex-row sm:items-center"
      >
        <div class="flex flex-wrap items-center gap-x-3 gap-y-1">
          <span>&copy; {{ year }} {{ siteName }}.</span>
          <span>{{ t('home.footer.allRightsReserved') }}</span>
        </div>
        <div class="flex items-center gap-3">
          <span class="inline-flex items-center gap-1.5">
            <span class="h-1.5 w-1.5 animate-pulse rounded-full bg-emerald-500"></span>
            {{ t('publicFooter.status.operational') }}
          </span>
          <span class="hidden h-3 w-px bg-gray-300 dark:bg-dark-700 sm:inline-block"></span>
          <span class="hidden sm:inline">{{ t('publicFooter.builtWith') }}</span>
        </div>
      </div>
    </div>
  </footer>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import BrandLogo from '@/components/common/BrandLogo.vue'
import Icon from '@/components/icons/Icon.vue'
import { useAppStore, useAuthStore } from '@/stores'

const { t } = useI18n()
const appStore = useAppStore()
const authStore = useAuthStore()

const siteName = computed(
  () => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'Sub2API'
)
const siteLogo = computed(
  () => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || ''
)
const docUrl = computed(
  () => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || ''
)
const githubUrl = computed(() => {
  const settings = appStore.cachedPublicSettings as Record<string, unknown> | undefined
  const fromSettings = typeof settings?.github_url === 'string' ? (settings.github_url as string) : ''
  return fromSettings || 'https://github.com/lgl1996521/sub2api'
})
const contactEmail = computed(() => {
  const settings = appStore.cachedPublicSettings as Record<string, unknown> | undefined
  const fromSettings = typeof settings?.contact_email === 'string' ? (settings.contact_email as string) : ''
  return fromSettings
})

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))
const year = computed(() => new Date().getFullYear())
</script>
