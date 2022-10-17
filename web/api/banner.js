import service from '~/utils/request'

export const createBanner = (data) => {
  return service({
    url: '/api/v1/banner',
    method: 'post',
    data,
  })
}

export const updateBanner = (data) => {
  return service({
    url: '/api/v1/banner',
    method: 'put',
    data,
  })
}

export const deleteBanner = (params) => {
  return service({
    url: '/api/v1/banner',
    method: 'delete',
    params,
  })
}

export const getBanner = (params) => {
  return service({
    url: '/api/v1/banner',
    method: 'get',
    params,
  })
}

export const listBanner = (params) => {
  return service({
    url: '/api/v1/banner/list',
    method: 'get',
    params,
  })
}


