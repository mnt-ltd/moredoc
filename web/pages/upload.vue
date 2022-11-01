<template>
  <div class="page page-upload">
    <el-row>
      <el-col :span="24">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <strong>上传文档</strong>
          </div>
          <el-row :gutter="40">
            <el-col :span="14">
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
                <el-form-item label="默认售价（魔豆）">
                  <el-input-number
                    v-model="document.price"
                    :min="0"
                    :step="1"
                    :disabled="loading"
                  ></el-input-number>
                </el-form-item>
                <el-form-item label="同名覆盖">
                  <el-switch
                    v-model="document.overwrite"
                    active-color="#13ce66"
                    inactive-color="#ff4949"
                    active-text="是"
                    inactive-text="否"
                    :disabled="loading"
                  >
                  </el-switch>
                </el-form-item>

                <el-form-item>
                  <el-upload
                    ref="upload"
                    class="upload-demo"
                    drag
                    :action="'/api/v1/upload/document'"
                    :headers="{ authorization: `bearer ${token}` }"
                    :show-file-list="false"
                    :on-success="onSuccess"
                    :on-error="onError"
                    multiple
                    :disabled="loading"
                    :auto-upload="false"
                    :on-change="handleChange"
                    :file-list="fileList"
                  >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">
                      将文件拖到此处，或<em>点击上传</em>
                    </div>
                  </el-upload>
                  <el-table
                    v-if="fileList.length > 0"
                    :data="fileList"
                    style="width: 100%"
                    max-height="360"
                  >
                    <el-table-column prop="title" label="文件" min-width="180">
                      <template slot-scope="scope">
                        <el-input v-model="scope.row.title" :disabled="loading">
                          <template slot="append">{{
                            scope.row.ext
                          }}</template></el-input
                        >
                      </template>
                    </el-table-column>
                    <el-table-column prop="size" label="大小" width="100">
                      <template slot-scope="scope">
                        {{ formatBytes(scope.row.size) }}
                      </template>
                    </el-table-column>
                    <el-table-column
                      prop="price"
                      label="售价(魔豆)"
                      width="130"
                    >
                      <template slot-scope="scope">
                        <el-input-number
                          v-model="scope.row.price"
                          :min="0"
                          :step="1"
                          :disabled="loading"
                          controls-position="right"
                        ></el-input-number>
                      </template>
                    </el-table-column>
                    <el-table-column label="操作" width="70" fixed="right">
                      <template slot-scope="scope">
                        <el-button
                          size="mini"
                          type="text"
                          icon="el-icon-delete"
                          :disabled="loading"
                          @click="handleRemove(scope.$index)"
                        >
                          移除
                        </el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-form-item>
                <el-form-item style="margin-bottom: 0">
                  <el-progress
                    v-if="loading"
                    :percentage="percentAge"
                    :text-inside="true"
                    :stroke-width="12"
                    status="success"
                  ></el-progress>
                  <div v-if="loading" class="mgt-20px"></div>
                  <el-button
                    type="primary"
                    class="btn-block"
                    :loading="loading"
                    @click="onSubmit"
                  >
                    <span v-if="loading">上传中...</span>
                    <span v-else>确定上传</span>
                  </el-button>
                </el-form-item>
              </el-form>
            </el-col>
            <el-col :span="10" class="upload-tips">
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
                      >50.00 MB</span
                    >
                    。
                  </li>
                  <li>3. 支持批量上传</li>
                  <li>
                    4.
                    <span class="el-link el-link--danger">同名覆盖</span>
                    表示相同名称的文档（含扩展名），直接用新文档覆盖替换
                  </li>
                  <li>
                    3. 目前支持的文档类型：
                    <div>
                      <img src="/static/images/word_24.png" alt="Word文档" />
                      doc，docx，rtf，wps，odt
                    </div>
                    <div>
                      <img src="/static/images/ppt_24.png" alt="PPT文档" />
                      ppt，pptx，pps，ppsx，dps，odp，pot
                    </div>
                    <div>
                      <img src="/static/images/excel_24.png" alt="Excel文档" />
                      xls，xlsx，et，ods
                    </div>
                    <div>
                      <img src="/static/images/other_24.png" alt="其他文档" />
                      epub，umd，chm，mobi
                    </div>
                    <div>
                      <img src="/static/images/text_24.png" alt="TXT文档" /> txt
                    </div>
                    <div>
                      <img src="/static/images/pdf_24.png" alt="PDF文档" />
                      pdf
                    </div>
                  </li>
                  <li>
                    4. 上传遇到问题需要帮助？请查看
                    <nuxt-link
                      to="/article/help"
                      class="el-link el-link--default"
                      >文库帮助</nuxt-link
                    >
                    和
                    <nuxt-link
                      to="/article/feedback"
                      class="el-link el-link--default"
                      >意见反馈</nuxt-link
                    >
                  </li>
                  <li>
                    5. 为营造绿色网络环境，严禁上传含有淫秽色情及低俗信息等文档
                  </li>
                  <li>
                    6.
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
import { mapGetters } from 'vuex'
import { formatBytes } from '~/utils/utils'
export default {
  name: 'PageUpload',
  data() {
    return {
      document: {
        category_id: [],
        price: 0,
        overwrite: false,
      },
      loading: false,
      percentAge: 0,
      fileList: [],
      allowExt: [
        '.doc',
        '.docx',
        '.rtf',
        '.wps',
        '.odt',
        '.ppt',
        '.pptx',
        '.pps',
        '.ppsx',
        '.dps',
        '.odp',
        '.pot',
        '.xls',
        '.xlsx',
        '.et',
        '.ods',
        '.epub',
        '.umd',
        '.chm',
        '.mobi',
        '.txt',
        '.pdf',
      ],
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  computed: {
    ...mapGetters('category', ['categoryTrees']),
    ...mapGetters('user', ['token']),
  },
  async created() {},
  methods: {
    formatBytes,
    handleChange(file, files) {
      const filesMap = {}
      const filesList = []
      files.forEach((item) => {
        const name = item.name.toLowerCase()
        item.price = this.document.price || 0
        item.title = item.name.substring(0, item.name.lastIndexOf('.'))
        item.ext = item.name.substring(item.name.lastIndexOf('.')).toLowerCase()
        if (!filesMap[name] && this.allowExt.includes(item.ext)) {
          filesMap[name] = item
          filesList.push(item)
        }
      })
      this.fileList = filesList
    },
    handleRemove(index) {
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
          this.loading = true
          this.$refs.upload.submit()
        }
      })
    },
    onError(err) {
      console.log(err)
      this.$message.error(err.message)
    },
    onSuccess(res, file, fileList) {
      const length = fileList.length
      const successItems = fileList.filter(
        (item) => item.response && item.response.code === 200
      )
      const percentAge = parseInt((successItems.length / length) * 100)
      if (percentAge === 100) {
        // 上传成功100%之后，提交表单数据
        this.$message.success('上传成功')
        this.loading = false
        this.percentAge = 0
        // this.fileList = []
        // this.$refs.upload.clearFiles()
      } else {
        this.percentAge = percentAge
      }
    },
  },
}
</script>
<style lang="scss">
.page-upload {
  .el-table {
    .el-input-number {
      width: 120px;
    }
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
</style>
