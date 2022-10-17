<template>
  <div class="com-form-attachment">
    <el-form
      ref="formAttachment"
      label-position="top"
      label-width="80px"
      :model="attachment"
    >
      <el-form-item
        label="名称"
        prop="name"
        :rules="[
          { required: true, trigger: 'blur', message: '请输入附件名称' },
        ]"
      >
        <el-input
          v-model="attachment.name"
          placeholder="请输入附件名称"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item label="是否合法">
        <el-switch
          v-model="attachment.is_approved"
          style="display: block"
          active-color="#13ce66"
          inactive-color="#ff4949"
          active-text="是"
          inactive-text="否"
        >
        </el-switch>
      </el-form-item>
      <el-form-item label="描述">
        <el-input
          v-model="attachment.description"
          type="textarea"
          rows="3"
          placeholder="请输入附件相关描述或备注"
        ></el-input>
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
import { updateAttachment } from '~/api/attachment'
export default {
  name: 'FormAttachment',
  props: {
    initAttachment: {
      type: Object,
      default: () => {
        return {
          id: 0,
          name: '',
          is_approved: 1,
          description: '',
        }
      },
    },
  },
  data() {
    return {
      loading: false,
      attachment: {
        id: 0,
        name: '',
        description: '',
        is_approved: 1,
      },
    }
  },
  watch: {
    initAttachment: {
      handler(val) {
        const attachment = { ...this.attachment, ...val }
        attachment.is_approved = attachment.is_approved === 1
        this.attachment = attachment
      },
      immediate: true,
    },
  },
  created() {
    const attachment = { ...this.attachment, ...this.initAttachment }
    attachment.is_approved = attachment.is_approved === 1
    this.attachment = attachment
  },
  methods: {
    onSubmit() {
      this.$refs.formAttachment.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const attachment = { ...this.attachment }
        attachment.is_approved = attachment.is_approved ? 1 : 0
        if (this.attachment.id > 0) {
          const res = await updateAttachment(attachment)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        }
        this.loading = false
      })
    },
    clearValidate() {
      this.$refs.formAttachment.clearValidate()
    },
    resetFields() {
      this.$refs.formAttachment.resetFields()
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
