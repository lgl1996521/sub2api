<template>
  <header
    class="sticky top-0 z-40 border-b border-gray-200/60 bg-white/75 backdrop-blur-xl dark:border-dark-800/60 dark:bg-dark-900/75"
  >
    <nav class="mx-auto flex h-16 max-w-6xl items-center gap-4 px-4 sm:px-6">
      <!-- Brand -->
      <router-link to="/home" class="flex shrink-0 items-center">
        <BrandLogo :site-name="siteName" :custom-logo="siteLogo" />
      </router-link>

      <!-- Desktop nav -->
      <div class="hidden flex-1 items-center justify-center gap-1 md:flex">
        <a
          v-for="item in navItems"
          :key="item.key"
          :href="item.href"
          :target="item.external ? '_blank' : undefined"
          :rel="item.external ? 'noopener noreferrer' : undefined"
          class="relative inline-flex items-center gap-1.5 rounded-lg px-3 py-2 text-sm font-medium transition-colors"
          :class="
            isActive(item)
              ? 'text-primary-600 dark:text-primary-400'
              : 'text-gray-600 hover:bg-gray-100 hover:text-gray-900 dark:text-gray-300 dark:hover:bg-dark-800 dark:hover:text-white'
          "
          @click="(ev) => handleNavClick(ev, item)"
        >
          <span>{{ t(item.key) }}</span>
          <span
            v-if="isActive(item)"
            class="absolute inset-x-3 -bottom-[1px] h-0.5 rounded-full bg-gradient-to-r from-primary-500 to-brand-500"
          ></span>
        </a>
      </div>

      <!-- Right actions -->
      <div class="ml-auto flex items-center gap-1.5 sm:gap-2">
        <LocaleSwitcher />

        <button
          type="button"
          @click="toggleTheme"
          class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-900 dark:text-gray-300 dark:hover:bg-dark-800 dark:hover:text-white"
          :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
        >
          <Icon v-if="isDark" name="sun" size="md" />
          <Icon v-else name="moon" size="md" />
        </button>

        <a
          v-if="docUrl"
          :href="docUrl"
          target="_blank"
          rel="noopener noreferrer"
          class="hidden rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-900 dark:text-gray-300 dark:hover:bg-dark-800 dark:hover:text-white sm:inline-flex"
          :title="t('home.viewDocs')"
        >
          <Icon name="book" size="md" />
        </a>

        <router-link
          v-if="isAuthenticated"
          :to="dashboardPath"
          class="inline-flex items-center gap-1.5 rounded-full bg-gradient-to-r from-primary-500 to-brand-500 py-1 pl-1 pr-3 text-xs font-medium text-white shadow-sm transition-all hover:shadow-md hover:shadow-primary-500/30"
        >
          <span
            class="flex h-5 w-5 items-center justify-center rounded-full bg-white/25 text-[10px] font-semibold"
          >
            {{ userInitial }}
          </span>
          <span>{{ t('home.dashboard') }}</span>
        </router-link>
        <template v-else>
          <router-link
            to="/login"
            class="hidden rounded-lg px-3 py-1.5 text-xs font-medium text-gray-700 transition-colors hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-dark-800 sm:inline-flex"
          >
            {{ t('home.login') }}
          </router-link>
          <router-link
            to="/register"
            class="inline-flex items-center gap-1 rounded-full bg-gradient-to-r from-primary-500 to-brand-500 px-3 py-1.5 text-xs font-medium text-white shadow-sm transition-all hover:shadow-md hover:shadow-primary-500/30"
          >
            {{ t('publicHeader.signUp') }}
          </router-link>
        </template>

        <!-- Mobile nav toggle -->
        <button
          type="button"
          class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 dark:text-gray-300 dark:hover:bg-dark-800 md:hidden"
          @click="mobileOpen = !mobileOpen"
          :aria-label="t('publicHeader.toggleMenu')"
        >
          <Icon :name="mobileOpen ? 'x' : 'menu'" size="md" />
        </button>
      </div>
    </nav>

    <!-- Mobile menu -->
    <transition name="slide-down">
      <div
        v-if="mobileOpen"
        class="border-t border-gray-200/60 bg-white/95 backdrop-blur-xl dark:border-dark-800/60 dark:bg-dark-900/95 md:hidden"
      >
        <div class="mx-auto flex max-w-6xl flex-col gap-1 px-4 py-3">
          <a
            v-for="item in navItems"
            :key="item.key"
            :href="item.href"
            :target="item.external ? '_blank' : undefined"
            :rel="item.external ? 'noopener noreferrer' : undefined"
            class="rounded-lg px-3 py-2 text-sm font-medium transition-colors"
            :class="
              isActive(item)
                ? 'bg-primary-50 text-primary-600 dark:bg-primary-900/20 dark:text-primary-400'
                : 'text-gray-700 hover:bg-gray-100 dark:text-gray-200 dark:hover:bg-dark-800'
            "
            @click="(ev) => handleNavClick(ev, item)"
          >
            {{ t(item.key) }}
          </a>
        </div>
      </div>
    </transition>
  </header>
