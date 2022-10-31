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
              <nuxt-link
                to="/findpassword"
                title="找回密码"
                class="el-link el-link--info"
                ><small>忘记密码？</small></nuxt-link
              >
              <nuxt-link
                to="/register"
                title="注册用户"
                class="float-right el-link el-link--info"
                ><small>注册用户</small></nuxt-link
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
        <div
          v-for="(category, index) in categoryTrees"
          :key="'tree-' + category.id"
        >
          <el-col v-if="index < 4" :span="6">
            <el-card class="box-card" shadow="never">
              <div slot="header" class="clearfix">
                <strong>{{ category.title }}</strong>
              </div>
              <nuxt-link
                v-for="child in category.children"
                :key="'child-' + child.id"
                class="el-link el-link--default"
                :to="{ path: '/category', query: { id: child.id } }"
                >{{ child.title }}</nuxt-link
              >
            </el-card>
          </el-col>
        </div>
      </el-row>
    </div>
    <el-row :gutter="20" class="category-item">
      <el-col
        v-for="category in categoryTrees"
        :key="'card-cate-' + category.id"
        :md="12"
        :sm="12"
        :xs="24"
      >
        <el-card class="box-card mgt-20px" shadow="never">
          <div slot="header" class="clearfix">
            <strong>{{ category.title }}</strong>
            <nuxt-link :to="{ path: '/category', query: { id: category.id } }"
              ><el-button style="float: right; padding: 3px 0" type="text"
                >更多</el-button
              ></nuxt-link
            >
          </div>
          <div>
            <div class="card-body-left hidden-xs-only">
              <nuxt-link
                :to="{ path: '/category', query: { id: category.id } }"
              >
                <el-image :src="category.cover">
                  <div slot="error" class="image-slot">
                    <img
                      src="/static/images/cover-news.png"
                      :alt="category.title"
                    />
                  </div>
                </el-image>
              </nuxt-link>
            </div>
            <div class="card-body-right">
              <nuxt-link class="el-link el-link--default" to="/document/"
                >Docker — 从入门到实战-BookStack.CN</nuxt-link
              >
              <nuxt-link class="el-link el-link--default" to="/document/"
                >MongoDB简明教程</nuxt-link
              >
              <nuxt-link class="el-link el-link--default" to="/document/"
                >TypeScript 官方文档</nuxt-link
              >
              <nuxt-link class="el-link el-link--default" to="/document/"
                >DolphinPHP1.3.0完全开发手册-基于ThinkPHP5.0.20的快速开发框架-05221135</nuxt-link
              >
              <nuxt-link class="el-link el-link--default" to="/document/">
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
import { mapGetters } from 'vuex'
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
  computed: {
    ...mapGetters('category', ['categoryTrees']),
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
        padding: 15px 0;
      }
      a {
        display: inline-block;
        padding: 5px 0;
        text-decoration: none;
        margin-right: 10px;
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
        .image-slot {
          height: 145px;
          overflow: hidden;
        }
        img {
          width: 180px;
          height: 145px;
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
