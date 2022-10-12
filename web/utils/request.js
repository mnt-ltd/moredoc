import axios from 'axios' // 引入axios
import store from '~/store/index'

const service = axios.create({
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// http request 拦截器
service.interceptors.request.use(
  (config) => {
    const token = store().getters['user/token'] || ''
    if (token) config.headers.authorization = `Bearer ${token}`
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// http response 拦截器
service.interceptors.response.use(
  (response) => {
    return response
  },
  (error) => {
    if (error.response.status === 401) {
      store().commit('user/logout')
    }
    // let message = error.response.data.message || error.response.statusText
    // Message({
    //   showClose: true,
    //   message: message,
    //   type: 'error',
    // })
    return error.response
  }
)

export default service
