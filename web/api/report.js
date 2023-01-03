import service from '~/utils/request'

export const createReport = (data) => {
  return service({
    url: '/api/v1/report',
    method: 'post',
    data,
  })
}

export const updateReport = (data) => {
  return service({
    url: '/api/v1/report',
    method: 'put',
    data,
  })
}

export const deleteReport = (params) => {
  return service({
    url: '/api/v1/report',
    method: 'delete',
    params,
  })
}

export const listReport = (params) => {
  return service({
    url: '/api/v1/report/list',
    method: 'get',
    params,
  })
}


