<script setup lang="ts">
useSeoMeta({
  title: '小红书 - 标记我的生活',
  description: '小红书 - 标记我的生活，发现全世界的好东西',
  ogTitle: '小红书',
  twitterCard: 'summary_large_image',
})
</script>

<template>
  <UApp>
    <AppHeader />
    <main class="min-h-[calc(100vh-120px)] pb-20 bg-gray-50">
      <NuxtPage />
    </main>
    <BottomNav />
  </UApp>
</template>
