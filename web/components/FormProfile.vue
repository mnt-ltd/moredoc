<template>
  <div class="com-form-profile">
    <el-form label-position="left" label-width="80px" :model="profile">
      <el-form-item label="用户名">
        <el-input
          v-model="profile.username"
          placeholder="请输入您的登录用户名"
          :disabled="true"
        ></el-input>
      </el-form-item>
      <el-form-item label="真实姓名" prop="realname">
        <el-input v-model="profile.realname"></el-input>
      </el-form-item>
      <el-form-item label="电子邮箱">
        <el-input v-model="profile.email"></el-input>
      </el-form-item>
      <el-form-item label="联系电话">
        <el-input v-model="profile.mobile"></el-input>
      </el-form-item>
      <el-form-item label="联系地址">
        <el-input
          v-model="profile.address"
          type="textarea"
          :rows="3"
        ></el-input>
      </el-form-item>
      <el-form-item label="个性签名">
        <el-input
          v-model="profile.signature"
          type="textarea"
          :rows="3"
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
    ...mapActions('user', ['updateUser']),
    async setProfile() {
      const res = await this.updateUser(this.profile)
      if (res.status === 200) {
        this.$emit('success', res)
      }
    },
  },
}
</script>
