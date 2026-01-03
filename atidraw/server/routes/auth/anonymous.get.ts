export default eventHandler(async (event) => {
  const runtimeConfig = useRuntimeConfig()

  // Allow anonymous sessions regardless of third-party OAuth configuration
  const fingerprint = await getRequestFingerprint(event, {
    userAgent: true,
    xForwardedFor: true,
  })
  await setUserSession(event, {
    user: {
      provider: 'anonymous',
      id: fingerprint || 'anonymous',
      name: `${event.context.cf?.city}, ${event.context.cf?.country}`,
      avatar: '',
      url: '',
    },
  })

  return sendRedirect(event, '/draw')
})
