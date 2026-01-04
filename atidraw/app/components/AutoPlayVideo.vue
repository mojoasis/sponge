<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'

interface Props {
  src: string
  poster?: string
  active: boolean
}

const props = defineProps<Props>()

const videoRef = ref<HTMLVideoElement | null>(null)
let observer: IntersectionObserver | null = null

function safePlay(muted = true) {
  const video = videoRef.value
  if (!video) return

  video.muted = muted
  const p = video.play()
  if (p && typeof p.catch === 'function') {
    p.catch(() => {})
  }
}

function pause() {
  videoRef.value?.pause()
}

onMounted(() => {
  nextTick(() => {
    const video = videoRef.value
    if (!video) return

    observer = new IntersectionObserver(
      entries => {
        entries.forEach(entry => {
          if (entry.isIntersecting) {
            safePlay(true)
          } else {
            pause()
          }
        })
      },
      { threshold: 0.6 }
    )

    observer.observe(video)
  })
})

onBeforeUnmount(() => {
  observer?.disconnect()
})

watch(
  () => props.active,
  val => {
    if (!videoRef.value) return
    if (val) {
      safePlay(true)
    } else {
      pause()
    }
  }
)

function handleUserPlay() {
  safePlay(false)
}
</script>

<template>
  <div class="relative w-full h-full bg-black">
    <video
      ref="videoRef"
      :src="src"
      :poster="poster"
      class="w-full h-full object-contain bg-black"
      preload="metadata"
      playsinline
      loop
    />

    <!-- 覆盖层：点击解锁声音 -->
    <div
      class="absolute inset-0 flex items-center justify-center cursor-pointer"
      @click.stop="handleUserPlay"
    >
      <div
        class="w-14 h-14 flex items-center justify-center rounded-full
               bg-black/40 backdrop-blur-md border border-white/20"
      >
        ▶
      </div>
    </div>
  </div>
</template>
