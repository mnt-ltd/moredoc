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
        <form-register></form-register>
      </el-card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'RegisterPage',
  data() {
    return {}
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
    ...mapGetters('user', ['user']),
  },
  created() {
    // 已登录，回到个人中心
    if (this.user.id && this.user.id > 0) {
      this.$router.replace({ name: 'user-id', params: { id: this.user.id } })
    }
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
    width: 520px;
    max-width: 100%;
    margin: 100px auto;
    margin-right: 0;
    .el-card__body {
      padding-bottom: 0;
    }
  }
}
</style>
