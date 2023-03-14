<template>
  <div class="com-form-comment-check">
    <el-form
      ref="form"
      :model="icomment"
      class="form-comment-check"
      label-position="top"
    >
      <el-form-item prop="content" label="评论内容">
        <el-input
          v-model="icomment.content"
          type="textarea"
          :placeholder="placeholder"
          :autosize="{ minRows: 4, maxRows: 6 }"
          disabled
        />
      </el-form-item>
      <el-form-item label="审核状态">
        <el-radio-group v-model="icomment.status">
          <el-radio :label="0">待审核</el-radio>
          <el-radio :label="1">审核通过</el-radio>
          <el-radio :label="2">审核拒绝</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="onSubmit"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { updateComment } from '~/api/comment'
export default {
  name: 'FormCommentCheck',
  props: {
    comment: {
      type: Object,
      default: () => {},
    },
    placeholder: {
      type: String,
      default: '请输入评论内容',
    },
  },
  data() {
    return {
      icomment: {
        id: 0,
        content: '',
        status: 0,
      },
    }
  },
  watch: {
    comment: {
      handler(val) {
        this.icomment = { status: 0, ...val }
      },
      immediate: true,
    },
  },
  methods: {
    async onSubmit() {
      const res = await updateComment(this.icomment)
      if (res.status === 200) {
        this.$message.success('更新成功')
        this.$emit('success')
      } else {
        this.$message.error(res.data.message)
      }
    },
  },
}
</script>
<style lang="scss"></style>
