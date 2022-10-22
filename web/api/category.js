import service from '~/utils/request'

export const createCategory = (data) => {
  return service({
    url: '/api/v1/category',
    method: 'post',
    data,
  })
}

export const updateCategory = (data) => {
  return service({
    url: '/api/v1/category',
    method: 'put',
    data,
  })
}

export const deleteCategory = (params) => {
  return service({
    url: '/api/v1/category',
    method: 'delete',
    params,
  })
}

export const getCategory = (params) => {
  return service({
    url: '/api/v1/category',
    method: 'get',
    params,
  })
}

export const listCategory = (params) => {
  return service({
    url: '/api/v1/category/list',
    method: 'get',
    params,
  })
}


