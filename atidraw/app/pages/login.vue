<script setup lang="ts">
import type { ApiResponse, AuthResponseData } from '../../shared/types/api'

const { loggedIn } = useUserSession()

const username = ref('')
const password = ref('')
const loading = ref(false)
const error = ref('')
const isRegister = ref(false)

// 如果已登录，跳转到首页
watch(loggedIn, (value) => {
  if (value) {
    navigateTo('/')
  }
}, { immediate: true })

/**
 * 处理登录
 */
async function handleLogin() {
  error.value = ''
  
  // 参数验证
  if (!username.value || !password.value) {
    error.value = '请输入用户名和密码'
    return
  }

  loading.value = true

  try {
    const result = await $fetch<ApiResponse<AuthResponseData>>('/api/auth/login', {
      method: 'POST',
      body: {
        username: username.value,
        password: password.value,
      },
    })

    // 检查响应格式
    if (result.code === 200 && result.data) {
      // 登录成功后跳转到首页
      await navigateTo('/')
    } else {
      error.value = result.msg || '登录失败'
    }
  } catch (err: any) {
    error.value = err.data?.message || err.message || '登录失败，请检查用户名和密码'
  } finally {
    loading.value = false
  }
}

/**
 * 处理注册
 */
async function handleRegister() {
  error.value = ''
  
  // 参数验证
  if (!username.value || !password.value) {
    error.value = '请输入用户名和密码'
    return
  }

  loading.value = true

  try {
    const result = await $fetch<ApiResponse<AuthResponseData>>('/api/auth/register', {
      method: 'POST',
      body: {
        username: username.value,
        password: password.value,
      },
    })

    // 检查响应格式
    if (result.code === 200 && result.data) {
      // 注册成功后跳转到首页
      await navigateTo('/')
    } else {
      error.value = result.msg || '注册失败'
    }
  } catch (err: any) {
    error.value = err.data?.message || err.message || '注册失败，请检查用户名和密码'
  } finally {
    loading.value = false
  }
}

/**
 * 切换登录/注册模式
 */
function toggleMode() {
  isRegister.value = !isRegister.value
  error.value = ''
  username.value = ''
  password.value = ''
}

/**
 * 提交表单
 */
function handleSubmit() {
  if (isRegister.value) {
    handleRegister()
  } else {
    handleLogin()
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-red-50 to-red-100 px-4">
    <div class="w-full max-w-md">
      <div class="bg-white rounded-lg shadow-lg p-8">
        <h1 class="text-3xl font-bold text-center text-red-500 mb-8">小红书</h1>

        <!-- 登录/注册切换标签 -->
        <div class="flex mb-6 border-b border-gray-200">
          <button
            :class="[
              'flex-1 py-2 text-center font-medium transition-colors',
              !isRegister 
                ? 'text-red-500 border-b-2 border-red-500' 
                : 'text-gray-500 hover:text-gray-700'
            ]"
            @click="toggleMode"
          >
            登录
          </button>
          <button
            :class="[
              'flex-1 py-2 text-center font-medium transition-colors',
              isRegister 
                ? 'text-red-500 border-b-2 border-red-500' 
                : 'text-gray-500 hover:text-gray-700'
            ]"
            @click="toggleMode"
          >
            注册
          </button>
        </div>

        <form @submit.prevent="handleSubmit" class="space-y-4">
          <!-- 用户名输入 -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              用户名
            </label>
            <input
              v-model="username"
              type="text"
              placeholder="输入用户名"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent outline-none transition"
              required
            />
          </div>

          <!-- 密码输入 -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              密码
            </label>
            <input
              v-model="password"
              type="password"
              placeholder="输入密码"
              class="w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-red-500 focus:border-transparent outline-none transition"
              required
            />
          </div>

          <!-- 错误提示 -->
          <div v-if="error" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg text-sm">
            {{ error }}
          </div>

          <!-- 提交按钮 -->
          <button
            type="submit"
            :disabled="loading"
            class="w-full bg-red-500 hover:bg-red-600 disabled:bg-gray-400 text-white font-medium py-2 rounded-lg transition duration-200"
          >
            {{ loading ? (isRegister ? '注册中...' : '登录中...') : (isRegister ? '注册' : '登录') }}
          </button>
        </form>
      </div>
    </div>
  </div>
</template>
