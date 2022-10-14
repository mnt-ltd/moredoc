import { Message } from 'element-ui'
import { login, getUser, updateUser } from '~/api/user'
export const user = {
  namespaced: true,
  state: {
    user: {
      id: 0,
      username: '',
      realname: '',
      email: '',
      mobile: '',
      avatar: '',
      address: '',
      signature: '',
    },
    token: '',
  },
  mutations: {
    setUser(state, user) {
      state.user = user
    },
    mergeUser(state, user) {
      state.user = Object.assign(state.user, user)
    },
    setToken(state, token) {
      state.token = token
    },
    logout(state) {
      state.user = {}
      state.token = ''
      localStorage.clear()
    },
  },
  actions: {
    // 获取用户信息
    async getUser({ commit }) {
      const res = await getUser()
      if (res.status === 200) {
        commit('setUser', res.data.user)
      }
      return res
    },
    async updateUser({ commit }, profile) {
      const res = await updateUser(profile)
      if (res.status === 200) {
        commit('mergeUser', profile)
      } else {
        Message({
          type: 'error',
          message: res.data.message || '修改失败',
        })
      }
      return res
    },
    async login({ commit }, loginInfo) {
      const res = await login(loginInfo)
      if (res.status === 200) {
        commit('setUser', res.data.user)
        commit('setToken', res.data.token)
      } else {
        Message({
          type: 'error',
          message: res.data.message || '登录失败',
        })
      }
      return res
    },
    logout({ commit }) {
      commit('logout')
    },
  },
  getters: {
    user(state) {
      return state.user
    },
    token(state) {
      return state.token
    },
  },
}
