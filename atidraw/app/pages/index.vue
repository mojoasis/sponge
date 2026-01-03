<script setup lang="ts">
  import { ref, computed, onMounted, watch } from 'vue'
  import { useRoute, useRuntimeConfig } from '#app'
  import type { ApiResponse, FeedResponseData, VideoFeedItem } from '../../shared/types/api'
  import { UseTimeAgo, vInfiniteScroll } from '@vueuse/components'
  
  const route = useRoute()
  const config = useRuntimeConfig()
  
  const ossBaseUrl = (config.public.ossUrl as string) || 'http://124.70.84.161:6900/'
  
  function getOssUrl(path?: string | null): string {
    if (!path) return ''
    if (/^https?:\/\//i.test(path)) return path
    const base = ossBaseUrl.replace(/\/+$/, '')
    const cleanPath = path.replace(/^\/+/, '')
    return `${base}/${cleanPath}`
  }
  
  // --- 状态管理 ---
  const feedList = ref<VideoFeedItem[]>([])
  const loading = ref(false)
  const hasMore = ref(true)
  const latestTime = ref<number>(Date.now())
  const page = ref(1)
  const pageSize = 10
  // 确保类型兼容，如果是数字 ID 请改为 number | null
  const activeVideoId = ref<string | number | null>(null)
  
  // --- 数据加载逻辑 ---
  const currentCategory = computed(() => {
    return typeof route.query.channel_id === 'string' ? route.query.channel_id : 'recommend'
  })
  
  async function loadFeed(isRefresh = false) {
    if (loading.value) return
    if (!isRefresh && !hasMore.value) return
    
    loading.value = true
  
    try {
      const query = {
        latestTime: isRefresh ? Date.now() : latestTime.value,
        page: isRefresh ? 1 : page.value,
        size: pageSize,
        channel_id: currentCategory.value,
      }
  
      const result = await $fetch<ApiResponse<FeedResponseData>>('/api/feed', {
        method: 'GET',
        query,
      })
  
      if (result.code === 200 && result.data) {
        const newList = result.data.list || []
        if (isRefresh) {
          feedList.value = newList
          page.value = 2
        } else {
          feedList.value.push(...newList)
          page.value++
        }
        hasMore.value = result.data.hasMore ?? false
        if (result.data.nextLatestTime) {
          latestTime.value = result.data.nextLatestTime
        }
      }
    } catch (error) {
      console.error('加载feed失败:', error)
    } finally {
      setTimeout(() => {
        loading.value = false
      }, 200)
    }
  }
  
  function handleLoadMore() {
    if (!loading.value && hasMore.value) {
      loadFeed(false)
    }
  }
  
  onMounted(() => {
    loadFeed(true)
  })
  
  watch(currentCategory, () => {
    activeVideoId.value = null
    hasMore.value = true
    loadFeed(true)
  })
  
  function formatNumber(num?: number): string {
    if (!num) return '0'
    return num >= 10000 ? `${(num / 10000).toFixed(1)}万` : num.toString()
  }
  
  // 处理播放逻辑
  function handlePlay(id: string | number) {
    activeVideoId.value = id
  }
  </script>
  
  <template>
    <div class="bg-gray-50 min-h-screen">
      <div class="max-w-2xl mx-auto bg-white shadow-sm">
        <div class="divide-y divide-gray-100">
          <div v-for="item in feedList" :key="item.id" class="py-4">
            <div class="flex items-center gap-3 px-4 mb-3">
              <UAvatar :src="getOssUrl(item.author?.avatar)" size="sm" />
              <div class="flex-1">
                <div class="text-sm font-bold">{{ item.author?.nickName || '用户' }}</div>
                <div class="text-[11px] text-gray-400">
                  <UseTimeAgo v-if="item.publishTime" :time="new Date(item.publishTime)" v-slot="{ timeAgo }">
                    {{ timeAgo }}
                  </UseTimeAgo>
                </div>
              </div>
            </div>
  
            <div class="px-4">
              <p v-if="item.title" class="text-sm mb-3 text-gray-800 leading-relaxed">{{ item.title }}</p>
              
              <div class="relative rounded-xl overflow-hidden bg-black aspect-video mb-4 shadow-inner group">
                <template v-if="item.videoUrl && activeVideoId === item.id">
                  <video
                    :src="getOssUrl(item.videoUrl)"
                    class="w-full h-full"
                    controls
                    autoplay
                    muted
                    playsinline
                  ></video>
                </template>
                
                <template v-else>
                  <img
                    :src="getOssUrl(item.coverUrl)"
                    class="w-full h-full object-cover"
                    loading="lazy"
                  />
                  <div 
                    v-if="item.videoUrl" 
                    class="absolute inset-0 flex items-center justify-center bg-black/20 cursor-pointer"
                    @click.stop="handlePlay(item.id)"
                  >
                    <div class="w-16 h-16 flex items-center justify-center rounded-full bg-white/20 backdrop-blur-md hover:scale-110 transition-all duration-300">
                      <UIcon name="i-ph-play-fill" class="text-4xl text-white" />
                    </div>
                  </div>
                </template>
              </div>
  
              <div class="flex gap-6 text-gray-500 pb-2">
                <button class="flex items-center gap-1.5 hover:text-red-500 transition-colors">
                  <UIcon name="i-ph-heart" class="text-2xl" />
                  <span class="text-xs">{{ formatNumber(item.stats?.likes) }}</span>
                </button>
                <button class="flex items-center gap-1.5 hover:text-blue-500 transition-colors">
                  <UIcon name="i-ph-chat-circle" class="text-2xl" />
                  <span class="text-xs">{{ formatNumber(item.stats?.comments) }}</span>
                </button>
              </div>
            </div>
          </div>
  
          <div
            v-infinite-scroll="[handleLoadMore, { distance: 100, interval: 500 }]"
            class="py-10 text-center"
          >
            <div v-if="loading" class="flex items-center justify-center gap-2 text-gray-400">
              <UIcon name="i-ph-circle-notched" class="animate-spin text-xl" />
              <span class="text-sm font-medium">正在努力加载...</span>
            </div>
            <div v-else-if="!hasMore" class="text-sm text-gray-300 font-light italic">
              — 到底啦 —
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <style scoped>
  /* 优化视频过渡显示，防止切换时闪烁 */
  video {
    background-color: #000;
  }
  </style>