// 对Date的扩展，将 Date 转化为指定格式的String
// 月(M)、日(d)、小时(h)、分(m)、秒(s)、季度(q) 可以用 1-2 个占位符，
// 年(y)可以用 1-4 个占位符，毫秒(S)只能用 1 个占位符(是 1-3 位的数字)
// (new Date()).Format("yyyy-MM-dd hh:mm:ss.S") ==> 2006-07-02 08:09:04.423
// (new Date()).Format("yyyy-M-d h:m:s.S")      ==> 2006-7-2 8:9:4.18
// eslint-disable-next-line no-extend-native
Date.prototype.Format = function (fmt) {
  const o = {
    'M+': this.getMonth() + 1, // 月份
    'd+': this.getDate(), // 日
    'h+': this.getHours(), // 小时
    'm+': this.getMinutes(), // 分
    's+': this.getSeconds(), // 秒
    'q+': Math.floor((this.getMonth() + 3) / 3), // 季度
    S: this.getMilliseconds(), // 毫秒
  }
  if (/(y+)/.test(fmt)) {
    fmt = fmt.replace(
      RegExp.$1,
      (this.getFullYear() + '').substr(4 - RegExp.$1.length)
    )
  }
  for (const k in o) {
    if (new RegExp('(' + k + ')').test(fmt)) {
      fmt = fmt.replace(
        RegExp.$1,
        RegExp.$1.length === 1 ? o[k] : ('00' + o[k]).substr(('' + o[k]).length)
      )
    }
  }
  return fmt
}

export function formatTimeToStr(times, pattern) {
  let d = new Date(times).Format('yyyy-MM-dd hh:mm:ss')
  if (pattern) {
    d = new Date(times).Format(pattern)
  }
  return d.toLocaleString()
}

export function formatDatetime(time) {
  if (typeof time === 'string' && time !== '') {
    const date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  }
  return '-'
}
export function formatDate(time) {
  if (typeof time === 'string' && time !== '') {
    const date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd')
  }
  return '-'
}

export function formatOnlyMonthDate(time) {
  if (typeof time === 'string' && time !== '') {
    const date = new Date(time)
    return formatTimeToStr(date, 'MM-dd')
  }
  return '-'
}

export function formatYear(time) {
  if (typeof time === 'string' && time !== '') {
    const date = new Date(time)
    return formatTimeToStr(date, 'yyyy')
  }
  return '-'
}

export function formatRelativeTime(time) {
  if (!(typeof time === 'string' && time !== '')) {
    return '刚刚'
  }
  const timestamp = parseInt(new Date(time).getTime() / 1000)
  const now = parseInt(new Date().getTime() / 1000)
  const diff = now - timestamp
  const minute = 60
  const hour = minute * 60
  const day = hour * 24
  const month = day * 30

  const monthC = diff / month
  const dayC = diff / day
  const hourC = diff / hour
  const minC = diff / minute

  if (monthC > 12) {
    return parseInt(monthC / 12) + ' 年前'
  } else if (monthC >= 1) {
    return parseInt(monthC) + ' 月前'
  } else if (dayC >= 1) {
    return parseInt(dayC) + ' 天前'
  } else if (hourC >= 1) {
    return parseInt(hourC) + ' 小时前'
  } else if (minC >= 1) {
    return parseInt(minC) + ' 分钟前'
  }
  return '刚刚'
}

export function formatBytes(bytes, decimals = 2) {
  if (!+bytes) return '0 Bytes'

  const k = 1024
  const dm = decimals < 0 ? 0 : decimals
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']

  const i = Math.floor(Math.log(bytes) / Math.log(k))

  return `${(bytes / Math.pow(k, i)).toFixed(dm)} ${sizes[i]}`
}

// categoryToTrees 分类转树形结构
export function categoryToTrees(categories, withDisabled = true) {
  const result = []
  const map = {}
  try {
    // 避免修改原对象。比如第一次调用，categories加了children属性，第二次，同样继续追加children属性的值
    const cates = JSON.parse(JSON.stringify(categories))
    cates.forEach((item) => {
      if (withDisabled) {
        item.disabled = !item.enable
      }
      map[item.id] = item
    })
    cates.forEach((item) => {
      const parent = map[item.parent_id]
      if (parent) {
        if (parent.disabled) item.disabled = true
        ;(parent.children || (parent.children = [])).push(item)
      } else {
        // 如果父级ID不存在，不能直接判断为根节点，还要确认下其父级ID是否存在，存在父级ID的不是根节点
        // eslint-disable-next-line no-lonely-if
        if (!item.parent_id) {
          result.push(item)
        }
      }
    })
  } catch (error) {
    console.log(error)
  }

  return result
}

