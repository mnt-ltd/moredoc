import { getSettings } from '~/api/config'
export const setting = {
  namespaced: true,
  state: {
    settings: {
      system: {},
      footer: {},
      security: {},
      display: {},
    },
  },
  mutations: {
    setSettings(state, settings) {
      state.settings = settings
    },
  },
  actions: {
    async getSettings({ commit }) {
      const res = await getSettings()
      if (res.status === 200) {
        commit('setSettings', res.data)
      }
      return res
    },
  },
  getters: {
    settings(state) {
      return state.settings
    },
  },
}
