import service from '~/utils/request'

export const updatePermission = (data) => {
  return service({
    url: '/api/v1/permission',
    method: 'put',
    data,
  })
}

export const getPermission = (params) => {
  return service({
    url: '/api/v1/permission',
    method: 'get',
    params,
  })
}

export const listPermission = (params) => {
  return service({
    url: '/api/v1/permission/list',
    method: 'get',
    params,
  })
}


