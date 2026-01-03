import type { ApiResponse } from '~/shared/types/api'

/**
 * 客户端API请求composable
 * 自动从session获取token并添加到请求头
 */
export const useApi = () => {
  const { data: session } = useUserSession()
  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase || 'http://localhost:18888'

  /**
   * 统一的API请求函数（客户端使用）
   * @param url API地址
   * @param options 请求选项
   * @returns Promise<ApiResponse<T>>
   */
  const apiRequest = async <T = any>(
    url: string,
    options: {
      method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
      body?: any
      headers?: Record<string, string>
      requireAuth?: boolean // 是否需要认证，默认false
    } = {},
  ): Promise<ApiResponse<T>> => {
    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      'Accept': '*/*',
      ...options.headers,
    }

    // 自动从session获取token
    const token = (session.value as any)?.token

    // 如果需要认证但没有token，抛出错误
    if (options.requireAuth && !token) {
      throw new Error('需要登录认证')
    }

    // 如果有token，添加到Authorization header
    if (token) {
      headers.Authorization = `Bearer ${token}`
    }

    try {
      const response = await $fetch<ApiResponse<T>>(url, {
        method: options.method || 'GET',
        baseURL: apiBase,
        headers,
        body: options.body,
      })

      // 检查响应格式
      if (!response || typeof response !== 'object' || !('code' in response)) {
        throw new Error('无效的API响应格式')
      }

      // 检查业务状态码
      if (response.code !== 200) {
        throw new Error(response.msg || '请求失败')
      }

      return response
    } catch (error: any) {
      // 处理网络错误或业务错误
      if (error.data) {
        throw new Error(error.data.message || error.data.msg || '请求失败')
      }
      throw new Error(error.message || '网络请求失败')
    }
  }

  return {
    apiRequest,
  }
}

