import axios from 'axios' // 引入axios
import qs from 'qs'
import store from '~/store/index'

const service = axios.create({
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json',
  },
  paramsSerializer(params) {
    // 序列化参数，防止params = {status: [1,2]} 这样的参数解析成 status[]=1&status[]=2，后端无法解析。
    // 使用当前qs序列化，上述参数会被处理为 status=1&status=2
    return qs.stringify(params, { arrayFormat: 'repeat' })
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
    return error.response
  }
)

export default service
