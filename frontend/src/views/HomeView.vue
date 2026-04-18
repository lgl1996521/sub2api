<template>
  <!-- Custom Home Content: Full Page Mode -->
  <div v-if="homeContent" class="min-h-screen">
    <!-- iframe mode -->
    <iframe
      v-if="isHomeContentUrl"
      :src="homeContent.trim()"
      class="h-screen w-full border-0"
      allowfullscreen
    ></iframe>
    <!-- HTML mode - SECURITY: homeContent is admin-only setting, XSS risk is acceptable -->
    <div v-else v-html="homeContent"></div>
  </div>

  <!-- Default Home Page -->
  <div
    v-else
    class="relative flex min-h-screen flex-col overflow-hidden bg-gradient-to-br from-gray-50 via-primary-50/30 to-gray-100 dark:from-dark-950 dark:via-dark-900 dark:to-dark-950"
  >
    <!-- Background Decorations -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div
        class="absolute -right-40 -top-40 h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"
      ></div>
      <div
        class="absolute -bottom-40 -left-40 h-96 w-96 rounded-full bg-primary-500/15 blur-3xl"
      ></div>
      <div
        class="absolute left-1/3 top-1/4 h-72 w-72 rounded-full bg-primary-300/10 blur-3xl"
      ></div>
      <div
        class="absolute bottom-1/4 right-1/4 h-64 w-64 rounded-full bg-primary-400/10 blur-3xl"
      ></div>
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(20,184,166,0.03)_1px,transparent_1px),linear-gradient(90deg,rgba(20,184,166,0.03)_1px,transparent_1px)] bg-[size:64px_64px]"
      ></div>
    </div>

    <!-- Header -->
    <header class="relative z-20 px-6 py-4">
      <nav class="mx-auto flex max-w-6xl items-center justify-between">
        <!-- Logo -->
        <div class="flex items-center">
          <div class="h-10 w-10 overflow-hidden rounded-xl shadow-md">
            <img :src="siteLogo || '/logo.png'" alt="Logo" class="h-full w-full object-contain" />
          </div>
        </div>

        <!-- Nav Actions -->
        <div class="flex items-center gap-3">
          <!-- Language Switcher -->
          <LocaleSwitcher />

          <!-- Doc Link -->
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:text-dark-400 dark:hover:bg-dark-800 dark:hover:text-white"
            :title="t('home.viewDocs')"
          >
            <Icon name="book" size="md" />
          </a>

          <!-- Theme Toggle -->
          <button
            @click="toggleTheme"
            class="rounded-lg p-2 text-gray-500 transition-colors hover:bg-gray-100 hover:text-gray-700 dark:text-dark-400 dark:hover:bg-dark-800 dark:hover:text-white"
            :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          >
            <Icon v-if="isDark" name="sun" size="md" />
            <Icon v-else name="moon" size="md" />
          </button>

          <!-- Login / Dashboard Button -->
          <router-link
            v-if="isAuthenticated"
            :to="dashboardPath"
            class="inline-flex items-center gap-1.5 rounded-full bg-gray-900 py-1 pl-1 pr-2.5 transition-colors hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700"
          >
            <span
              class="flex h-5 w-5 items-center justify-center rounded-full bg-gradient-to-br from-primary-400 to-primary-600 text-[10px] font-semibold text-white"
            >
              {{ userInitial }}
            </span>
            <span class="text-xs font-medium text-white">{{ t('home.dashboard') }}</span>
            <svg
              class="h-3 w-3 text-gray-400"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4.5 19.5l15-15m0 0H8.25m11.25 0v11.25"
              />
            </svg>
          </router-link>
          <router-link
            v-else
            to="/login"
            class="inline-flex items-center rounded-full bg-gray-900 px-3 py-1 text-xs font-medium text-white transition-colors hover:bg-gray-800 dark:bg-gray-800 dark:hover:bg-gray-700"
          >
            {{ t('home.login') }}
          </router-link>
        </div>
      </nav>
    </header>

    <!-- Main Content -->
    <main class="relative z-10 flex-1 px-6 py-16">
      <div class="mx-auto max-w-6xl">
        <!-- Hero Section - Left/Right Layout -->
        <div class="mb-12 flex flex-col items-center justify-between gap-12 lg:flex-row lg:gap-16">
          <!-- Left: Text Content -->
          <div class="flex-1 text-center lg:text-left">
            <h1
              class="mb-4 text-4xl font-bold text-gray-900 dark:text-white md:text-5xl lg:text-6xl"
            >
              {{ siteName }}
            </h1>
            <p class="mb-8 text-lg text-gray-600 dark:text-dark-300 md:text-xl">
              {{ siteSubtitle }}
            </p>

            <!-- CTA Button -->
            <div>
              <router-link
                :to="isAuthenticated ? dashboardPath : '/login'"
                class="btn btn-primary px-8 py-3 text-base shadow-lg shadow-primary-500/30"
              >
                {{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}
                <Icon name="arrowRight" size="md" class="ml-2" :stroke-width="2" />
              </router-link>
            </div>
          </div>

          <!-- Right: Terminal Animation -->
          <div class="flex flex-1 justify-center lg:justify-end">
            <div class="terminal-container">
              <div class="terminal-window">
                <!-- Window header -->
                <div class="terminal-header">
                  <div class="terminal-buttons">
                    <span class="btn-close"></span>
                    <span class="btn-minimize"></span>
                    <span class="btn-maximize"></span>
                  </div>
                  <span class="terminal-title">terminal</span>
                </div>
                <!-- Terminal content -->
                <div class="terminal-body">
                  <div class="code-line line-1">
                    <span class="code-prompt">$</span>
                    <span class="code-cmd">curl</span>
                    <span class="code-flag">-X POST</span>
                    <span class="code-url">/v1/messages</span>
                  </div>
                  <div class="code-line line-2">
                    <span class="code-comment"># Routing to upstream...</span>
                  </div>
                  <div class="code-line line-3">
                    <span class="code-success">200 OK</span>
                    <span class="code-response">{ "content": "Hello!" }</span>
                  </div>
                  <div class="code-line line-4">
                    <span class="code-prompt">$</span>
                    <span class="cursor"></span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Feature Tags - Centered -->
        <div class="mb-12 flex flex-wrap items-center justify-center gap-4 md:gap-6">
          <div
            class="inline-flex items-center gap-2.5 rounded-full border border-gray-200/50 bg-white/80 px-5 py-2.5 shadow-sm backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/80"
          >
            <Icon name="swap" size="sm" class="text-primary-500" />
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{
              t('home.tags.subscriptionToApi')
            }}</span>
          </div>
          <div
            class="inline-flex items-center gap-2.5 rounded-full border border-gray-200/50 bg-white/80 px-5 py-2.5 shadow-sm backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/80"
          >
            <Icon name="shield" size="sm" class="text-primary-500" />
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{
              t('home.tags.stickySession')
            }}</span>
          </div>
          <div
            class="inline-flex items-center gap-2.5 rounded-full border border-gray-200/50 bg-white/80 px-5 py-2.5 shadow-sm backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/80"
          >
            <Icon name="chart" size="sm" class="text-primary-500" />
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{
              t('home.tags.realtimeBilling')
            }}</span>
          </div>
        </div>

        <!-- Features Grid -->
        <div class="mb-12 grid gap-6 md:grid-cols-3">
          <!-- Feature 1: Unified Gateway -->
          <div
            class="group rounded-2xl border border-gray-200/50 bg-white/60 p-6 backdrop-blur-sm transition-all duration-300 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/50 dark:bg-dark-800/60"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-blue-500 to-blue-600 shadow-lg shadow-blue-500/30 transition-transform group-hover:scale-110"
            >
              <Icon name="server" size="lg" class="text-white" />
            </div>
            <h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('home.features.unifiedGateway') }}
            </h3>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ t('home.features.unifiedGatewayDesc') }}
            </p>
          </div>

          <!-- Feature 2: Account Pool -->
          <div
            class="group rounded-2xl border border-gray-200/50 bg-white/60 p-6 backdrop-blur-sm transition-all duration-300 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/50 dark:bg-dark-800/60"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-primary-500 to-primary-600 shadow-lg shadow-primary-500/30 transition-transform group-hover:scale-110"
            >
              <svg
                class="h-6 w-6 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="1.5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M18 18.72a9.094 9.094 0 003.741-.479 3 3 0 00-4.682-2.72m.94 3.198l.001.031c0 .225-.012.447-.037.666A11.944 11.944 0 0112 21c-2.17 0-4.207-.576-5.963-1.584A6.062 6.062 0 016 18.719m12 0a5.971 5.971 0 00-.941-3.197m0 0A5.995 5.995 0 0012 12.75a5.995 5.995 0 00-5.058 2.772m0 0a3 3 0 00-4.681 2.72 8.986 8.986 0 003.74.477m.94-3.197a5.971 5.971 0 00-.94 3.197M15 6.75a3 3 0 11-6 0 3 3 0 016 0zm6 3a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0zm-13.5 0a2.25 2.25 0 11-4.5 0 2.25 2.25 0 014.5 0z"
                />
              </svg>
            </div>
            <h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('home.features.multiAccount') }}
            </h3>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ t('home.features.multiAccountDesc') }}
            </p>
          </div>

          <!-- Feature 3: Billing & Quota -->
          <div
            class="group rounded-2xl border border-gray-200/50 bg-white/60 p-6 backdrop-blur-sm transition-all duration-300 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/50 dark:bg-dark-800/60"
          >
            <div
              class="mb-4 flex h-12 w-12 items-center justify-center rounded-xl bg-gradient-to-br from-purple-500 to-purple-600 shadow-lg shadow-purple-500/30 transition-transform group-hover:scale-110"
            >
              <svg
                class="h-6 w-6 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="1.5"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z"
                />
              </svg>
            </div>
            <h3 class="mb-2 text-lg font-semibold text-gray-900 dark:text-white">
              {{ t('home.features.balanceQuota') }}
            </h3>
            <p class="text-sm leading-relaxed text-gray-600 dark:text-dark-400">
              {{ t('home.features.balanceQuotaDesc') }}
            </p>
          </div>
        </div>

        <!-- Supported Providers -->
        <div class="mb-8 text-center">
          <h2 class="mb-3 text-2xl font-bold text-gray-900 dark:text-white">
            {{ t('home.providers.title') }}
          </h2>
          <p class="text-sm text-gray-600 dark:text-dark-400">
            {{ t('home.providers.description') }}
          </p>
        </div>

        <div class="mb-16 flex flex-wrap items-center justify-center gap-4">
          <!-- Claude - Supported -->
          <div
            class="flex items-center gap-2 rounded-xl border border-primary-200 bg-white/60 px-5 py-3 ring-1 ring-primary-500/20 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/60"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-orange-400 to-orange-500"
            >
              <span class="text-xs font-bold text-white">C</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.claude') }}</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- GPT - Supported -->
          <div
            class="flex items-center gap-2 rounded-xl border border-primary-200 bg-white/60 px-5 py-3 ring-1 ring-primary-500/20 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/60"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-green-500 to-green-600"
            >
              <span class="text-xs font-bold text-white">G</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">GPT</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- Gemini - Supported -->
          <div
            class="flex items-center gap-2 rounded-xl border border-primary-200 bg-white/60 px-5 py-3 ring-1 ring-primary-500/20 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/60"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-blue-500 to-blue-600"
            >
              <span class="text-xs font-bold text-white">G</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.gemini') }}</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- Antigravity - Supported -->
          <div
            class="flex items-center gap-2 rounded-xl border border-primary-200 bg-white/60 px-5 py-3 ring-1 ring-primary-500/20 backdrop-blur-sm dark:border-primary-800 dark:bg-dark-800/60"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-rose-500 to-pink-600"
            >
              <span class="text-xs font-bold text-white">A</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.antigravity') }}</span>
            <span
              class="rounded bg-primary-100 px-1.5 py-0.5 text-[10px] font-medium text-primary-600 dark:bg-primary-900/30 dark:text-primary-400"
              >{{ t('home.providers.supported') }}</span
            >
          </div>
          <!-- More - Coming Soon -->
          <div
            class="flex items-center gap-2 rounded-xl border border-gray-200/50 bg-white/40 px-5 py-3 opacity-60 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/40"
          >
            <div
              class="flex h-8 w-8 items-center justify-center rounded-lg bg-gradient-to-br from-gray-500 to-gray-600"
            >
              <span class="text-xs font-bold text-white">+</span>
            </div>
            <span class="text-sm font-medium text-gray-700 dark:text-dark-200">{{ t('home.providers.more') }}</span>
            <span
              class="rounded bg-gray-100 px-1.5 py-0.5 text-[10px] font-medium text-gray-500 dark:bg-dark-700 dark:text-dark-400"
              >{{ t('home.providers.soon') }}</span
            >
          </div>
        </div>

        <!-- Section: Model Health Status -->
        <section v-if="health.enabled && health.platforms.length" class="mb-16">
          <div class="mb-6 text-center">
            <h2 class="mb-2 text-2xl font-bold text-gray-900 dark:text-white">
              {{ t('home.health.title') }}
            </h2>
            <p class="text-sm text-gray-600 dark:text-dark-400">
              {{ t('home.health.description') }}
              <span v-if="health.collected_at" class="text-xs text-gray-400 dark:text-dark-500">
                · {{ t('home.health.lastUpdated') }} {{ formatTime(health.collected_at) }}
              </span>
            </p>
          </div>
          <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            <div
              v-for="item in health.platforms"
              :key="item.platform"
              class="relative rounded-2xl border border-gray-200/50 bg-white/70 p-5 backdrop-blur-sm dark:border-dark-700/50 dark:bg-dark-800/70"
            >
              <div class="mb-3 flex items-center justify-between">
                <h3 class="text-base font-semibold text-gray-900 dark:text-white">
                  {{ item.display_name }}
                </h3>
                <span
                  class="inline-flex items-center gap-1.5 rounded-full px-2.5 py-0.5 text-xs font-medium"
                  :class="healthBadgeClass(item.status)"
                >
                  <span
                    class="h-1.5 w-1.5 rounded-full"
                    :class="healthDotClass(item.status)"
                  ></span>
                  {{ t('home.health.status.' + item.status) }}
                </span>
              </div>
              <div class="grid grid-cols-3 gap-2 text-center">
                <div>
                  <div class="text-lg font-semibold text-gray-900 dark:text-white">
                    {{ item.available_count }}
                  </div>
                  <div class="text-[11px] text-gray-500 dark:text-dark-400">
                    {{ t('home.health.available') }}
                  </div>
                </div>
                <div>
                  <div class="text-lg font-semibold text-amber-500">
                    {{ item.rate_limit_count }}
                  </div>
                  <div class="text-[11px] text-gray-500 dark:text-dark-400">
                    {{ t('home.health.rateLimited') }}
                  </div>
                </div>
                <div>
                  <div class="text-lg font-semibold text-rose-500">
                    {{ item.error_count }}
                  </div>
                  <div class="text-[11px] text-gray-500 dark:text-dark-400">
                    {{ t('home.health.errored') }}
                  </div>
                </div>
              </div>
              <div class="mt-3 h-1.5 overflow-hidden rounded-full bg-gray-100 dark:bg-dark-700">
                <div
                  class="h-full rounded-full transition-all"
                  :class="healthBarClass(item.status)"
                  :style="{ width: healthRatio(item) + '%' }"
                ></div>
              </div>
              <div class="mt-2 text-right text-[11px] text-gray-400 dark:text-dark-500">
                {{ t('home.health.ofTotal', { total: item.total_accounts }) }}
              </div>
            </div>
          </div>
        </section>

        <!-- Section: Subscription Plans -->
        <section v-if="plans.length" class="mb-16">
          <div class="mb-6 text-center">
            <h2 class="mb-2 text-2xl font-bold text-gray-900 dark:text-white">
              {{ t('home.plans.title') }}
            </h2>
            <p class="text-sm text-gray-600 dark:text-dark-400">
              {{ t('home.plans.description') }}
            </p>
          </div>
          <div class="grid gap-5 sm:grid-cols-2 lg:grid-cols-3">
            <div
              v-for="plan in plans"
              :key="plan.id"
              class="group flex flex-col rounded-2xl border border-gray-200/60 bg-white/80 p-6 shadow-sm backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:shadow-xl hover:shadow-primary-500/10 dark:border-dark-700/60 dark:bg-dark-800/80"
            >
              <div class="mb-3 flex items-center justify-between">
                <span
                  class="rounded-md bg-primary-100 px-2 py-0.5 text-[11px] font-medium uppercase tracking-wide text-primary-700 dark:bg-primary-900/30 dark:text-primary-300"
                >
                  {{ plan.group_platform || plan.group_name }}
                </span>
                <span
                  v-if="plan.rate_multiplier && plan.rate_multiplier !== 1"
                  class="text-[11px] text-gray-500 dark:text-dark-400"
                >
                  × {{ plan.rate_multiplier }}
                </span>
              </div>
              <h3 class="mb-1 text-lg font-semibold text-gray-900 dark:text-white">
                {{ plan.name }}
              </h3>
              <p
                v-if="plan.description"
                class="mb-4 text-sm text-gray-600 dark:text-dark-400"
              >
                {{ plan.description }}
              </p>
              <div class="mb-4 flex items-baseline gap-2">
                <span class="text-3xl font-bold text-gray-900 dark:text-white">
                  ¥{{ formatPrice(plan.price) }}
                </span>
                <span
                  v-if="plan.original_price && plan.original_price > plan.price"
                  class="text-sm text-gray-400 line-through"
                >
                  ¥{{ formatPrice(plan.original_price) }}
                </span>
                <span class="text-xs text-gray-500 dark:text-dark-400">
                  / {{ plan.validity_days }} {{ planValidityUnit(plan) }}
                </span>
              </div>
              <ul
                v-if="plan.features.length"
                class="mb-4 space-y-1.5 text-sm text-gray-600 dark:text-dark-400"
              >
                <li
                  v-for="(feat, idx) in plan.features.slice(0, 4)"
                  :key="idx"
                  class="flex items-start gap-2"
                >
                  <Icon name="check" size="sm" class="mt-0.5 shrink-0 text-primary-500" />
                  <span>{{ feat }}</span>
                </li>
              </ul>
              <div class="mt-auto">
                <router-link
                  :to="isAuthenticated ? '/payment/checkout?plan_id=' + plan.id : '/login'"
                  class="btn btn-primary w-full"
                >
                  {{ t('home.plans.subscribe') }}
                </router-link>
              </div>
            </div>
          </div>
        </section>

        <!-- Section: Server Lines -->
        <section v-if="serverLines.length" class="mb-16">
          <div class="mb-6 text-center">
            <h2 class="mb-2 text-2xl font-bold text-gray-900 dark:text-white">
              {{ t('home.lines.title') }}
            </h2>
            <p class="text-sm text-gray-600 dark:text-dark-400">
              {{ t('home.lines.description') }}
            </p>
          </div>
          <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            <div
              v-for="line in serverLines"
              :key="line.id"
              class="group relative flex flex-col rounded-2xl border p-5 backdrop-blur-sm transition-all duration-300 hover:-translate-y-0.5 hover:shadow-xl"
              :class="
                selectedLineId === line.id
                  ? 'border-primary-400 bg-primary-50/60 ring-2 ring-primary-400/50 dark:border-primary-500 dark:bg-primary-900/20'
                  : 'border-gray-200/60 bg-white/70 dark:border-dark-700/60 dark:bg-dark-800/70'
              "
            >
              <div class="mb-3 flex items-start justify-between gap-3">
                <div class="flex items-center gap-2">
                  <span class="text-2xl leading-none">{{ regionEmoji(line.region) }}</span>
                  <div>
                    <h3 class="text-base font-semibold text-gray-900 dark:text-white">
                      {{ line.name }}
                    </h3>
                    <p class="text-[11px] uppercase tracking-wide text-gray-500 dark:text-dark-400">
                      {{ regionLabel(line.region) }}
                    </p>
                  </div>
                </div>
                <span
                  class="inline-flex items-center gap-1.5 rounded-full px-2.5 py-0.5 text-xs font-medium"
                  :class="lineBadgeClass(line.status)"
                >
                  <span class="h-1.5 w-1.5 rounded-full" :class="lineDotClass(line.status)"></span>
                  {{ t('home.lines.status.' + line.status) }}
                </span>
              </div>
              <p
                v-if="line.description"
                class="mb-3 text-sm text-gray-600 dark:text-dark-400"
              >
                {{ line.description }}
              </p>
              <div class="mb-4 flex items-center justify-between text-sm">
                <span class="text-gray-500 dark:text-dark-400">
                  {{ t('home.lines.latency') }}
                </span>
                <span
                  class="font-mono font-medium"
                  :class="latencyColorClass(line.latency_ms, line.status)"
                >
                  {{ line.status === 'unknown' ? '—' : line.latency_ms + ' ms' }}
                </span>
              </div>
              <div class="mt-auto flex items-center gap-2">
                <button
                  type="button"
                  class="btn btn-secondary flex-1 text-xs"
                  :class="{
                    'btn-primary': selectedLineId === line.id
                  }"
                  @click="selectLine(line)"
                >
                  {{
                    selectedLineId === line.id
                      ? t('home.lines.selected')
                      : t('home.lines.useLine')
                  }}
                </button>
                <a
                  :href="line.url"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="rounded-lg border border-gray-200 p-2 text-gray-500 transition-colors hover:bg-gray-50 dark:border-dark-700 dark:text-dark-400 dark:hover:bg-dark-700/50"
                  :title="line.url"
                >
                  <Icon name="link" size="sm" />
                </a>
              </div>
            </div>
          </div>
        </section>
      </div>
    </main>

    <!-- Footer -->
    <footer class="relative z-10 border-t border-gray-200/50 px-6 py-8 dark:border-dark-800/50">
      <div
        class="mx-auto flex max-w-6xl flex-col items-center justify-center gap-4 text-center sm:flex-row sm:text-left"
      >
        <p class="text-sm text-gray-500 dark:text-dark-400">
          &copy; {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}
        </p>
        <div class="flex items-center gap-4">
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            {{ t('home.docs') }}
          </a>
          <a
            :href="githubUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="text-sm text-gray-500 transition-colors hover:text-gray-700 dark:text-dark-400 dark:hover:text-white"
          >
            GitHub
          </a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onBeforeUnmount, reactive } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'
