<script setup lang="ts">
const route = useRoute()

interface NavItem {
  path: string
  icon: string
  label: string
  name: string
}

const navItems: NavItem[] = [
  { path: '/', icon: 'i-ph-house', label: '发现', name: 'index' },
  { path: '/publish', icon: 'i-ph-plus-circle', label: '发布', name: 'publish' },
  { path: '/notifications', icon: 'i-ph-bell', label: '通知', name: 'notifications' },
  { path: '/profile', icon: 'i-ph-user', label: '我', name: 'profile' },
]

const isActive = (path: string): boolean => {
  if (path === '/') {
    return route.path === '/'
  }
  return route.path.startsWith(path)
}
</script>

<template>
  <nav class="fixed bottom-0 left-0 right-0 bg-white border-t border-gray-200 z-50 safe-area-inset-bottom shadow-lg">
    <div class="flex items-center justify-around h-16 px-2 max-w-2xl mx-auto">
      <NuxtLink
        v-for="item in navItems"
        :key="item.path"
        :to="item.path"
        class="flex flex-col items-center justify-center flex-1 h-full transition-colors relative"
        :class="isActive(item.path) ? 'text-red-500' : 'text-gray-500'"
      >
        <UIcon
          :name="item.icon"
          class="text-2xl mb-1"
        />
        <span class="text-xs font-medium">{{ item.label }}</span>
        <div
          v-if="isActive(item.path)"
          class="absolute top-0 left-1/2 transform -translate-x-1/2 w-8 h-0.5 bg-red-500 rounded-full"
        />
      </NuxtLink>
    </div>
  </nav>
</template>

<style scoped>
.safe-area-inset-bottom {
  padding-bottom: env(safe-area-inset-bottom);
}
</style>

