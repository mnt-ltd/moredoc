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
        <el-card class="text-center mgt-20px" shadow="never">
          <nuxt-link to="/upload">
            <el-button type="warning" class="btn-block" icon="el-icon-upload"
              >上传文档</el-button
            >
          </nuxt-link>
        </el-card>
        <el-card
          class="box-card mgt-20px hidden-xs-only login-form"
          shadow="never"
        >
          <el-row>
            <el-col :span="8">
              <el-avatar
                :size="64"
                src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
              ></el-avatar>
            </el-col>
            <el-col :span="16">
              <h3>欢迎您，游客</h3>
              <div class="help-block login-tips">登录可体验更多功能</div>
            </el-col>
          </el-row>
          <div class="line"></div>
          <div>
            <ul>
              <li>分享知识，获取收益</li>
              <li>发表点评，抒发见解</li>
              <li>下载文档，畅享阅读</li>
              <li>身份认证，彰显尊贵</li>
            </ul>
          </div>
          <div class="btn-login">
            <nuxt-link to="/login"
              ><el-button class="btn-block" type="primary"
                >马上登录</el-button
              ></nuxt-link
            >
          </div>
          <div class="help-block">
            <el-row>
              <el-col :span="12">
                <nuxt-link to="/findpassword" class="el-link el-link--default"
                  ><small>找回密码</small></nuxt-link
                >
              </el-col>
              <el-col :span="12" class="text-right">
                <nuxt-link to="/register" class="el-link el-link--default"
                  ><small>注册账户</small></nuxt-link
                >
              </el-col>
            </el-row>
          </div>
        </el-card>
      </el-col>
      <el-col :span="18" class="latest-recommend" keep-alive>
        <el-card shadow="never">
          <div slot="header">最新推荐</div>
          <el-row :gutter="20">
            <el-col
              v-for="item in recommends"
              :key="'recommend' + item.id"
              :span="4"
            >
              <nuxt-link :to="`/document/${item.id}`">
                <el-image
                  :src="
                    item.attachment && item.attachment.hash
                      ? `/view/cover/${item.attachment.hash}`
                      : ''
                  "
                  :alt="item.title"
                >
                  <div slot="error" class="image-slot">
                    <img src="/static/images/default-cover.png" />
                  </div>
                </el-image>
                <div class="el-link el-link--default">{{ item.title }}</div>
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
                :to="`/category/${child.id}`"
                >{{ child.title }}</nuxt-link
              >
            </el-card>
          </el-col>
        </div>
      </el-row>
    </div>
    <el-row :gutter="20" class="category-item">
      <el-col
        v-for="item in documents"
        :key="'card-cate-' + item.category_id"
        :span="12"
      >
        <el-card class="box-card mgt-20px" shadow="never">
          <div slot="header" class="clearfix">
            <strong>{{ item.category_name }}</strong>
            <nuxt-link :to="`/category/${item.category_id}`"
              ><el-button style="float: right; padding: 3px 0" type="text"
                >更多</el-button
              ></nuxt-link
            >
          </div>
          <div>
            <div class="card-body-left hidden-xs-only">
              <nuxt-link :to="`/category/${item.category_id}`">
                <el-image :src="item.category_cover">
                  <div slot="error" class="image-slot">
                    <img
                      src="/static/images/cover-news.png"
                      :alt="item.category_name"
                    />
                  </div>
                </el-image>
              </nuxt-link>
            </div>
            <div class="card-body-right">
              <nuxt-link
                v-for="doc in item.document"
                :key="'c-' + item.category_id + 'd' + doc.id"
                class="el-link el-link--default"
                :to="`/document/${doc.id}`"
                >{{ doc.title }}</nuxt-link
              >
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
import { listDocument, listDocumentForHome } from '~/api/document'
export default {
  name: 'IndexPage',
  data() {
    return {
      banners: [],
      recommends: [],
      documents: [],
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
    await Promise.all([
      this.getRecommendDocuments(),
      this.listBanner(),
      this.getDocuments(),
    ])
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
    async getRecommendDocuments() {
      const res = await listDocument({
        field: ['id', 'title'],
        is_recommend: true,
        order: 'recommend_at desc',
        limit: 12,
      })
      if (res.status === 200) {
        this.recommends = res.data.document || []
      }
    },
    async getDocuments() {
      const res = await listDocumentForHome({
        limit: 5,
      })
      if (res.status === 200) {
        this.documents = res.data.document || []
      } else {
        console.log(res)
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

  .login-form {
    h3 {
      margin-top: 8px;
    }
    .line {
      border-top: 1px solid #efefef;
      margin: 15px 0;
    }
    ul,
    li {
      margin: 0;
      padding: 0;
    }
    ul {
      margin: 10px 0;
    }
    li {
      margin-left: 20px;
      line-height: 200%;
      color: #555;
      font-size: 15px;
    }
    .btn-login {
      margin: 15px 0 5px;
    }
    .login-tips {
      margin-top: -10px;
      font-size: 14px;
    }
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
      .el-image {
        border: 2px solid #efefef;
        border-radius: 5px;
        height: 160px;
        img {
          width: 100%;
          transition: transform 0.3s ease 0s;
          &:hover {
            transform: scale(1.2);
          }
        }
      }

      div.el-link {
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
