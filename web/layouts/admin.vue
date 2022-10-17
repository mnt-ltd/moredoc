<template>
  <el-container class="layout-admin">
    <el-aside
      :class="isCollapse ? 'layout-aside-collapsed' : ''"
      :width="'240px'"
    >
      <div class="quickstart-upload">
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
          :default-openeds="['/admin/documemnt', '/admin/user']"
          :default-active="$route.path"
          :collapse="isCollapse"
          class="layout-admin-menu"
        >
          <el-menu-item index="/admin/dashboard">
            <i class="el-icon-monitor"></i>
            <span slot="title">面板</span>
          </el-menu-item>

          <el-submenu index="/admin/documemnt">
            <template slot="title">
              <i class="el-icon-document-copy"></i>
              <span slot="title">文档管理</span>
            </template>
            <el-menu-item index="/admin/document/category">
              <i class="el-icon-s-grid"></i>
              <span>分类管理</span>
            </el-menu-item>
            <el-menu-item index="/admin/document/list">
              <i class="el-icon-tickets"></i>
              <span>文档列表</span>
            </el-menu-item>
            <el-menu-item index="/admin/document/recycle">
              <i class="el-icon-delete"></i>
              <span>回收站</span>
            </el-menu-item>
          </el-submenu>
          <el-submenu index="/admin/user">
            <template slot="title">
              <i class="el-icon-user"></i>
              <span slot="title">用户管理</span>
            </template>
            <el-menu-item index="/admin/user/list">
              <i class="el-icon-user"></i>
              <span>用户管理</span>
            </el-menu-item>
            <el-menu-item index="/admin/user/group">
              <i class="el-icon-magic-stick"></i>
              <span>角色管理</span>
            </el-menu-item>
            <el-menu-item index="/admin/user/permission">
              <i class="el-icon-circle-check"></i>
              <span>权限管理</span>
            </el-menu-item>
          </el-submenu>
          <el-menu-item index="/admin/banner">
            <i class="el-icon-picture-outline"></i>
            <span slot="title">横幅管理</span>
          </el-menu-item>
          <el-menu-item index="/admin/friendlink">
            <i class="el-icon-link"></i>
            <span slot="title">友链管理</span>
          </el-menu-item>
          <!-- <el-menu-item index="/admin/comment">
            <i class="el-icon-chat-dot-square"></i>
            <span slot="title">评论管理</span>
          </el-menu-item> -->
          <el-menu-item index="/admin/attachment">
            <i class="el-icon-paperclip"></i>
            <span slot="title">附件管理</span>
          </el-menu-item>
          <el-menu-item index="/admin/config">
            <i class="el-icon-setting"></i>
            <span slot="title">系统设置</span>
          </el-menu-item>
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
            <i class="el-icon-user"></i>
            <span>{{ user.username }}</span>
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
      <!-- <el-footer>
        <span>© 2019-2022</span>
        <span>Powered by</span>
        <a href="https://mnt.ltd" target="_blank" title="MOREDOC"
          >MOREDOC · 魔刀文库</a
        >
      </el-footer> -->
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
    ...mapActions('user', ['logout']),
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
          this.logout()
          this.$router.replace({ path: '/' })
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
  .quickstart-upload {
    padding: 0;
    .el-button {
      height: 59px;
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
