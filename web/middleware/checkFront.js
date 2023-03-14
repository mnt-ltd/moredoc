export default function ({ store, route, redirect }) {
  const settings = store.getters['setting/settings']
  const user = store.getters['user/user']
  const permissions = store.getters['user/permissions'] || []

  if (settings.security && settings.security.login_required && !user.id) {
    // 未登录，且开启了登录访问限制
    if (
      !(
        route.name === 'login' ||
        route.name === 'register' ||
        route.name === 'findpassword'
      )
    ) {
      redirect('/login')
      return
    }
  }

  if (settings.security && settings.security.is_close) {
    // 1. 用户未登录，跳转到登录页面
    if (user.id === 0 && route.name !== 'login') {
      redirect('/login')
      return
    }

    // 用户已登录，如果不是管理员
    if (user.id !== 0 && permissions.length === 0 && route.name !== 'login') {
      redirect('/login')
      return
    }
  }
}
