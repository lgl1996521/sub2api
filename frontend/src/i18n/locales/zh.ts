export default {
  // Navigation
  publicHeader: {
    home: '首页',
    models: '模型广场',
    plans: '套餐价格',
    quickstart: '快速开始',
    faq: '常见问题',
    docs: '文档',
    signUp: '注册',
    login: '登录',
    dashboard: '控制台',
    toggleMenu: '切换菜单'
  },

  // Home Page
  home: {
    viewOnGithub: '在 GitHub 上查看',
    viewDocs: '查看文档',
    docs: '文档',
    switchToLight: '切换到浅色模式',
    switchToDark: '切换到深色模式',
    dashboard: '控制台',
    login: '登录',
    getStarted: '立即开始',
    goToDashboard: '进入控制台',

    // Hero
    heroTitle: 'AI API 国内解决方案',
    heroDescription: '即刻体验 OpenAI、Claude、Gemini 等全球顶尖 AI 模型；智能路由技术，降低封号风险，国内直连。',

    // Stats
    stats: {
      serviceDays: '稳定服务时间',
      developers: '服务开发者',
      uptime: '稳定运行率'
    },

    // Features
    features: {
      zeroRisk: '零封号风险',
      zeroRiskDesc: '企业级号池代理技术，我们来解决网络访问与稳定性。',
      oneClick: '一键安装',
      oneClickDesc: '脚本自动配置 Key 与节点，三步即可开始使用。',
      flexibleBilling: '灵活计费',
      flexibleBillingDesc: '包月套餐适合高频使用，按量付费用多少扣多少。',
      transparentCache: '缓存命中率透明',
      transparentCacheDesc: '后台清晰展示缓存命中率，降低成本，拒绝套路。',
      specialFeatures: '特有功能',
      specialFeaturesDesc: '支持多模型负载均衡、自动重试、智能路由。',
      usageStats: '使用统计',
      usageStatsDesc: '实时追踪用量与费用，记录可追溯，计费更透明。'
    },

    // Pricing
    pricing: {
      title: '灵活的套餐选择',
      subtitle: '选择适合您的套餐方案 所有套餐均可叠加使用',
      monthly: '包月套餐',
      paygo: 'PayGO套餐',
      basic: {
        name: '基础月卡',
        desc: '性价比之选',
        price: '¥99',
        period: '/月',
        rate: '≈ 0.33¥/$',
        quota: '月总额度 $300',
        daily: '每日0点重置到 $37.5',
        weekly: '周限额度 $75'
      },
      pro: {
        name: '专业月卡',
        desc: '畅享无忧',
        price: '¥249',
        period: '/月',
        rate: '≈ 0.28¥/$',
        quota: '月总额度 $900',
        daily: '每日0点重置到 $112.5',
        weekly: '周限额度 $225',
        support: '优先客服支持'
      },
      enterprise: {
        name: '企业月卡',
        desc: '超大杯',
        price: '¥699',
        period: '/月',
        rate: '≈ 0.23¥/$',
        quota: '月总额度 $3000',
        daily: '每日0点重置到 $375',
        weekly: '周限额度 $750',
        support: '专属客服支持'
      },
      paygoDesc: '充值 1 元 = 1 美元额度，用多少扣多少，无时间限制',
      paygoRatio: '人民币:美元',
      paygoSave: '节省费用'
    },

    // Quick Start
    quickstart: {
      title: '三步开始使用',
      subtitle: '购买套餐 → 安装客户端 → 运行命令，保持官方体验。',
      step1: {
        title: '获取 Key',
        desc: '注册并购买套餐，获得专属 API Key 与节点配置。'
      },
      step2: {
        title: '配置客户端',
        desc: '使用官方 SDK 或 CLI，修改 baseURL 和 apiKey 即可。'
      },
      step3: {
        title: '开始使用',
        desc: '运行 API 调用，支持官方所有参数，快速进入开发流。'
      },
      cta: '立即注册并开始使用'
    },

    // FAQ
    faq: {
      title: '常见问题',
      q1: {
        question: '如何开始使用 Sub2API？',
        answer: '注册账号 → 购买套餐或充值 → 在控制台获取 API Key → 按照文档修改您的客户端配置即可开始使用。'
      },
      q2: {
        question: '支持哪些模型？',
        answer: '我们支持 OpenAI（GPT-4/GPT-5）、Anthropic（Claude 系列）、Google（Gemini 系列）等主流模型的全部版本，新模型上线通常在 24 小时内支持。'
      },
      q3: {
        question: '计费方式是怎样的？',
        answer: '我们提供两种计费方式：包月套餐（每日重置额度，适合高频用户）和 PayGO 按量付费（充值后按实际用量扣费，永不过期）。'
      },
      q4: {
        question: '是否支持开发票？',
        answer: '支持。累计充值满 200 元后可申请开具增值税普通发票（电子），请联系客服处理。'
      }
    },

    // CTA
    cta: {
      title: '立即开始使用 Sub2API',
      subtitle: '简单三步：购买套餐 → 获取 Key → 开始开发',
      button: '免费注册'
    },

    // Footer
    footer: {
      tagline: 'AI API 国内解决方案，让开发者无障碍使用全球顶尖 AI 模型。',
      product: '产品',
      support: '支持',
      about: '关于',
      rights: '© 2025 Sub2API. All rights reserved.'
    }
  },

  // Models Page
  models: {
    plaza: {
      title: '模型广场',
      subtitle: '探索我们支持的所有 AI 模型，点击卡片查看详情',
      searchPlaceholder: '搜索模型...',
      filterAll: '全部',
      filterChat: '对话',
      filterEmbedding: '嵌入',
      filterImage: '图像',
      filterAudio: '音频',
      filterVideo: '视频',
      filterRealtime: '实时',
      filterSearch: '搜索',
      filterTTS: '语音合成',
      filterSTT: '语音识别'
    }
  },

  // Common
  common: {
    loading: '加载中...',
    error: '出错了',
    retry: '重试',
    cancel: '取消',
    confirm: '确认',
    save: '保存',
    delete: '删除',
    edit: '编辑',
    create: '创建',
    search: '搜索',
    filter: '筛选',
    sort: '排序',
    more: '更多',
    less: '收起',
    expand: '展开',
    collapse: '折叠',
    close: '关闭',
    open: '打开',
    copy: '复制',
    copied: '已复制',
    download: '下载',
    upload: '上传',
    preview: '预览',
    submit: '提交',
    reset: '重置',
    back: '返回',
    next: '下一步',
    previous: '上一步',
    finish: '完成',
    success: '成功',
    failed: '失败',
    warning: '警告',
    info: '提示'
  }
}
