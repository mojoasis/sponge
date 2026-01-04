<script setup lang="ts">
import { ref, watch, onMounted, onBeforeUnmount, nextTick } from 'vue'

interface Props {
  playUrl: string
  poster?: string
  active: boolean
}

const props = defineProps<Props>()

const videoRef = ref<HTMLVideoElement | null>(null)
const showOverlay = ref(false)
const playing = ref(false)

let hideTimer: number | null = null
let observer: IntersectionObserver | null = null

function play(muted = true) {
  const v = videoRef.value
  if (!v) return

  v.muted = muted
  const p = v.play()
  playing.value = true
  p?.catch(() => {})
}

function pause() {
  videoRef.value?.pause()
  playing.value = false
}

function togglePlay() {
  if (!videoRef.value) return
  videoRef.value.paused ? play(false) : pause()
  flashOverlay()
}

function flashOverlay() {
  showOverlay.value = true
  if (hideTimer) window.clearTimeout(hideTimer)
  hideTimer = window.setTimeout(() => {
    showOverlay.value = false
  }, 800)
}

onMounted(() => {
  nextTick(() => {
    const v = videoRef.value
    if (!v) return

    observer = new IntersectionObserver(
      entries => {
        entries.forEach(e => {
          if (e.isIntersecting) {
            play(true)
          } else {
            pause()
          }
        })
      },
      { threshold: 0.6 }
    )

    observer.observe(v)
  })
})

onBeforeUnmount(() => {
  observer?.disconnect()
  if (hideTimer) window.clearTimeout(hideTimer)
})

watch(
  () => props.active,
  val => {
    val ? play(true) : pause()
  }
)
</script>

<template>
  <div class="relative w-full h-full bg-black" @click="togglePlay">
    <video
      ref="videoRef"
      :src="playUrl"
      :poster="poster"
      class="w-full h-full object-cover"
      preload="metadata"
      playsinline
      loop
    />

    <!-- 播放提示（淡入淡出） -->
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
