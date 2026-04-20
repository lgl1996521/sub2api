<template>
  <header class="sticky top-0 z-50 border-b border-gray-800/50 bg-gray-950/80 backdrop-blur-xl">
    <div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-6">
      <!-- Logo -->
      <router-link to="/" class="flex items-center gap-2">
        <span class="text-xl font-bold text-white">Sub2API</span>
      </router-link>

      <!-- Desktop Nav -->
      <nav class="hidden items-center gap-1 md:flex">
        <router-link
          v-for="item in navItems"
          :key="item.key"
          :to="item.href"
          class="rounded-lg px-4 py-2 text-sm font-medium text-gray-400 transition-colors hover:bg-gray-800 hover:text-white"
          :class="{ 'text-primary-400': isActive(item.href) }"
        >
          {{ t(item.key) }}
        </router-link>
      </nav>

      <!-- Right Actions -->
      <div class="flex items-center gap-2">
        <!-- Theme Toggle -->
        <button
          type="button"
          class="rounded-lg p-2 text-gray-400 transition-colors hover:bg-gray-800 hover:text-white"
          :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          @click="toggleTheme"
        >
          <Icon v-if="isDark" name="sun" size="sm" />
          <Icon v-else name="moon" size="sm" />
        </button>

        <!-- Language Switcher -->
        <div class="relative">
          <button
            type="button"
            class="flex items-center gap-1 rounded-lg p-2 text-sm font-medium text-gray-400 transition-colors hover:bg-gray-800 hover:text-white"
            @click="showLangMenu = !showLangMenu"
          >
            <Icon name="globe" size="sm" />
            <span class="hidden sm:inline">{{ currentLocaleLabel }}</span>
          </button>
          <div
            v-if="showLangMenu"
            class="absolute right-0 mt-2 w-32 overflow-hidden rounded-lg border border-gray-800 bg-gray-900 shadow-xl"
          >
            <button
              v-for="lang in availableLocales"
              :key="lang.code"
              class="w-full px-4 py-2 text-left text-sm text-gray-400 transition-colors hover:bg-gray-800 hover:text-white"
              :class="{ 'text-primary-400': locale === lang.code }"
              @click="changeLocale(lang.code)"
            >
              {{ lang.label }}
            </button>
          </div>
        </div>

        <!-- Auth Buttons -->
        <template v-if="!isAuthenticated">
          <router-link
            to="/login"
            class="hidden rounded-lg px-4 py-2 text-sm font-medium text-gray-400 transition-colors hover:text-white sm:block"
          >
            {{ t('publicHeader.login') }}
          </router-link>
          <router-link
            to="/register"
            class="rounded-full bg-primary-500 px-4 py-2 text-sm font-medium text-white transition-all hover:bg-primary-600"
          >
            {{ t('publicHeader.signUp') }}
          </router-link>
        </template>

        <template v-else>
          <router-link
            :to="dashboardPath"
            class="inline-flex items-center gap-2 rounded-full bg-primary-500 px-4 py-2 text-sm font-medium text-white transition-all hover:bg-primary-600"
          >
            <span class="flex h-6 w-6 items-center justify-center rounded-full bg-white/20 text-xs font-semibold">
              {{ userInitial }}
            </span>
            <span class="hidden sm:inline">{{ t('publicHeader.dashboard') }}</span>
          </router-link>
        </template>

        <!-- Mobile Menu Button -->
        <button
          type="button"
          class="rounded-lg p-2 text-gray-400 transition-colors hover:bg-gray-800 hover:text-white md:hidden"
          @click="showMobileMenu = !showMobileMenu"
        >
          <Icon :name="showMobileMenu ? 'x' : 'menu'" size="md" />
        </button>
      </div>
    </div>

    <!-- Mobile Menu -->
    <div
      v-if="showMobileMenu"
      class="border-t border-gray-800 bg-gray-950 px-6 py-4 md:hidden"
    >
      <nav class="flex flex-col gap-2">
        <router-link
          v-for="item in navItems"
          :key="item.key"
          :to="item.href"
          class="rounded-lg px-4 py-2 text-sm font-medium text-gray-400 transition-colors hover:bg-gray-800 hover:text-white"
          :class="{ 'text-primary-400': isActive(item.href) }"
          @click="showMobileMenu = false"
        >
          {{ t(item.key) }}
        </router-link>
        <div class="my-2 border-t border-gray-800"></div>
        <template v-if="!isAuthenticated">
          <router-link
            to="/login"
            class="rounded-lg px-4 py-2 text-sm font-medium text-gray-400 transition-colors hover:bg-gray-800 hover:text-white"
            @click="showMobileMenu = false"
          >
            {{ t('publicHeader.login') }}
          </router-link>
        </template>
      </nav>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute } from 'vue-router'
import Icon from '@/components/icons/Icon.vue'
import { useAuthStore } from '@/stores'

const { t, locale } = useI18n()
const route = useRoute()
const authStore = useAuthStore()

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))

const userInitial = computed(() => {
  const user = authStore.user
  if (!user?.email) return ''
  return user.email.charAt(0).toUpperCase()
})

const isDark = ref(document.documentElement.classList.contains('dark'))

const navItems = [
  { key: 'publicHeader.home', href: '/' },
  { key: 'publicHeader.models', href: '/models' },
  { key: 'publicHeader.plans', href: '/#pricing' },
  { key: 'publicHeader.quickstart', href: '/#quickstart' },
  { key: 'publicHeader.faq', href: '/#faq' }
]

const availableLocales = [
  { code: 'zh', label: '中文' },
  { code: 'en', label: 'English' }
]

const currentLocaleLabel = computed(() => {
  return availableLocales.find(l => l.code === locale.value)?.label || '中文'
})

const showLangMenu = ref(false)
const showMobileMenu = ref(false)

function isActive(href: string) {
  if (href.startsWith('/#')) {
    return route.path === '/' && route.hash === href.slice(1)
  }
  return route.path === href
}

function toggleTheme() {
  const html = document.documentElement
  if (html.classList.contains('dark')) {
    html.classList.remove('dark')
    localStorage.setItem('theme', 'light')
    isDark.value = false
  } else {
    html.classList.add('dark')
    localStorage.setItem('theme', 'dark')
    isDark.value = true
  }
}

function changeLocale(code: string) {
  locale.value = code
  localStorage.setItem('locale', code)
  showLangMenu.value = false
}

// Close menus on click outside
document.addEventListener('click', (e) => {
  const target = e.target as HTMLElement
  if (!target.closest('.relative')) {
    showLangMenu.value = false
  }
})
</script>
