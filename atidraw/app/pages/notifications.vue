<script setup lang="ts">
interface Notification {
  id: number
  type: 'like' | 'comment' | 'follow'
  user: {
    name: string
    avatar: string
  }
  content: string
  time: string
  post?: {
    title: string
    image: string
  }
}

const notifications = ref<Notification[]>([
  {
    id: 1,
    type: 'like',
    user: {
      name: '小红薯',
      avatar: '',
    },
    content: '赞了你的笔记',
    time: '2小时前',
    post: {
      title: '今日穿搭分享',
      image: '',
    },
  },
  {
    id: 2,
    type: 'comment',
    user: {
      name: '美食达人',
      avatar: '',
    },
    content: '评论了你的笔记：太棒了！',
    time: '5小时前',
    post: {
      title: '周末美食',
      image: '',
    },
  },
  {
    id: 3,
    type: 'follow',
    user: {
      name: '旅行者',
      avatar: '',
    },
    content: '关注了你',
    time: '1天前',
  },
])

function getNotificationIcon(type: Notification['type']) {
  switch (type) {
    case 'like':
      return 'i-ph-heart-fill'
    case 'comment':
      return 'i-ph-chat-circle-fill'
    case 'follow':
      return 'i-ph-user-plus'
    default:
      return 'i-ph-bell'
  }
}

function getNotificationColor(type: Notification['type']) {
  switch (type) {
    case 'like':
      return 'text-red-500'
    case 'comment':
      return 'text-blue-500'
    case 'follow':
      return 'text-green-500'
    default:
      return 'text-gray-500'
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-2xl mx-auto px-4 py-6">
      <h2 class="text-2xl font-bold mb-6 text-gray-800">通知</h2>
      
      <div class="space-y-3">
        <div
          v-for="notification in notifications"
          :key="notification.id"
          class="bg-white rounded-lg shadow-sm p-4 flex items-start gap-3"
        >
          <UAvatar
            :src="notification.user.avatar"
            size="md"
            icon="i-ph-user"
            class="flex-shrink-0"
          />
          
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-2 mb-1">
              <UIcon
                :name="getNotificationIcon(notification.type)"
                :class="getNotificationColor(notification.type)"
                class="text-lg"
              />
              <span class="font-semibold text-gray-800">{{ notification.user.name }}</span>
              <span class="text-gray-600 text-sm">{{ notification.content }}</span>
            </div>
            
            <p
              v-if="notification.post"
              class="text-sm text-gray-500 mb-2"
            >
              {{ notification.post.title }}
            </p>
            
            <span class="text-xs text-gray-400">{{ notification.time }}</span>
          </div>
        </div>
      </div>
      
      <div
        v-if="notifications.length === 0"
        class="text-center py-12"
      >
        <UIcon
          name="i-ph-bell-slash"
          class="text-6xl text-gray-300 mb-4 mx-auto"
        />
        <p class="text-gray-500">暂无通知</p>
      </div>
    </div>
  </div>
</template>

