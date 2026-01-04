<script setup lang="ts">
import type { VideoFeedItem } from '~/shared/types/api'
defineProps<{ item: VideoFeedItem; active: boolean }>()
</script>

<template>
  <div class="w-full h-[100svh] relative">
    <FeedVideo
      :play-url="item.playUrl"
      :poster="item.coverUrl"
      :active="active"
    />

    <!-- 右侧操作区（预留） -->
    <div class="absolute right-3 bottom-20 text-white space-y-4">
      ❤️
      💬
      🔁
    </div>

    <!-- 标题 -->
    <div class="absolute left-3 bottom-6 text-white text-sm">
      {{ item.title }}
    </div>
  </div>
</template>
