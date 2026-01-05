/**
 * OSS URL处理 Composable
 * 提供统一的URL拼接功能
 */
// @ts-ignore
import { useRuntimeConfig } from '#imports'
export const useOssUrl = () => {
  const config = useRuntimeConfig()
  const ossBaseUrl = config.public.ossUrl || ''

  /**
   * 拼接OSS URL
   * @param path 相对路径或完整URL
   * @returns 完整的URL
   */
  const getOssUrl = (path?: string | null): string => {
    if (!path) {
      return ''
    }

    // 如果已经是完整URL，直接返回
    if (/^https?:\/\//i.test(path)) {
      return path
    }

    // 确保baseUrl以/结尾
    const base = ossBaseUrl.endsWith('/') ? ossBaseUrl : `${ossBaseUrl}/`

    // 确保path不以/开头（避免双斜杠）
    const cleanPath = path.startsWith('/') ? path.slice(1) : path

    return `${base}${cleanPath}`
  }

  /**
   * 批量处理URL
   * @param paths URL路径数组
   * @returns 处理后的URL数组
   */
  const getOssUrls = (paths: (string | null | undefined)[]): string[] => {
    return paths
      .map(path => getOssUrl(path))
      .filter(url => url !== '')
  }

  return {
    ossBaseUrl,
    getOssUrl,
    getOssUrls,
  }
}