</template>

<script setup lang="ts">
import { computed, ref, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router'
import BrandLogo from '@/components/common/BrandLogo.vue'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'
import { useAuthStore, useAppStore } from '@/stores'

interface NavItem {
  key: string
  /** `/home#section` for hash targets, `/path` for router, or full URL for external. */
  href: string
  external?: boolean
  /** Optional hash name when the target is on `/home`. */
  hash?: string
}

const { t } = useI18n()
const route = useRoute()
const authStore = useAuthStore()
const appStore = useAppStore()

const siteName = computed(
  () => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'Sub2API'
)
const siteLogo = computed(
  () => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || ''
)
const docUrl = computed(
  () => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || ''
)

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))
const userInitial = computed(() => {
  const user = authStore.user
  if (!user || !user.email) return ''
  return user.email.charAt(0).toUpperCase()
})

const mobileOpen = ref(false)
const isDark = ref(document.documentElement.classList.contains('dark'))

const navItems = computed<NavItem[]>(() => {
  const items: NavItem[] = [
    { key: 'publicHeader.home', href: '/home', hash: '' },
    { key: 'publicHeader.models', href: '/models' },
    { key: 'publicHeader.modelStatus', href: '/home#health', hash: 'health' },
    { key: 'publicHeader.plans', href: '/home#plans', hash: 'plans' },
    { key: 'publicHeader.lines', href: '/home#lines', hash: 'lines' }
  ]
  if (docUrl.value) {
    items.push({ key: 'publicHeader.docs', href: docUrl.value, external: true })
  }
  return items
})

function isActive(item: NavItem): boolean {
  if (item.external) return false
  if (item.href === '/models') return route.path === '/models'
  if (item.href.startsWith('/home')) {
    return route.path === '/home' || route.path === '/'
  }
  return route.path === item.href
}

function handleNavClick(ev: MouseEvent, item: NavItem) {
  mobileOpen.value = false
  if (item.external) return
  if (item.hash !== undefined && (route.path === '/home' || route.path === '/')) {
    ev.preventDefault()
    const el = item.hash ? document.getElementById(item.hash) : null
    if (el) {
      el.scrollIntoView({ behavior: 'smooth', block: 'start' })
      history.replaceState(null, '', item.hash ? '#' + item.hash : '/home')
    } else {
      window.scrollTo({ top: 0, behavior: 'smooth' })
    }
  }
}

function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

onMounted(() => {
  // Sync initial theme state in case this header is mounted after theme init.
  isDark.value = document.documentElement.classList.contains('dark')
})
</script>

<style scoped>
.slide-down-enter-active,
.slide-down-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}
.slide-down-enter-from,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>
