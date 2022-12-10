<template>
  <div class="com-form-document-recommend">
    <el-form
      ref="formDocument"
      label-position="top"
      label-width="80px"
      :model="document"
    >
      <el-form-item label="文档" prop="title">
        <el-input v-model="document.title" :disabled="true"></el-input>
      </el-form-item>
      <el-form-item label="推荐状态" prop="recommend_at">
        <div v-if="document.recommend_at">
          <el-radio-group v-model="document.type">
            <el-radio-button :label="0">取消推荐</el-radio-button>
            <el-radio-button :label="1">推荐</el-radio-button>
            <el-radio-button :label="2">重新推荐</el-radio-button>
          </el-radio-group>
          <el-alert
            class="mgt-20px"
            title="重新推荐，可让文档的推荐排序重新变靠前"
            type="warning"
            :closable="false"
          >
          </el-alert>
        </div>

        <el-switch
          v-else
          v-model="document.type"
          style="display: block"
          active-color="#13ce66"
          inactive-color="#ff4949"
          active-text="设为推荐"
          inactive-text="未推荐"
          :active-value="1"
          :inactive-value="0"
        >
        </el-switch>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          :loading="loading"
          @click="onSubmit"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { setDocumentRecommend } from '~/api/document'

export default {
  name: 'FormDocument',
  props: {
    initDocument: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      document: {
        type: 0,
      },
    }
  },
  watch: {
    initDocument: {
      handler(val) {
        const document = { ...val }
        this.initDocumentWithType(document)
      },
      immediate: true,
    },
  },
  created() {
    const document = { ...this.initDocument }
    this.initDocumentWithType(document)
  },
  methods: {
    initDocumentWithType(document) {
      if (document.recommend_at) {
        // 已推荐
        document.type = 1
      } else {
        // 未推荐
        document.type = 0
      }
      this.document = document
    },
    async onSubmit() {
      const req = {
        id: [this.document.id],
        type: this.document.type,
      }
      const res = await setDocumentRecommend(req)
      if (res.status === 200) {
        this.$message.success('操作成功')
        this.$emit('success')
      } else {
        this.$message.error(res.msg)
      }
      this.loading = false
    },
  },
}
</script>
<style lang="scss"></style>
