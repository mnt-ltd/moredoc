<template>
  <div class="page page-upload">
    <el-row>
      <el-col :span="24">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <strong>上传文档</strong>
          </div>
          <el-row :gutter="40">
            <el-col :span="14" class="part-left">
              <el-form
                ref="form"
                :model="document"
                label-position="top"
                label-width="80px"
              >
                <el-form-item
                  label="文档分类"
                  prop="category_id"
                  :rules="[
                    {
                      required: true,
                      trigger: 'blur',
                      message: '请选择文档分类',
                    },
                  ]"
                >
                  <el-cascader
                    v-model="document.category_id"
                    :options="categoryTrees"
                    :filterable="true"
                    :disabled="loading"
                    :props="{
                      checkStrictly: true,
                      expandTrigger: 'hover',
                      label: 'title',
                      value: 'id',
                    }"
                    placeholder="请选择文档分类"
                  ></el-cascader>
                </el-form-item>
                <el-form-item
                  :label="`默认售价（${
                    settings.system.credit_name || '魔豆'
                  }）`"
                  prop="price"
                >
                  <el-input-number
                    v-model="document.price"
                    :min="0"
                    :step="1"
                    :disabled="loading"
                  ></el-input-number>
                </el-form-item>
                <el-form-item>
                  <el-upload
                    ref="upload"
                    drag
                    multiple
                    :action="'/api/v1/upload/document'"
                    :headers="{ authorization: `bearer ${token}` }"
                    :show-file-list="false"
                    :disabled="loading || !canIUploadDocument"
                    :auto-upload="false"
                    :on-change="onChange"
                    :file-list="fileList"
                  >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">
                      将文件拖到此处，或<em>点击上传</em>
                    </div>
                  </el-upload>
                  <vxe-table
                    v-if="fileList.length > 0"
                    :data="fileList"
                    style="width: 100%"
                    max-height="480"
                    stripe
                    border="inner"
                    :column-config="{resizable: true}"
                  >
                    <vxe-column type="seq" width="60"></vxe-column>
                    <vxe-column field="title" title="文件" min-width="180">
                      <template #default="{row}">
                        <el-input v-model="row.title" :disabled="loading">
                          <template slot="append">{{
                            row.ext
                          }}</template></el-input
                        >
                        <div v-if="row.error">
                          <el-progress
                            :key="row.name"
                            :percentage="row.percentage"
                            status="exception"
                          ></el-progress>
                          <small class="el-link el-link--danger error-tips">{{
                            row.error
                          }}</small>
                        </div>
                        <el-progress
                          v-else-if="row.percentage > 0"
                          :percentage="row.percentage"
                        ></el-progress>
                      </template>
                    </vxe-column>
                    <vxe-column field="size" title="大小" width="100" sortable>
                      <template #default="{row}">
                        <span>{{ formatBytes(row.size) }}</span>
                      </template>
                    </vxe-column>
                    <vxe-column field="price" :title="`售价(${settings.system.credit_name || '魔豆'})`" :width="130" sortable>
                      <template #default="{row}">
                        <el-input-number
                          v-model="row.price"
                          :min="0"
                          :step="1"
                          :disabled="loading"
                          controls-position="right"
                        ></el-input-number>
                      </template>
                    </vxe-column>
                    <vxe-column width="100" fixed="right">
                      <template #header>
                        操作 (<el-button
                          type="text"
                          size="mini"
                          :disabled="loading"
                          @click="clearAllFiles"
                          >清空</el-button
                        >)
                      </template>
                      <template #default="{rowIndex}">
                        <el-button
                          size="mini"
                          type="text"
                          icon="el-icon-delete"
                          :disabled="loading"
                          @click="handleRemove(rowIndex)"
                        >
                          移除
                        </el-button>
                      </template>
                    </vxe-column>
                  </vxe-table>
                </el-form-item>
                <el-form-item style="margin-bottom: 0">
                  <el-button
                    v-if="canIUploadDocument"
                    type="primary"
                    class="btn-block"
                    :loading="loading"
                    @click="onSubmit"
                  >
                    <span v-if="loading">请勿刷新页面，文档上传中...</span>
                    <span v-else>确定上传</span>
                  </el-button>
                  <el-button
                    v-else
                    type="primary"
                    icon="el-icon-hot-water"
                    class="btn-block"
                    disabled
                  >
                    <span v-if="user.id > 0">您暂无权限上传文档</span>
                    <span v-else>您未登录，请先登录</span>
                  </el-button>
                </el-form-item>
              </el-form>
            </el-col>
            <el-col :span="10" class="upload-tips part-right">
              <div><strong>温馨提示</strong></div>
              <div class="help-block">
                <ul>
                  <li>
                    1. 带有
                    <span class="el-link el-link--danger">*</span> 为必填项。
                  </li>
                  <li>
                    <!-- 应该从管理后台的配置中查询 -->
                    2. 允许上传的最大单个文档大小为：<span
                      class="el-link el-link--primary"
                      >{{
                        settings.security.max_document_size.toFixed(2) ||
                        '50.00'
                      }}
                      MB</span
                    >
                    。
                  </li>
                  <li>3. 支持批量上传</li>
                  <!-- <li>
                    4.
                    <span class="el-link el-link--danger">同名覆盖</span>
                    表示相同名称的文档（含扩展名），直接用新文档文件替换，以达到更新文档文件的目的。
                  </li> -->
                  <li>
                    4. 目前支持的文档类型：
                    <div v-if="wordExt.length > 0">
                      <img src="/static/images/word_24.png" alt="Word文档" />
                      {{ wordExt.join('，') }}
                    </div>
                    <div v-if="pptExt.length > 0">
                      <img src="/static/images/ppt_24.png" alt="PPT文档" />
                      {{ pptExt.join('，') }}
                    </div>
                    <div v-if="excelExt.length > 0">
                      <img src="/static/images/excel_24.png" alt="Excel文档" />
                      {{ excelExt.join('，') }}
                    </div>
                    <div v-if="otherExt.length > 0">
                      <img src="/static/images/other_24.png" alt="其他文档" />
                      {{ otherExt.join('，') }}
                    </div>
                    <div v-if="allowExt.includes('.txt')">
                      <img src="/static/images/text_24.png" alt="TXT文档" />
                      .txt
                    </div>
                    <div v-if="allowExt.includes('.pdf')">
                      <img src="/static/images/pdf_24.png" alt="PDF文档" />
                      .pdf
                    </div>
                  </li>
                  <li>
                    5. 上传遇到问题需要帮助？请查看
                    <nuxt-link
                      to="/article/help"
                      class="el-link el-link--primary"
                      >文库帮助</nuxt-link
                    >
                    和
                    <nuxt-link
                      to="/article/feedback"
                      class="el-link el-link--primary"
                      >意见反馈</nuxt-link
                    >
                  </li>
                  <li>
                    6. 为营造绿色网络环境，严禁上传含有淫秽色情及低俗信息等文档
                  </li>
                  <li>
                    7.
                    对于涉嫌侵权和违法违规的文档，本站有权在不提前通知的情况下对文档进行删除，您在本站上传文档，表示认同该条款
                  </li>
                </ul>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import { formatBytes } from '~/utils/utils'
import { createDocument } from '~/api/document'
import { canIUploadDocument } from '~/api/user'
import { uploadDocument } from '~/api/attachment'
export default {
  data() {
    return {
      canIUploadDocument: false,
      document: {
        category_id: [],
        price: 0,
        overwrite: false,
      },
      maxDocumentSize: 50 * 1024 * 1024,
      fileList: [],
      filesMap: {},
      loading: false,
      percentAge: 0,
      allowExt: [
        '.doc',
        '.docx',
        '.rtf',
        '.wps',
        '.odt',
        '.dot',
        '.ppt',
        '.pptx',
        '.pps',
        '.ppsx',
        '.dps',
        '.odp',
        '.pot',
        '.xls',
        '.xlsx',
        '.csv',
        '.tsv',
        '.et',
        '.ods',
        '.epub',
        '.umd',
        '.chm',
        '.mobi',
        '.txt',
        '.pdf',
      ],
      wordExtEnum: ['.doc', '.docx', '.rtf', '.wps', '.odt', '.dot'],
      pptExtEnum: ['.ppt', '.pptx', '.pps', '.ppsx', '.dps', '.odp', '.pot'],
      excelExtEnum: ['.xls', '.xlsx', '.csv', '.tsv', '.et', '.ods'],
      otherExtEnum: ['.epub', '.umd', '.chm', '.mobi'],
      wordExt: [],
      pptExt: [],
      excelExt: [],
      otherExt: [],
      totalFiles: 0, // 总个数
      totalFailed: 0, // 失败个数
      totalSuccess: 0, // 成功个数
      totalDone: 0, // 完成个数
    }
  },
  head() {
    return {
      title: '上传文档 - ' + this.settings.system.title || 'MOREDOC · 魔豆文库',
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: `上传文档,${this.settings.system.sitename},${this.settings.system.keywords}`,
        },
        {
          hid: 'description',
          name: 'description',
          content: this.settings.system.description,
        },
      ],
    }
  },
  computed: {
    ...mapGetters('user', ['token', 'user']),
    ...mapGetters('category', ['categoryTrees']),
    ...mapGetters('setting', ['settings']),
  },
  async created() {
    const res = await canIUploadDocument()
    if (res.status === 200) {
      this.canIUploadDocument = true
    }
    try {
      this.maxDocumentSize =
        (this.settings.security.max_document_size || 50) * 1024 * 1024
    } catch (error) {
      console.log(error)
    }

    try {
      this.allowExt =
        this.settings.security.document_allowed_ext || this.allowExt
    } catch (error) {
      console.log(error)
    }
    this.allowExt.map((ext) => {
      if (this.wordExtEnum.includes(ext)) {
        this.wordExt.push(ext)
      } else if (this.pptExtEnum.includes(ext)) {
        this.pptExt.push(ext)
      } else if (this.excelExtEnum.includes(ext)) {
        this.excelExt.push(ext)
      } else if (this.otherExtEnum.includes(ext)) {
        this.otherExt.push(ext)
      }
    })
  },
  methods: {
    formatBytes,
    ...mapActions('user', ['getUser']),
    onChange(file) {
      const name = file.name.toLowerCase()
      const ext = file.name.substring(file.name.lastIndexOf('.')).toLowerCase()
      if (!this.allowExt.includes(ext)) {
        this.$message.warning(`${file.name} 不支持的文件格式，忽略该文件`)
        return
      }

      if (file.size > this.maxDocumentSize) {
        this.$message.warning(
          `${file.name} 文件大小${formatBytes(
            file.size
          )} 超过限制（最大${formatBytes(this.maxDocumentSize)}），忽略该文件`
        )
        return
      }

      // 文件不能大于指定的文件大小
      if (
        !this.filesMap[name] &&
        this.allowExt.includes(ext) &&
        file.size <= this.maxDocumentSize
      ) {
        const item = {
          ...file,
          title: file.name.substring(0, file.name.lastIndexOf('.')),
          ext,
          price: this.document.price || 0,
          progressStatus: 'success',
          error: '',
          percentage: 0,
          attachment_id: 0,
        }
        this.filesMap[name] = item
        this.fileList.push(item)
        this.totalFiles = this.fileList.length
      }
    },
    handleRemove(index) {
      this.filesMap[this.fileList[index].name] = null
      this.fileList.splice(index, 1)
    },
    onSubmit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          // 1. 验证文件是否存在
          if (this.fileList.length === 0) {
            this.$message.error('请先选择要上传的文档')
            return
          }

          this.totalFiles = this.fileList.length
          this.totalDone = 0
          this.loading = true
          try {
            // 取消之前上传的请求，不然一直pending，新请求会没法发送
            window.uploadDocumentCancel.map((c) => c())
            window.uploadDocumentCancel = []
          } catch (error) {}

          // chrome 等浏览器同一域名下最多只能同时发起 6 个请求，所以这里将 fileList 拆分成多个数组，每个数组的长度为 2，以便控制并发，每次只同时上传 2 个文件
          const fileList = this.fileList.reduce((prev, cur, index) => {
            const i = Math.floor(index / 2)
            prev[i] = prev[i] || []
            prev[i].push(cur)
            return prev
          }, [])
          fileList.reduce(async (prev, cur) => {
            await prev
            await Promise.all(
              cur.map(async (file) => {
                await this.uploadDocument(file)
              })
            )
          }, Promise.resolve())
        }
      })
    },
    clearAllFiles() {
      if (this.loading) {
        return
      }
      this.fileList = []
      this.filesMap = {}
      this.$refs.upload.clearFiles()
    },
    async uploadDocument(file) {
      if (file.percentage === 100 && file.attachment_id) {
        // 不用再次上传
        this.createDocument(file)
        this.totalDone++
        return
      }
      file.error = ''
      file.progressStatus = 'success'

      const formData = new FormData()
      formData.append('file', file.raw)

      try {
        const res = await uploadDocument(formData, {
          onUploadProgress: (progressEvent) => {
            file.percentage = parseInt(
              (progressEvent.loaded / progressEvent.total) * 100
            )
          },
          // timeout: 1000 * 6,
        })
        if (res.status === 200) {
          file.attachment_id = res.data.data.id || 0
          this.createDocument(file)
          this.totalSuccess++
        } else {
          file.progressStatus = 'exception'
          file.error = res.data.message || res.statusText
          this.$message.error(`《${file.name}》${file.error}`)
          this.totalFailed++
        }
      } catch (error) {
        file.progressStatus = 'exception'
        file.error = '上传失败或超时，请重试'
        this.$message.error(`《${file.name}》${file.error}`)
        this.totalFailed++
      }

      this.totalDone++
      if (this.totalDone === this.totalFiles) {
        this.loading = false
      }
    },
    async createDocument(doc) {
      const createDocumentRequest = {
        overwrite: this.document.overwrite,
        category_id: this.document.category_id,
        document: [
          {
            title: doc.title,
            price: doc.price,
            attachment_id: doc.attachment_id,
          },
        ],
      }
      const res = await createDocument(createDocumentRequest)
      if (res.status === 200) {
        // 从 fileList 中剔除 attachment_id 与当前文档相同的文档
        this.$message.success(`《${doc.title}》上传成功`)
        this.fileList = this.fileList.filter((item) => {
          return item.attachment_id !== doc.attachment_id && doc.attachment_id
        })

        // 过滤 filesMap 中的文档
        this.filesMap = Object.keys(this.filesMap).reduce((acc, key) => {
          if (this.filesMap[key].attachment_id !== doc.attachment_id) {
            acc[key] = this.filesMap[key]
          }
          return acc
        }, {})
      } else {
        this.$message.error(`《${doc.title}》上传失败 ` + res.data.message)
      }
    },
  },
}
</script>
<style lang="scss">
.page-upload {
  .vxe-table{
    .el-input-number{
      width: 100%;
    }
    .vxe-header--column{
      .vxe-cell{
        white-space: normal;
      }
    }
  }
  .el-table {
    .el-input-number {
      width: 120px;
    }
  }
  .el-progress {
    position: absolute;
    width: 100%;
    bottom: -1px;
  }
  .error-tips {
    font-size: 12px;
  }
  .upload-tips {
    line-height: 180%;
    border-left: 1px dashed rgb(252, 155, 91);
    ul,
    li {
      list-style: none;
      margin: 0;
      padding: 0;
    }
    li {
      margin-bottom: 10px;
    }
    .el-link {
      top: -2px;
    }
    img {
      position: relative;
      top: 7px;
    }
  }
}
@media screen and (max-width: $mobile-width) {
  .page-upload {
    .part-left {
      width: 100% !important;
      .el-upload {
        display: block;
        .el-upload-dragger {
          width: 100% !important;
        }
      }
    }
    .part-right {
      width: 100% !important;
      margin-top: 20px;
      li {
        margin-bottom: 0;
      }
    }
  }
}
</style>
