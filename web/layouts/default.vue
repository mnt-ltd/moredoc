<template>
  <el-container class="layout-default">
    <el-header v-if="$route.path !== '/search'">
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
          <el-menu-item index="/">
            <nuxt-link to="/">首页</nuxt-link>
          </el-menu-item>
          <el-menu-item
            v-for="item in categoryTrees"
            :key="'c-' + item.id"
            :index="`/category/${item.id}`"
          >
            <nuxt-link :to="`/category/${item.id}`">{{ item.title }}</nuxt-link>
          </el-menu-item>
          <el-menu-item index="/login" class="float-right">
            <nuxt-link to="/login"><i class="el-icon-user"></i> 登录</nuxt-link>
          </el-menu-item>
        </el-menu>
      </div>
    </el-header>
    <el-main>
      <nuxt />
    </el-main>
    <el-footer>
      <div class="footer-friendlink">
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
            title="魔刀文库"
            :href="settings.system.domain"
          >
            {{ settings.system.sitename }}
            <span v-if="settings.system.copyright_start_year == currentYear"
              >©{{ currentYear }}</span
            >
            <span v-else>
              ©{{ settings.system.copyright_start_year }} - {{ currentYear }}
            </span>
          </el-link>
          <span>|</span>
          <el-link
            :underline="false"
            type="white"
            target="_blank"
            title="站点地图"
            href="sitemap.xml"
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
          <el-link
            :underline="false"
            type="primary"
            target="_blank"
            href="https://mnt.ltd/?prod=moredoc"
            title="MOREDOC"
            >Powered By
            <strong class="el-link--primary">MOREDOC</strong></el-link
          >
        </div>
      </div>
    </el-footer>
  </el-container>
</template>
<script>
import { mapGetters, mapActions } from 'vuex'
import { listFriendlink } from '~/api/friendlink'
import { categoryToTrees } from '~/utils/utils'
export default {
  data() {
    return {
      search: {
        wd: '',
      },
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
        '魔刀文库',
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
    ...mapGetters('user', ['user', 'token']),
    ...mapGetters('setting', ['settings']),
    ...mapGetters('category', ['categories']),
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
  mounted() {},
  methods: {
    ...mapActions('category', ['getCategories']),
    ...mapActions('setting', ['getSettings']),
    onSearch() {
      this.$router.push({
        path: '/search',
        query: {
          wd: this.search.wd,
        },
      })
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
  },
}
</script>
<style lang="scss">
.layout-default {
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
    }
  }
  background-color: $background-grey-light;
  .logo {
    &.is-active {
      border-color: transparent !important;
    }
    img {
      margin-top: -4px;
      height: 42px;
    }
  }
}
.page {
  width: $default-width;
  max-width: $max-width;
  margin: 0 auto;
  overflow-x: hidden;
}
</style>
