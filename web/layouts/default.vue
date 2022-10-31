<template>
  <el-container class="layout-default">
    <el-header v-if="$route.path !== '/search'">
      <div>
        <el-menu
          :default-active="$route.path"
          class="float-left"
          mode="horizontal"
          :router="true"
        >
          <el-menu-item class="logo">
            <img src="/static/images/logo.png" alt="魔刀文库" />
          </el-menu-item>
          <el-menu-item index="/">
            <template slot="title">
              <!-- <i class="el-icon-s-home"></i> -->
              <span>首页</span>
            </template>
          </el-menu-item>
          <el-menu-item index="/category"> 分类 </el-menu-item>
          <!-- <el-submenu index="/category">
            <template slot="title">分类</template>
            <el-menu-item index="/category/1">选项1</el-menu-item>
            <el-menu-item index="/category/2">选项2</el-menu-item>
            <el-menu-item index="/category/3">选项3</el-menu-item>
            <el-submenu index="/category/4">
              <template slot="title">选项4</template>
              <el-menu-item index="/category/5">选项1</el-menu-item>
              <el-menu-item index="/category/6">选项2</el-menu-item>
              <el-menu-item index="/category/7">选项3</el-menu-item>
            </el-submenu>
          </el-submenu> -->
          <el-menu-item index="/upload">上传</el-menu-item>
          <el-menu-item index="/user">我的</el-menu-item>
        </el-menu>
        <el-form
          :inline="true"
          :model="search"
          class="float-right nav-search-form"
        >
          <el-form-item>
            <el-input
              v-model="search.wd"
              placeholder="Search..."
              suffix-icon="el-icon-search"
            ></el-input>
          </el-form-item>
        </el-form>
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
          <el-link :underline="false" type="white" href="/article/about"
            >关于我们</el-link
          >
          <el-link :underline="false" type="white" href="/article/about"
            >文库协议</el-link
          >
          <el-link :underline="false" type="white" href="/article/about"
            >联系我们</el-link
          >
          <el-link :underline="false" type="white" href="/article/about"
            >意见反馈</el-link
          >
          <el-link :underline="false" type="white" href="/article/about"
            >免责声明</el-link
          >
        </div>
        <div>
          <el-link
            :underline="false"
            type="white"
            title="魔刀文库"
            href="/article/about"
          >
            魔刀文库 ©2022
          </el-link>
          <span>|</span>
          <el-link :underline="false" type="white" href="/article/about"
            >站点地图</el-link
          >
        </div>
        <div>
          <el-link :underline="false" type="white" href="/article/about"
            >粤ICP备18004373号-4</el-link
          >
        </div>
        <div>
          <el-link :underline="false" type="primary" href="/article/about"
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
export default {
  data() {
    return {
      search: {
        wd: '',
      },
      friendlinks: [],
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
  async created() {
    const [res] = await Promise.all([
      listFriendlink({
        enable: true,
        field: ['id', 'title', 'link'],
      }),
      this.getCategories(),
    ])
    if (res.status === 200) {
      this.friendlinks = res.data.friendlink
    }
  },
  mounted() {},
  methods: {
    ...mapActions('category', ['getCategories']),
  },
}
</script>
<style lang="scss">
.layout-default {
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
  // font-family: Lato, Helvetica, Arial, sans-serif;
  // font-family: 'Source Sans Pro', -apple-system, BlinkMacSystemFont, 'Segoe UI',
  //   Roboto, 'Helvetica Neue', Arial, sans-serif;
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
  a {
    color: #303133;
  }
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
    & > div {
      margin: 0 auto;
      width: $default-width;
      max-width: $max-width;
    }
    .el-menu.el-menu--horizontal {
      border-bottom: 0;
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
    padding-left: 0;
    img {
      margin-top: -5px;
      height: 50px;
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
