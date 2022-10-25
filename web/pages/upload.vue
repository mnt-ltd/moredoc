<template>
  <div class="page page-upload">
    <el-row>
      <el-col :span="24">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <strong>上传文档</strong>
          </div>
          <el-row :gutter="20">
            <el-col :span="14">
              <el-form
                ref="form"
                :model="document"
                label-position="top"
                label-width="80px"
              >
                <el-form-item label="分类">
                  <el-input v-model="document.title"></el-input>
                </el-form-item>
                <el-form-item label="售价">
                  <el-input v-model="document.price"></el-input>
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
            <el-col :span="10"> </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
  data() {
    return {
      document: {
        name: '',
      },
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  async created() {},
  methods: {},
}
</script>
