// server/api/auth/login.post.ts
import type { AuthResponseData, AuthRequest } from '#shared/types/api'
import { apiRequest } from '../../utils/api'

export default defineEventHandler(async (event) => {
  const body = await readBody<AuthRequest>(event)

  // 参数验证
  if (!body.username || !body.password) {
    throw createError({
      statusCode: 400,
      message: '用户名和密码不能为空',
    })
  }

  try {
    // 使用统一的API请求工具
    const res = await apiRequest<AuthResponseData>(event, '/gateway/user/v1/login', {
      method: 'POST',
      body: {
        username: body.username,
        password: body.password,
      },
    })

    const { token, user } = res.data

    // 设置用户会话，存储完整的用户信息和token
    await setUserSession(event, {
      user: {
        provider: 'local',
        id: String(user.id),
        name: user.nickName || user.userName,
        userName: user.userName,
        nickName: user.nickName,
        avatar: user.avatar,
        email: user.email,
        signature: user.signature,
        backgroundImage: user.backgroundImage,
        phone: user.phone,
      },
      token,
    })

    // 返回统一格式的响应
    return {
      code: 200,
      data: {
        token,
        user,
      },
      msg: '登录成功',
    }
  }
  catch (error: any) {
    throw createError({
      statusCode: 401,
      message: error.message || '登录失败',
    })
  }
})
