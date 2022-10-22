<template>
  <div class="com-form-category">
    <el-form
      ref="formCategory"
      label-position="top"
      label-width="80px"
      :model="category"
    >
      <el-form-item label="上级分类">
        <el-cascader
          v-model="category.parent_id"
          :options="trees"
          :filterable="true"
          :props="{
            checkStrictly: true,
            expandTrigger: 'hover',
            label: 'title',
            value: 'id',
          }"
          clearable
          placeholder="请选择上级分类"
        ></el-cascader>
      </el-form-item>
      <el-form-item
        label="名称"
        prop="title"
        :rules="[{ required: true, trigger: 'blur', message: '请输入名称' }]"
      >
        <el-input
          v-model="category.title"
          :placeholder="
            category.id > 0
              ? '请输入分类名称'
              : '请输入分类名称，多个分类名称换行输入，重复的分类名称自动跳过...'
          "
          :type="category.id > 0 ? 'text' : 'textarea'"
          :rows="5"
          clearable
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="排序(值越大越靠前)">
            <el-input-number
              v-model.number="category.sort"
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
              v-model="category.enable"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch> </el-form-item
        ></el-col>
      </el-row>
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
import { createCategory, updateCategory } from '~/api/category'
export default {
  name: 'FormCategory',
  props: {
    initCategory: {
      type: Object,
      default: () => {
        return {}
      },
    },
    trees: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
  data() {
    return {
      loading: false,
      category: {
        id: 0,
        title: '',
        sort: 0,
        enable: false,
      },
    }
  },
  watch: {
    initCategory: {
      handler(val) {
        if (!val.sort) val.sort = 0
        this.category = val
      },
      immediate: true,
    },
  },
  created() {
    this.category = { title: '', sort: 0, ...this.initCategory }
  },
  methods: {
    onSubmit() {
      this.$refs.formCategory.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const category = { ...this.category }
        if (this.category.id > 0) {
          const res = await updateCategory(category)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createCategory(category)
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
      this.$refs.formCategory.clearValidate()
    },
    resetFields() {
      this.$refs.formCategory.resetFields()
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
