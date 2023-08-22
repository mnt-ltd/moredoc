<template>
  <el-container class="layout-default" v-loading="loading">
    <el-header v-if="$route.name !== 'search' || isMobile">
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
          <el-submenu
            index="channel"
            class="hidden-xs-only"
            v-show="$route.path !== '/' || navigations.length > 0"
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
          <template v-if="navigations.length === 0">
            <el-menu-item
              v-for="(item, index) in categoryTrees"
              :key="'c-' + item.id"
              :index="`/category/${item.id}`"
              v-show="$route.path === '/' && index < 6"
              class="hidden-xs-only"
            >
              <nuxt-link :to="`/category/${item.id}`">{{
                item.title
              }}</nuxt-link>
            </el-menu-item>
          </template>
          <template v-else>
            <template v-for="item in navigations">
              <el-submenu
                :key="'nav-' + item.id"
                :index="`nav-${item.id}`"
                v-if="item.children && item.children.length > 0"
                class="hidden-xs-only"
              >
                <template slot="title">{{ item.title }}</template>
                <NavigationLink
                  :hiddenXS="true"
                  v-for="child in item.children || []"
                  :key="'child-' + child.id"
                  :navigation="child"
                />
              </el-submenu>
              <NavigationLink
                v-else
                :navigation="item"
                :key="'nav-' + item.id"
                :hiddenXS="true"
              />
            </template>
          </template>
          <el-menu-item
            index="searchbox"
            class="nav-searchbox hidden-xs-only"
            :class="navigations.length <= 2 ? 'nav-searchbox-large' : ''"
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
            class="float-right nav-ucenter hidden-xs-only"
          >
            <el-dropdown trigger="hover" @command="handleDropdown">
              <span class="el-dropdown-link">
                <user-avatar class="nav-user-avatar" :user="user" :size="42" />
                <span>{{ user.username }}</span
                ><i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="ucenter"
                  ><i class="fa fa-home"></i> 个人主页</el-dropdown-item
                >
                <el-dropdown-item command="me"
                  ><i class="fa fa-user-o"></i> 个人中心</el-dropdown-item
                >
                <!-- <el-dropdown-item command="profile"
                  ><i class="fa fa-edit"></i> 个人资料</el-dropdown-item
                > -->
                <el-dropdown-item command="upload"
                  ><i class="el-icon-upload2 dropdown-upload"></i
                  >上传文档</el-dropdown-item
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
          <el-menu-item
            v-else
            index="/login"
            class="float-right hidden-xs-only"
          >
            <nuxt-link to="/login"><i class="el-icon-user"></i> 登录</nuxt-link>
          </el-menu-item>
          <el-menu-item
            v-if="isMobile"
            class="menu-drawer float-right"
            index="menuDrawer"
            @click="showMenuDrawer"
          >
            <i class="el-icon-s-operation"></i>
          </el-menu-item>
        </el-menu>
      </div>
    </el-header>
    <el-main>
      <nuxt />
    </el-main>
    <el-footer>
      <div
        v-if="$route.path === '/' && friendlinks.length > 0"
        class="footer-friendlink"
      >
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
        <div v-if="settings.system.icp || settings.system.sec_icp">
          <el-link
            :underline="false"
            v-if="settings.system.icp"
            type="white"
            target="_blank"
            :title="settings.system.icp"
            href="https://beian.miit.gov.cn/"
            >{{ settings.system.icp }}</el-link
          >
          <el-link
            :underline="false"
            v-if="settings.system.sec_icp"
            type="white"
            target="_blank"
            :title="settings.system.sec_icp"
            :href="`http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=${settings.system.sec_icp.replace(
              /[^\d]/g,
              ''
            )}`"
            >{{ settings.system.sec_icp }}</el-link
          >
        </div>
        <div v-if="settings.display.copyright_statement">
          <div
            class="el-link el-link--default"
            v-html="settings.display.copyright_statement"
          ></div>
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
          CE
          <span>{{ settings.system.version }}</span>
        </div>
      </div>
    </el-footer>
    <el-drawer
      :visible.sync="menuDrawerVisible"
      size="60%"
      :with-header="false"
      class="menu-drawer-box"
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
      <ul class="navs">
        <li>
          <div
            @click="goToLink('/login')"
            class="el-link el-link--default login-link"
          >
            <user-avatar :size="38" :user="user" class="user-avatar" />
            <span v-if="user.id > 0">{{ user.username }}</span>
            <span v-else>登录注册</span>
          </div>
        </li>
        <template v-if="user.id > 0">
          <li cass="mgt-20px">
            <el-button
              v-if="sign.id > 0"
              :key="'sign-' + sign.id"
              class="btn-block"
              type="success"
              size="medium"
              disabled
            >
              <i class="fa fa-calendar-check-o" aria-hidden="true"></i>
              今日已签到
            </el-button>
            <el-button
              v-else
              :key="'sign-0'"
              class="btn-block"
              type="success"
              size="medium"
              @click="signToday"
            >
              <i class="fa fa-calendar-plus-o"></i>
              每日签到</el-button
            >
          </li>
          <li>
            <div
              @click="goToLink(`/user/${user.id}`)"
              class="el-link el-link--default"
            >
              <i class="fa fa-home"></i> &nbsp;个人主页
            </div>
          </li>
          <li>
            <div class="el-link el-link--default" @click="goToLink(`/me`)">
              <i class="fa fa-user-o"></i> &nbsp;个人中心
            </div>
          </li>
          <li>
            <div @click="goToLink(`/upload`)" class="el-link el-link--default">
              <i class="el-icon-upload2"></i> 上传文档
            </div>
          </li>
          <li>
            <nuxt-link to="/admin" class="el-link el-link--default"
              ><i class="el-icon-box"></i> &nbsp;管理后台</nuxt-link
            >
          </li>
          <li>
            <div class="el-link el-link--default" @click="logout">
              <i class="fa fa-sign-out"></i> &nbsp;退出登录
            </div>
          </li></template
        >
      </ul>
      <el-collapse v-model="activeCollapse">
        <el-collapse-item name="categories">
          <template slot="title"
            ><i class="el-icon-menu"></i> &nbsp; <span>频道分类</span>
          </template>
          <ul>
            <li
              v-for="item in categoryTrees"
              :key="'collapse-sub-cate-' + item.id"
            >
              <div
                class="el-link el-link--default"
                @click="goToLink(`/category/${item.id}`)"
              >
                {{ item.title }}
              </div>
            </li>
          </ul>
        </el-collapse-item>
      </el-collapse>
      <el-menu :default-active="$route.path" class="el-menu-mobile">
        <template v-for="item in navigations">
          <el-submenu
            :key="'nav-' + item.id"
            :index="`nav-${item.id}`"
            v-if="item.children && item.children.length > 0"
          >
            <template slot="title">{{ item.title }}</template>
            <NavigationLink
              v-for="child in item.children || []"
              :key="'child-' + child.id"
              :navigation="child"
            />
          </el-submenu>
          <NavigationLink v-else :navigation="item" :key="'nav-' + item.id" />
        </template>
      </el-menu>
    </el-drawer>
  </el-container>
</template>
<script>
import { mapGetters, mapActions } from 'vuex'
import UserAvatar from '~/components/UserAvatar.vue'
import FormUserinfo from '~/components/FormUserinfo.vue'
import { listFriendlink } from '~/api/friendlink'
import { listNavigation } from '~/api/navigation'
import { categoryToTrees, requireLogin } from '~/utils/utils'
import { getSignedToday, signToday } from '~/api/user'

export default {
  components: { UserAvatar, FormUserinfo },
  middleware: ['checkFront', 'analytic'],
  data() {
    return {
      search: {
        wd: '',
      },
      friendlinks: [],
      timeouter: null,
      currentYear: new Date().getFullYear(),
      categoryTrees: [],
      menuDrawerVisible: false,
      sign: { id: 0 },
      activeCollapse: 'categories',
      loading: false,
      navigations: [],
    }
  },
  head() {
    return {
      title:
        this.settings.system.title || this.settings.system.sitename || '文库',
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
    ...mapGetters('user', ['user', 'token', 'allowPages', 'permissions']),
    ...mapGetters('setting', ['settings']),
    ...mapGetters('category', ['categories']),
  },
  async created() {
    this.loading = true
    await Promise.all([
      this.getCategories(),
      this.getSettings(),
      this.listNavigation(),
      this.listFriendlink(),
    ])

    this.categoryTrees = categoryToTrees(this.categories).filter(
      (item) => item.enable
    )

    this.loopUpdate()
    this.loading = false
    if (requireLogin(this.settings, this.user, this.$route, this.permissions)) {
      this.$router.push('/login')
      return
    }
  },
  methods: {
    ...mapActions('category', ['getCategories']),
    ...mapActions('setting', ['getSettings']),
    ...mapActions('user', ['logout', 'getUser', 'checkAndRefreshUser']),
    async showMenuDrawer() {
      this.getSignedToday()
      this.menuDrawerVisible = true
    },
    goToLink(link) {
      this.menuDrawerVisible = false
      this.$router.push(link)
    },
    async getSignedToday() {
      const res = await getSignedToday()
      if (res.status === 200) {
        this.sign = res.data || this.sign
      }
    },
    async listFriendlink() {
      const res = await listFriendlink({
        enable: true,
        field: ['id', 'title', 'link'],
      })
      if (res.status === 200) {
        this.friendlinks = res.data.friendlink || []
      }
    },
    async signToday() {
      const res = await signToday()
      if (res.status === 200) {
        const sign = res.data || { id: 1 }
        this.sign = sign
        this.getUser()
        this.$message.success(
          `签到成功，获得 ${sign.award || 0} ${
            this.settings.system.credit_name || '魔豆'
          }奖励`
        )
      } else {
        this.$message.error(res.message || res.data.message)
      }
    },
    async listNavigation() {
      const res = await listNavigation({ page: 1, size: 10000 })
      if (res.status === 200) {
        let navigations = res.data.navigation || []
        navigations = categoryToTrees(navigations).filter((item) => item.enable)
        this.navigations = navigations
      }
    },
    onSearch() {
      if (!this.search.wd) return
      this.menuDrawerVisible = false
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
        this.checkAndRefreshUser()
        // 更新系统配置信息
        this.getSettings()
        // 更新分类信息
        this.getCategories()
        // 递归
        this.loopUpdate()
      }, 1000 * 60) // 每分钟更新一次
    },
    async handleDropdown(command) {
      console.log('handleDropdown', command)
      switch (command) {
        case 'logout':
          const res = await this.logout()
          console.log(res)
          location.reload()
          break
        case 'upload':
          this.$router.push('/upload')
          break
        case 'ucenter':
          this.$router.push(`/user/${this.user.id}`)
          break
        case 'me':
          this.$router.push(`/me`)
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
  .el-breadcrumb__inner a,
  .el-breadcrumb__inner.is-link {
    font-weight: normal;
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
    .logo {
      &.is-active {
        border-color: transparent !important;
      }
      img {
        margin-top: -4px;
        height: 42px;
      }
    }
    & > div {
      margin: 0 auto;
      width: $default-width;
      max-width: $max-width;
    }
    .el-menu--horizontal > .el-submenu .el-submenu__title {
      padding-top: 1px;
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
    .menu-drawer {
      display: none;
    }
    .nav-searchbox {
      padding: 0 25px !important;
      top: -2px;
      &.nav-searchbox-large {
        .el-input {
          width: 300px;
        }
      }
      &.is-active {
        border-color: transparent;
      }
      .el-input {
        width: 200px;
      }
    }
    a {
      text-decoration: none;
      height: 60px;
      line-height: 60px;
      display: inline-block;
      // padding: 0 20px;
      padding: 0 15px;
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

  .nav-ucenter {
    &.is-active {
      border-color: transparent !important;
    }
    .el-dropdown-link {
      line-height: 60px;
      display: inline-block;
      font-weight: 400;
      font-size: 1.2em;
      margin-top: -8px;
      .nav-user-avatar {
        position: relative;
        top: -2px;
      }
    }
  }
}
.dropdown-upload {
  font-size: 17px;
  margin-left: -2px;
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
  .el-menu-item {
    a {
      display: block;
    }
  }
}

.el-dialog__header {
  padding: 20px 20px 10px;
}
.el-dialog__body {
  padding: 1px 20px;
}

.el-menu-mobile {
  margin-top: 15px;
  margin-left: -15px;
  border-right: 0;
  & > li {
    padding-left: 0;
  }
  .el-submenu__title {
    height: 32px;
    line-height: 32px;
    font-size: 15px;
  }
  .el-menu-item {
    height: 32px;
    line-height: 32px;
  }
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
    .menu-drawer-box {
      .el-drawer__body {
        padding: 15px;
      }
      .navs {
        padding: 0;
        list-style: none;
        li {
          line-height: 30px;
          .login-link {
            max-width: 100%;
            .user-avatar {
              margin-right: 5px;
              top: 4px;
            }
            span {
              display: inline-block;
              position: relative;
              top: 0;
              white-space: nowrap;
              overflow: hidden;
              text-overflow: ellipsis;
            }
          }
          .el-link {
            font-size: 14px;
            display: block;
          }
        }
        & > li {
          margin: 5px 0;
        }
        .el-icon-upload2 {
          font-size: 17px;
        }
      }
      .el-collapse {
        .el-icon-menu {
          font-size: 16px;
        }
        .el-collapse-item__content {
          padding-bottom: 0;
          margin-top: -15px;
        }
        ul {
          padding-left: 30px;
          li {
            line-height: 30px;
          }
        }
      }
    }
    .el-header {
      .el-menu.el-menu--horizontal {
        border-bottom: 0;
        width: 100%;
        max-width: unset;
        min-width: unset;
        .float-right {
          float: right;
          a {
            padding: 0 15px;
          }
        }
      }
      a {
        padding: 0 15px;
      }
      .menu-drawer {
        padding: 0 15px;
        display: inline-block;
        &.is-active {
          border-bottom: 0;
        }
        .el-icon-s-operation {
          font-size: 22px;
        }
      }
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
