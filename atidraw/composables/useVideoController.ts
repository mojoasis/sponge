import { ref } from 'vue'

const currentVideo = ref<HTMLVideoElement | null>(null)

export function useVideoController() {
  function requestPlay(video: HTMLVideoElement, muted = true) {
    if (currentVideo.value && currentVideo.value !== video) {
      currentVideo.value.pause()
      currentVideo.value.currentTime = 0
    }
    currentVideo.value = video
    video.muted = muted
    video.play().catch(() => {})
  }

  function pause(video: HTMLVideoElement) {
    if (currentVideo.value === video) {
      video.pause()
      currentVideo.value = null
    }
  }

  return { requestPlay, pause }
}
