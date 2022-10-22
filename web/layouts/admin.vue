<template>
  <el-container class="layout-admin">
    <el-aside
      :class="isCollapse ? 'layout-aside-collapsed' : ''"
      :width="'240px'"
    >
      <div class="logo" @click="gohome">
        <img src="/static/images/default-logo-icon.png" alt="" />
        <span>魔刀文库系统</span>
      </div>
      <transition
        :duration="{ enter: 800, leave: 800 }"
        mode="out-in"
        name="el-fade-in-linear"
      >
        <el-menu
          :router="true"
          :default-active="$route.path"
          :collapse="isCollapse"
          class="layout-admin-menu"
        >
          <el-menu-item index="/admin/dashboard">
            <i class="el-icon-monitor"></i>
            <span slot="title">面板</span>
          </el-menu-item>
          <template v-for="menu in menus">
            <el-submenu
              v-if="menu.children"
              v-show="allowPages.includes(menu.page)"
              :key="'submenu-' + menu.page"
              :index="menu.page"
            >
              <template slot="title">
                <i :class="menu.icon"></i>
                <span slot="title">{{ menu.title }}</span>
              </template>
              <el-menu-item
                v-for="child in menu.children"
                v-show="allowPages.includes(child.page)"
                :key="child.page"
                :index="child.page"
              >
                <i :class="child.icon"></i>
                <span>{{ child.title }}</span>
              </el-menu-item>
            </el-submenu>
            <el-menu-item
              v-else
              v-show="allowPages.includes(menu.page)"
              :key="'menu-' + menu.page"
              :index="menu.page"
            >
              <i :class="menu.icon"></i>
              <span slot="title">{{ menu.title }}</span>
            </el-menu-item>
          </template>
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
    </el-container>
    <el-dialog title="个人资料" :visible.sync="formProfileVisible" width="30%">
      <FormProfile @success="profileSuccess" />
    </el-dialog>
    <el-dialog title="个人资料" :visible.sync="formPasswordVisible" width="30%">
      <FormPassword @success="passwordSuccess" />
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
      formProfileVisible: false,
      formPasswordVisible: false,
      isCollapse: false,
      menus: [
        {
          page: '/admin/document',
          title: '文档管理',
          icon: 'el-icon-document-copy',
          children: [
            {
              page: '/admin/document/category',
              title: '分类管理',
              icon: 'el-icon-s-grid',
            },
            {
              page: '/admin/document/list',
              title: '文档列表',
              icon: 'el-icon-tickets',
            },
            {
              page: '/admin/document/recycle',
              title: '回收站',
              icon: 'el-icon-delete',
            },
          ],
        },
        {
          page: '/admin/user',
          title: '用户管理',
          icon: 'el-icon-user',
          children: [
            {
              page: '/admin/user/list',
              title: '用户管理',
              icon: 'el-icon-user',
            },
            {
              page: '/admin/user/group',
              title: '角色管理',
              icon: 'el-icon-magic-stick',
            },
            {
              page: '/admin/user/permission',
              title: '权限管理',
              icon: 'el-icon-circle-check',
            },
          ],
        },
        {
          page: '/admin/banner',
          title: '横幅管理',
          icon: 'el-icon-picture-outline',
        },
        {
          page: '/admin/friendlink',
          title: '友链管理',
          icon: 'el-icon-link',
        },
        {
          page: '/admin/comment',
          title: '评论管理',
          icon: 'el-icon-chat-dot-square',
        },
        {
          page: '/admin/attachment',
          title: '附件管理',
          icon: 'el-icon-paperclip',
        },
        {
          page: '/admin/config',
          title: '系统设置',
          icon: 'el-icon-setting',
        },
      ],
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库',
    }
  },
  computed: {
    ...mapGetters('user', ['user', 'token', 'permissions', 'allowPages']),
  },
  created() {
    this.getUserPermissions()
  },
  mounted() {
    const screenWidth = document.body.clientWidth
    if (screenWidth < 1000) {
      this.isCollapse = !this.isCollapse
    }
  },
  methods: {
    ...mapActions('user', ['logout', 'getUserPermissions']),
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
    gohome() {
      this.$router.push('/')
    },
  },
}
</script>
<style lang="scss">
.layout-admin {
  .logo {
    cursor: pointer;
  }
}
.layout-aside-collapsed {
  width: 64px !important;
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
