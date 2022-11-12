<template>
  <div class="com-form-group-permission">
    <el-alert
      title="风险提示：当前权限仅针对管理组，普通用户请不要设置此授权！"
      show-icon
      type="warning"
      :closable="false"
    >
    </el-alert>
    <el-form label-position="top" label-width="80px" :model="groupPermission">
      <el-form-item>
        <el-checkbox
          v-model="isCheckedAll"
          :indeterminate="isIndeterminate"
          @change="checkedAll"
          >全选</el-checkbox
        >
        <el-tree
          ref="tree"
          :data="permissionTrees"
          show-checkbox
          node-key="id"
          default-expand-all
          :default-checked-keys="groupPermission.permission_id"
          @check-change="handleCheckChange"
        >
        </el-tree>
      </el-form-item>
      <el-form-item class="btn-fixed">
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
import { mapActions } from 'vuex'
import { listPermission } from '~/api/permission'
import { permissionsToTree } from '~/utils/permission'
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
      permissionTrees: [],
      isCheckedAll: false,
      isIndeterminate: true,
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
    ...mapActions('user', ['getUserPermissions']),
    async onSubmit() {
      this.loading = true
      const res = await updateGroupPermission({
        group_id: this.groupPermission.group_id,
        permission_id: this.$refs.tree.getCheckedKeys(),
      })
      if (res.status === 200) {
        this.groupPermission.permission_id = this.$refs.tree.getCheckedKeys()
        this.$message.success('设置成功')
        this.getUserPermissions()
        this.$emit('success')
      } else {
        this.resetChecked()
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
          this.resetChecked()
          this.$message.error(resPermissions.data.message)
        }
        if (resGroupPermissions.status !== 200) {
          this.resetChecked()
          this.$message.error(resGroupPermissions.data.message)
        }

        if (
          resPermissions.status === 200 &&
          resGroupPermissions.status === 200
        ) {
          const trees = permissionsToTree(resPermissions.data.permission)
          this.permissionTrees = trees
          this.permissions = resPermissions.data.permission || []
          this.groupPermission.permission_id =
            resGroupPermissions.data.permission_id || []
        }
      }
    },
    // 全选
    checkedAll(yes) {
      this.$refs.tree.setCheckedKeys(
        yes ? this.permissions.map((item) => item.id) : []
      )
    },
    handleCheckChange() {
      const checkedKeys = this.$refs.tree.getCheckedKeys()
      let keysLength = 0
      this.permissionTrees.forEach((item) => {
        keysLength++
        if (item.children) {
          keysLength += item.children.length
        }
      })
      this.isCheckedAll = checkedKeys.length === keysLength

      // 中间状态
      this.isIndeterminate =
        checkedKeys.length > 0 && checkedKeys.length < keysLength
    },
    resetChecked() {
      this.$refs.tree.setCheckedKeys(this.groupPermission.permission_id)
    },
  },
}
</script>
<style lang="scss">
.com-form-group-permission {
  .el-form {
    padding-bottom: 80px;
  }
  .btn-fixed {
    position: absolute;
    bottom: -22px;
    z-index: 99;
    background: #fff;
    width: 100%;
    margin-left: -20px;
    padding: 20px;
    box-sizing: border-box;
  }
}
</style>
