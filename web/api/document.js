import service from '~/utils/request'

export const createDocument = (data) => {
  return service({
    url: '/api/v1/document',
    method: 'post',
    data,
  })
}

export const updateDocument = (data) => {
  return service({
    url: '/api/v1/document',
    method: 'put',
    data,
  })
}

export const deleteDocument = (params) => {
  return service({
    url: '/api/v1/document',
    method: 'delete',
    params,
  })
}

export const getDocument = (params) => {
  return service({
    url: '/api/v1/document',
    method: 'get',
    params,
  })
}

export const listDocument = (params) => {
  return service({
    url: '/api/v1/document/list',
    method: 'get',
    params,
  })
}


