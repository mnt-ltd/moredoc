<template>
  <div class="com-upload-image">
    <el-upload
      class="image-uploader"
      :action="action"
      :headers="{ authorization: `bearer ${token}` }"
      :show-file-list="false"
      :on-success="success"
      :on-error="onError"
      accept="image/*"
      :multiple="false"
      :disabled="disabled"
    >
      <el-image
        v-if="disabled"
        :src="image"
        :style="'width:' + width + ';height:' + height"
      >
        <div slot="error" class="image-slot">
          <img :src="errorImage" />
        </div>
      </el-image>
      <el-tooltip v-else content="点击上传图片" placement="top">
        <el-image :src="image" :style="'width:' + width + ';height:' + height">
          <div slot="error" class="image-slot">
            <img :src="errorImage" />
          </div>
        </el-image>
      </el-tooltip>
    </el-upload>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  name: 'FormNotebookCover',
  props: {
    action: {
      type: String,
      default: '',
    },
    width: {
      type: String,
      default: 'auto',
    },
    height: {
      type: String,
      default: 'auto',
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    image: {
      type: String,
      default: '',
    },
    errorImage: {
      type: String,
      default: '/static/images/blank.png',
    },
  },
  data() {
    return {
      profile: {},
    }
  },
  computed: {
    ...mapGetters('user', ['token']),
  },
  methods: {
    success(res) {
      this.$emit('success', res)
    },
    onError(err) {
      try {
        const message = JSON.parse(err.message)
        this.$message.error(message.message)
      } catch (e) {
        this.$message.error(err)
      }
    },
  },
}
</script>
<style lang="scss">
.com-upload-image {
  .el-image {
    img {
      width: 100%;
      height: auto;
    }
  }
}
</style>
