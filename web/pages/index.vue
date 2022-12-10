<template>
  <div class="page page-index">
    <div class="searchbox">
      <el-carousel :interval="3000" arrow="always" :height="'360px'">
        <a
          v-for="banner in banners"
          :key="'banner-' + banner.id"
          :href="banner.url"
          target="_blank"
          :title="banner.title"
        >
          <el-carousel-item
            :style="
              'background: url(' + banner.path + ') center center no-repeat;'
            "
          >
          </el-carousel-item>
        </a>
      </el-carousel>
      <el-form :model="search" class="search-form" @submit.native.prevent>
        <el-form-item>
          <el-input
            v-model="search.wd"
            size="large"
            placeholder="搜索文档..."
            @keydown.native.enter="onSearch"
          >
            <el-button
              slot="append"
              icon="el-icon-search"
              @click="onSearch"
            ></el-button>
          </el-input>
        </el-form-item>
        <el-form-item>
          <span>大家在搜:</span>
          <a href="#"><el-tag size="small">Java教程</el-tag></a>
          <a href="#"><el-tag size="small">PHP教程</el-tag></a>
          <a href="#"><el-tag size="small">GPTChat</el-tag></a>
          <a href="#"><el-tag size="small">开源中国</el-tag></a>
          <a href="#"><el-tag size="small">小学语文</el-tag></a>
        </el-form-item>
      </el-form>
    </div>
    <el-row :gutter="20" class="mgt-20px">
      <el-col :span="6" class="float-right">
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
              <el-button class="btn-block" type="primary" @click="login"
                >立即登录</el-button
              >
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
      <el-col :span="18" class="latest-recommend">
        <el-card shadow="never">
          <div slot="header">最新推荐</div>
          <el-row :gutter="20">
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/entgo-0.11-zh/uploads/202210/171f825ac77e9e82.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  Kong Gateway v3.0.x Documentation
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/istio-1.16-zh/uploads/202212/172cb1a7fc49b727.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">Istio v1.16 中文文档</div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/traefik-2.9-en/uploads/202211/172b2cf6a8b5eb8b.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  Traefik v2.9 Documentation
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/kubevela-1.6-zh/uploads/202212/172cb4a60512d3c9.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  KubeVela v1.6 中文文档
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/explore-python/uploads/201812/1573e6514b85abb3.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">Python 之旅</div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/python-3.11.0-zh/uploads/202211/17285eb115212b0b.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  Python 3.11.0 官方文档(全)
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/entgo-0.11-zh/uploads/202210/171f825ac77e9e82.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  Kong Gateway v3.0.x Documentation
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/istio-1.16-zh/uploads/202212/172cb1a7fc49b727.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">Istio v1.16 中文文档</div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/traefik-2.9-en/uploads/202211/172b2cf6a8b5eb8b.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  Traefik v2.9 Documentation
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/kubevela-1.6-zh/uploads/202212/172cb4a60512d3c9.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  KubeVela v1.6 中文文档
                </div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/explore-python/uploads/201812/1573e6514b85abb3.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">Python 之旅</div>
              </nuxt-link>
            </el-col>
            <el-col :span="4">
              <nuxt-link to="/">
                <img
                  src="https://static.sitestack.cn/projects/python-3.11.0-zh/uploads/202211/17285eb115212b0b.png/cover"
                  alt=""
                />
                <div class="el-link el-link--default">
                  Python 3.11.0 官方文档(全)
                </div>
              </nuxt-link>
            </el-col>
          </el-row>
        </el-card>
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
        :span="12"
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
import { listBanner } from '~/api/banner'
export default {
  name: 'IndexPage',
  data() {
    return {
      user: {
        username: '',
        password: '',
      },
      banners: [],
      search: {
        wd: '',
      },
    }
  },
  head() {
    return {
      // title: 'MOREDOC · 魔刀文库，开源文库系统',
      script: [],
    }
  },
  computed: {
    ...mapGetters('category', ['categoryTrees']),
  },
  async created() {
    await this.listBanner()
  },
  methods: {
    async listBanner() {
      const res = await listBanner({
        enable: true,
        field: ['id', 'title', 'path', 'url'],
        type: 0, // 0，网站横幅
      })
      if (res.status === 200) {
        this.banners = res.data.banner
      } else {
        console.log(res)
      }
    },
    onSearch() {
      if (this.search.wd) {
        const wd = this.search.wd
        this.search.wd = ''
        this.$router.push({ path: '/search', query: { wd } })
      }
    },
    login() {
      // 跳转到登录页面，先串通页面
      this.$router.push('/login')
    },
  },
}
</script>
<style lang="scss">
.page-index {
  width: 100%;
  max-width: 100%;
  margin-top: -20px;
  .searchbox {
    position: relative;
    margin-bottom: 20px;
    a {
      display: inline-block;
    }
    .el-carousel__item {
      background-size: cover !important;
    }
    // 搜索表单垂直居中显示
    .search-form {
      position: absolute;
      z-index: 99;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 640px;
      color: #fff;
      .el-form-item {
        margin-bottom: 0;
      }
      .el-tag {
        margin-left: 5px;
      }
      .el-input__icon {
        color: #666;
      }
      .el-input__inner {
        border-right: 0;
        height: 45px;
        line-height: 45px;
        font-size: 15px;
        &:focus {
          border-color: #dcdfe6;
        }
      }
      .el-input-group__append {
        background-color: #fff;
        border-left: 0;
      }
    }
  }
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

  .latest-recommend {
    .el-card__body {
      padding-bottom: 0;
    }
    a {
      text-decoration: none;
      display: block;
      margin-bottom: 20px;
      &:hover {
        color: #409eff;
      }
      img {
        width: 100%;
        border: 1px solid #efefef;
        border-radius: 5px;
      }
      div {
        height: 40px;
        overflow: hidden;
        margin-bottom: 0px;
        text-overflow: ellipsis;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        word-break: break-word;
        font-size: 13px;
        line-height: 20px;
      }
    }
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
          border-radius: 5px;
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
