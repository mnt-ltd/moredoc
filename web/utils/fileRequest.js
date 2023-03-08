import axios from 'axios' // 引入axios
import store from '~/store/index'
const cancelToken = axios.CancelToken // 取消请求
window.uploadDocumentCancel = [] // 取消上传文件请求

const fileService = axios.create({
  timeout: 6000000, // 文件上传超时时间，100分钟
  headers: {
    'Content-Type': 'multipart/form-data',
  },
})

fileService.interceptors.request.use(
  (config) => {
    const token = store().getters['user/token'] || ''
    if (token) config.headers.authorization = `Bearer ${token}`
    config.cancelToken = new cancelToken((c) => {
      window.uploadDocumentCancel.push(c)
    })
    return config
  },
  (error) => {
    console.log(error)
    return Promise.reject(error)
  }
)

export default fileService
