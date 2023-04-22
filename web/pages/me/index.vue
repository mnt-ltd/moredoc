<template>
  <div class="page page-me">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="never">
          <user-card :user="user" />
        </el-card>
      </el-col>
      <el-col :span="18">
        <el-card shadow="never"></el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  name: 'PageMe',
  computed: {
    user() {
      return this.$store.state.user
    },
  },
}
</script>
