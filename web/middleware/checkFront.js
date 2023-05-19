import { requireLogin } from '~/utils/utils'
export default function ({ store, route, redirect }) {
  store.dispatch('user/refreshUser')
  const settings = store.getters['setting/settings']
  const user = store.getters['user/user']
  const permissions = store.getters['user/permissions'] || []
  if (requireLogin(settings, user, route, permissions)) {
    return redirect('/login')
  }
}
