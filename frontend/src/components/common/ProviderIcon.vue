<template>
  <span
    class="inline-flex items-center justify-center overflow-hidden"
    :class="[sizeClass, roundedClass, bgClass]"
    :aria-label="provider"
  >
    <!-- OpenAI: hex-knot logo simplified -->
    <svg
      v-if="variant === 'openai'"
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <path
        d="M44.7 20.2a12 12 0 0 0-1-9.8 12 12 0 0 0-13-5.8A12 12 0 0 0 9.5 9.3 12 12 0 0 0 4 19.5a12 12 0 0 0 1 9.8A12 12 0 0 0 18 35.1 12 12 0 0 0 38.5 38.7 12 12 0 0 0 44 28.5a12 12 0 0 0 .7-8.3ZM27.9 41.4a9 9 0 0 1-5.7-2l.3-.2 9.4-5.4a1.6 1.6 0 0 0 .8-1.3V19.3l4 2.3v10.9a9 9 0 0 1-8.8 8.9ZM8.7 33.5a9 9 0 0 1-1.1-6l.3.2 9.4 5.4a1.5 1.5 0 0 0 1.5 0L30.3 27v4.6a.1.1 0 0 1-.1.1l-9.5 5.5a9 9 0 0 1-12-3.7Zm-2.5-20.6a9 9 0 0 1 4.7-4l0 .4v10.8a1.5 1.5 0 0 0 .7 1.3l11.4 6.6-4 2.3-9.5-5.5a9 9 0 0 1-3.3-11.9Zm32.6 7.6-11.5-6.6 4-2.3 9.5 5.5a9 9 0 0 1-1.4 16.3v-11.2a1.6 1.6 0 0 0-.6-1.3ZM42.7 16l-.3-.2-9.4-5.5a1.5 1.5 0 0 0-1.5 0L20 16.9V12.3a.1.1 0 0 1 0-.2l9.6-5.5a9 9 0 0 1 13.1 9.3Zm-24.7 9.4-4-2.3V12.2a9 9 0 0 1 14.7-6.9l-.3.2-9.4 5.4a1.6 1.6 0 0 0-.8 1.3Zm2.2-4.7 5-2.9 5.1 2.9v5.8L25.2 29.6l-5.1-2.9Z"
        fill="currentColor"
      />
    </svg>

    <!-- Anthropic: stylised "A" / starburst -->
    <svg
      v-else-if="variant === 'anthropic'"
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <path
        d="M16 7L6 41h7.5l2-7h12l2 7H37L27 7h-11Zm5.5 6l4 13.5h-8L21.5 13Z"
        fill="currentColor"
      />
      <path
        d="M34 7l8 34h-6l-8-34h6Z"
        fill="currentColor"
        opacity="0.75"
      />
    </svg>

    <!-- Gemini / Google: 4-point sparkle -->
    <svg
      v-else-if="variant === 'gemini'"
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <path
        d="M24 4c0 9-7 16-16 16 0 0 0 0 0 0v8c9 0 16 7 16 16v0c0-9 7-16 16-16v-8c-9 0-16-7-16-16Z"
        fill="currentColor"
      />
    </svg>

    <!-- Antigravity / planet-ring -->
    <svg
      v-else-if="variant === 'antigravity'"
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <ellipse
        cx="24"
        cy="24"
        rx="20"
        ry="6"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        transform="rotate(-20 24 24)"
      />
      <circle cx="24" cy="24" r="10" fill="currentColor" />
    </svg>

    <!-- Codex: keycap "{}" -->
    <svg
      v-else-if="variant === 'codex'"
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <path
        d="M19 8c-6 0-8 3-8 8v4c0 2-2 4-4 4s4 2 4 4v4c0 5 2 8 8 8"
        stroke="currentColor"
        stroke-width="3"
        fill="none"
        stroke-linecap="round"
      />
      <path
        d="M29 8c6 0 8 3 8 8v4c0 2 2 4 4 4s-4 2-4 4v4c0 5-2 8-8 8"
        stroke="currentColor"
        stroke-width="3"
        fill="none"
        stroke-linecap="round"
      />
    </svg>

    <!-- Bedrock: mountain -->
    <svg
      v-else-if="variant === 'bedrock'"
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <path
        d="M4 38L18 14l8 12 4-6 10 18H4Z"
        fill="currentColor"
      />
    </svg>

    <!-- Fallback: S2 glyph -->
    <svg
      v-else
      viewBox="0 0 48 48"
      xmlns="http://www.w3.org/2000/svg"
      class="h-3/5 w-3/5"
      aria-hidden="true"
    >
      <path
        d="M31 16c-1.3-2.3-3.9-3.8-7.1-3.8-4.3 0-7.4 2.5-7.4 6 0 3.1 2.2 4.9 6.1 5.8l2.6 0.6c3.1 0.7 4.5 1.8 4.5 3.6 0 2.2-2 3.7-5 3.7-2.9 0-5-1.2-6.1-3.5"
        stroke="currentColor"
        stroke-width="3.2"
        stroke-linecap="round"
        fill="none"
      />
    </svg>
  </span>
</template>

<script setup lang="ts">
import { computed } from 'vue'

type ProviderVariant =
  | 'openai'
  | 'anthropic'
  | 'gemini'
  | 'google'
  | 'antigravity'
  | 'codex'
  | 'bedrock'
  | 'default'

type Props = {
  provider: string
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl'
  rounded?: 'full' | 'xl' | 'lg' | 'md' | 'none'
  /** Tailwind bg/text utility, e.g. `bg-orange-500 text-white`. Overrides auto. */
  classOverride?: string
}

const props = withDefaults(defineProps<Props>(), {
  size: 'md',
  rounded: 'xl',
  classOverride: ''
})

const variant = computed<ProviderVariant>(() => {
  switch ((props.provider || '').toLowerCase()) {
    case 'openai':
      return 'openai'
    case 'anthropic':
    case 'claude':
      return 'anthropic'
    case 'gemini':
    case 'google':
      return 'gemini'
    case 'antigravity':
      return 'antigravity'
    case 'codex':
      return 'codex'
    case 'bedrock':
      return 'bedrock'
    default:
      return 'default'
  }
})

const sizeClass = computed(() => {
  switch (props.size) {
    case 'xs':
      return 'h-5 w-5'
    case 'sm':
      return 'h-7 w-7'
    case 'md':
      return 'h-9 w-9'
    case 'lg':
      return 'h-12 w-12'
    case 'xl':
      return 'h-16 w-16'
    default:
      return 'h-9 w-9'
  }
})

const roundedClass = computed(() => {
  switch (props.rounded) {
    case 'full':
      return 'rounded-full'
    case 'xl':
      return 'rounded-xl'
    case 'lg':
      return 'rounded-lg'
    case 'md':
      return 'rounded-md'
    default:
      return ''
  }
})

const bgClass = computed(() => {
  if (props.classOverride) return props.classOverride
  switch (variant.value) {
    case 'openai':
      return 'bg-[#10a37f] text-white'
    case 'anthropic':
      return 'bg-[#d97757] text-white'
    case 'gemini':
      return 'bg-gradient-to-br from-[#4285f4] via-[#9b72ff] to-[#d96570] text-white'
    case 'antigravity':
      return 'bg-gradient-to-br from-indigo-500 to-purple-600 text-white'
    case 'codex':
      return 'bg-gradient-to-br from-slate-700 to-slate-900 text-white'
    case 'bedrock':
      return 'bg-gradient-to-br from-amber-500 to-orange-600 text-white'
    default:
      return 'bg-gradient-to-br from-primary-500 to-brand-500 text-white'
  }
})
</script>
