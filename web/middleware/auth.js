export default function ({ store, route, redirect }) {
  const token = store.getters['user/token']

  // 如果已登录，则不允许在访问登录页
  if (route.name === 'admin-login' && token) {
    redirect('/admin')
  }

  // 如果未登录，则不允许访问 /admin 前缀的页面
  if (
    !token &&
    route.name !== 'admin-login' &&
    route.path.indexOf('/admin') === 0
  ) {
    redirect('/admin/login')
  }
}
