import { requireLogin } from '~/utils/utils'
export default function ({ store, route, redirect }) {
  store.dispatch('user/refreshUser')
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
