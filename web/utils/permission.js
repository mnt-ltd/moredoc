// 这里前端维护映射关系，后端只需要返回权限列表即可
const cumstomPermissionMap = {
  'api.v1.UserAPI': {
    label: '用户管理',
    path: 'ListUser',
    children: [],
  },
  'api.v1.FriendlinkAPI': {
    label: '友链管理',
    path: 'ListFriendlink',
    children: [],
  },
  'api.v1.AttachmentAPI': {
    label: '附件管理',
    path: 'ListAttachment',
    children: [],
  },
  'api.v1.BannerAPI': {
    label: '横幅管理',
    path: 'ListBanner',
    children: [],
  },
  'api.v1.GroupAPI': {
    label: '角色管理',
    path: 'ListGroup',
    children: [],
  },
  'api.v1.PermissionAPI': {
    label: '权限管理',
    path: 'ListPermission',
    children: [],
  },
  'api.v1.ConfigAPI': {
    label: '系统设置',
    path: 'ListConfig',
    children: [],
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
