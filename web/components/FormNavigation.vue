<template>
  <div class="com-form-navigation">
    <el-form
      ref="formNavigation"
      label-position="top"
      label-width="80px"
      :model="navigation"
    >
      <el-form-item
        label="名称"
        prop="title"
        :rules="[{ required: true, trigger: 'blur', message: '请输入名称' }]"
      >
        <el-input
          v-model="navigation.title"
          placeholder="请输入名称"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item
        label="地址"
        prop="href"
        :rules="[
          {
            required: true,
            trigger: 'blur',
            message: '请输入导航地址，如 https://mnt.ltd',
          },
        ]"
      >
        <el-input
          v-model="navigation.href"
          placeholder="请输入导航地址，如 https://mnt.ltd"
          clearable
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="排序(值越大越靠前)">
            <el-input-number
              v-model.number="navigation.sort"
              clearable
              :min="0"
              :step="1"
              placeholder="请输入排序值"
            ></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否启用">
            <el-switch
              v-model="navigation.enable"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch> </el-form-item
        ></el-col>
      </el-row>

      <el-form-item label="描述">
        <el-input
          v-model="navigation.description"
          type="textarea"
          rows="3"
          placeholder="请输入导航相关描述或备注"
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
import { createNavigation, updateNavigation } from '~/api/navigation'
export default {
  name: 'FormNavigation',
  props: {
    initNavigation: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      navigation: {},
    }
  },
  watch: {
    initNavigation: {
      handler(val) {
        this.navigation = { ...val }
      },
      immediate: true,
    },
  },
  created() {
    this.navigation = { ...this.initNavigation }
  },
  methods: {
    onSubmit() {
      this.$refs.formNavigation.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const navigation = { ...this.navigation }
        if (this.navigation.id > 0) {
          const res = await updateNavigation(navigation)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createNavigation(navigation)
          if (res.status === 200) {
            this.$message.success('新增成功')
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
      this.$refs.formNavigation.clearValidate()
    },
    resetFields() {
      this.navigation = {
        id: 0,
        title: '',
        href: '',
        sort: 0,
        enable: true,
        description: '',
      }
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
