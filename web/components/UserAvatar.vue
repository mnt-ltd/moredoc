<template>
  <div class="com-user-avatar">
    <el-avatar
      :size="size"
      :alt="user.username"
      :src="user.avatar"
      @error="errorAvatar"
    >
      <img src="/static/images/avatar.png" />
    </el-avatar>
  </div>
</template>
<script>
export default {
  name: 'UserAvatar',
  props: {
    size: {
      type: Number,
      default: 80,
    },
    src: {
      type: String,
      default: '',
    },
    user: {
      type: Object,
      default: () => {},
    },
  },
  methods: {
    errorAvatar() {
      return true
    },
  },
}
</script>
<style lang="scss">
.com-user-avatar {
  display: inline-block;
  position: relative;
  .el-avatar {
    border: 2px solid #ddd;
    padding: 3px;
    box-sizing: border-box;
    background-color: #fff;
    &:hover {
      border: 2px solid #409eff;
    }
    img {
      // height: 100%;
      // width: 100%;
      // border-radius: 50%;
      margin: 0 auto;
    }
  }
}
</style>
