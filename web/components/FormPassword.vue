<template>
  <div class="com-form-password">
    <el-form
      ref="formPassword"
      label-position="top"
      label-width="80px"
      :model="profile"
    >
      <el-form-item label="用户名">
        <el-input
          v-model="profile.username"
          placeholder="请输入您的登录用户名"
          :disabled="true"
        ></el-input>
      </el-form-item>
      <el-form-item
        label="原密码"
        prop="old_password"
        :rules="[
          { required: true, trigger: 'blur', message: '请输入您的原密码' },
        ]"
      >
        <el-input v-model="profile.old_password" type="password"></el-input>
      </el-form-item>
      <el-form-item
        label="新密码"
        prop="new_password"
        :rules="[
          { required: true, trigger: 'blur', message: '请输入您的新密码' },
        ]"
      >
        <el-input v-model="profile.new_password" type="password"></el-input>
      </el-form-item>
      <el-form-item
        label="确认密码"
        prop="repeat_password"
        :rules="[
          { required: true, trigger: 'blur', message: '请再次输入您的新密码' },
        ]"
      >
        <el-input v-model="profile.repeat_password" type="password"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="setPassword"
          >修改密码</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
import { updateUserPassword } from '~/api/user'
export default {
  name: 'FormProfile',
  data() {
    return {
      profile: {
        username: '',
        old_password: '',
        new_password: '',
        repeat_password: '',
      },
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
  },
  created() {
    this.profile = {
      ...this.profile,
      username: this.user.username,
    }
  },
  methods: {
    setPassword() {
      this.$refs.formPassword.validate(async (valid) => {
        if (valid) {
          if (this.profile.new_password !== this.profile.repeat_password) {
            this.$message.error('新密码和确认密码不一致')
            return
          }
          const res = await updateUserPassword({
            id: this.user.id,
            old_password: this.profile.old_password,
            new_password: this.profile.new_password,
          })
          if (res.status === 200) {
            this.$message.success('密码修改成功')
            this.$refs.formPassword.resetFields()
            this.$emit('success', res)
          } else {
            this.$message.error(res.data.message)
          }
        }
      })
    },
  },
}
</script>
