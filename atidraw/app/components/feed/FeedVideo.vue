<script setup lang="ts">
import { ref, watch, onMounted, onBeforeUnmount, nextTick } from 'vue'
import {useVideoController} from "~~/composables/useVideoController";

interface Props {
  playUrl: string
  poster?: string
  active: boolean
  preloadUrl?: string
}

const props = defineProps<Props>()
const { requestPlay, pause } = useVideoController()

const videoRef = ref<HTMLVideoElement | null>(null)
const fitMode = ref<'cover' | 'contain'>('cover')
const playing = ref(false)
const showOverlay = ref(false)

let observer: IntersectionObserver | null = null
let hideTimer: number | null = null

function handleMetadataLoaded() {
  const v = videoRef.value
  if (!v) return

  const ratio = v.videoWidth / v.videoHeight
  fitMode.value = ratio > 1 ? 'contain' : 'cover'
}

function play(muted = true) {
  if (!videoRef.value) return
  requestPlay(videoRef.value, muted)
  playing.value = true
}

function stop() {
  if (!videoRef.value) return
  pause(videoRef.value)
  playing.value = false
}

function togglePlay() {
  if (!videoRef.value) return
  videoRef.value.paused ? play(false) : stop()
  flashOverlay()
}

function flashOverlay() {
  showOverlay.value = true
  if (hideTimer) clearTimeout(hideTimer)
  hideTimer = window.setTimeout(() => {
    showOverlay.value = false
  }, 700)
}

onMounted(() => {
  nextTick(() => {
    const v = videoRef.value
    if (!v) return

    observer = new IntersectionObserver(
      entries => {
        entries.forEach(e => {
          e.isIntersecting ? play(true) : stop()
        })
      },
      { threshold: 0.6 }
    )

    observer.observe(v)
  })
})

onBeforeUnmount(() => {
  observer?.disconnect()
  if (hideTimer) clearTimeout(hideTimer)
})

watch(
  () => props.active,
  val => {
    val ? play(true) : stop()
  }
)
</script>

<template>
  <div class="relative w-full h-full bg-black" @click="togglePlay">
    <video
      ref="videoRef"
      :src="playUrl"
      :poster="poster"
      :class="[
        'w-full h-full',
        fitMode === 'cover' ? 'object-cover' : 'object-contain'
      ]"
      preload="metadata"
      playsinline
      loop
      @loadedmetadata="handleMetadataLoaded"
    />

    <!-- 下一条预加载 -->
    <video
      v-if="preloadUrl"
      :src="preloadUrl"
      preload="metadata"
      class="hidden"
    />

    <!-- 播放状态提示 -->
    <transition name="fade">
      <div
        v-if="showOverlay"
        class="absolute inset-0 flex items-center justify-center"
      >
        <div
          class="w-16 h-16 rounded-full bg-black/40
                 flex items-center justify-center text-white text-2xl"
        >
          {{ playing ? '❚❚' : '▶' }}
        </div>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
