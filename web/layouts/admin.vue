<template>
  <el-container class="layout-admin">
    <el-aside :class="isCollapse ? 'layout-aside-collapsed' : ''">
      <div class="create-project">
        <el-tooltip
          v-if="isCollapse"
          class="item"
          effect="dark"
          content="上传文档"
          placement="right"
        >
          <el-button
            type="success"
            icon="el-icon-plus"
            class="btn-block"
            @click="showCreateProjectDialog"
          ></el-button>
        </el-tooltip>
        <el-button
          v-else
          type="success"
          icon="el-icon-plus"
          @click="showCreateProjectDialog"
          >上传文档</el-button
        >
      </div>
      <transition
        :duration="{ enter: 800, leave: 800 }"
        mode="out-in"
        name="el-fade-in-linear"
      >
        <el-menu
          :router="true"
          :default-openeds="[
            '/admin/settings',
            '/admin/me',
            '/admin/templates',
          ]"
          :collapse="isCollapse"
          class="layout-admin-menu"
        >
          <el-menu-item index="/admin/dashboard">
            <i class="el-icon-s-platform"></i>
            <span slot="title">面板</span></el-menu-item
          >
          <el-submenu index="/admin/me">
            <template slot="title"
              ><i class="el-icon-user-solid"></i>
              <span slot="title">我的</span></template
            >
            <el-menu-item index="/admin/projects">
              <i class="el-icon-data-analysis"></i>
              我的项目</el-menu-item
            >
            <!-- <el-menu-item index="/admin/shares">我的分享</el-menu-item> -->
            <el-menu-item index="/admin/mytemplates"
              ><i class="el-icon-tickets"></i> 我的模板</el-menu-item
            >
            <el-menu-item index="/admin/mycharts">
              <i class="el-icon-s-grid"></i> 我的模块</el-menu-item
            >
          </el-submenu>
          <el-submenu index="/admin/templates">
            <template slot="title"
              ><i class="el-icon-s-shop"></i>
              <span slot="title">应用中心</span></template
            >
            <el-menu-item index="/admin/templates"
              ><i class="el-icon-tickets"></i> 模板市场</el-menu-item
            >
            <el-menu-item index="/admin/charts"
              ><i class="el-icon-s-grid"></i> 图表模块</el-menu-item
            >
          </el-submenu>
          <el-submenu index="/admin/settings">
            <template slot="title"
              ><i class="el-icon-s-tools"></i>
              <span slot="title">系统管理</span></template
            >
            <el-menu-item index="/admin/settings"
              ><i class="el-icon-setting"></i> 系统设置</el-menu-item
            >
            <el-menu-item index="/admin/users"
              ><i class="el-icon-user"></i> 用户管理</el-menu-item
            >
          </el-submenu>
        </el-menu>
      </transition>
    </el-aside>
    <el-container>
      <el-header>
        <el-button
          v-if="isCollapse"
          class="fold"
          icon="el-icon-s-unfold"
          type="text"
          @click="isCollapse = false"
        ></el-button>
        <el-button
          v-else
          class="fold"
          icon="el-icon-s-fold"
          type="text"
          @click="isCollapse = true"
        ></el-button>
        <span>我的项目</span>
        <el-dropdown style="float: right" trigger="click" @command="command">
          <el-button>
            <i class="el-icon-user" style="margin-right: 15px"> </i>
            <span>{{ user.realname }}</span>
          </el-button>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item command="profile"> 个人资料 </el-dropdown-item>
            <el-dropdown-item command="password"> 修改密码 </el-dropdown-item>
            <el-dropdown-item command="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-header>
      <el-main>
        <nuxt />
      </el-main>
      <el-footer>
        <span>© 2019-2022</span>
        <span>Powered by</span>
        <a href="https://mnt.ltd" target="_blank" title="MOREDOC">MOREDOC · 魔刀文库</a>
      </el-footer>
    </el-container>
    <el-dialog title="个人资料" :visible.sync="formProfileVisible" width="30%">
      <FormProfile @success="profileSuccess" />
    </el-dialog>
    <el-dialog title="个人资料" :visible.sync="formPasswordVisible" width="30%">
      <FormPassword @success="passwordSuccess" />
    </el-dialog>
    <el-dialog
      title="创建项目"
      :visible.sync="createProjectVisible"
      width="30%"
    >
      <form-set-project @success="createProjectSuccess"></form-set-project>
    </el-dialog>
  </el-container>
</template>
<script>
import { mapGetters, mapActions } from 'vuex'
import FormProfile from '~/components/FormProfile.vue'
import FormPassword from '~/components/FormPassword.vue'
export default {
  components: {
    FormProfile,
    FormPassword,
  },
  middleware: ['auth'],
  data() {
    return {
      createProjectVisible: false,
      formProfileVisible: false,
      formPasswordVisible: false,
      isCollapse: false,
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库',
    }
  },
  computed: {
    ...mapGetters('user', ['user', 'token']),
  },
  watch: {
    // async $route(to, from) {
    //   this.activePath = to.path
    //   this.$refs.mobileMenu.close('mobileMenu')
    // },
  },
  mounted() {
    const screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isCollapse = !this.isCollapse
    }
  },
  methods: {
    ...mapActions('user', ['Logout']),
    async createProject() {},
    showCreateProjectDialog() {
      this.createProjectVisible = true
    },
    createProjectSuccess(e) {
      this.createProjectVisible = false
      this.$router.push({
        path: '/admin/projects',
        query: { _t: new Date().getTime() },
      })
    },
    profileSuccess() {
      this.formProfileVisible = false
    },
    passwordSuccess() {
      this.formPasswordVisible = false
    },
    command(cmd) {
      switch (cmd) {
        case 'profile':
          this.formProfileVisible = true
          break
        case 'password':
          this.formPasswordVisible = true
          break
        case 'logout':
          this.Logout()
          this.$router.replace({ path: '/admin/login' })
          this.$message.success('退出成功')
          break
      }
    },
  },
}
</script>
<style lang="scss">
.layout-aside-collapsed {
  width: 60px !important;
  overflow: hidden;
  .create-project {
    padding: 0;
    .el-button {
      height: 60px;
      line-height: 30px;
      padding: 0 5px;
      border-radius: 0;
      span {
        display: block;
      }
    }
  }
}
</style>
