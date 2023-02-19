import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'
import { user } from '~/store/module/user'
import { setting } from '~/store/module/setting'
import { category } from '~/store/module/category'
import { device } from '~/store/module/device'

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user', 'category', 'setting', 'device'],
  key: 'moredoc',
})

const store = () =>
  new Vuex.Store({
    modules: {
      user,
      category,
      setting,
      device,
    },
    plugins: [vuexLocal.plugin],
  })

export default store
