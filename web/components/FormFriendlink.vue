<template>
  <div class="com-form-friendlink">
    <el-form
      ref="formFriendlink"
      label-position="top"
      label-width="80px"
      :model="friendlink"
    >
      <el-form-item
        label="名称"
        prop="title"
        :rules="[{ required: true, trigger: 'blur', message: '请输入名称' }]"
      >
        <el-input
          v-model="friendlink.title"
          placeholder="请输入名称"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item
        label="地址"
        prop="link"
        :rules="[
          {
            required: true,
            trigger: 'blur',
            message: '请输入友链地址，如 https://mnt.ltd',
          },
        ]"
      >
        <el-input
          v-model="friendlink.link"
          placeholder="请输入友链地址，如 https://mnt.ltd"
          clearable
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="排序(值越大越靠前)">
            <el-input-number
              v-model.number="friendlink.sort"
              clearable
              :min="0"
              :step="1"
              placeholder="请输入排序值"
            ></el-input-number>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="是否启用">
            <el-switch
              v-model="friendlink.enable"
              style="display: block"
              active-color="#13ce66"
              inactive-color="#ff4949"
              active-text="是"
              inactive-text="否"
            >
            </el-switch> </el-form-item
        ></el-col>
      </el-row>

      <el-form-item label="描述">
        <el-input
          v-model="friendlink.description"
          type="textarea"
          rows="3"
          placeholder="请输入友链相关描述或备注"
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
import { createFriendlink, updateFriendlink } from '~/api/friendlink'
export default {
  name: 'FormFriendlink',
  props: {
    initFriendlink: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      friendlink: {},
    }
  },
  watch: {
    initFriendlink: {
      handler(val) {
        this.friendlink = { ...val }
      },
      immediate: true,
    },
  },
  created() {
    this.friendlink = { ...this.initFriendlink }
  },
  methods: {
    onSubmit() {
      this.$refs.formFriendlink.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const friendlink = { ...this.friendlink }
        if (this.friendlink.id > 0) {
          const res = await updateFriendlink(friendlink)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createFriendlink(friendlink)
          if (res.status === 200) {
            this.$message.success('新增成功')
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
      this.$refs.formFriendlink.clearValidate()
    },
    resetFields() {
      this.friendlink = {
        id: 0,
        title: '',
        link: '',
        sort: 0,
        enable: true,
        description: '',
      }
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
