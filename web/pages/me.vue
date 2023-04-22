<template>
  <div class="page page-me">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="never">
          <user-card :user="user" />
          <nuxt-link to="/upload">
            <el-button
              type="primary"
              icon="el-icon-upload2"
              class="btn-block mgt-20px"
              >上传文档</el-button
            >
          </nuxt-link>
          <el-menu
            class="mgt-20px"
            :router="true"
            :default-active="defaultActive.value"
          >
            <el-menu-item
              v-for="item in tabs"
              :key="item.value"
              :index="item.value"
            >
              <i :class="item.icon"></i>
              <span slot="title">{{ item.label }}</span>
            </el-menu-item>
          </el-menu>
        </el-card>
      </el-col>
      <el-col :span="18">
        <el-card shadow="never">
          <div slot="header">{{ defaultActive.label }}</div>
          <div class="nuxt-child">
            <nuxt-child :user="user" />
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
export default {
  name: 'PageMe',
  computed: {
    ...mapGetters('user', ['user']),
    ...mapGetters('setting', ['settings']),
  },
  data() {
    return {
      defaultActive: {
        value: '/me',
        label: '我的动态',
      },
      tabs: [
        {
          label: '我的动态',
          value: '/me',
          icon: 'el-icon-magic-stick',
        },
        {
          label: '我的文档',
          value: '/me/document',
          icon: 'el-icon-document',
        },
        {
          label: '我的收藏',
          value: '/me/favorite',
          icon: 'el-icon-star-off',
        },
        {
          label: '我的下载',
          value: '/me/download',
          icon: 'el-icon-download',
        },
        {
          label: '安全设置',
          value: '/me/profile',
          icon: 'fa fa-shield',
        },
      ],
    }
  },
  head() {
    return {
      title: `${this.defaultActive.label} - ${this.user.username} - ${this.settings.system.sitename}`,
    }
  },
  watch: {
    '$route.path': {
      handler(val) {
        // val 去掉最后的斜杠
        val = val.replace(/\/$/, '')
        const item = this.tabs.find((item) => item.value === val)
        this.defaultActive = {
          value: val,
          label: item.label || '我的动态',
        }
      },
      immediate: true,
    },
  },
}
</script>
<style lang="scss">
.page-me {
  .el-menu {
    border-right: 0;
  }
  .fa-shield {
    margin-right: 5px;
    width: 24px;
    text-align: center;
    font-size: 18px;
    vertical-align: middle;
  }
  .nuxt-child {
    min-height: calc(100vh - 190px);
  }
}
</style>
