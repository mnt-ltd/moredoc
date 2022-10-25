<template>
  <div class="page page-login">
    <el-card shadow="never">
      <div slot="header" class="clearfix">
        <span>用户登录</span>
      </div>
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
    margin: 100px auto;
    margin-right: 0;
    .el-card__body {
      padding-bottom: 0;
    }
  }
}
</style>