import {
  publicAPI,
  type PublicPlan,
  type PublicHealthResponse,
  type PublicServerLine,
  type PlatformHealthStatus,
  type ServerLineStatus
} from '@/api/public'

const { t, locale } = useI18n()

const authStore = useAuthStore()
const appStore = useAppStore()

// Site settings - directly from appStore (already initialized from injected config)
const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'Sub2API')
const siteLogo = computed(() => appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '')
const siteSubtitle = computed(() => appStore.cachedPublicSettings?.site_subtitle || 'AI API Gateway Platform')
const docUrl = computed(() => appStore.cachedPublicSettings?.doc_url || appStore.docUrl || '')
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')

// Check if homeContent is a URL (for iframe display)
const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

// Theme
const isDark = ref(document.documentElement.classList.contains('dark'))

// GitHub URL
const githubUrl = 'https://github.com/Wei-Shaw/sub2api'

// Auth state
const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => isAdmin.value ? '/admin/dashboard' : '/dashboard')
const userInitial = computed(() => {
  const user = authStore.user
  if (!user || !user.email) return ''
  return user.email.charAt(0).toUpperCase()
})

// Current year for footer
const currentYear = computed(() => new Date().getFullYear())

// Toggle theme
function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

// Initialize theme
function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (
    savedTheme === 'dark' ||
    (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)
  ) {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
}

