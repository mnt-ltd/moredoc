<template>
  <div class="page-login">
    <el-card shadow="never">
      <div slot="header" class="clearfix">
        <span>管理员登录</span>
      </div>
      <form-login></form-login>
    </el-card>
  </div>
</template>

<script>
export default {
  // 这样也行，多个中间件的时候就这样写
  middleware: ['auth'],
  head() {
    return {
      title: `管理员登录 - MOREDOC · 魔刀文库`,
    }
  },
}
</script>
