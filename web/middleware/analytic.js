export default function ({ store, route, redirect }) {
  const settings = store.getters['setting/settings']
  window._hmt = window._hmt || []
  window.loaded_hmt = window.loaded_hmt || false
  if (window.loaded_hmt) {
    window._hmt.push(['_trackPageview', route.fullPath])
    return
  }
  try {
    eval(settings.system.analytics)
    window._hmt = _hmt
    window.loaded_hmt = true
  } catch (error) {
    console.log(error)
  }
}
