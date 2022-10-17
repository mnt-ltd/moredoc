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

  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`
}
