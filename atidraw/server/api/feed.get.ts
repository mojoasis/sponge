// server/api/feed.get.ts
import type { FeedResponseData } from '#shared/types/api'
import { apiRequest } from '../utils/api'

export default defineEventHandler(async (event) => {
  const query = await getQuery<{
    latestTime?: string
    page?: string
    size?: string
    channel_id?: string
  }>(event)

  const latestTime = query.latestTime ? Number.parseInt(query.latestTime) : Date.now()
  const page = query.page ? Number.parseInt(query.page) : 1
  const size = query.size ? Number.parseInt(query.size) : 10

  try {
    // 构建查询参数
    const queryParams = new URLSearchParams({
      latestTime: latestTime.toString(),
      page: page.toString(),
      size: size.toString(),
    })

    // 使用统一的API请求工具，如果有token会自动携带，但不强制要求认证
    const res = await apiRequest<FeedResponseData>(
      event,
      `/gateway/video/v1/feed?${queryParams.toString()}`,
      {
        method: 'GET',
        requireAuth: false, // 不需要认证，但如果有token会自动携带
      },
    )

    // 返回统一格式的响应
    return {
      code: 200,
      data: res.data,
      msg: '查询成功',
    }
  }
  catch (error: any) {
    // 记录详细错误信息用于调试
    console.error('[Feed API Error]', {
      message: error.message,
      statusCode: error.statusCode,
      data: error.data,
      stack: error.stack,
    })

    // 根据 Nuxt 4 的建议，使用 message 而不是 statusMessage
    const statusCode = error.statusCode || (error.message?.includes('需要登录') ? 401 : 500)
    const message = error.message || '获取feed失败'

    throw createError({
      statusCode,
      message,
    })
  }
})
