import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'
import { user } from '~/store/module/user'
import { setting } from '~/store/module/setting'
import { category } from '~/store/module/category'

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user', 'category', 'setting'],
  key: 'moredoc',
})

const store = () =>
  new Vuex.Store({
    modules: {
      user,
      category,
      setting,
    },
    plugins: [vuexLocal.plugin],
  })

export default store
