<template>
  <div
    class="page page-login"
    :style="
      settings.system.login_background
        ? 'background:url(' +
          settings.system.login_background +
          ') no-repeat center center'
        : ''
    "
  >
    <div>
      <el-card
        shadow="never"
        :class="settings.security.is_close ? 'close-box' : ''"
      >
        <div slot="header" class="clearfix">
          <span v-if="user.id > 0 && settings.security.is_close">网站关闭</span>
          <span v-else>用户登录</span>
        </div>
        <div v-if="settings.security.is_close" class="close-tips">
          <div v-html="settings.security.close_statement"></div>
        </div>
        <form-login
          v-if="!(user.id > 0 && settings.security.is_close)"
          :redirect="redirect"
        ></form-login>
      </el-card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  // 已登录用户，直接跳转到个人中心
  middleware: ['checklogin'],
  head() {
    return {
      title: `用户登录 - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: `用户登录,${this.settings.system.sitename},${this.settings.system.keywords}`,
        },
        {
          hid: 'description',
          name: 'description',
          content: `${this.settings.system.description}`,
        },
      ],
    }
  },
  data() {
    return {
      redirect: this.$route.query.redirect || '/me',
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
    ...mapGetters('user', ['user']),
  },
}
</script>
<style lang="scss">
.page-login {
  width: 100%;
  margin-top: -20px;
  margin-bottom: -20px;
  background-size: cover !important;
  & > div {
    width: $default-width;
    margin: 0 auto;
  }
  .el-card {
    width: 640px;
    max-width: 100%;
    margin: 100px auto;
    margin-right: 0;
    &.close-box {
      margin-right: auto;
      width: 640px;
      .close-tips {
        margin-bottom: 20px;
        border: 1px dashed #f60;
        padding: 20px;
        border-radius: 4px;
        line-height: 180%;
        font-size: 15px;
      }
    }
    .el-card__body {
      padding-bottom: 0;
    }
  }
}
@media screen and (max-width: $mobile-width) {
  .page-login {
    background: none !important;
    & > div {
      width: 100%;
      margin: 0 auto;
    }
    .el-card {
      width: 100%;
      margin: 20px auto;
    }
  }
}
</style>
