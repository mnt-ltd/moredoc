export const device = {
  namespaced: true,
  state: {
    device: {
      isMobile: false,
      isPad: false,
      isPC: false,
      width: 0,
    },
  },
  mutations: {
    setDevice(state, width) {
      state.device.width = width
      if (width < 768) {
        state.device.isMobile = true
        state.device.isPad = false
        state.device.isPC = false
      } else if (width < 992) {
        state.device.isMobile = false
        state.device.isPad = true
        state.device.isPC = false
      } else {
        state.device.isMobile = false
        state.device.isPad = false
        state.device.isPC = true
      }
    },
  },
  actions: {
    async setDeviceWidth({ commit }, width) {
      commit('setDevice', width)
    },
  },
  getters: {
    isMobile(state) {
      return state.device.isMobile
    },
    isPad(state) {
      return state.device.isPad
    },
    isPC(state) {
      return state.device.isPC
    },
    width() {
      return state.device.width
    },
  },
}
