<template>
  <div class="com-form-userinfo">
    <el-row>
      <el-col :span="10" class="text-center">
        <upload-image
          :error-image="'/static/images/avatar.png'"
          :width="'64px'"
          :action="'/api/v1/upload/avatar'"
          :image="user.avatar"
          @success="getUser"
          class="edit-avatar"
        />
        <!-- 上传成功之后，重新获取用户资料 -->
        <div>
          <h3>{{ user.username }}</h3>
        </div>
      </el-col>
      <el-col :span="14">
        <el-descriptions class="margin-top" :column="1">
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-time"></i>
              <span> 注册时间</span>
            </template>
            {{ formatDatetime(user.created_at) }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-time"></i>
              <span> 最后登录</span>
            </template>
            {{ formatDatetime(user.login_at) }}
          </el-descriptions-item>
          <el-descriptions-item>
            <template slot="label">
              <i class="el-icon-location-outline"></i>
              <span> 登录IP</span>
            </template>
            {{ user.last_login_ip || '-' }}
          </el-descriptions-item>
        </el-descriptions>
      </el-col>
    </el-row>
    <el-tabs v-model="activeTab" @tab-click="handleClick">
      <el-tab-pane label="个人资料" name="profile">
        <el-form label-width="80px">
          <el-form-item label="用户名">
            <el-input v-model="user.username" disabled></el-input>
          </el-form-item>
          <el-form-item label="真实姓名">
            <el-input v-model="user.realname" disabled></el-input>
          </el-form-item>
          <el-form-item label="身份证号">
            <el-input v-model="user.identity" disabled></el-input>
          </el-form-item>
          <el-form-item label="联系邮箱">
            <el-input v-model="user.email" disabled></el-input>
          </el-form-item>
          <el-form-item label="联系电话">
            <el-input v-model="user.mobile" disabled></el-input>
          </el-form-item>
          <el-form-item label="联系地址">
            <el-input
              v-model="user.address"
              type="textarea"
              disabled
            ></el-input>
          </el-form-item>
          <el-form-item label="个性签名">
            <el-input
              v-model="user.signature"
              type="textarea"
              disabled
            ></el-input>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <el-tab-pane label="修改资料" name="updateprofile">
        <form-profile />
      </el-tab-pane>
      <el-tab-pane label="修改密码" name="password">
        <form-password />
      </el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import UploadImage from './UploadImage.vue'
import FormProfile from './FormProfile.vue'
import FormPassword from './FormPassword.vue'
import { formatDatetime } from '~/utils/utils'

export default {
  name: 'FormUserinfo',
  components: { UploadImage, FormProfile, FormPassword },
  data() {
    return {
      activeTab: 'profile',
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
  },
  created() {
    this.getUser()
  },
  methods: {
    formatDatetime,
    ...mapActions('user', ['getUser']),
    handleClick(tab, event) {
      console.log(tab, event)
    },
  },
}
</script>

<style lang="scss">
.com-form-userinfo {
  .el-descriptions-item__label {
    span {
      margin-left: 5px;
    }
  }
  .el-descriptions__body .el-descriptions__table {
    color: #888;
  }
  .edit-avatar {
    position: relative;
    &::after {
      font-family: element-icons !important;
      content: '\e78c';
      position: absolute;
      top: 7px;
      margin-left: 10px;
    }
  }
}
@media screen and (max-width: $mobile-width) {
  .com-form-userinfo {
    .el-descriptions-item__container {
      display: block;
    }
  }
}
</style>
