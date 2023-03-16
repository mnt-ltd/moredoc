<template>
  <div class="com-form-banner">
    <el-form
      ref="formBanner"
      label-position="top"
      label-width="80px"
      :model="banner"
    >
      <el-form-item
        label="图片"
        prop="path"
        :rules="[
          { required: true, message: '请上传横幅图片', trigger: 'blur' },
        ]"
      >
        <UploadImage
          :action="'/api/v1/upload/banner'"
          :image="banner.path"
          :error-image="'/static/images/banner.png'"
          @success="success"
        />
      </el-form-item>
      <el-form-item label="名称" prop="title">
        <el-input
          v-model="banner.title"
          clearable
          placeholder="请输入横幅名称"
        ></el-input>
      </el-form-item>
      <el-form-item label="链接" prop="url">
        <el-input
          v-model="banner.url"
          clearable
          placeholder="请输入链接地址，链接地址为空点击横幅不会跳转"
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="8">
          <el-form-item label="类型" prop="type">
            <el-select
              v-model="banner.type"
              clearable
              placeholder="请选择横幅类型"
            >
              <el-option
                v-for="opt in bannerTypeOptions"
                :key="'type-' + opt.value"
                :label="opt.label"
                :value="opt.value"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="排序(值越大越靠前)" prop="sort">
            <el-input-number
              v-model="banner.sort"
              :min="0"
              :step="1"
            ></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="8">
          <el-form-item label="是否启用" prop="enable">
            <el-switch
              v-model="banner.enable"
              style="display: block; margin-top: 8px"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="描述">
        <el-input
          v-model="banner.description"
          type="textarea"
          rows="5"
          placeholder="请输入附件相关描述或备注"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          :loading="loading"
          @click="onSubmit"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import UploadImage from './UploadImage.vue'
import { createBanner, updateBanner } from '~/api/banner'
import { bannerTypeOptions } from '~/utils/enum'

export default {
  name: 'FormBanner',
  components: { UploadImage },
  props: {
    initBanner: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      banner: {},
      bannerTypeOptions,
    }
  },
  watch: {
    initBanner: {
      handler(val) {
        this.banner = { ...val }
        if (!this.banner.type) this.banner.type = 0
      },
      immediate: true,
    },
  },
  created() {
    this.banner = { ...this.initBanner }
    if (!this.banner.type) this.banner.type = 0
  },
  methods: {
    onSubmit() {
      this.$refs.formBanner.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const banner = { ...this.banner }
        if (this.banner.id > 0) {
          const res = await updateBanner(banner)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createBanner(banner)
          if (res.status === 200) {
            this.$message.success('添加成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        }
        this.loading = false
      })
    },
    clearValidate() {
      this.$refs.formBanner.clearValidate()
    },
    resetFields() {
      this.banner = {
        id: 0,
        title: '',
        sort: 0,
        description: '',
        path: '',
        type: 0,
        enable: true,
        url: '',
      }
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
    success(res) {
      this.banner.path = res.data.path
    },
  },
}
</script>
