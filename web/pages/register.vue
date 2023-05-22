<template>
  <div
    class="page page-register"
    :style="
      settings.system.register_background
        ? 'background:url(' +
          settings.system.register_background +
          ') no-repeat center center'
        : ''
    "
  >
    <div>
      <el-card shadow="never">
        <div slot="header" class="clearfix">
          <span>用户注册</span>
        </div>
        <form-register :redirect="redirect"></form-register>
      </el-card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  // 已登录用户，直接跳转到个人中心
  middleware: ['checklogin'],
  data() {
    return {
      redirect: this.$route.query.redirect || '/me',
    }
  },
  head() {
    return {
      title: `用户注册 - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: `用户注册,${this.settings.system.sitename},${this.settings.system.keywords}`,
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
  methods: {},
}
</script>
<style lang="scss">
.page-register {
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
    .el-card__body {
      padding-bottom: 0;
    }
  }
}

@media screen and (max-width: $mobile-width) {
  .page-register {
    background: none !important;
    & > div {
      width: 100%;
      margin: 0;
    }
    .el-card {
      width: 100%;
      margin: 20px 0;
    }
  }
}
</style>
