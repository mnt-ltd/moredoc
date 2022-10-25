<template>
  <div class="page-login">
    <el-card shadow="never">
      <div slot="header" class="clearfix">
        <span>用户登录</span>
      </div>
      <!-- TODO: 登录成功之后，跳转到指定页面 -->
      <form-login :redirect="'/admin'"></form-login>
    </el-card>
  </div>
</template>

<script>
export default {
  head() {
    return {
      title: `用户登录 - MOREDOC · 魔刀文库`,
    }
  },
}
</script>
<style lang="scss">
.page-login {
  .el-card {
    width: 450px;
    max-width: 100%;
    margin: 20vh auto;
  }
}
</style>