// ---- Homepage dynamic sections: health / plans / server lines ----

const health = reactive<PublicHealthResponse>({
  enabled: false,
  collected_at: undefined,
  platforms: []
})

const plans = ref<PublicPlan[]>([])
const serverLines = ref<PublicServerLine[]>([])
const selectedLineId = ref<string>(localStorage.getItem('preferredServerLine') || '')

let healthTimer: ReturnType<typeof setInterval> | null = null
let linesTimer: ReturnType<typeof setInterval> | null = null

async function loadHealth() {
  try {
    const resp = await publicAPI.getModelHealth()
    const data = (resp.data ?? resp) as unknown as PublicHealthResponse
    health.enabled = Boolean(data.enabled)
    health.collected_at = data.collected_at
    health.platforms = data.platforms ?? []
  } catch {
    // Public endpoint, ignore transient failures.
  }
}

async function loadPlans() {
  try {
    const resp = await publicAPI.getPlans()
    const data = (resp.data ?? resp) as unknown as PublicPlan[]
    plans.value = Array.isArray(data) ? data : []
  } catch {
    plans.value = []
  }
}

async function loadServerLines() {
  try {
    const resp = await publicAPI.getServerLines()
    const data = (resp.data ?? resp) as unknown as PublicServerLine[]
    serverLines.value = Array.isArray(data) ? data : []
    if (selectedLineId.value && !serverLines.value.some((l) => l.id === selectedLineId.value)) {
      selectedLineId.value = ''
      localStorage.removeItem('preferredServerLine')
    }
  } catch {
    serverLines.value = []
  }
}

