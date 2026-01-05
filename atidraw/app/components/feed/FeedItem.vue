<script setup lang="ts">
import type { VideoFeedItem } from '#shared/types/api'
import { computed } from 'vue'

const props = defineProps<{
  item: VideoFeedItem
  active: boolean
  nextPlayUrl?: string
}>()

const playUrl = computed(() => props.item.playUrl || '')
const posterUrl = computed(() => props.item.coverUrl || '')
</script>

<template>
  <div class="w-full h-[100svh] relative">
    <FeedVideo
      :play-url="playUrl"
      :poster="posterUrl"
      :active="active"
      :preload-url="nextPlayUrl"
    />

    <!-- 标题 -->
    <div
      class="absolute left-4 bottom-6 text-white z-10
             text-sm leading-snug max-w-[80%]"
    >
      {{ item.title }}
    </div>
  </div>
</template>
