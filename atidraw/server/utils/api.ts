import type { ApiResponse } from '#shared/types/api'

/**
 * 统一的API请求工具函数（服务端使用）
 * @param event H3事件对象，用于获取session中的token
 * @param url API地址
 * @param options 请求选项
 * @returns Promise<ApiResponse<T>>
 */
export async function apiRequest<T = any>(
  event: any,
  url: string,
  options: {
    method?: 'GET' | 'POST' | 'PUT' | 'DELETE'
    body?: any
    headers?: Record<string, string>
    token?: string
    requireAuth?: boolean // 是否需要认证，默认false
  } = {},
): Promise<ApiResponse<T>> {
  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase || 'http://81.70.142.31:18888'

  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    'Accept': '*/*',
    ...options.headers,
  }

  // 自动从session获取token，或使用传入的token
  let token = options.token
  if (!token && event) {
    try {
      const session = await getUserSession(event)
      token = (session as any)?.token
    }
    catch {
      // 如果获取session失败，忽略
    }
  }

  // 如果需要认证但没有token，抛出错误
  if (options.requireAuth && !token) {
    const error = new Error('需要登录认证') as any
    error.statusCode = 401
    throw error
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
  }
  catch (error: any) {
    // 处理网络错误或业务错误
    const apiError = error as any
    // 保留原始状态码
    if (apiError.statusCode) {
      const err = new Error(apiError.data?.message || apiError.data?.msg || apiError.message || '请求失败') as any
      err.statusCode = apiError.statusCode
      err.data = apiError.data
      throw err
    }
    // 处理业务错误
    if (apiError.data) {
      const err = new Error(apiError.data.message || apiError.data.msg || '请求失败') as any
      err.statusCode = apiError.status || apiError.statusCode || 500
      err.data = apiError.data
      throw err
    }
    // 其他错误
    const err = new Error(apiError.message || '网络请求失败') as any
    err.statusCode = apiError.statusCode || 500
    throw err
  }
}
