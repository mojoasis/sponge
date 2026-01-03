/**
 * 统一的API响应格式
 */
export interface ApiResponse<T = any> {
  code: number
  data: T
  msg: string
}

/**
 * 用户信息类型
 */
export interface UserInfo {
  id: string
  userName: string
  nickName?: string
  avatar?: string
  signature?: string
  backgroundImage?: string
  phone?: string
  email?: string
}

/**
 * 登录/注册响应数据
 */
export interface AuthResponseData {
  token: string
  user: UserInfo
}

/**
 * 登录/注册请求参数
 */
export interface AuthRequest {
  username: string
  password: string
}

/**
 * 视频Feed项
 */
export interface VideoFeedItem {
  id: string
  title?: string
  description?: string
  coverUrl?: string
  videoUrl?: string
  author?: {
    id: string
    userName?: string
    nickName?: string
    avatar?: string
  }
  stats?: {
    likes?: number
    comments?: number
    shares?: number
  }
  publishTime?: number
  latestTime?: number
}

/**
 * Feed响应数据
 */
export interface FeedResponseData {
  list: VideoFeedItem[]
  hasMore: boolean
  nextLatestTime?: number
}

