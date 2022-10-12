import Vue from 'vue'
import Element from 'element-ui'
import hotkeys from 'hotkeys-js'

// import locale from 'element-ui/lib/locale/lang/en'
// Vue.use(Element, { locale })

import zhLocale from 'element-ui/lib/locale/lang/zh-CN'

Vue.use(Element, { zhLocale })
// 以便光标在输入框时快捷键同样有效
hotkeys.filter = (e) => {
  return true
}

// Vue2 引入快捷键
Vue.prototype.$hotkeys = hotkeys
