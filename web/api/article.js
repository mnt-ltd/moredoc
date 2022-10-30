import service from '~/utils/request'

export const createArticle = (data) => {
  return service({
    url: '/api/v1/article',
    method: 'post',
    data,
  })
}

export const updateArticle = (data) => {
  return service({
    url: '/api/v1/article',
    method: 'put',
    data,
  })
}

export const deleteArticle = (params) => {
  return service({
    url: '/api/v1/article',
    method: 'delete',
    params,
  })
}

export const getArticle = (params) => {
  return service({
    url: '/api/v1/article',
    method: 'get',
    params,
  })
}

export const listArticle = (params) => {
  return service({
    url: '/api/v1/article/list',
    method: 'get',
    params,
  })
}


