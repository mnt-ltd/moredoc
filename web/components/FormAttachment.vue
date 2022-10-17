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
          v-model="attachment.enable"
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
          rows="5"
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
        return {}
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
      },
    }
  },
  watch: {
    initAttachment: {
      handler(val) {
        this.attachment = val
      },
      immediate: true,
    },
  },
  created() {
    this.attachment = this.initAttachment
  },
  methods: {
    onSubmit() {
      this.$refs.formAttachment.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const attachment = { ...this.attachment }
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
