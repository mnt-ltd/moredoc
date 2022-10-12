<template>
  <div class="com-form-login">
    <el-form label-position="top" label-width="80px" :model="user">
      <el-form-item label="用户名">
        <el-input
          v-model="user.username"
          placeholder="请输入您的登录用户名"
        ></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input
          v-model="user.password"
          placeholder="请输入您的登录密码"
          type="password"
          @keydown.native.enter="execLogin"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="execLogin"
          >立即登录</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { mapActions } from 'vuex'
import {getUserCaptcha} from '~/api/user'
export default {
  name: 'FormLogin',
  data() {
    return {
      user: {
        username: '',
        password: '',
      },
    }
  },
  async created(){
    const res = await getUserCaptcha({type:'login'})
    console.log(res)
  },
  methods: {
    ...mapActions('user', ['Login']),
    async execLogin() {
      await this.Login(this.user)
    },
  },
}
</script>
