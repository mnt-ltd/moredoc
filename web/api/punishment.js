import service from '~/utils/request'

export const createPunishment = (data) => {
  return service({
    url: '/api/v1/punishment',
    method: 'post',
    data,
  })
}

export const updatePunishment = (data) => {
  return service({
    url: '/api/v1/punishment',
    method: 'put',
    data,
  })
}

export const getPunishment = (params) => {
  return service({
    url: '/api/v1/punishment',
    method: 'get',
    params,
  })
}

export const listPunishment = (params) => {
  return service({
    url: '/api/v1/punishment/list',
    method: 'get',
    params,
  })
}

// 取消惩罚
export const cancelPunishment = (data) => {
  return service({
    url: '/api/v1/punishment/cancel',
    method: 'put',
    data,
  })
}
