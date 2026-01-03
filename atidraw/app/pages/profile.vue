<script setup lang="ts">
// 在 script setup 中添加默认值
const { loggedIn, user } = useUserSession()

// 确保 user 对象有默认值，使用后端返回的完整字段
const currentUser = computed(() => {
  if (!user.value) {
    return {
      name: '未登录用户',
      email: '点击登录',
      avatar: '',
      userName: '',
      nickName: '',
      signature: '',
      phone: '',
      backgroundImage: '',
    }
  }
  return user.value as typeof user.value & { backgroundImage?: string }
})

const router = useRouter()

const stats = ref({
  notes: 128,
  likes: 2560,
  followers: 520,
  following: 89,
})

const tabs = ref([
  { id: 'notes', label: '笔记', icon: 'i-ph-grid-four' },
  { id: 'likes', label: '赞过', icon: 'i-ph-heart' },
  { id: 'collections', label: '收藏', icon: 'i-ph-bookmark' },
])

const activeTab = ref('notes')

// 处理登录按钮点击事件
const handleLoginClick = () => {
  navigateTo('/login')
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-2xl mx-auto">
      <!-- 用户信息卡片 -->
      <div class="bg-white">
        <!-- 背景图 -->
        <div
          v-if="loggedIn && currentUser.backgroundImage"
          class="h-32 bg-cover bg-center"
          :style="{ backgroundImage: `url(${currentUser.backgroundImage})` }"
        />
        <div
          v-else
          class="h-32 bg-gradient-to-br from-red-100 to-red-200"
        />
        
        <div class="px-4 pt-6 pb-4">
          <div class="flex items-start gap-4 mb-4 -mt-16">
            <UAvatar
              :src="currentUser.avatar"
              size="xl"
              icon="i-ph-user"
              class="w-20 h-20 border-4 border-white shadow-lg"
            />
            <div class="flex-1 pt-16">
              <!-- 昵称或用户名 -->
              <h2 class="text-xl font-bold text-gray-800 mb-1">
                {{ currentUser.nickName || currentUser.userName || currentUser.name || '未登录用户' }}
              </h2>
              <!-- 用户名（如果与昵称不同） -->
              <p
                v-if="currentUser.userName && currentUser.userName !== currentUser.nickName"
                class="text-sm text-gray-500 mb-1"
              >
                @{{ currentUser.userName }}
              </p>
              <!-- 个性签名 -->
              <p
                v-if="currentUser.signature"
                class="text-sm text-gray-600 mb-2"
              >
                {{ currentUser.signature }}
              </p>
              <!-- 联系信息 -->
              <div class="flex flex-col gap-1 text-xs text-gray-500">
                <p v-if="currentUser.email">
                  <UIcon name="i-ph-envelope" class="inline mr-1" />
                  {{ currentUser.email }}
                </p>
                <p v-if="currentUser.phone">
                  <UIcon name="i-ph-phone" class="inline mr-1" />
                  {{ currentUser.phone }}
                </p>
              </div>
            </div>
            <UButton
              v-if="loggedIn"
              icon="i-ph-gear"
              variant="ghost"
              size="sm"
              class="mt-16"
            />
          </div>

          <!-- 统计数据 -->
          <div class="flex items-center justify-around py-4 border-t border-gray-100">
            <div class="text-center">
              <div class="text-lg font-bold text-gray-800">{{ stats.notes }}</div>
              <div class="text-xs text-gray-500">笔记</div>
            </div>
            <div class="text-center">
              <div class="text-lg font-bold text-gray-800">{{ stats.followers }}</div>
              <div class="text-xs text-gray-500">粉丝</div>
            </div>
            <div class="text-center">
              <div class="text-lg font-bold text-gray-800">{{ stats.following }}</div>
              <div class="text-xs text-gray-500">关注</div>
            </div>
            <div class="text-center">
              <div class="text-lg font-bold text-gray-800">{{ stats.likes }}</div>
              <div class="text-xs text-gray-500">获赞</div>
            </div>
          </div>
        </div>
      </div>

      <!-- 登录提示 -->
      <div
        v-if="!loggedIn"
        class="bg-white mt-4 mx-4 rounded-lg shadow-sm p-6 text-center"
      >
        <UIcon
          name="i-ph-user-circle"
          class="text-6xl text-gray-300 mb-4 mx-auto"
        />
        <h3 class="text-lg font-semibold mb-2 text-gray-800">登录后查看个人主页</h3>
        <p class="text-gray-500 mb-4">登录后可以查看你的笔记、粉丝和关注</p>
        <!-- 使用方法处理点击事件 -->
        <UButton
          @click="handleLoginClick"
          label="立即登录"
          size="lg"
        />
      </div>

      <!-- 标签页 -->
      <div
        v-else
        class="mt-4"
      >
        <div class="flex items-center bg-white border-b border-gray-200">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            class="flex-1 flex items-center justify-center gap-2 py-3 text-sm font-medium transition-colors"
            :class="activeTab === tab.id
              ? 'text-red-500 border-b-2 border-red-500'
              : 'text-gray-600'"
            @click="activeTab = tab.id"
          >
            <UIcon :name="tab.icon" class="text-lg" />
            <span>{{ tab.label }}</span>
          </button>
        </div>

        <!-- 内容区域 -->
        <div class="p-4">
          <div
            v-if="activeTab === 'notes'"
            class="grid grid-cols-3 gap-2"
          >
            <div
              v-for="i in 9"
              :key="i"
              class="aspect-square bg-gray-200 rounded-lg"
            />
          </div>

          <div
            v-else-if="activeTab === 'likes'"
            class="text-center py-12"
          >
            <UIcon
              name="i-ph-heart"
              class="text-6xl text-gray-300 mb-4 mx-auto"
            />
            <p class="text-gray-500">暂无赞过的内容</p>
          </div>

          <div
            v-else
            class="text-center py-12"
          >
            <UIcon
              name="i-ph-bookmark"
              class="text-6xl text-gray-300 mb-4 mx-auto"
            />
            <p class="text-gray-500">暂无收藏的内容</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
