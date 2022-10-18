<template>
  <div class="com-form-permission">
    <el-form
      ref="formPermission"
      label-position="top"
      label-width="80px"
      :model="permission"
    >
      <el-row :gutter="20">
        <el-col :span="6">
          <el-form-item label="Method" prop="method">
            <el-input v-model="permission.method" :disabled="true"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="18">
          <el-form-item label="API" prop="path">
            <el-input v-model="permission.path" :disabled="true"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item
        label="名称"
        prop="title"
        :rules="[
          { required: true, trigger: 'blur', message: '请输入权限名称' },
        ]"
      >
        <el-input
          v-model="permission.title"
          placeholder="请输入权限名称"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item label="描述">
        <el-input
          v-model="permission.description"
          type="textarea"
          rows="5"
          placeholder="请输入权限相关描述或备注"
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
import { updatePermission } from '~/api/permission'
export default {
  name: 'FormPermission',
  props: {
    initPermission: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      permission: {
        id: 0,
        title: '',
        description: '',
      },
    }
  },
  watch: {
    initPermission: {
      handler(val) {
        this.permission = val
      },
      immediate: true,
    },
  },
  created() {
    this.permission = this.initPermission
  },
  methods: {
    onSubmit() {
      this.$refs.formPermission.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const permission = { ...this.permission }
        if (this.permission.id > 0) {
          const res = await updatePermission(permission)
          if (res.status === 200) {
            this.$message.success('修改成功')
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
      this.$refs.formPermission.clearValidate()
    },
    resetFields() {
      this.$refs.formPermission.resetFields()
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
