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
          >立即登录</el-button
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
    }
  },
  created() {
    this.loadCaptcha()
  },
  methods: {
    ...mapActions('user', ['Login']),
    async execLogin() {
      await this.Login(this.user)
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
