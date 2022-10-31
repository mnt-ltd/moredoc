import { listCategory } from '~/api/category'
import { categoryToTrees } from '~/utils/utils'
export const category = {
  namespaced: true,
  state: {
    categories: [],
  },
  mutations: {
    setCategories(state, categories) {
      state.categories = categories
    },
  },
  actions: {
    async getCategories({ commit }) {
      const res = await listCategory({
        field: ['id', 'title', 'parent_id', 'cover', 'doc_count', 'enable'],
      })
      if (res.status === 200) {
        commit('setCategories', res.data.category)
      }
      return res
    },
  },
  getters: {
    categories(state) {
      return state.categories
    },
    categoryTrees(state) {
      return categoryToTrees(state.categories)
    },
  },
}
