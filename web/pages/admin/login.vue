<template>
  <div class="page-login">
    <el-card shadow="never">
      <div slot="header" class="clearfix">
        <span>用户登录</span>
      </div>
      <form-login></form-login>
    </el-card>
  </div>
</template>

<script>
export default {
  // 这样也行，多个中间件的时候就这样写
  name: 'IndexPage',
  middleware: ['auth'],
  head() {
    return {
      title: `用户登录 - MOREDOC · 魔刀文库`,
    }
  },
}
</script>
