<template>
  <div class="com-form-user">
    <el-form ref="user" label-position="top" label-width="80px" :model="user">
      <el-form-item
        label="用户名"
        prop="username"
        :rules="[{ required: true, message: '请输入用户名', trigger: 'blur' }]"
      >
        <el-input
          v-model="user.username"
          placeholder="请输入用户名"
          :disabled="user.id > 0 ? true : false"
        ></el-input>
      </el-form-item>
      <el-form-item
        label="密码"
        prop="password"
        :rules="
          user.id > 0
            ? []
            : [{ required: true, message: '请输入用户密码', trigger: 'blur' }]
        "
      >
        <el-input
          v-model="user.password"
          :placeholder="
            user.id > 0 ? '输入密码表示修改用户密码' : '请输入用户密码'
          "
          type="password"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item
        label="邮箱"
        prop="email"
        :rules="[
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' },
        ]"
      >
        <el-input v-model="user.email" placeholder="请输入邮箱"></el-input>
      </el-form-item>
      <el-form-item
        label="角色"
        prop="group_id"
        :rules="[{ required: true, message: '请选择角色', trigger: 'blur' }]"
      >
        <el-select
          v-model="user.group_id"
          multiple
          filterable
          placeholder="请选择角色"
        >
          <el-option
            v-for="group in groups"
            :key="'group-' + group.id"
            :label="group.title"
            :value="group.id"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="setUser"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { addUser, setUser } from '~/api/user'
export default {
  name: 'FormUser',
  props: {
    groups: {
      type: Array,
      default: () => {
        return []
      },
    },
    initUser: {
      type: Object,
      default: () => {
        return {
          id: 0,
          email: '',
          username: '',
          password: '',
          group_id: [],
        }
      },
    },
  },
  data() {
    return {
      user: { id: 0 },
    }
  },
  watch: {
    initUser: {
      handler(val) {
        this.user = val
      },
      immediate: true,
    },
  },
  created() {
    this.user = this.initUser
  },
  methods: {
    setUser() {
      this.$refs.user.validate(async (valid) => {
        if (valid) {
          if (this.user.id > 0) {
            const res = await setUser({
              id: this.user.id,
              username: this.user.username,
              password: this.user.password,
              group_id: this.user.group_id,
            })
            if (res.status === 200) {
              this.$message.success('设置成功')
              this.$emit('success')
            } else {
              this.$message.error(res.data.message)
            }
          } else {
            const res = await addUser(this.user)
            if (res.status === 200) {
              this.$message.success('新增成功')
              this.$emit('success')
            } else {
              this.$message.error(res.data.message)
            }
          }
        }
      })
    },
    reset() {
      this.user = { id: 0 }
      this.$refs.user.resetFields()
      this.$refs.user.clearValidate()
    },
  },
}
</script>
<style lang="scss">
.com-form-user {
  .el-select {
    width: 100%;
  }
}
</style>
