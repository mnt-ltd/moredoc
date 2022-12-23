<template>
  <!-- 更新当前用户自身资料 -->
  <div class="com-form-profile">
    <el-form ref="profile" label-width="80px" :model="profile">
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
        label="联系邮箱"
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
          :rows="3"
        ></el-input>
      </el-form-item>

      <el-form-item class="btn-setprofile">
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
import { mapActions, mapGetters } from 'vuex'
export default {
  name: 'FormProfile',
  data() {
    return {
      profile: {},
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
  },
  created() {
    this.profile = {
      ...this.user,
    }
  },
  methods: {
    ...mapActions('user', ['updateUserProfile']),
    setProfile() {
      this.$refs.profile.validate(async (valid) => {
        if (valid) {
          const res = await this.updateUserProfile(this.profile)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.$emit('success', res)
          }
        }
      })
    },
  },
}
</script>
<style lang="scss">
.com-form-profile {
  .btn-setprofile {
    .el-form-item__content {
      margin-left: 0 !important;
    }
  }
}
</style>
