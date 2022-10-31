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
                  ></el-input-number>
                </el-form-item>
                <el-form-item label="同名覆盖">
                  <el-switch
                    v-model="document.overwrite"
                    active-color="#13ce66"
                    inactive-color="#ff4949"
                    active-text="是"
                    inactive-text="否"
                  >
                  </el-switch>
                </el-form-item>

                <el-form-item>
                  <el-upload
                    class="upload-demo"
                    drag
                    action="https://jsonplaceholder.typicode.com/posts/"
                    multiple
                  >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">
                      将文件拖到此处，或<em>点击上传</em>
                    </div>
                    <!-- <div slot="tip" class="el-upload__tip">
                    只能上传jpg/png文件，且不超过500kb
                  </div> -->
                  </el-upload>
                </el-form-item>
                <el-form-item style="margin-bottom: 0">
                  <el-button type="primary" class="btn-block"
                    >确定上传</el-button
                  >
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
export default {
  name: 'PageUpload',
  data() {
    return {
      document: {
        category_id: [],
        price: 0,
        overwrite: false,
      },
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  computed: {
    ...mapGetters('category', ['categoryTrees']),
  },
  async created() {},
  methods: {},
}
</script>
<style lang="scss">
.page-upload {
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
