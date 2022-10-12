<template>
  <div class="com-form-profile">
    <el-form label-position="top" label-width="80px" :model="profile">
      <el-form-item label="用户名">
        <el-input
          v-model="profile.username"
          placeholder="请输入您的登录用户名"
          :disabled="true"
        ></el-input>
      </el-form-item>
      <el-form-item
        label="真实姓名"
        prop="realname"
        :rules="[
          { required: true, trigger: 'blur', message: '请输入您的真实姓名' },
        ]"
      >
        <el-input v-model="profile.realname"></el-input>
      </el-form-item>
      <el-form-item label="电子邮箱">
        <el-input v-model="profile.email"></el-input>
      </el-form-item>
      <el-form-item label="联系电话">
        <el-input v-model="profile.mobile"></el-input>
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
    ...mapActions('user', ['setUserProfile']),
    async setProfile() {
      const res = await this.setUserProfile(this.profile)
      if (res.status === 200) {
        this.$emit('success', res)
      }
    },
  },
}
</script>
