<template>
  <div class="com-form-find-password-step-one">
    <el-form label-position="top" label-width="80px" :model="user">
      <el-form-item label="电子邮箱">
        <el-input
          v-model="user.email"
          placeholder="请输入您注册账户时的电子邮箱"
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
          @click="execFindPassword"
          :loading="loading"
          :disabled="disabled"
          >立即提交</el-button
        >
        <nuxt-link to="/register" title="" class="el-link el-link--default"
          >注册账户</nuxt-link
        >
        <nuxt-link
          to="/login"
          title="登录账户"
          class="el-link el-link--default float-right"
          >登录账户</nuxt-link
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getUserCaptcha, findPasswordStepOne } from '~/api/user'
export default {
  name: 'FormFindPasswordStepOne',
  props: {
    redirect: {
      type: String,
      default: '',
    },
  },
  data() {
    return {
      user: {
        email: '',
        captcha: '',
        captcha_id: '',
      },
      captcha: {
        enable: false,
      },
      loading: false,
      disabled: false,
    }
  },
  created() {
    this.loadCaptcha()
  },
  methods: {
    async execFindPassword() {
      this.loading = true
      const res = await findPasswordStepOne(this.user)
      if (res.status === 200) {
        this.$message.success('提交成功，请查看您的邮箱')
        this.disabled = true
      } else {
        this.loadCaptcha()
        this.$message.error(res.data.message || '请求失败')
      }
      this.loading = false
    },
    async loadCaptcha() {
      const res = await getUserCaptcha({ type: 'find_password', t: Date.now() })
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
