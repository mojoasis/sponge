<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { useRuntimeConfig } from '#app'
import type { ApiResponse, FeedResponseData, VideoFeedItem } from '#shared/types/api'

/* ---------------- 基础配置 ---------------- */

const config = useRuntimeConfig()
const ossBaseUrl = (config.public.ossUrl as string) || ''

function getOssUrl(path?: string | null) {
  if (!path) return ''
  if (/^https?:\/\//i.test(path)) return path
  return `${ossBaseUrl.replace(/\/$/, '')}/${path.replace(/^\//, '')}`
}

/* ---------------- Feed 状态 ---------------- */

const feedList = ref<VideoFeedItem[]>([])
const activeId = ref<string | number | null>(null)

let observer: IntersectionObserver | null = null

/* ---------------- 数据加载（首次） ---------------- */

async function loadFeed() {
  const res = await $fetch<ApiResponse<FeedResponseData>>('/api/feed', {
    method: 'GET',
    query: {
      page: 1,
      size: 20,
      channel_id: 'recommend'
    }
  })

  if (res.code === 200 && res.data?.list) {
    feedList.value = res.data.list
    observeItems()
  }
}

/* ---------------- IntersectionObserver ---------------- */
/**
 * 规则：谁占据屏幕中部，谁 active
 */
function observeItems() {
  nextTick(() => {
    observer?.disconnect()

    observer = new IntersectionObserver(
      entries => {
        entries.forEach(entry => {
          if (!entry.isIntersecting) return
          const id = entry.target.getAttribute('data-id')
          if (id) activeId.value = id
        })
      },
      {
        threshold: 0.6
      }
    )

    document
      .querySelectorAll('.feed-item')
      .forEach(el => observer?.observe(el))
  })
}

/* ---------------- 生命周期 ---------------- */

onMounted(() => {
  loadFeed()
})

onBeforeUnmount(() => {
  observer?.disconnect()
})
</script>

<template>
  <!-- 整体容器：移动端纵向滑动 -->
  <div
    class="w-full h-[100svh] overflow-y-scroll snap-y snap-mandatory bg-black"
  >
    <!-- 单条 Feed：一屏 -->
    <div
      v-for="item in feedList"
      :key="item.id"
      class="feed-item w-full h-[100svh] snap-start relative"
      :data-id="item.id"
    >
      <!-- 视频 -->
      <FeedItem
        :item="{
          ...item,
          playUrl: getOssUrl(item.playUrl),
          coverUrl: getOssUrl(item.coverUrl)
        }"
        :active="String(activeId) === String(item.id)"
      />
    </div>
  </div>
</template>

<style scoped>
/* 禁止横向滚动 */
::-webkit-scrollbar {
  display: none;
}
</style>