const extMapIcon = {
  '.pdf': 'pdf',
  '.doc': 'word',
  '.docx': 'word',
  '.rtf': 'word',
  '.wps': 'word',
  '.odt': 'word',
  '.dot': 'word',
  '.ppt': 'ppt',
  '.pptx': 'ppt',
  '.pps': 'ppt',
  '.ppsx': 'ppt',
  '.dps': 'ppt',
  '.odp': 'ppt',
  '.pot': 'ppt',
  '.xls': 'excel',
  '.xlsx': 'excel',
  '.et': 'excel',
  '.ods': 'excel',
  '.csv': 'excel',
  '.tsv': 'excel',
  '.txt': 'text',
  '.epub': 'epub',
  '.mobi': 'mobi',
  '.chm': 'chm',
  '.umd': 'umd',
}

export function getIcon(ext) {
  return extMapIcon[ext] || 'other'
}

// 解析 $route.query 中的数组
export function parseQueryIntArray(query, keys) {
  const result = {}
  keys.forEach((key) => {
    if (typeof query[key] === 'object') {
      result[key] = (query[key] || []).map((item) => parseInt(item))
    } else if (query[key]) {
      result[key] = [parseInt(query[key]) || 0]
    }
  })
  return result
}

export function parseQueryBoolArray(query, keys) {
  const result = {}
  keys.forEach((key) => {
    if (typeof query[key] === 'object') {
      result[key] = (query[key] || []).map((item) => item === 'true')
    } else if (query[key]) {
      result[key] = [query[key] === 'true']
    }
  })
  return result
}

// 是否需要登录。针对关闭站点访问、或登录访问限制
export function requireLogin(settings, user, route, permissions = []) {
  if (settings.security && settings.security.login_required && !user.id) {
    // 未登录，且开启了登录访问限制
    if (
      !(
        route.name === 'login' ||
        route.name === 'register' ||
        route.name === 'findpassword'
      )
    ) {
      return true
    }
  }

  if (settings.security && settings.security.is_close) {
    // 1. 用户未登录，跳转到登录页面
    if (user.id === 0 && route.name !== 'login') {
      return true
    }

    // 用户已登录，如果不是管理员
    if (user.id !== 0 && permissions.length === 0 && route.name !== 'login') {
      return true
    }
  }
  return false
}

export function genLinkHTML(title, href) {
  return `<a href="${href}" target="_blank" class="el-link el-link--primary">${(
    title || '-'
  )
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')}</a>`
}

// 获取时间区间
export function genTimeDuration(duration) {
  // { label: '最近一天', value: 'day' },
  // { label: '最近一周', value: 'week' },
  // { label: '最近一个月', value: 'month' },
  // { label: '最近三个月', value: 'three_month' },
  // { label: '最近半年', value: 'half_year' },
  // { label: '最近一年', value: 'year' },
  const fmt = 'yyyy-MM-dd hh:mm:ss'
  const start = new Date()
  switch (duration) {
    case 'day':
      // 最近一天
      start.setTime(start.getTime() - 3600 * 1000 * 24)
      break
    case 'week':
      // 最近一周
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
      break
    case 'month':
      // 最近一个月
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 31)
      break
    case 'three_month':
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 92)
    // 最近三个月
    case 'half_year':
      // 最近半年
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 183)
      break
    case 'year':
      // 最近一年
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 365)
      break
    default:
      return []
  }
  return [start.Format(fmt), new Date().Format(fmt)]
}

export function genPrevPage(hash, pageNO, ext, enableGZIP){
  if (!ext){
    ext=".svg"
  }
  if (ext===".svg" && enableGZIP){
    ext=".gzip.svg"
  }
  return  `/view/page/${hash}/${pageNO}${ext}`
}