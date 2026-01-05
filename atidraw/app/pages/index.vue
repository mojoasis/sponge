<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { useRuntimeConfig } from '#app'
import type { ApiResponse, FeedResponseData, VideoFeedItem } from '#shared/types/api'

const config = useRuntimeConfig()
const ossBaseUrl = (config.public.ossUrl as string) || ''

function getOssUrl(path: string): string {
  if (!path) return ''
  if (/^https?:\/\//i.test(path)) return path
  return `${ossBaseUrl.replace(/\/$/, '')}/${path.replace(/^\//, '')}`
}

const feedList = ref<VideoFeedItem[]>([])
const activeId = ref<string | number | null>(null)

let observer: IntersectionObserver | null = null

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
      { threshold: 0.6 }
    )

    document
      .querySelectorAll('.feed-item')
      .forEach(el => observer?.observe(el))
  })
}

onMounted(loadFeed)
onBeforeUnmount(() => observer?.disconnect())
</script>

<template>
  <!-- 移动端统一视口容器 -->
  <div class="video-page snap-y snap-mandatory overflow-y-scroll bg-black">
    <div
      v-for="(item, index) in feedList"
      :key="item.id"
      class="feed-item snap-start"
      :data-id="item.id"
    >
      <FeedItem
        :item="{
          ...item,
          playUrl: getOssUrl(<string>item.playUrl),
          coverUrl: getOssUrl(<string>item.coverUrl)
        }"
        :active="String(activeId) === String(item.id)"
        :next-play-url="
          feedList[index + 1]?.playUrl
            ? getOssUrl(<string>feedList[index + 1]?.playUrl)
            : ''
        "
      />
    </div>
  </div>
</template>

<style scoped>
.video-page {
  height: 100vh;
  width: 100vw;
  background: black;
  //overflow: hidden;
  overflow-y: auto;
  overflow-x: hidden;
  padding-top: env(safe-area-inset-top);
  padding-bottom: env(safe-area-inset-bottom);
}

.feed-item {
  height: 100vh;
  width: 100vw;
}

::-webkit-scrollbar {
  display: none;
}
</style>
