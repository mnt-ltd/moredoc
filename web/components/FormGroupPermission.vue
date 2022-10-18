<template>
  <div class="com-form-group-permission">
    <el-form label-position="top" label-width="80px" :model="groupPermission">
      <el-checkbox-group v-model="groupPermission.permission_id">
        <el-checkbox
          v-for="item in permissions"
          :key="'permission-' + item.id"
          :label="item.id"
          >{{ item.title || item.method + ':' + item.path }}</el-checkbox
        >
      </el-checkbox-group>
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
import { listPermission } from '~/api/permission'
import { getGroupPermission, updateGroupPermission } from '~/api/group'
export default {
  name: 'FormGroupPermission',
  props: {
    groupId: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      loading: false,
      groupPermission: {
        group_id: 0,
        permission_id: [],
      },
      permissions: [],
    }
  },
  watch: {
    groupId: {
      handler(val) {
        this.groupPermission.group_id = val
        this.loadAllPermissions()
      },
      immediate: true,
    },
  },
  created() {
    this.loadAllPermissions()
    this.permission = this.initPermission
  },
  methods: {
    async onSubmit() {
      this.loading = true
      const res = await updateGroupPermission(this.groupPermission)
      if (res.status === 200) {
        this.$message.success('设置成功')
        this.$emit('success')
      } else {
        this.$message.error(res.data.message)
      }
      this.loading = false
    },
    async loadAllPermissions() {
      if (this.groupPermission.group_id > 0) {
        this.groupPermission.permission_id = [] // 重置授权信息
        const [resPermissions, resGroupPermissions] = await Promise.all([
          listPermission(),
          getGroupPermission({ id: this.groupPermission.group_id }),
        ])
        if (resPermissions.status !== 200) {
          this.$message.error(resPermissions.data.message)
        }
        if (resGroupPermissions.status !== 200) {
          this.$message.error(resGroupPermissions.data.message)
        }

        if (
          resPermissions.status === 200 &&
          resGroupPermissions.status === 200
        ) {
          this.permissions = resPermissions.data.permission || []
          this.groupPermission.permission_id =
            resGroupPermissions.data.permission_id || []
        }
      }
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
<style lang="scss">
.com-form-group-permission {
  .el-checkbox {
    margin-bottom: 20px;
  }
}
</style>
