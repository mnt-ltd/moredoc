import service from '~/utils/request'

export const createNavigation = (data) => {
  return service({
    url: '/api/v1/navigation',
    method: 'post',
    data,
  })
}

export const updateNavigation = (data) => {
  return service({
    url: '/api/v1/navigation',
    method: 'put',
    data,
  })
}

export const deleteNavigation = (params) => {
  return service({
    url: '/api/v1/navigation',
    method: 'delete',
    params,
  })
}

export const getNavigation = (params) => {
  return service({
    url: '/api/v1/navigation',
    method: 'get',
    params,
  })
}

export const listNavigation = (params) => {
  return service({
    url: '/api/v1/navigation/list',
    method: 'get',
    params,
  })
}


