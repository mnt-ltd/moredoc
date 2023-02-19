<template>
  <el-container class="layout-default">
    <el-header v-if="$route.name !== 'search'">
      <div>
        <el-menu
          :default-active="$route.path"
          class="float-left"
          mode="horizontal"
        >
          <el-menu-item class="logo" index="/">
            <nuxt-link to="/"
              ><img
                :src="settings.system.logo || '/static/images/logo.png'"
                :alt="settings.system.sitename"
            /></nuxt-link>
          </el-menu-item>
          <el-menu-item index="/" class="hidden-xs-only">
            <nuxt-link to="/">首页</nuxt-link>
          </el-menu-item>
          <el-menu-item
            v-for="(item, index) in categoryTrees"
            :key="'c-' + item.id"
            :index="`/category/${item.id}`"
            v-show="$route.path === '/' && index < 6"
            class="hidden-xs-only"
          >
            <nuxt-link :to="`/category/${item.id}`">{{ item.title }}</nuxt-link>
          </el-menu-item>
          <el-submenu
            index="channel"
            class="hidden-xs-only"
            v-show="$route.path !== '/'"
          >
            <template slot="title">频道分类</template>
            <el-menu-item
              v-for="item in categoryTrees"
              :key="'sub-cate-' + item.id"
              class="channel-category"
              :index="`/category/${item.id}`"
            >
              <nuxt-link
                class="el-link el-link--default"
                :to="`/category/${item.id}`"
                >{{ item.title }}</nuxt-link
              >
            </el-menu-item>
          </el-submenu>
          <el-menu-item
            index="searchbox"
            class="nav-searchbox hidden-xs-only"
            v-show="$route.path !== '/'"
          >
            <el-input
              v-model="search.wd"
              class="search-input"
              size="large"
              placeholder="搜索文档..."
              @keyup.enter.native="onSearch"
            >
              <i
                class="el-icon-search el-input__icon"
                @click="onSearch"
                slot="suffix"
              >
              </i>
            </el-input>
          </el-menu-item>
          <el-menu-item
            v-if="user.id > 0"
            index="ucenter"
            class="float-right nav-ucenter"
          >
            <el-dropdown trigger="hover" @command="handleDropdown">
              <span class="el-dropdown-link">
                <user-avatar class="nav-user-avatar" :user="user" :size="42" />
                <span>{{ user.username }}</span
                ><i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="ucenter"
                  ><i class="fa fa-user-o"></i> 个人主页</el-dropdown-item
                >
                <el-dropdown-item command="profile"
                  ><i class="fa fa-edit"></i> 个人资料</el-dropdown-item
                >
                <el-dropdown-item command="upload"
                  ><i class="fa fa-cloud-upload"></i>上传文档</el-dropdown-item
                >
                <el-dropdown-item v-if="allowPages.length > 0" command="admin">
                  <i class="el-icon-box"></i> 管理后台</el-dropdown-item
                >
                <el-dropdown-item command="logout"
                  ><i class="fa fa-sign-out"></i> 退出登录</el-dropdown-item
                >
              </el-dropdown-menu>
            </el-dropdown>
          </el-menu-item>
          <el-menu-item v-else index="/login" class="float-right">
            <nuxt-link to="/login"><i class="el-icon-user"></i> 登录</nuxt-link>
          </el-menu-item>
        </el-menu>
      </div>
    </el-header>
    <el-main>
      <nuxt />
    </el-main>
    <el-footer>
      <div v-if="$route.path === '/'" class="footer-friendlink">
        <el-card class="box-card" shadow="never">
          <div slot="header" class="clearfix">
            <strong>友情链接</strong>
          </div>
          <a
            v-for="link in friendlinks"
            :key="'fl-' + link.id"
            :underline="false"
            :href="link.link"
            class="el-link el-link--default"
            target="_blank"
            >{{ link.title }}</a
          >
        </el-card>
      </div>
      <div class="footer-links">
        <div>
          <el-link
            v-if="settings.footer.about"
            :underline="false"
            type="white"
            target="_blank"
            :href="settings.footer.about"
            >关于我们</el-link
          >
          <el-link
            v-if="settings.footer.agreement"
            :href="settings.footer.agreement"
            :underline="false"
            target="_blank"
            type="white"
            >文库协议</el-link
          >
          <el-link
            v-if="settings.footer.contact"
            :underline="false"
            type="white"
            target="_blank"
            :href="settings.footer.contact"
            >联系我们</el-link
          >
          <el-link
            v-if="settings.footer.feedback"
            :underline="false"
            type="white"
            :href="settings.footer.feedback"
            target="_blank"
            >意见反馈</el-link
          >
          <el-link
            v-if="settings.footer.copyright"
            :underline="false"
            type="white"
            :href="settings.footer.copyright"
            target="_blank"
            >免责声明</el-link
          >
        </div>
        <div>
          <el-link
            v-if="settings.system.domain"
            :underline="false"
            type="white"
            :title="settings.system.sitename || ''"
            :href="settings.system.domain"
          >
            {{ settings.system.sitename }}
          </el-link>
          <span class="copyright-year"
            ><span v-if="settings.system.copyright_start_year == currentYear"
              >©{{ currentYear }}</span
            >
            <span v-else>
              ©{{ settings.system.copyright_start_year }} - {{ currentYear }}
            </span></span
          >
          <span>|</span>
          <el-link
            :underline="false"
            type="white"
            target="_blank"
            title="站点地图"
            href="/sitemap.xml"
            >站点地图</el-link
          >
        </div>
        <div v-if="settings.system.icp">
          <el-link
            :underline="false"
            type="white"
            target="_blank"
            :title="settings.system.icp"
            href="https://beian.miit.gov.cn/"
            >{{ settings.system.icp }}</el-link
          >
        </div>
        <div>
          Powered By
          <el-link
            :underline="false"
            type="primary"
            target="_blank"
            href="https://mnt.ltd/#services"
            title="MOREDOC"
            class="powered-by"
            >MOREDOC</el-link
          >
          <span>{{ settings.system.version }}</span>
        </div>
      </div>
    </el-footer>
    <el-dialog
      title="个人资料"
      :visible.sync="userinfoDialogVisible"
      :width="isMobile ? '95%' : '520px'"
    >
      <form-userinfo v-if="userinfoDialogVisible" />
    </el-dialog>
  </el-container>
