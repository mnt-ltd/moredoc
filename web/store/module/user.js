import { Message } from 'element-ui'
import {
  login,
  getUser,
  updateUserProfile,
  logout,
  getUserPermissions,
  register,
} from '~/api/user'
import { permissionsToTree } from '~/utils/permission'
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
    permissions: [],
    allowPages: [],
  },
  mutations: {
    setUser(state, user) {
      state.user = user
    },
    mergeUser(state, user) {
      state.user = { ...state.user, ...user }
    },
    setToken(state, token) {
      state.token = token
    },
    logout(state) {
      state.user = {}
      state.token = ''
      state.permissions = []
      state.allowPages = []
      localStorage.clear()
    },
    setPermissions(state, permissions) {
      state.permissions = permissions
    },
    setAllowPages(state, pages) {
      state.allowPages = pages
    },
  },
  actions: {
    // 获取用户信息
    async getUser({ commit }) {
      const res = await getUser()
      if (res.status === 200) {
        commit('setUser', res.data)
      }
      return res
    },
    async updateUserProfile({ commit }, profile) {
      const res = await updateUserProfile(profile)
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
    async register({ commit, dispatch }, registerInfo) {
      const res = await register(registerInfo)
      if (res.status !== 200) {
        Message({
          type: 'error',
          message: res.data.message || '注册失败',
        })
        return res
      }
      commit('setUser', res.data.user)
      commit('setToken', res.data.token)
      // 获取用户权限
      await dispatch('getUserPermissions')
      return res
    },
    async login({ commit, dispatch }, loginInfo) {
      const res = await login(loginInfo)
      if (res.status !== 200) {
        Message({
          type: 'error',
          message: res.data.message || '登录失败',
        })
        return res
      }
      commit('setUser', res.data.user)
      commit('setToken', res.data.token)
      // 获取用户权限
      await dispatch('getUserPermissions')
      return res
    },
    async logout({ commit }) {
      const res = await logout()
      commit('logout')
      return res
    },
    async getUserPermissions({ commit }) {
      const res = await getUserPermissions()
      if (res.status === 200) {
        commit('setPermissions', res.data.permission)
        const allowPages = []
        try {
          const trees = permissionsToTree(res.data.permission)
          trees.forEach((tree) => {
            if (tree.pages && tree.id && tree.id > 0) {
              allowPages.push(...tree.pages)
            }
          })
        } catch (error) {}
        commit('setAllowPages', allowPages)
      } else {
        Message({
          type: 'error',
          message: res.data.message || '获取权限失败',
        })
      }
      return res
    },
    checkAndRefreshUser({ commit, state }) {
      try {
        const moredoc = JSON.parse(localStorage.getItem('moredoc'))
        if (state.token !== moredoc.user.token) {
          // 以 localStorage 存储的信息为准
          console.log('exec checkAndRefreshUser')
          commit('setUser', moredoc.user.user || {})
          commit('setToken', moredoc.user.token || '')
          commit('setPermissions', moredoc.user.permissions || [])
          commit('setAllowPages', moredoc.user.allowPages || [])
        }
      } catch (error) {
        console.log(error)
      }
    },
  },
  getters: {
    user(state) {
      return state.user
    },
    token(state) {
      return state.token
    },
    permissions(state) {
      return state.permissions
    },
    allowPages(state) {
      return state.allowPages
    },
  },
}