function selectLine(line: PublicServerLine) {
  if (selectedLineId.value === line.id) {
    selectedLineId.value = ''
    localStorage.removeItem('preferredServerLine')
    return
  }
  selectedLineId.value = line.id
  try {
    localStorage.setItem('preferredServerLine', line.id)
    localStorage.setItem('preferredServerLineUrl', line.url)
  } catch {
    // ignore quota errors
  }
}

function healthBadgeClass(status: PlatformHealthStatus): string {
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

function healthDotClass(status: PlatformHealthStatus): string {
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

function healthBarClass(status: PlatformHealthStatus): string {
  switch (status) {
    case 'healthy':
      return 'bg-emerald-500'
    case 'degraded':
      return 'bg-amber-500'
    case 'down':
      return 'bg-rose-500'
    default:
      return 'bg-gray-300 dark:bg-dark-600'
  }
}

function healthRatio(item: { available_count: number; total_accounts: number }): number {
  if (!item.total_accounts) return 0
  return Math.min(100, Math.round((item.available_count / item.total_accounts) * 100))
}

function lineBadgeClass(status: ServerLineStatus): string {
  return healthBadgeClass(status as PlatformHealthStatus)
}

function lineDotClass(status: ServerLineStatus): string {
  return healthDotClass(status as PlatformHealthStatus)
}

function latencyColorClass(latency: number, status: ServerLineStatus): string {
  if (status === 'down' || status === 'unknown') return 'text-rose-500'
  if (latency >= 1000) return 'text-amber-500'
  if (latency >= 400) return 'text-amber-600'
  return 'text-emerald-600 dark:text-emerald-400'
}

function regionEmoji(region: string): string {
  const key = (region || '').toLowerCase()
  switch (key) {
    case 'cn':
    case 'china':
    case 'domestic':
      return '🇨🇳'
    case 'hk':
      return '🇭🇰'
    case 'tw':
      return '🇹🇼'
    case 'jp':
      return '🇯🇵'
    case 'sg':
      return '🇸🇬'
    case 'us':
    case 'usa':
      return '🇺🇸'
    case 'eu':
      return '🇪🇺'
    case 'intl':
    case 'global':
    case 'international':
      return '🌐'
    default:
      return '🌐'
  }
}

function regionLabel(region: string): string {
  const key = (region || '').toLowerCase()
  const known = ['cn', 'hk', 'tw', 'jp', 'sg', 'us', 'eu', 'intl', 'global']
  if (known.includes(key)) {
    return t('home.lines.region.' + key)
  }
  return region
}

function formatPrice(value: number | null | undefined): string {
  if (value == null) return '0'
  if (Number.isInteger(value)) return String(value)
  return value.toFixed(2)
}

function planValidityUnit(plan: PublicPlan): string {
  const unit = (plan.validity_unit || 'day').toLowerCase()
  const key = 'home.plans.validity.' + unit
  const translated = t(key)
  return translated === key ? unit : translated
}

function formatTime(value?: string): string {
  if (!value) return ''
  try {
    const d = new Date(value)
    return d.toLocaleTimeString(locale.value || undefined, {
      hour: '2-digit',
      minute: '2-digit'
    })
  } catch {
    return value
  }
}

onMounted(() => {
  initTheme()

  // Check auth state
  authStore.checkAuth()

  // Ensure public settings are loaded (will use cache if already loaded from injected config)
  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }

  // Load dynamic homepage sections (best-effort, public endpoints).
  loadHealth()
  loadPlans()
  loadServerLines()

  healthTimer = setInterval(loadHealth, 60_000)
  linesTimer = setInterval(loadServerLines, 90_000)
})

onBeforeUnmount(() => {
  if (healthTimer) clearInterval(healthTimer)
  if (linesTimer) clearInterval(linesTimer)
})
</script>

<style scoped>
/* Terminal Container */
.terminal-container {
  position: relative;
  display: inline-block;
}

/* Terminal Window */
.terminal-window {
  width: 420px;
  background: linear-gradient(145deg, #1e293b 0%, #0f172a 100%);
  border-radius: 14px;
  box-shadow:
    0 25px 50px -12px rgba(0, 0, 0, 0.4),
    0 0 0 1px rgba(255, 255, 255, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
  overflow: hidden;
  transform: perspective(1000px) rotateX(2deg) rotateY(-2deg);
  transition: transform 0.3s ease;
}

.terminal-window:hover {
  transform: perspective(1000px) rotateX(0deg) rotateY(0deg) translateY(-4px);
}

/* Terminal Header */
.terminal-header {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  background: rgba(30, 41, 59, 0.8);
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.terminal-buttons {
  display: flex;
  gap: 8px;
}

.terminal-buttons span {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.btn-close {
  background: #ef4444;
}
.btn-minimize {
  background: #eab308;
}
.btn-maximize {
  background: #22c55e;
}

.terminal-title {
  flex: 1;
  text-align: center;
  font-size: 12px;
  font-family: ui-monospace, monospace;
  color: #64748b;
  margin-right: 52px;
}

/* Terminal Body */
.terminal-body {
  padding: 20px 24px;
  font-family: ui-monospace, 'Fira Code', monospace;
  font-size: 14px;
  line-height: 2;
}

.code-line {
  display: flex;
  align-items: center;
  gap: 8px;
  flex-wrap: wrap;
  opacity: 0;
  animation: line-appear 0.5s ease forwards;
}

.line-1 {
  animation-delay: 0.3s;
}
.line-2 {
  animation-delay: 1s;
}
.line-3 {
  animation-delay: 1.8s;
}
.line-4 {
  animation-delay: 2.5s;
}

@keyframes line-appear {
  from {
    opacity: 0;
    transform: translateY(5px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.code-prompt {
  color: #22c55e;
  font-weight: bold;
}
.code-cmd {
  color: #38bdf8;
}
.code-flag {
  color: #a78bfa;
}
.code-url {
  color: #14b8a6;
}
.code-comment {
  color: #64748b;
  font-style: italic;
}
.code-success {
  color: #22c55e;
  background: rgba(34, 197, 94, 0.15);
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}
.code-response {
  color: #fbbf24;
}

/* Blinking Cursor */
.cursor {
  display: inline-block;
  width: 8px;
  height: 16px;
  background: #22c55e;
  animation: blink 1s step-end infinite;
}

@keyframes blink {
  0%,
  50% {
    opacity: 1;
  }
  51%,
  100% {
    opacity: 0;
  }
}

/* Dark mode adjustments */
:deep(.dark) .terminal-window {
  box-shadow:
    0 25px 50px -12px rgba(0, 0, 0, 0.6),
    0 0 0 1px rgba(20, 184, 166, 0.2),
    0 0 40px rgba(20, 184, 166, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.1);
}
</style>