</template>
<script>
import { mapGetters, mapActions } from 'vuex'
import UserAvatar from '~/components/UserAvatar.vue'
import FormUserinfo from '~/components/FormUserinfo.vue'
import { listFriendlink } from '~/api/friendlink'
import { categoryToTrees } from '~/utils/utils'
export default {
  components: { UserAvatar, FormUserinfo },
  middleware: ['analytic'],
  data() {
    return {
      search: {
        wd: '',
      },
      userinfoDialogVisible: false,
      friendlinks: [],
      timeouter: null,
      currentYear: new Date().getFullYear(),
      categoryTrees: [],
    }
  },
  head() {
    return {
      title:
        this.settings.system.title ||
        this.settings.system.sitename ||
        '魔豆文库',
      keywords: this.settings.system.keywords,
      description: this.settings.system.description,
      // favicon
      link: [
        {
          rel: 'icon',
          type: 'image/x-icon',
          href: this.settings.system.favicon,
        },
      ],
    }
  },
  computed: {
    ...mapGetters('user', ['user', 'token', 'allowPages']),
    ...mapGetters('setting', ['settings']),
    ...mapGetters('category', ['categories']),
    ...mapGetters('device', ['isMobile']),
  },
  async created() {
    const [res] = await Promise.all([
      listFriendlink({
        enable: true,
        field: ['id', 'title', 'link'],
      }),
      this.getCategories(),
      this.getSettings(),
    ])
    if (res.status === 200) {
      this.friendlinks = res.data.friendlink
    }
    this.categoryTrees = categoryToTrees(this.categories).filter(
      (item) => item.enable
    )
    this.loopUpdate()
  },
  mounted() {
    window.addEventListener('resize', this.handleResize)
  },
  methods: {
    ...mapActions('category', ['getCategories']),
    ...mapActions('setting', ['getSettings']),
    ...mapActions('user', ['logout']),
    ...mapActions('device', ['setDeviceWidth']),
    handleResize() {
      console.log('handleResize', window.innerWidth)
      this.setDeviceWidth(window.innerWidth)
    },
    onSearch() {
      if (!this.search.wd) return
      let wd = this.search.wd
      this.$router.push({
        path: '/search',
        query: {
          wd: wd,
        },
      })
      this.search.wd = ''
    },
    loopUpdate() {
      clearTimeout(this.timeouter)
      this.timeouter = setTimeout(() => {
        // 更新系统配置信息
        this.getSettings()
        // 更新分类信息
        this.getCategories()
        // 递归
        this.loopUpdate()
      }, 1000 * 60) // 每分钟更新一次
    },
    handleDropdown(command) {
      console.log('handleDropdown', command)
      switch (command) {
        case 'logout':
          this.logout()
          break
        case 'upload':
          this.$router.push('/upload')
          break
        case 'ucenter':
          this.$router.push(`/user/${this.user.id}`)
          break
        case 'profile':
          this.userinfoDialogVisible = true
          break
        case 'admin':
          this.$router.push('/admin')
          break
        default:
          break
      }
    },
  },
}
</script>
<style lang="scss">
.layout-default {
  background-color: $background-grey-light;
  min-width: $min-width !important;
  .el-table th {
    height: 45px;
    line-height: 45px;
    padding: 1px 0 5px;
    &.el-table__cell {
      background-color: #f7fbff;
      color: #000;
      font-weight: normal;
      &.el-table-column--selection > .cell {
        padding-left: 14px;
      }
    }
    &.el-table__cell.is-leaf {
      border-bottom: 0;
    }
  }
  .el-main {
    padding-left: 0;
    padding-right: 0;
  }
  .nav-search-form {
    margin-top: 10px;
    .el-form-item {
      margin-right: 20px;
    }
  }
  .el-link {
    font-size: 15px;
  }
  .el-rate {
    display: inline-block;
    .el-rate__icon {
      margin-right: 0;
    }
    .el-rate__text {
      margin-left: 5px;
      color: #bdc3c7 !important;
    }
  }
  padding-top: 60px;
  .el-card {
    border-radius: 5px;
    border: 0;
  }
  .el-header {
    padding: 0;
    background: #fff;
    position: fixed;
    width: 100%;
    top: 0;
    z-index: 100;
    overflow: hidden;
    border-bottom: 1px solid $background-grey-light;
    & > div {
      margin: 0 auto;
      width: $default-width;
      max-width: $max-width;
    }
    .el-menu.el-menu--horizontal {
      border-bottom: 0;
      width: $default-width;
      max-width: $max-width;
      min-width: $min-width;
      .float-right {
        float: right;
        a {
          padding: 0 15px;
        }
      }
    }
    .nav-searchbox {
      padding: 0 25px !important;
      top: -2px;
      &.is-active {
        border-color: transparent;
      }
      .el-input {
        width: 360px;
      }
    }
    a {
      text-decoration: none;
      height: 60px;
      line-height: 60px;
      display: inline-block;
      padding: 0 20px;
    }
    .el-menu-item {
      padding: 0;
      [class^='el-icon-'] {
        font-size: 15px;
        margin-right: 2px;
      }
    }
  }
  .el-footer {
    padding: 0;
    background-color: #fff;
    .footer-friendlink {
      margin: 0 auto;
      padding: 10px 0;
      width: $default-width;
      max-width: $max-width;
      .el-link {
        margin-right: 10px;
      }
    }
    .footer-links {
      background-color: #666;
      color: #fff;
      padding: 40px 0;
      & > div {
        margin: 0 auto;
        width: $default-width;
        max-width: $max-width;
        text-align: center;
      }
      .el-link {
        color: #fff;
        margin: 10px 5px;
        .el-link--primary {
          color: #409eff;
        }
      }
      .copyright-year {
        font-size: 15px;
        position: relative;
        top: 1px;
        margin-left: -5px;
        margin-right: 5px;
      }
      .powered-by {
        font-size: 15px;
        position: relative;
        top: -2px;
        color: #409eff !important;
        margin-left: 0;
        margin-right: 0;
      }
    }
  }
  .logo {
    &.is-active {
      border-color: transparent !important;
    }
    img {
      margin-top: -4px;
      height: 42px;
    }
  }
  .nav-ucenter {
    &.is-active {
      border-color: transparent !important;
    }
    .el-dropdown-link {
      line-height: 60px;
      display: inline-block;
      font-weight: 400;
      font-size: 1.2em;
      margin-top: -2px;
      .nav-user-avatar {
        position: relative;
        top: -3px;
      }
    }
  }
}
.page {
  width: $default-width;
  min-width: $min-width !important;
  // max-width: $max-width;
  margin: 0 auto;
  overflow-x: hidden;
}
.btn-search {
  cursor: pointer;
}
.channel-category {
  &.is-active {
    background-color: #f2f6fc !important;
    a {
      color: #409eff;
    }
  }
  a {
    display: block;
  }
}
.el-menu--popup {
  min-width: 115px;
}

.el-dialog__header {
  padding: 20px 20px 10px;
}
.el-dialog__body {
  padding: 1px 20px;
}

// =======================
// 移动端样式
// =======================
@media screen and (max-width: $mobile-width) {
  .layout-default {
    min-width: 0 !important;
    .el-card {
      border-radius: 0;
    }
    .el-footer {
      .footer-links {
        font-size: 13px !important;
        padding: 30px 0;
        .el-link {
          font-size: 13px;
        }
        .copyright-year {
          font-size: 13px;
        }
      }
    }
  }
  .page {
    width: 100%;
    min-width: auto !important;
  }
  .el-card__header,
  .el-card__body {
    padding: 15px;
  }
  .el-dialog__header {
    padding: 15px 15px 5px;
  }
  .el-dialog__body {
    padding: 1px 15px;
  }
}
</style>
