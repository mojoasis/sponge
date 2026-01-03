/**
 * URL处理工具函数
 */

/**
 * 判断是否为完整URL（包含协议）
 */
function isAbsoluteUrl(url: string): boolean {
  return /^https?:\/\//i.test(url)
}

/**
 * 拼接OSS URL
 * @param baseUrl OSS基础URL
 * @param path 相对路径或完整URL
 * @returns 完整的URL
 */
export function joinOssUrl(baseUrl: string, path?: string | null): string {
  // 如果path为空，返回空字符串
  if (!path) {
    return ''
  }

  // 如果已经是完整URL，直接返回
  if (isAbsoluteUrl(path)) {
    return path
  }

  // 确保baseUrl以/结尾
  const base = baseUrl.endsWith('/') ? baseUrl : `${baseUrl}/`
  
  // 确保path不以/开头（避免双斜杠）
  const cleanPath = path.startsWith('/') ? path.slice(1) : path

  return `${base}${cleanPath}`
}

/**
 * 批量处理URL数组
 * @param baseUrl OSS基础URL
 * @param paths URL路径数组
 * @returns 处理后的URL数组
 */
export function joinOssUrls(baseUrl: string, paths: (string | null | undefined)[]): string[] {
  return paths
    .map(path => joinOssUrl(baseUrl, path))
    .filter(url => url !== '')
}

