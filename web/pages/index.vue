<template>
  <div class="page page-index">
    <el-row :gutter="20">
      <el-col :span="6" :xs="24" class="float-right">
        <el-card class="text-center stat-info" shadow="never">
          <el-row>
            <el-col :span="12">
              <small>收录文档</small>
              <div>
                <span class="el-link el-link--primary">3235</span>
              </div>
            </el-col>
            <el-col :span="12">
              <small>注册用户</small>
              <div><span class="el-link el-link--primary">872</span></div>
            </el-col>
          </el-row>
        </el-card>
        <el-card
          class="box-card mgt-20px hidden-xs-only login-form"
          shadow="never"
        >
          <el-form ref="form" :model="user">
            <el-form-item>
              <el-input
                v-model="user.username"
                placeholder="请输入登录用户名"
              ></el-input>
            </el-form-item>
            <el-form-item>
              <el-input
                v-model="user.password"
                placeholder="请输入登录密码"
              ></el-input>
            </el-form-item>
            <el-form-item style="margin-bottom: 0">
              <el-button class="btn-block" type="primary">立即登录</el-button>
            </el-form-item>
            <el-form-item style="margin-bottom: 5px">
              <nuxt-link to="/" class="el-link el-link--info"
                ><small>忘记密码？</small></nuxt-link
              >
              <nuxt-link to="/" class="float-right el-link el-link--info"
                ><small>注册会员</small></nuxt-link
              >
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="18" :xs="24">
        <el-carousel :interval="5000" arrow="always" :height="'323px'">
          <el-carousel-item v-for="item in 4" :key="item">
            <h3>{{ item }}</h3>
          </el-carousel-item>
        </el-carousel>
      </el-col>
    </el-row>
    <div class="categories mgt-20px">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-card class="box-card" shadow="never">
            <div slot="header" class="clearfix">
              <strong>编程开发</strong>
            </div>
            <nuxt-link to="/">前端开发</nuxt-link>
            <nuxt-link to="/">后端开发</nuxt-link>
            <nuxt-link to="/">移动开发</nuxt-link>
            <nuxt-link to="/">游戏开发</nuxt-link>
            <nuxt-link to="/">硬件开发</nuxt-link>
            <nuxt-link to="/">开发工具</nuxt-link>
          </el-card>
        </el-col>

        <el-col :span="6">
          <el-card class="box-card" shadow="never">
            <div slot="header" class="clearfix">
              <strong>设计·创作</strong>
            </div>
            <nuxt-link to="/">云计算</nuxt-link>
            <nuxt-link to="/">大数据</nuxt-link>
            <nuxt-link to="/">数据库</nuxt-link>
            <nuxt-link to="/">云平台</nuxt-link>
            <nuxt-link to="/">网络/安全</nuxt-link>
            <nuxt-link to="/">人工智能</nuxt-link>
          </el-card>
        </el-col>

        <el-col :span="6">
          <el-card class="box-card" shadow="never">
            <div slot="header" class="clearfix">
              <strong>系统运维</strong>
            </div>
            <nuxt-link to="/">架构</nuxt-link>
            <nuxt-link to="/">服务器</nuxt-link>
            <nuxt-link to="/">操作系统</nuxt-link>
            <nuxt-link to="/">游戏开发</nuxt-link>
            <nuxt-link to="/">网络/安全</nuxt-link>
            <nuxt-link to="/">考试认证</nuxt-link>
          </el-card>
        </el-col>

        <el-col :span="6">
          <el-card class="box-card" shadow="never">
            <div slot="header" class="clearfix">
              <strong>云计算·大数据</strong>
            </div>
            <nuxt-link to="/">云计算</nuxt-link>
            <nuxt-link to="/">大数据</nuxt-link>
            <nuxt-link to="/">数据库</nuxt-link>
            <nuxt-link to="/">云平台</nuxt-link>
            <nuxt-link to="/">网络/安全</nuxt-link>
            <nuxt-link to="/">人工智能</nuxt-link>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <el-row :gutter="20" class="category-item">
      <el-col v-for="i in 6" :key="'card-' + i" :md="12" :sm="12" :xs="24">
        <el-card class="box-card mgt-20px" shadow="never">
          <div slot="header" class="clearfix">
            <strong>卡片名称</strong>
            <el-button style="float: right; padding: 3px 0" type="text"
              >更多</el-button
            >
          </div>
          <div>
            <div class="card-body-left hidden-xs-only">
              <img src="/static/images/cover-news.png" alt="" />
            </div>
            <div class="card-body-right">
              <nuxt-link to="/document/"
                >Docker — 从入门到实战-BookStack.CN</nuxt-link
              >
              <nuxt-link to="/document/">MongoDB简明教程</nuxt-link>
              <nuxt-link to="/document/">TypeScript 官方文档</nuxt-link>
              <nuxt-link to="/document/"
                >DolphinPHP1.3.0完全开发手册-基于ThinkPHP5.0.20的快速开发框架-05221135</nuxt-link
              >
              <nuxt-link to="/document/">
                ThinkPHP5.1完全开发手册-09081747
              </nuxt-link>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
  data() {
    return {
      user: {
        username: '',
        password: '',
      },
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  async created() {},
  methods: {},
}
</script>
<style lang="scss">
.page-index {
  width: 100%;
  max-width: 100%;
  & > .el-row {
    width: $default-width;
    max-width: $max-width;
    margin: 0 auto !important;
  }
  .stat-info {
    color: #888;
    font-size: 18px;
    small {
      font-size: 13px;
    }
    .el-card__body {
      padding: 5px 0;
      .el-col {
        padding: 8px 0;
        &:first-child {
          border-right: 1px solid #efefef;
        }
      }
    }
  }

  .categories {
    background-color: #fff;
    .el-row {
      margin: 0 auto !important;
      width: $default-width;
      max-width: $max-width;
      .el-card__header {
        padding-left: 0;
        border-bottom: 0;
        padding-bottom: 0;
      }
      .el-card__body {
        padding: 20px 0;
      }
      a {
        display: inline-block;
        padding: 5px 0;
        text-decoration: none;
        margin-right: 5px;
      }
    }
  }

  .login-form .el-card__body {
    padding-bottom: 0;
  }

  .el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.6;
    text-align: center;
    line-height: 320px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n + 1) {
    background-color: #d3dce6;
  }

  .category-item {
    .el-card__body > div {
      display: flex;
      flex-direction: row;
      justify-content: space-between;
      .card-body-left {
        width: 180px;
        padding-right: 20px;
        img {
          width: 180px;
        }
      }
      .card-body-right {
        width: 100%;
        margin-top: -5px;
        box-sizing: border-box;
        padding-right: 200px;
        a {
          text-decoration: none;
          display: block;
          line-height: 30px;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
        }
      }
    }
  }
}
</style>
