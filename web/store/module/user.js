import { Message } from 'element-ui'
import { login, getUser, setUserProfile } from '~/api/user'
export const user = {
  namespaced: true,
  state: {
    user: {
      username: '',
      realname: '',
      email: '',
      mobile: '',
      avatar: '',
      status: false,
      limit: 0,
    },
    token: '',
  },
  mutations: {
    setUser(state, user) {
      state.user = user
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
    async GetUser({ commit }) {
      const res = await getUser()
      if (res.status === 200) {
        commit('setUser', res.data.data.user)
      }
      return res
    },
    async setUserProfile({ commit }, profile) {
      const res = await setUserProfile(profile)
      if (res.status === 200) {
        commit('setUser', res.data.data)
        Message({
          type: 'success',
          message: '修改成功',
        })
      } else {
        Message({
          type: 'error',
          message: res.data.message || '修改失败',
        })
      }
      return res
    },
    async Login({ commit }, loginInfo) {
      const res = await login(loginInfo)
      if (res.status === 200) {
        commit('setUser', res.data.data.user)
        commit('setToken', res.data.data.token)
        Message({
          type: 'success',
          message: '登录成功',
        })
        location.reload()
      } else {
        Message({
          type: 'error',
          message: res.data.message || '登录失败',
        })
      }
    },
    Logout({ commit }) {
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
