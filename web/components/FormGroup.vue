<template>
  <div class="com-form-group">
    <el-form
      ref="formGroup"
      label-position="top"
      label-width="80px"
      :model="group"
    >
      <el-form-item
        label="名称"
        prop="title"
        :rules="[{ required: true, trigger: 'blur', message: '请输入名称' }]"
      >
        <el-input
          v-model="group.title"
          placeholder="请输入名称"
          clearable
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="排序(值越大越靠前)">
            <el-input-number
              v-model.number="group.sort"
              clearable
              :min="0"
              :step="1"
              placeholder="请输入排序值"
            ></el-input-number>
          </el-form-item>
          <el-form-item label="是否为默认用户组">
            <el-switch
              v-model="group.is_default"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch>
          </el-form-item>
          <el-form-item label="允许上传文档">
            <el-switch
              v-model="group.enable_upload"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch>
          </el-form-item>
          <el-form-item label="评论需审核">
            <el-switch
              v-model="group.enable_comment_approval"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch>
          </el-form-item>
          <el-form-item label="是否在用户名后展示">
            <el-switch
              v-model="group.is_display"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="分组颜色">
            <chrome-picker v-model="colors"></chrome-picker> </el-form-item
        ></el-col>
      </el-row>

      <el-form-item label="描述">
        <el-input
          v-model="group.description"
          type="textarea"
          rows="5"
          placeholder="请输入分组描述"
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
import { Chrome } from 'vue-color'
import { createGroup, updateGroup } from '~/api/group'
export default {
  name: 'FormGroup',
  components: {
    'chrome-picker': Chrome,
  },
  props: {
    initGroup: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      colors: {
        hsl: { h: 154.92214648092607, s: 0, l: 0, a: 1 },
        hex: '#000000',
        hex8: '#000000FF',
        rgba: { r: 0, g: 0, b: 0, a: 1 },
        hsv: { h: 154.92214648092607, s: 0, v: 0, a: 1 },
        oldHue: 154.92214648092607,
        source: 'hex',
        a: 1,
      },
      group: {
        sort: 0,
      },
    }
  },
  watch: {
    initGroup: {
      handler(val) {
        if (!val.sort) val.sort = 0
        this.group = val
        this.colors = val.color || '#000000FF'
      },
      immediate: true,
    },
  },
  created() {
    this.group = this.initGroup
    if (!this.initGroup.sort) this.group.sort = 0
    this.colors = this.initGroup.color || '#000000FF'
  },
  methods: {
    onSubmit() {
      this.group.color = this.colors.hex8
      this.$refs.formGroup.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const group = { ...this.group }
        if (group.id > 0) {
          const res = await updateGroup(group)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.$emit('success')
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createGroup(group)
          if (res.status === 200) {
            this.$message.success('新增成功')
            this.$emit('success')
          } else {
            this.$message.error(res.data.message)
          }
        }
        this.loading = false
      })
    },
  },
}
</script>
