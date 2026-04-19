<template>
  <div class="flex items-center gap-2.5">
    <!-- Custom site logo overrides SVG when admin uploaded one -->
    <img
      v-if="customLogo"
      :src="customLogo"
      :alt="siteName"
      class="h-9 w-9 rounded-xl object-contain shadow-md"
    />
    <!-- Default: indigo→purple gradient SVG mark -->
    <span
      v-else
      class="relative flex h-9 w-9 items-center justify-center overflow-hidden rounded-xl shadow-md"
      :class="glow ? 'shadow-primary-500/40' : ''"
    >
      <svg
        viewBox="0 0 48 48"
        xmlns="http://www.w3.org/2000/svg"
        class="h-full w-full"
        aria-hidden="true"
      >
        <defs>
          <linearGradient :id="gradientId" x1="0" y1="0" x2="1" y2="1">
            <stop offset="0%" stop-color="#6366f1" />
            <stop offset="100%" stop-color="#a855f7" />
          </linearGradient>
          <linearGradient :id="gradientId + '-soft'" x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stop-color="#ffffff" stop-opacity="0.22" />
            <stop offset="100%" stop-color="#ffffff" stop-opacity="0" />
          </linearGradient>
        </defs>
        <!-- Rounded square background -->
        <rect x="0" y="0" width="48" height="48" rx="12" :fill="'url(#' + gradientId + ')'" />
        <rect
          x="0"
          y="0"
          width="48"
          height="24"
          rx="12"
          :fill="'url(#' + gradientId + '-soft)'"
        />
        <!-- Stylized "S2" mark: two curved arcs + "2" -->
        <path
          d="M31 16c-1.3-2.3-3.9-3.8-7.1-3.8-4.3 0-7.4 2.5-7.4 6 0 3.1 2.2 4.9 6.1 5.8l2.6 0.6c3.1 0.7 4.5 1.8 4.5 3.6 0 2.2-2 3.7-5 3.7-2.9 0-5-1.2-6.1-3.5"
          stroke="#ffffff"
          stroke-width="2.6"
          stroke-linecap="round"
          fill="none"
        />
        <circle cx="35" cy="33" r="3" fill="#ffffff" />
      </svg>
    </span>

    <div v-if="showText" class="flex flex-col leading-tight">
      <span
        class="text-base font-bold tracking-tight"
        :class="textClass"
      >
        {{ siteName }}
      </span>
      <span
        v-if="subtitle"
        class="text-[10px] uppercase tracking-[0.12em] text-gray-500 dark:text-dark-400"
      >
        {{ subtitle }}
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type Props = {
  siteName?: string
  subtitle?: string
  customLogo?: string
  showText?: boolean
  glow?: boolean
  gradientText?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  siteName: 'Sub2API',
  subtitle: '',
  customLogo: '',
  showText: true,
  glow: false,
  gradientText: true
})

// Stable-per-instance gradient id so multiple logos don't collide.
const gradientId = computed(
  () => 'brand-logo-gradient-' + Math.random().toString(36).slice(2, 8)
)

const textClass = computed(() =>
  props.gradientText
    ? 'bg-gradient-to-r from-primary-500 to-brand-500 bg-clip-text text-transparent'
    : 'text-gray-900 dark:text-white'
)
</script>
