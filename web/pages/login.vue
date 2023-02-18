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
      <el-card shadow="never">
        <div slot="header" class="clearfix">
          <span>用户登录</span>
        </div>
        <form-login></form-login>
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
  computed: {
    ...mapGetters('setting', ['settings']),
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
    width: 520px;
    max-width: 100%;
    margin: 100px auto;
    margin-right: 0;
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
