export const userStatusOptions = [
  { label: '正常', value: 0 },
  { label: '禁用', value: 1 },
  { label: '审核中', value: 2 },
  { label: '拒绝', value: 3 },
  { label: '忽略', value: 4 },
]

export const attachmentTypeOptions = [
  { label: '未知', value: 0 },
  { label: '头像', value: 1 },
  { label: '文档', value: 2 },
  { label: '文章', value: 3 },
  { label: '评论', value: 4 },
  { label: '横幅', value: 5 },
  { label: '分类封面', value: 6 },
  { label: '配置', value: 7 },
]

export const documentStatusOptions = [
  { label: '待转换', value: 0, type: 'info' },
  { label: '转换中', value: 1, type: 'primary' },
  { label: '已转换', value: 2, type: 'success' },
  { label: '转换失败', value: 3, type: 'warning' },
  { label: '已禁用', value: 4, type: 'danger' },
  { label: '重新转换', value: 5, type: 'info' },
]

// 0 网站横幅，1 小程序横幅，2 APP横幅
export const bannerTypeOptions = [
  { label: '网站横幅', value: 0, type: 'primary' },
  { label: '小程序横幅', value: 1, type: 'success' },
  { label: 'APP横幅', value: 2, type: 'warning' },
]

export const boolOptions = [
  { label: '是', value: true, type: 'success' },
  { label: '否', value: false, type: 'danger' },
]

export const reportOptions = [
  { label: '垃圾广告', value: 1 },
  { label: '淫秽色情', value: 2 },
  { label: '虚假中奖', value: 3 },
  { label: '敏感信息', value: 4 },
  { label: '人身攻击', value: 5 },
  { label: '骚扰他人', value: 6 },
]

export const methodOptions = [
  {
    label: 'GET',
    value: 'GET',
    type: 'success',
  },
  {
    label: 'POST',
    value: 'POST',
    type: 'primary',
  },
  {
    label: 'PUT',
    value: 'PUT',
    type: 'warning',
  },
  {
    label: 'DELETE',
    value: 'DELETE',
    type: 'danger',
  },
  {
    label: 'GRPC',
    value: 'GRPC',
    type: 'primary',
  },
]

export const datetimePickerOptions = {
  shortcuts: [
    {
      text: '最近一周',
      onClick(picker) {
        const end = new Date()
        const start = new Date()
        start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
        picker.$emit('pick', [start, end])
      },
    },
    {
      text: '最近一个月',
      onClick(picker) {
        const end = new Date()
        const start = new Date()
        start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
        picker.$emit('pick', [start, end])
      },
    },
    {
      text: '最近三个月',
      onClick(picker) {
        const end = new Date()
        const start = new Date()
        start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
        picker.$emit('pick', [start, end])
      },
    },
    {
      text: '最近半年',
      onClick(picker) {
        const end = new Date()
        const start = new Date()
        start.setTime(start.getTime() - 3600 * 1000 * 24 * 183)
        picker.$emit('pick', [start, end])
      },
    },
    {
      text: '最近一年',
      onClick(picker) {
        const end = new Date()
        const start = new Date()
        start.setTime(start.getTime() - 3600 * 1000 * 24 * 365)
        picker.$emit('pick', [start, end])
      },
    },
  ],
}

export const datetimePickerPunishmentOptions = {
  shortcuts: [
    {
      text: '1小时',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000)
        picker.$emit('pick', end)
      },
    },
    {
      text: '12小时',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 12)
        picker.$emit('pick', end)
      },
    },
    {
      text: '1天',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 24)
        picker.$emit('pick', end)
      },
    },
    {
      text: '2天',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 24 * 2)
        picker.$emit('pick', end)
      },
    },
    {
      text: '1周',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 24 * 7)
        picker.$emit('pick', end)
      },
    },
    {
      text: '1月',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 24 * 30)
        picker.$emit('pick', end)
      },
    },
    {
      text: '半年',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 24 * 183)
        picker.$emit('pick', [start, end])
      },
    },
    {
      text: '1年',
      onClick(picker) {
        const end = new Date()
        end.setTime(end.getTime() + 3600 * 1000 * 24 * 365)
        picker.$emit('pick', [start, end])
      },
    },
  ],
}

export const punishmentTypeOptions = [
  { label: '禁用账户', value: 1, type: 'danger' },
  { label: '禁止评论', value: 2, type: 'warning' },
  { label: '禁止上传', value: 3, type: 'warning' },
  { label: '禁止下载', value: 4, type: 'warning' },
  { label: '禁止收藏', value: 5, type: 'warning' },
]
