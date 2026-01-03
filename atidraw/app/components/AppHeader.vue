<script setup lang="ts">
const { loggedIn, user, clear } = useUserSession()
const route = useRoute()

interface Category {
  name: string
  id: string
}

const categories: Category[] = [
  { name: '推荐', id: 'recommend' },
  { name: '穿搭', id: 'fashion' },
  { name: '美食', id: 'food' },
  { name: '彩妆', id: 'makeup' },
  { name: '影视', id: 'movie' },
  { name: '职场', id: 'career' },
  { name: '情感', id: 'emotion' },
  { name: '家居', id: 'home' },
  { name: '游戏', id: 'game' },
  { name: '旅行', id: 'travel' },
  { name: '健身', id: 'fitness' },
]

// 只在首页显示分类标签
const showCategories = computed(() => route.path === '/')

const currentCategory = ref<string>(
  typeof route.query.channel_id === 'string' 
    ? route.query.channel_id 
    : 'recommend'
)

function selectCategory(categoryId: string) {
  currentCategory.value = categoryId
  navigateTo({
    query: { channel_id: categoryId },
  })
}
</script>

<template>
  <header class="sticky top-0 z-40 bg-white border-b border-gray-200">
    <!-- 顶部导航栏 -->
    <div class="flex items-center justify-between h-14 px-4">
      <div class="flex items-center gap-2">
        <h1 class="text-xl font-bold text-red-500">小红书</h1>
      </div>
      <div class="flex items-center gap-3">
        <UButton
          icon="i-ph-magnifying-glass"
          color="gray"
          variant="ghost"
          size="sm"
        />
        <UButton
          icon="i-ph-chat-circle"
          color="gray"
          variant="ghost"
          size="sm"
        />
      </div>
    </div>
    
    <!-- 分类标签栏 - 只在首页显示 -->
    <div
      v-if="showCategories"
      class="flex items-center gap-2 px-4 py-2 overflow-x-auto scrollbar-hide"
    >
      <button
        v-for="category in categories"
        :key="category.id"
        class="px-4 py-1.5 rounded-full text-sm whitespace-nowrap transition-colors"
        :class="currentCategory === category.id 
          ? 'bg-red-500 text-white font-medium' 
          : 'bg-gray-100 text-gray-700 hover:bg-gray-200'"
        @click="selectCategory(category.id)"
      >
        {{ category.name }}
      </button>
    </div>
  </header>
</template>

<style scoped>
.scrollbar-hide {
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.scrollbar-hide::-webkit-scrollbar {
  display: none;
}
</style>

