export default function ({ store, route, redirect }) {
  const settings = store.getters['setting/settings']
  try {
    var _hmt = _hmt || []
    eval(settings.system.analytics)
    if (_hmt) {
      _hmt.push(['_trackPageview', route.fullPath])
    }
  } catch (error) {
    console.log(error)
  }
}
