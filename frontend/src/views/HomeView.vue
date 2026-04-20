<template>
  <div class="relative flex min-h-screen flex-col overflow-hidden bg-gradient-to-br from-gray-950 via-slate-900 to-gray-950 dark:from-gray-950 dark:via-slate-900 dark:to-gray-950">
    <!-- Background Decorations -->
    <div class="pointer-events-none absolute inset-0 overflow-hidden">
      <div class="absolute -right-40 -top-40 h-96 w-96 rounded-full bg-primary-500/10 blur-3xl"></div>
      <div class="absolute -bottom-40 -left-40 h-96 w-96 rounded-full bg-brand-500/10 blur-3xl"></div>
      <div class="absolute left-1/3 top-1/4 h-72 w-72 rounded-full bg-accent-500/5 blur-3xl"></div>
      <div class="absolute bottom-1/4 right-1/4 h-64 w-64 rounded-full bg-primary-400/5 blur-3xl"></div>
      <div class="absolute inset-0 bg-[linear-gradient(rgba(20,184,166,0.03)_1px,transparent_1px),linear-gradient(90deg,rgba(20,184,166,0.03)_1px,transparent_1px)] bg-[size:64px_64px]"></div>
    </div>

    <!-- Main Content -->
    <main class="relative z-10 flex-1 px-6 py-16">
      <div class="mx-auto max-w-6xl">
        <!-- Hero Section -->
        <section class="mb-20 flex flex-col items-center justify-between gap-12 lg:flex-row lg:gap-16">
          <!-- Left: Text Content -->
          <div class="flex-1 text-center lg:text-left">
            <!-- Terminal Badge -->
            <div class="mb-6 inline-flex items-center gap-2 rounded-full border border-gray-700/50 bg-gray-900/50 px-4 py-2 backdrop-blur-sm">
              <span class="text-sm font-mono text-gray-400">$</span>
              <span class="text-sm font-mono text-primary-400">sub2api</span>
              <span class="text-sm font-mono text-gray-500">--mirror</span>
              <span class="text-sm font-mono text-brand-400">--china</span>
            </div>

            <h1 class="mb-4 text-4xl font-bold text-white md:text-5xl lg:text-6xl">
              AI API <span class="bg-gradient-to-r from-primary-400 to-brand-400 bg-clip-text text-transparent">国内解决方案</span>
            </h1>
            <p class="mb-8 text-lg text-gray-400 md:text-xl">
              即刻体验 OpenAI、Claude、Gemini 等全球顶尖 AI 模型；智能路由技术，降低封号风险，国内直连。
            </p>

            <!-- CTA Buttons -->
            <div class="mb-8 flex flex-wrap items-center justify-center gap-4 lg:justify-start">
              <router-link
                :to="isAuthenticated ? dashboardPath : '/register'"
                class="inline-flex items-center gap-2 rounded-full bg-gradient-to-r from-primary-500 to-brand-500 px-6 py-3 font-medium text-white shadow-lg shadow-primary-500/25 transition-all hover:shadow-primary-500/40"
              >
                {{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}
                <Icon name="arrowRight" size="sm" />
              </router-link>
              <a
                href="https://docs.sub2api.com"
                target="_blank"
                rel="noopener noreferrer"
                class="inline-flex items-center gap-2 rounded-full border border-gray-700 bg-gray-900/50 px-6 py-3 font-medium text-gray-300 transition-all hover:border-gray-600 hover:bg-gray-800"
              >
                <Icon name="book" size="sm" />
                官方文档
              </a>
            </div>

            <!-- Stats -->
            <div class="flex flex-wrap items-center justify-center gap-8 lg:justify-start">
              <div class="text-center">
                <div class="text-2xl font-bold text-white">{{ serviceDays }}天</div>
                <div class="text-sm text-gray-500">稳定服务时间</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-bold text-white">20,000+</div>
                <div class="text-sm text-gray-500">服务开发者</div>
              </div>
              <div class="text-center">
                <div class="text-2xl font-bold text-white">99.9%</div>
                <div class="text-sm text-gray-500">稳定运行率</div>
              </div>
            </div>
          </div>

          <!-- Right: Terminal Window -->
          <div class="flex flex-1 justify-center lg:justify-end">
            <div class="w-full max-w-lg overflow-hidden rounded-xl border border-gray-700/50 bg-gray-900/90 shadow-2xl backdrop-blur-sm">
              <!-- Window Header -->
              <div class="flex items-center gap-2 border-b border-gray-800 px-4 py-3">
                <div class="flex gap-1.5">
                  <span class="h-3 w-3 rounded-full bg-rose-500/80"></span>
                  <span class="h-3 w-3 rounded-full bg-amber-500/80"></span>
                  <span class="h-3 w-3 rounded-full bg-emerald-500/80"></span>
                </div>
                <span class="ml-3 font-mono text-xs text-gray-500">quickstart.sh</span>
              </div>
              <!-- Terminal Content -->
              <div class="p-4 font-mono text-sm">
                <div class="mb-2 text-gray-500"># 安装 Sub2API CLI（推荐）</div>
                <div class="mb-2">
                  <span class="text-primary-400">$</span>
                  <span class="text-gray-300"> curl -fsSL https://sub2api.com/install.sh | bash</span>
                </div>
                <div class="mb-2 text-gray-500"># 一键安装（自动配置 Token / 节点）</div>
                <div class="mb-2">
                  <span class="text-primary-400">$</span>
                  <span class="text-gray-300"> API_KEY="&lt;你的 Key&gt;" \</span>
                </div>
                <div class="mb-2 pl-4 text-gray-300">BASE_URL="&lt;节点&gt;" \</div>
                <div class="mb-2 pl-4 text-gray-300">bash -c "$(curl -fsSL sub2api.com/install.sh)"</div>
                <div class="mb-2 text-gray-500"># 开始对话</div>
                <div>
                  <span class="text-primary-400">$</span>
                  <span class="text-gray-300"> sub2api "帮我写一个 Python 爬虫"</span>
                </div>
              </div>
            </div>
          </div>
        </section>

        <!-- Integrations Section -->
        <section id="integrations" class="mb-20 scroll-mt-24">
          <div class="mb-8 text-center">
            <div class="mb-2 text-xs font-semibold uppercase tracking-wider text-primary-400">Integrations</div>
            <h2 class="text-2xl font-bold text-white md:text-3xl">兼容官方 SDK，统一入口接入</h2>
          </div>
          <div class="grid gap-4 sm:grid-cols-3">
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-emerald-500 to-teal-600 text-xl font-bold text-white">O</div>
              <div class="font-semibold text-white">OpenAI</div>
              <div class="text-sm text-gray-500">GPT-4 / GPT-5 全系列</div>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-orange-500 to-amber-600 text-xl font-bold text-white">C</div>
              <div class="font-semibold text-white">Claude</div>
              <div class="text-sm text-gray-500">Anthropic 全系列</div>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-gradient-to-br from-blue-500 to-indigo-600 text-xl font-bold text-white">G</div>
              <div class="font-semibold text-white">Gemini</div>
              <div class="text-sm text-gray-500">Google AI 全系列</div>
            </div>
          </div>
        </section>

        <!-- Features Section -->
        <section id="features" class="mb-20 scroll-mt-24">
          <div class="mb-8 text-center">
            <div class="mb-2 text-xs font-semibold uppercase tracking-wider text-primary-400">Features</div>
            <h2 class="mb-2 text-2xl font-bold text-white md:text-3xl">开发者工具，专为国内链路优化</h2>
            <p class="text-gray-500">更少的花销，更快的路径：镜像、代理、计费、统计，一处解决。</p>
          </div>
          <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-primary-500/20 text-primary-400">
                <Icon name="shield" size="md" />
              </div>
              <h3 class="mb-2 font-semibold text-white">零封号风险</h3>
              <p class="text-sm text-gray-500">企业级号池代理技术，我们来解决网络访问与稳定性。</p>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-brand-500/20 text-brand-400">
                <Icon name="download" size="md" />
              </div>
              <h3 class="mb-2 font-semibold text-white">一键安装</h3>
              <p class="text-sm text-gray-500">脚本自动配置 Key 与节点，三步即可开始使用。</p>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-emerald-500/20 text-emerald-400">
                <Icon name="creditCard" size="md" />
              </div>
              <h3 class="mb-2 font-semibold text-white">灵活计费</h3>
              <p class="text-sm text-gray-500">包月套餐适合高频使用，按量付费用多少扣多少。</p>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-amber-500/20 text-amber-400">
                <Icon name="chart" size="md" />
              </div>
              <h3 class="mb-2 font-semibold text-white">缓存命中率透明</h3>
              <p class="text-sm text-gray-500">后台清晰展示缓存命中率，降低成本，拒绝套路。</p>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-purple-500/20 text-purple-400">
                <Icon name="sparkles" size="md" />
              </div>
              <h3 class="mb-2 font-semibold text-white">特有功能</h3>
              <p class="text-sm text-gray-500">支持多模型负载均衡、自动重试、智能路由。</p>
            </div>
            <div class="group cursor-pointer rounded-xl border border-gray-800 bg-gray-900/50 p-6 transition-all hover:border-primary-500/50 hover:bg-gray-800/50">
              <div class="mb-4 flex h-12 w-12 items-center justify-center rounded-lg bg-rose-500/20 text-rose-400">
                <Icon name="eye" size="md" />
              </div>
              <h3 class="mb-2 font-semibold text-white">使用统计</h3>
              <p class="text-sm text-gray-500">实时追踪用量与费用，记录可追溯，计费更透明。</p>
            </div>
          </div>
        </section>

        <!-- Comparison Table -->
        <section id="compare" class="mb-20 scroll-mt-24">
          <div class="mb-8 text-center">
            <div class="mb-2 text-xs font-semibold uppercase tracking-wider text-primary-400">Comparison</div>
            <h2 class="mb-2 text-2xl font-bold text-white md:text-3xl">OpenAI · Claude · Gemini</h2>
            <p class="text-gray-500">快速选型参考：更偏向的工作流与适用场景。</p>
          </div>
          <div class="overflow-x-auto rounded-xl border border-gray-800 bg-gray-900/50">
            <table class="w-full text-left text-sm">
              <thead>
                <tr class="border-b border-gray-800">
                  <th class="px-6 py-4 font-semibold text-gray-400">对比项</th>
                  <th class="px-6 py-4 font-semibold text-white">OpenAI</th>
                  <th class="px-6 py-4 font-semibold text-white">Claude</th>
                  <th class="px-6 py-4 font-semibold text-white">Gemini</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-800">
                <tr>
                  <td class="px-6 py-4 font-medium text-gray-400">定位</td>
                  <td class="px-6 py-4 text-gray-300">通用对话 / 代码生成 / 多模态</td>
                  <td class="px-6 py-4 text-gray-300">长文本 / 复杂推理 / 代码理解</td>
                  <td class="px-6 py-4 text-gray-300">检索 / 多模态 / Google 生态</td>
                </tr>
                <tr>
                  <td class="px-6 py-4 font-medium text-gray-400">常用模型</td>
                  <td class="px-6 py-4 text-gray-300">GPT-5 / GPT-4o</td>
                  <td class="px-6 py-4 text-gray-300">Claude 4 Opus / Sonnet</td>
                  <td class="px-6 py-4 text-gray-300">Gemini 2.5 Pro / Flash</td>
                </tr>
                <tr>
                  <td class="px-6 py-4 font-medium text-gray-400">适合场景</td>
                  <td class="px-6 py-4 text-gray-300">通用 AI 应用 / 快速原型</td>
                  <td class="px-6 py-4 text-gray-300">大型项目 / 文档分析 / 代码审查</td>
                  <td class="px-6 py-4 text-gray-300">搜索增强 / 多模态应用</td>
                </tr>
                <tr>
                  <td class="px-6 py-4 font-medium text-gray-400">本平台能力</td>
                  <td class="px-6 py-4 text-primary-400">国内直连 · 智能路由 · 用量统计</td>
                  <td class="px-6 py-4 text-primary-400">国内直连 · 代理号池 · 用量统计</td>
                  <td class="px-6 py-4 text-primary-400">国内直连 · 统一入口 · 用量统计</td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>

        <!-- Pricing Section -->
        <section id="pricing" class="mb-20 scroll-mt-24">
          <div class="mb-8 text-center">
            <div class="mb-2 text-xs font-semibold uppercase tracking-wider text-primary-400">Pricing</div>
            <h2 class="mb-2 text-2xl font-bold text-white md:text-3xl">灵活的套餐选择</h2>
            <p class="text-gray-500">选择适合您的套餐方案 所有套餐均可叠加使用</p>
          </div>

          <!-- Pricing Tabs -->
          <div class="mb-8 flex justify-center">
            <div class="inline-flex rounded-full border border-gray-800 bg-gray-900/50 p-1">
              <button
                type="button"
                class="rounded-full px-6 py-2 text-sm font-medium transition-all"
                :class="pricingTab === 'monthly' ? 'bg-primary-500 text-white' : 'text-gray-400 hover:text-white'"
                @click="pricingTab = 'monthly'"
              >
                包月套餐
              </button>
              <button
                type="button"
                class="rounded-full px-6 py-2 text-sm font-medium transition-all"
                :class="pricingTab === 'paygo' ? 'bg-primary-500 text-white' : 'text-gray-400 hover:text-white'"
                @click="pricingTab = 'paygo'"
              >
                PayGO套餐
              </button>
            </div>
          </div>

          <!-- Monthly Plans -->
          <div v-if="pricingTab === 'monthly'" class="grid gap-6 md:grid-cols-3">
            <div class="rounded-xl border border-gray-800 bg-gray-900/50 p-6">
              <div class="mb-2 inline-block rounded-full bg-primary-500/10 px-3 py-1 text-xs font-medium text-primary-400">限时</div>
              <h3 class="mb-1 text-xl font-bold text-white">基础月卡</h3>
              <p class="mb-4 text-sm text-gray-500">性价比之选</p>
              <div class="mb-4">
                <span class="text-3xl font-bold text-white">¥99</span>
                <span class="text-gray-500">/月</span>
                <span class="ml-2 text-sm text-gray-500">≈ 0.33¥/$</span>
              </div>
              <div class="mb-6 text-sm text-primary-400">月总额度 $300</div>
              <ul class="mb-6 space-y-3 text-sm text-gray-400">
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  每日0点重置到 $37.5
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  周限额度 $75
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  支持所有模型
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  有效期 30 天
                </li>
              </ul>
              <button class="w-full rounded-full bg-primary-500 py-2.5 font-medium text-white transition-all hover:bg-primary-600">
                立即购买
              </button>
            </div>

            <div class="relative rounded-xl border-2 border-primary-500 bg-gray-900/50 p-6">
              <div class="absolute -top-3 left-1/2 -translate-x-1/2 rounded-full bg-primary-500 px-3 py-1 text-xs font-medium text-white">推荐</div>
              <div class="mb-2 inline-block rounded-full bg-primary-500/10 px-3 py-1 text-xs font-medium text-primary-400">限时</div>
              <h3 class="mb-1 text-xl font-bold text-white">专业月卡</h3>
              <p class="mb-4 text-sm text-gray-500">畅享无忧</p>
              <div class="mb-4">
                <span class="text-3xl font-bold text-white">¥249</span>
                <span class="text-gray-500">/月</span>
                <span class="ml-2 text-sm text-gray-500">≈ 0.28¥/$</span>
              </div>
              <div class="mb-6 text-sm text-primary-400">月总额度 $900</div>
              <ul class="mb-6 space-y-3 text-sm text-gray-400">
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  每日0点重置到 $112.5
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  周限额度 $225
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  支持所有模型
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  有效期 30 天
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  优先客服支持
                </li>
              </ul>
              <button class="w-full rounded-full bg-primary-500 py-2.5 font-medium text-white transition-all hover:bg-primary-600">
                立即购买
              </button>
            </div>

            <div class="rounded-xl border border-gray-800 bg-gray-900/50 p-6">
              <div class="mb-2 inline-block rounded-full bg-primary-500/10 px-3 py-1 text-xs font-medium text-primary-400">限时</div>
              <h3 class="mb-1 text-xl font-bold text-white">企业月卡</h3>
              <p class="mb-4 text-sm text-gray-500">超大杯</p>
              <div class="mb-4">
                <span class="text-3xl font-bold text-white">¥699</span>
                <span class="text-gray-500">/月</span>
                <span class="ml-2 text-sm text-gray-500">≈ 0.23¥/$</span>
              </div>
              <div class="mb-6 text-sm text-primary-400">月总额度 $3000</div>
              <ul class="mb-6 space-y-3 text-sm text-gray-400">
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  每日0点重置到 $375
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  周限额度 $750
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  支持所有模型
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  有效期 30 天
                </li>
                <li class="flex items-center gap-2">
                  <Icon name="check" size="xs" class="text-emerald-500" />
                  专属客服支持
                </li>
              </ul>
              <button class="w-full rounded-full bg-primary-500 py-2.5 font-medium text-white transition-all hover:bg-primary-600">
                立即购买
              </button>
            </div>
          </div>

          <!-- PayGO Plans -->
          <div v-else class="rounded-xl border border-gray-800 bg-gray-900/50 p-8 text-center">
            <div class="mx-auto max-w-2xl">
              <h3 class="mb-4 text-2xl font-bold text-white">按量付费</h3>
              <p class="mb-6 text-gray-400">充值 1 元 = 1 美元额度，用多少扣多少，无时间限制</p>
              <div class="mb-8 flex items-center justify-center gap-4">
                <div class="text-center">
                  <div class="text-4xl font-bold text-primary-400">1:1</div>
                  <div class="text-sm text-gray-500">人民币:美元</div>
                </div>
                <div class="text-2xl text-gray-700">=</div>
                <div class="text-center">
                  <div class="text-4xl font-bold text-emerald-400">86%</div>
                  <div class="text-sm text-gray-500">节省费用</div>
                </div>
              </div>
              <div class="flex flex-wrap justify-center gap-4">
                <button class="rounded-full bg-primary-500 px-8 py-3 font-medium text-white transition-all hover:bg-primary-600">
                  充值 100 元
                </button>
                <button class="rounded-full border border-gray-700 bg-gray-800 px-8 py-3 font-medium text-gray-300 transition-all hover:border-gray-600 hover:bg-gray-700">
                  充值 500 元
                </button>
                <button class="rounded-full border border-gray-700 bg-gray-800 px-8 py-3 font-medium text-gray-300 transition-all hover:border-gray-600 hover:bg-gray-700">
                  充值 1000 元
                </button>
              </div>
            </div>
          </div>
        </section>

        <!-- Quick Start Section -->
        <section id="quickstart" class="mb-20 scroll-mt-24">
          <div class="mb-8 text-center">
            <div class="mb-2 text-xs font-semibold uppercase tracking-wider text-primary-400">Quickstart</div>
            <h2 class="mb-2 text-2xl font-bold text-white md:text-3xl">三步开始使用</h2>
            <p class="text-gray-500">购买套餐 → 安装客户端 → 运行命令，保持官方体验。</p>
          </div>
          <div class="grid gap-6 md:grid-cols-3">
            <div class="rounded-xl border border-gray-800 bg-gray-900/50 p-6 text-center transition-all hover:border-primary-500/50">
              <div class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-primary-500/20 text-lg font-bold text-primary-400">01</div>
              <h3 class="mb-2 font-semibold text-white">获取 Key</h3>
              <p class="text-sm text-gray-500">注册并购买套餐，获得专属 API Key 与节点配置。</p>
            </div>
            <div class="rounded-xl border border-gray-800 bg-gray-900/50 p-6 text-center transition-all hover:border-primary-500/50">
              <div class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-primary-500/20 text-lg font-bold text-primary-400">02</div>
              <h3 class="mb-2 font-semibold text-white">配置客户端</h3>
              <p class="text-sm text-gray-500">使用官方 SDK 或 CLI，修改 baseURL 和 apiKey 即可。</p>
            </div>
            <div class="rounded-xl border border-gray-800 bg-gray-900/50 p-6 text-center transition-all hover:border-primary-500/50">
              <div class="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-primary-500/20 text-lg font-bold text-primary-400">03</div>
              <h3 class="mb-2 font-semibold text-white">开始使用</h3>
              <p class="text-sm text-gray-500">运行 API 调用，支持官方所有参数，快速进入开发流。</p>
            </div>
          </div>
          <div class="mt-8 text-center">
            <router-link
              to="/register"
              class="inline-flex items-center gap-2 rounded-full bg-gradient-to-r from-primary-500 to-brand-500 px-8 py-3 font-medium text-white shadow-lg shadow-primary-500/25 transition-all hover:shadow-primary-500/40"
            >
              <Icon name="rocket" size="sm" />
              立即注册并开始使用
            </router-link>
          </div>
        </section>

        <!-- FAQ Section -->
        <section id="faq" class="mb-20 scroll-mt-24">
          <div class="mb-8 text-center">
            <div class="mb-2 text-xs font-semibold uppercase tracking-wider text-primary-400">FAQ</div>
            <h2 class="text-2xl font-bold text-white md:text-3xl">常见问题</h2>
          </div>
          <div class="mx-auto max-w-3xl space-y-3">
            <details class="group rounded-xl border border-gray-800 bg-gray-900/50 open:border-primary-500/50">
              <summary class="flex cursor-pointer items-center justify-between p-4 font-medium text-white">
                <span>如何开始使用 Sub2API？</span>
                <Icon name="chevronDown" size="sm" class="transition-transform group-open:rotate-180" />
              </summary>
              <div class="border-t border-gray-800 p-4 text-sm text-gray-400">
                注册账号 → 购买套餐或充值 → 在控制台获取 API Key → 按照文档修改您的客户端配置即可开始使用。
              </div>
            </details>
            <details class="group rounded-xl border border-gray-800 bg-gray-900/50 open:border-primary-500/50">
              <summary class="flex cursor-pointer items-center justify-between p-4 font-medium text-white">
                <span>支持哪些模型？</span>
                <Icon name="chevronDown" size="sm" class="transition-transform group-open:rotate-180" />
              </summary>
              <div class="border-t border-gray-800 p-4 text-sm text-gray-400">
                我们支持 OpenAI（GPT-4/GPT-5）、Anthropic（Claude 系列）、Google（Gemini 系列）等主流模型的全部版本，新模型上线通常在 24 小时内支持。
              </div>
            </details>
            <details class="group rounded-xl border border-gray-800 bg-gray-900/50 open:border-primary-500/50">
              <summary class="flex cursor-pointer items-center justify-between p-4 font-medium text-white">
                <span>计费方式是怎样的？</span>
                <Icon name="chevronDown" size="sm" class="transition-transform group-open:rotate-180" />
              </summary>
              <div class="border-t border-gray-800 p-4 text-sm text-gray-400">
                我们提供两种计费方式：包月套餐（每日重置额度，适合高频用户）和 PayGO 按量付费（充值后按实际用量扣费，永不过期）。
              </div>
            </details>
            <details class="group rounded-xl border border-gray-800 bg-gray-900/50 open:border-primary-500/50">
              <summary class="flex cursor-pointer items-center justify-between p-4 font-medium text-white">
                <span>是否支持开发票？</span>
                <Icon name="chevronDown" size="sm" class="transition-transform group-open:rotate-180" />
              </summary>
              <div class="border-t border-gray-800 p-4 text-sm text-gray-400">
                支持。累计充值满 200 元后可申请开具增值税普通发票（电子），请联系客服处理。
              </div>
            </details>
          </div>
        </section>

        <!-- CTA Section -->
        <section class="rounded-2xl bg-gradient-to-r from-primary-500/20 to-brand-500/20 p-8 text-center md:p-12">
          <h2 class="mb-4 text-2xl font-bold text-white md:text-3xl">立即开始使用 Sub2API</h2>
          <p class="mb-6 text-gray-400">简单三步：购买套餐 → 获取 Key → 开始开发</p>
          <router-link
            to="/register"
            class="inline-flex items-center gap-2 rounded-full bg-primary-500 px-8 py-3 font-medium text-white transition-all hover:bg-primary-600"
          >
            免费注册
            <Icon name="arrowRight" size="sm" />
          </router-link>
        </section>
      </div>
    </main>

    <!-- Footer -->
    <footer class="relative z-10 border-t border-gray-800 bg-gray-900/50 px-6 py-12">
      <div class="mx-auto max-w-6xl">
        <div class="grid gap-8 md:grid-cols-4">
          <div>
            <div class="mb-4 flex items-center gap-2">
              <span class="text-xl font-bold text-white">Sub2API</span>
            </div>
            <p class="text-sm text-gray-500">AI API 国内解决方案，让开发者无障碍使用全球顶尖 AI 模型。</p>
          </div>
          <div>
            <h4 class="mb-4 font-semibold text-white">产品</h4>
            <ul class="space-y-2 text-sm text-gray-500">
              <li><router-link to="/" class="hover:text-primary-400">首页</router-link></li>
              <li><router-link to="/models" class="hover:text-primary-400">模型广场</router-link></li>
              <li><router-link to="/#pricing" class="hover:text-primary-400">套餐价格</router-link></li>
              <li><router-link to="/#features" class="hover:text-primary-400">功能特性</router-link></li>
            </ul>
          </div>
          <div>
            <h4 class="mb-4 font-semibold text-white">支持</h4>
            <ul class="space-y-2 text-sm text-gray-500">
              <li><a href="https://docs.sub2api.com" target="_blank" class="hover:text-primary-400">官方文档</a></li>
              <li><router-link to="/#quickstart" class="hover:text-primary-400">快速开始</router-link></li>
              <li><router-link to="/#faq" class="hover:text-primary-400">常见问题</router-link></li>
              <li><a href="https://status.sub2api.com" target="_blank" class="hover:text-primary-400">服务状态</a></li>
            </ul>
          </div>
          <div>
            <h4 class="mb-4 font-semibold text-white">关于</h4>
            <ul class="space-y-2 text-sm text-gray-500">
              <li><router-link to="/terms" class="hover:text-primary-400">服务条款</router-link></li>
              <li><router-link to="/privacy" class="hover:text-primary-400">隐私政策</router-link></li>
            </ul>
          </div>
        </div>
        <div class="mt-8 border-t border-gray-800 pt-8 text-center text-sm text-gray-500">
          © 2025 Sub2API. All rights reserved.
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'
import { useAuthStore } from '@/stores'

const { t } = useI18n()
const authStore = useAuthStore()

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))

// Service days since launch
const serviceDays = computed(() => {
  const launchDate = new Date('2024-06-01')
  const now = new Date()
  return Math.floor((now.getTime() - launchDate.getTime()) / (1000 * 60 * 60 * 24))
})

// Pricing tab state
const pricingTab = ref<'monthly' | 'paygo'>('monthly')
</script>
