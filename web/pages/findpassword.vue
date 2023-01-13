<template>
  <div class="page page-findpassword">
    <el-card shadow="never">
      <div slot="header">找回密码</div>
      <form-find-password-step-two v-if="token" />
      <form-find-password-step-one v-else />
    </el-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        email: '',
      },
      token: this.$route.query.token,
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  async created() {},
  methods: {},
}
</script>
<style lang="scss">
.page-findpassword {
  .el-card {
    width: 720px;
    margin: 0 auto;
  }
}
</style>
