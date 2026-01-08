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
  // 禁用在线字体提供商以解决中国大陆无法访问谷歌字体的问题
  ui: {
    fonts: false,
  },
  // 配置运行时变量
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://81.70.142.31:18888/',
      ossUrl: process.env.NUXT_PUBLIC_OSS_URL || 'http://124.70.84.161:6900/',
    },
  },
  routeRules: {
    '/drawings/**': { isr: true },
  },
  future: { compatibilityVersion: 4 },
  compatibilityDate: '2025-10-14',
  // 添加静态站点配置
  nitro: {
    preset: 'static',
  },
  hub: {
    blob: true,
  },
  // Development modules
  eslint: {
    config: {
      stylistic: {
        quotes: 'single',
      },
    },
  },
})
