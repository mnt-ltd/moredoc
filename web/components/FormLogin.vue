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
      <el-form-item v-if="captcha.enable" label="验证码">
        <div v-if="captcha.type == 'audio'">
          <el-row :gutter="15">
            <el-col :span="20">
              <audio controls="controls" :src="captcha.captcha"></audio>
            </el-col>
            <el-col :span="4">
              <el-tooltip placement="top" content="刷新语音验证码">
                <el-button
                  icon="el-icon-refresh"
                  class="btn-audio-refresh"
                  @click="loadCaptcha"
                ></el-button>
              </el-tooltip>
            </el-col>
          </el-row>
        </div>
        <div v-else>
          <el-tooltip placement="right" content="点击可刷新验证码">
            <img :src="captcha.captcha" class="pointer" @click="loadCaptcha" />
          </el-tooltip>
        </div>
        <el-input v-model="user.captcha" placeholder="请输入验证码"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="execLogin"
          :loading="loading"
          >立即登录</el-button
        >
        <nuxt-link
          to="/findpassword"
          title="找回密码"
          class="el-link el-link--default"
          >找回密码</nuxt-link
        >
        <nuxt-link
          :to="{ name: 'register', query: { redirect } }"
          title="注册账户"
          class="el-link el-link--default float-right"
          >注册账户</nuxt-link
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { mapActions } from 'vuex'
import { getUserCaptcha } from '~/api/user'
export default {
  name: 'FormLogin',
  props: {
    redirect: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      user: {
        username: '',
        password: '',
        captcha: '',
        captcha_id: '',
      },
      captcha: {
        enable: false,
      },
      loading: false,
    }
  },
  created() {
    this.loadCaptcha()
  },
  methods: {
    ...mapActions('user', ['login']),
    async execLogin() {
      this.loading = true
      const res = await this.login(this.user)
      if (res.status === 200) {
        this.$message.success('登录成功')
        if (this.redirect) {
          this.$router.push(this.redirect)
        } else {
          this.$router.push({ name: 'index' })
        }
        this.loading = false
      } else {
        this.loadCaptcha()
        this.loading = false
      }
    },
    async loadCaptcha() {
      const res = await getUserCaptcha({ type: 'login', t: Date.now() })
      if (res.data.enable) {
        // 启用了验证码
        this.user = {
          ...this.user,
          captcha_id: res.data.id,
        }
        this.captcha = res.data
      }
    },
  },
}
</script>
<style scoped>
.btn-audio-refresh {
  vertical-align: -webkit-baseline-middle;
}
</style>
