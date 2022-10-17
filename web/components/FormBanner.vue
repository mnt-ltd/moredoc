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
          :width="'600px'"
          @success="success"
        />
      </el-form-item>
      <el-form-item
        label="名称"
        prop="title"
        :rules="[
          { required: true, message: '请输入横幅名称', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="banner.title"
          clearable
          placeholder="请输入横幅名称"
        ></el-input>
      </el-form-item>
      <el-form-item
        label="链接"
        prop="url"
        :rules="[
          {
            required: true,
            message: '请输入链接地址',
            trigger: 'blur',
          },
        ]"
      >
        <el-input
          v-model="banner.url"
          clearable
          placeholder="请输入链接地址"
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
          <el-form-item label="是否启用" prop="status">
            <el-switch
              v-model="banner.status"
              style="display: block; margin-top: 8px"
              active-color="#ff4949"
              inactive-color="#13ce66"
              active-text="否"
              inactive-text="是"
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
        return {
          id: 0,
          title: '',
          description: '',
        }
      },
    },
  },
  data() {
    return {
      loading: false,
      banner: {
        id: 0,
        title: '',
        description: '',
        path: '',
        type: 0,
        status: 0,
      },
      bannerTypeOptions,
    }
  },
  watch: {
    initBanner: {
      handler(val) {
        const banner = { ...this.banner, ...val }
        this.banner = banner
      },
      immediate: true,
    },
  },
  created() {
    const banner = { ...this.banner, ...this.initBanner }
    this.banner = banner
  },
  methods: {
    onSubmit() {
      this.$refs.formBanner.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const banner = { ...this.banner }
        banner.status = banner.status ? 1 : 0
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
      this.$refs.formBanner.resetFields()
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
    success(res) {
      this.banner.path = res.data.path
      console.log(res)
    },
  },
}
</script>
