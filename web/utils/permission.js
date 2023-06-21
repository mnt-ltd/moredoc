// 这里前端维护映射关系，后端只需要返回权限列表即可
// Key 为组，对象为权限相关名称
const cumstomPermissionMap = {
  'api.v1.DocumentAPI': {
    label: '文档管理',
    path: 'ListDocument',
    children: [],
    pages: ['/admin/document', '/admin/document/list'],
  },
  'api.v1.RecycleAPI': {
    label: '回收站管理',
    path: 'ListRecycleDocument',
    children: [],
    pages: ['/admin/document', '/admin/document/recycle'],
  },
  'api.v1.CategoryAPI': {
    label: '分类管理',
    path: 'ListCategory',
    children: [],
    pages: ['/admin/document', '/admin/document/category'],
  },
  'api.v1.UserAPI': {
    label: '用户管理',
    path: 'ListUser',
    children: [],
    pages: ['/admin/user', '/admin/user/list'],
  },
  'api.v1.FriendlinkAPI': {
    label: '友链管理',
    path: 'ListFriendlink',
    children: [],
    pages: ['/admin/friendlink'],
  },
  'api.v1.AttachmentAPI': {
    label: '附件管理',
    path: 'ListAttachment',
    children: [],
    pages: ['/admin/attachment'],
  },
  'api.v1.ReportAPI': {
    label: '举报管理',
    path: 'ListReport',
    children: [],
    pages: ['/admin/report'],
  },
  'api.v1.BannerAPI': {
    label: '横幅管理',
    path: 'ListBanner',
    children: [],
    pages: ['/admin/banner'],
  },
  'api.v1.GroupAPI': {
    label: '角色管理',
    path: 'ListGroup',
    children: [],
    pages: ['/admin/user', '/admin/user/group'],
  },
  'api.v1.PermissionAPI': {
    label: '权限管理',
    path: 'ListPermission',
    children: [],
    pages: ['/admin/user', '/admin/user/permission'],
  },
  'api.v1.ConfigAPI': {
    label: '系统设置',
    path: 'ListConfig',
    children: [],
    pages: ['/admin/config'],
  },
  'api.v1.ArticleAPI': {
    label: '文章管理',
    path: 'ListArticle',
    children: [],
    pages: ['/admin/article'],
  },
  'api.v1.CommentAPI': {
    label: '评论管理',
    path: 'ListComment',
    children: [],
    pages: ['/admin/comment'],
  },
  'api.v1.PunishmentAPI': {
    label: '惩罚管理',
    path: 'ListPunishment',
    children: [],
    pages: ['/admin/user/punishment'],
  },
  upload: {
    id: 0,
    label: '上传管理',
    children: [],
  },
}

// 权限树
export const permissionsToTree = (permissions) => {
  const tree = []
  const permissionMap = JSON.parse(JSON.stringify(cumstomPermissionMap))
  permissions.forEach((permission) => {
    const slice = permission.path.split('/')
    // GRPC和HTTP要分开处理。method===GRPC的，为grpc接口，其他均为HTTP
    if (permission.method === 'GRPC') {
      if (!permissionMap[slice[1]]) {
        permissionMap[slice[1]] = {
          children: [],
        }
      }
      if (slice[2] === permissionMap[slice[1]].path) {
        permissionMap[slice[1]].id = permission.id
      }
      permissionMap[slice[1]].children.push({
        ...permission,
        label: permission.title || permission.path,
      })
    } else {
      if (!permissionMap[slice[3]]) {
        permissionMap[slice[3]] = {
          children: [],
          label: slice[3],
        }
      }
      permissionMap[slice[3]].children.push({
        ...permission,
        label: permission.title || permission.path,
      })
    }
  })
  Object.keys(permissionMap).forEach((key) => {
    tree.push(permissionMap[key])
  })
  return tree
}
