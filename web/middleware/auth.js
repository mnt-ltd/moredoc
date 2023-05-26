export default function ({ store, route, redirect }) {
  const token = store.getters['user/token']
  const allowPages = store.getters['user/allowPages'] || []

  // 未登录或没有权限
  if (!token || allowPages.length === 0) {
    redirect('/')
    return
  }

  // 管理员均可访问的页面
  allowPages.push(
    '/admin',
    '/admin/index',
    '/admin/dashboard',
    '/admin/navigation'
  )

  // 去除route.path最后的斜杠
  let routePath = route.path
  if (routePath.endsWith('/')) {
    routePath = route.path.slice(0, -1)
  }

  // 没有特定页面的访问权限
  if (!allowPages.includes(routePath)) {
    redirect('/admin')
  }
}
