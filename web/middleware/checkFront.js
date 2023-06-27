import { requireLogin } from '~/utils/utils'
export default function ({ store, route, redirect, from }) {
  // Every time the route changes (fired on initialization too)
  // 如果是注册或者登录，则带个redirect参数，用于登录后跳转
  if (
    (route.name === 'login' || route.name === 'register') &&
    !(from.name === 'login' || from.name === 'register')
  ) {
    if (!route.query.redirect) {
      route.query.redirect = from.fullPath
      redirect(route)
      return
    }
  }

  store.dispatch('user/checkAndRefreshUser')
  const settings = store.getters['setting/settings']
  const user = store.getters['user/user']
  const permissions = store.getters['user/permissions'] || []
  if (requireLogin(settings, user, route, permissions)) {
    return redirect('/login')
  }

  // 如果访问的路由前缀是 /me，但是用户没有登录，那么跳转到 /login
  if (route.path.startsWith('/me') && !user.id) {
    return redirect('/login')
  }
}
