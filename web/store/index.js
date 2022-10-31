import Vue from 'vue'
import Vuex from 'vuex'
import VuexPersistence from 'vuex-persist'
import { user } from '~/store/module/user'
import { category } from '~/store/module/category'

Vue.use(Vuex)

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user', 'category'],
  key: 'moredoc',
})

const store = () =>
  new Vuex.Store({
    modules: {
      user,
      category,
    },
    plugins: [vuexLocal.plugin],
  })

export default store
