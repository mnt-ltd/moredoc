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
              v-model="friendlink.status"
              style="display: block"
              active-color="#ff4949"
              inactive-color="#13ce66"
              active-text="否"
              inactive-text="是"
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
      friendlink: {
        id: 0,
        sort: 0,
        status: 0,
        title: '',
        description: '',
      },
    }
  },
  watch: {
    initFriendlink: {
      handler(val) {
        const friendlink = { ...this.friendlink, ...val }
        friendlink.status = !!friendlink.status
        this.friendlink = friendlink
      },
      immediate: true,
    },
  },
  created() {
    this.friendlink = { ...this.friendlink, ...this.initFriendlink }
  },
  methods: {
    onSubmit() {
      this.$refs.formFriendlink.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const friendlink = { ...this.friendlink }
        friendlink.status = friendlink.status ? 1 : 0
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
      this.$refs.formFriendlink.resetFields()
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
