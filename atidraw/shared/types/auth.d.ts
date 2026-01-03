import type { UserInfo } from './api'

declare module '#auth-utils' {
  interface User {
    provider?: 'github' | 'google' | 'anonymous' | 'local'
    id: string
    name: string
    avatar?: string
    url?: string
    email?: string
    // 扩展字段，包含后端返回的所有用户信息
    userName?: string
    nickName?: string
    signature?: string
    backgroundImage?: string
    phone?: string
  }
  
  interface Session {
    user?: User
    token?: string
  }
}

export type { UserInfo }
