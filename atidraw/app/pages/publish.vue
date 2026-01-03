<script setup lang="ts">
// 在 script setup 中添加默认值
const { loggedIn, user } = useUserSession()
// 确保 user 对象有默认值
const currentUser = computed(() => {
  if (!user.value) {
    return {
      name: '未登录用户',
      email: '点击登录',
      avatar: ''
    }
  }
  return user.value
})
const toast = useToast()
const saving = ref(false)
const drawing = ref('')

function onDraw(dataURL: string) {
  drawing.value = dataURL
}

async function save(dataURL: string) {
  if (saving.value) return
  saving.value = true
  const blob = await fetch(dataURL).then(res => res.blob())
  const form = new FormData()
  form.append('drawing', new File([blob], `drawing.jpg`, { type: 'image/jpeg' }))

  await $fetch('/api/upload', {
    method: 'POST',
    body: form,
  })
    .then(() => {
      toast.add({
        title: '发布成功！',
        description: '你的作品已经分享给大家了',
        color: 'success',
      })
      navigateTo('/')
    }).catch((err) => {
      toast.add({
        title: '发布失败',
        description: err.data?.message || err.message,
        color: 'error',
      })
    })
  saving.value = false
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-2xl mx-auto px-4 py-6">
      <h2 class="text-2xl font-bold mb-6 text-gray-800">发布笔记</h2>

      <div v-if="loggedIn">
        <div class="bg-white rounded-lg shadow-sm p-6">
          <DrawPad
            :saving="saving"
            class="max-w-full"
            @save="save"
            @draw="onDraw"
          />
        </div>
      </div>

      <div
        v-else
        class="bg-white rounded-lg shadow-sm p-8 text-center"
      >
        <UIcon
          name="i-ph-user-circle"
          class="text-6xl text-gray-300 mb-4 mx-auto"
        />
        <h3 class="text-xl font-semibold mb-2 text-gray-800">登录后发布</h3>
        <p class="text-gray-500 mb-6">登录后可以发布你的作品，与大家分享</p>

        <div class="space-y-3 max-w-sm mx-auto">
          <UButton
            to="/login"
            label="登录后发布"
            icon="i-ph-user-circle"
            variant="solid"
            size="lg"
            block
          />
        </div>
      </div>
    </div>
  </div>
</template>

