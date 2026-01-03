// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: [
    '@nuxthub/core',
    '@nuxt/ui',
    '@nuxt/eslint',
    'nuxt-auth-utils',
  ],
  devtools: { enabled: true },
  css: ['~/assets/main.css'],
  routeRules: {
    '/drawings/**': { isr: true },
  },
  future: { compatibilityVersion: 4 },
  compatibilityDate: '2025-10-14',
  hub: {
    blob: true,
  },
  // 配置运行时变量
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:18888',
      ossUrl: process.env.NUXT_PUBLIC_OSS_URL || 'http://124.70.84.161:6900/',
    },
  },
  // Development modules
  eslint: {
    config: {
      stylistic: {
        quotes: 'single',
      },
    },
  },
  // 禁用在线字体提供商以解决中国大陆无法访问谷歌字体的问题
  ui: {
    fonts: false,
  },
})
