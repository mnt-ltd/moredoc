<template>
  <el-image
    :src="
      document.attachment && document.attachment.hash
        ? `/view/cover/${document.attachment.hash}`
        : ''
    "
    lazy
    :alt="document.title"
    :style="{ height: height }"
  >
    <div slot="error" class="image-slot">
      <img src="/static/images/default-cover.png" />
    </div>
  </el-image>
</template>
<script>
export default {
  name: 'DocumentCover',
  props: {
    document: {
      type: Object,
      default: () => ({}),
    },
    height: {
      type: String,
      default: '',
    },
  },
  data() {
    return {}
  },
}
</script>
<style lang="scss" scoped>
.el-image {
  border: 2px solid #efefef;
  border-radius: 5px;
}
img {
  width: 100%;
  transition: transform 0.3s ease 0s;
  &:hover {
    transform: scale(1.2);
  }
}
</style>
