<template>
  <!-- 更新指定用户的资料 -->
  <div class="com-form-user-profile">
    <el-form label-position="left" label-width="80px" :model="profile">
      <el-form-item label="用户名">
        <el-input
          v-model="profile.username"
          placeholder="请输入您的登录用户名"
          :disabled="true"
        ></el-input>
      </el-form-item>
      <el-form-item label="真实姓名" prop="realname">
        <el-input v-model="profile.realname" clearable></el-input>
      </el-form-item>
      <el-form-item label="身份证号">
        <el-input v-model="profile.identity" clearable></el-input>
      </el-form-item>
      <el-form-item
        label="电子邮箱"
        prop="email"
        :rules="[
          { required: true, message: '请输入电子邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入正确的电子邮箱', trigger: 'blur' },
        ]"
      >
        <el-input v-model="profile.email" clearable></el-input>
      </el-form-item>
      <el-form-item label="联系电话">
        <el-input v-model="profile.mobile" clearable></el-input>
      </el-form-item>
      <el-form-item label="联系地址">
        <el-input
          v-model="profile.address"
          clearable
          type="textarea"
          :rows="3"
        ></el-input>
      </el-form-item>
      <el-form-item label="个性签名">
        <el-input
          v-model="profile.signature"
          type="textarea"
          clearable
          :rows="5"
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="setProfile"
          >修改资料</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { updateUserProfile } from '~/api/user'

export default {
  name: 'FormUserProfile',
  props: {
    initUser: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      profile: {},
    }
  },
  watch: {
    initUser: {
      handler(val) {
        this.profile = val
      },
      immediate: true,
    },
  },
  created() {
    this.profile = {
      ...this.initUser,
    }
  },
  methods: {
    reset() {
      this.profile = {}
    },
    async setProfile() {
      const res = await updateUserProfile(this.profile)
      console.log(this.profile, res)
      if (res.status === 200) {
        this.$message.success('修改成功')
        this.$emit('success')
      } else {
        this.$message.error(res.data.message)
      }
    },
  },
}
</script>
