<template>
  <div class="page page-index">
    <div class="searchbox">
      <el-carousel :interval="3000" arrow="always" :height="isMobile ? '250px' : '360px'" @change="changeCarousel">
        <a
          v-for="(banner, index) in banners"
          :key="'banner-' + banner.id"
          :href="banner.url ? banner.url : 'javascript:;'"
          :target="banner.url ? '_blank' : ''"
          :title="banner.title"
        >
          <el-carousel-item
            :style="
              'background: url(' +  (carouselIndexes.indexOf(index)>-1 ? banner.path: '') + ') center center no-repeat;'
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
            <i
              slot="suffix"
              @click="onSearch"
              class="el-input__icon el-icon-search btn-search"
            ></i>
          </el-input>
        </el-form-item>
        <el-form-item v-if="settings.system.recommend_words">
          <span class="hidden-xs-only">大家在搜:</span>
          <nuxt-link
            v-for="word in settings.system.recommend_words"
            :key="'kw-' + word"
            :to="{
              path: '/search',
              query: { wd: word },
            }"
            ><el-tag size="small">{{ word }}</el-tag></nuxt-link
          >
        </el-form-item>
      </el-form>
    </div>
    <el-row :gutter="20" class="mgt-20px">
      <el-col :span="6" class="float-right right-at-recommend">
        <el-card class="text-center stat-info" shadow="never">
          <el-row>
            <el-col :span="settings.display.show_register_user_count ? 12 : 24">
              <small>收录文档</small>
              <div>
                <span class="el-link el-link--primary">{{
                  stats.document_count || 0
                }}</span>
              </div>
            </el-col>
            <el-col :span="12" v-if="settings.display.show_register_user_count">
              <small>注册用户</small>
              <div>
                <span class="el-link el-link--primary">{{
                  stats.user_count || 0
                }}</span>
              </div>
            </el-col>
          </el-row>
        </el-card>
        <el-card class="text-center mgt-20px hidden-xs-only" shadow="never">
          <nuxt-link to="/upload">
            <el-button type="warning" class="btn-block" icon="el-icon-upload"
              >上传文档</el-button
            >
          </nuxt-link>
        </el-card>
        <el-card
          v-if="user.id > 0"
          class="box-card mgt-20px hidden-xs-only login-form"
          shadow="never"
        >
          <el-row>
            <el-col :span="8">
              <nuxt-link :to="`/user/${user.id}`">
                <user-avatar :size="64" :user="user" />
              </nuxt-link>
            </el-col>
            <el-col :span="16">
              <nuxt-link
                class="el-link el-link--default"
                :to="`/user/${user.id}`"
                ><h3>{{ user.username }}</h3></nuxt-link
              >
              <div class="help-block login-tips">
                <span class="el-link el-link--default" @click="logout">
                  <i class="fa fa-sign-out"></i> &nbsp;<small>退出登录</small>
                </span>
              </div>
            </el-col>
          </el-row>
          <div class="line"></div>
          <el-row class="text-center user-count">
            <el-col :span="8">
              <div><small>文档</small></div>
              <span>{{ user.doc_count || 0 }}</span>
            </el-col>
            <el-col :span="8">
              <div><small>收藏</small></div>
              <span>{{ user.favorite_count || 0 }}</span>
            </el-col>
            <el-col :span="8">
              <div>
                <small>{{ settings.system.credit_name || '魔豆' }}</small>
              </div>
              <span>{{ user.credit_count || 0 }}</span>
            </el-col>
          </el-row>
          <el-button
            v-if="sign.id > 0"
            :key="'sign-' + sign.id"
            class="btn-block"
            type="success"
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
            @click="signToday"
          >
            <i class="fa fa-calendar-plus-o"></i>
            每日签到</el-button
          >
          <div class="mgt-20px">
            <div>个性签名</div>
            <div class="help-block user-signature">
              {{ user.signature || '暂无个性签名' }}
            </div>
          </div>
        </el-card>
        <el-card
          v-else
          class="box-card mgt-20px hidden-xs-only login-form"
          shadow="never"
        >
          <el-row>
            <el-col :span="8">
              <nuxt-link to="/login"
                ><user-avatar :size="64" :user="user"
              /></nuxt-link>
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
              v-for="(item, index) in recommends"
              :key="'recommend' + item.id"
              :span="4"
              :class="isMobile && index > 7 ? 'hidden-xs-only' : ''"
            >
              <nuxt-link :to="`/document/${item.id}`">
                <el-image
                  :src="
                    item.attachment && item.attachment.hash
                      ? `/view/cover/${item.attachment.hash}`
                      : ''
                  "
                  lazy
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
    <div
      class="categories mgt-20px"
      v-if="settings.display.show_index_categories"
    >
      <el-row :gutter="20">
        <div
          v-for="(category, index) in categoryTrees"
          :key="'tree-' + category.id"
        >
          <el-col v-if="index < 4" :span="6">
            <el-card class="box-card" shadow="never">
              <div slot="header" class="clearfix">
                <img
                  :src="category.icon || '/static/images/logo-icon.png'"
                  :alt="category.title"
                  class="category-icon"
                />
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
                <el-image lazy class="category-cover" :src="item.category_cover">
                  <div slot="error" class="image-slot">
                    <img
                      src="/static/images/default-category-cover.png"
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
              >
                <img
                  :src="`/static/images/${getIcon(doc.ext)}_24.png`"
                  :alt="`${getIcon(doc.ext)}文档`"
                />
                <span>{{ doc.title }}</span>
              </nuxt-link>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'
import UserAvatar from '~/components/UserAvatar.vue'
import { listBanner } from '~/api/banner'
import { listDocument, listDocumentForHome } from '~/api/document'
import { getSignedToday, signToday } from '~/api/user'
import { getStats } from '~/api/config'
import { getIcon } from '~/utils/utils'
export default {
  components: { UserAvatar },
  data() {
    return {
      banners: [],
      recommends: [],
      documents: [],
      search: {
        wd: '',
      },
      sign: {
        sign_at: 0,
      },
      stats: {
        document_count: '-',
        user_count: '-',
      },
      carouselIndexes: [0], // 跑马灯index，用于跑马灯图片的懒加载
    }
  },
  head() {
    return {
      title: this.settings.system.title || 'MOREDOC · 魔豆文库',
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: `${this.settings.system.sitename},${this.settings.system.keywords}`,
        },
        {
          hid: 'description',
          name: 'description',
          content: this.settings.system.description,
        },
      ],
    }
  },
  computed: {
    ...mapGetters('category', ['categoryTrees']),
    ...mapGetters('user', ['user']),
    ...mapGetters('setting', ['settings']),
  },
  async created() {
    await Promise.all([
      this.getRecommendDocuments(),
      this.listBanner(),
      this.getDocuments(),
      this.getSignedToday(),
      this.getStats(),
      this.getUser(),
    ])
  },
  methods: {
    ...mapActions('user', ['logout', 'getUser']),
    getIcon,
    async listBanner() {
      const res = await listBanner({
        enable: true,
        field: ['id', 'title', 'path', 'url'],
        type: 0, // 0，网站横幅
      })
      if (res.status === 200) {
        this.banners = res.data.banner
      }
    },
    onSearch() {
      if (this.search.wd) {
        location.href = '/search?wd=' + encodeURIComponent(this.search.wd)
        return
      }
    },
    async getSignedToday() {
      const res = await getSignedToday()
      if (res.status === 200) {
        this.sign = res.data || { id: 0 }
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
      }
    },
    async getStats() {
      const res = await getStats()
      if (res.status === 200) {
        this.stats = res.data || {}
      }
    },
    login() {
      // 跳转到登录页面，先串通页面
      this.$router.push('/login')
    },
    changeCarousel(index){
      let carouselIndexes = this.carouselIndexes
      if (carouselIndexes.indexOf(index) === -1) {
        carouselIndexes.push(index)
      }
      this.carouselIndexes = carouselIndexes
    }
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
    .category-icon {
      width: 22px;
      height: 22px;
      position: relative;
      top: 5px;
    }
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
        padding: 15px 0 20px;
        max-height: 80px;
        overflow: hidden;
        display: -webkit-box;
        // -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;
      }
      a {
        display: inline-block;
        padding: 2px 0 5px;
        text-decoration: none;
        margin-right: 10px;
        margin-bottom: 5px;
      }
    }
  }

  .login-form {
    h3 {
      margin-top: 5px;
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

    .user-count {
      margin: 20px 0;
      font-size: 13px;
      color: #999;
      .el-col:nth-child(2) {
        border-left: 1px solid #efefef;
        border-right: 1px solid #efefef;
      }
      span {
        display: block;
        margin-top: 5px;
        font-size: 16px;
        color: #409eff;
      }
    }

    .user-signature {
      text-align: left;
      text-indent: 2em;
      margin-top: 10px;
      height: 41px;
      font-size: 14px;
      line-height: 23px;
      overflow: hidden;
      text-overflow: ellipsis;
      word-break: break-all;
      display: -webkit-box;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
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
        width: 115px;
        max-width: 100%;
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
        .category-cover{
          height: 145px;
          width: 180px;
          overflow: hidden;
        }
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
          img {
            display: none;
          }
        }
      }
    }
  }
}

// =================================
// 移动端样式
// =================================
@media screen and (max-width: $mobile-width) {
  .page-index {
    .searchbox {
      margin-bottom: 15px;
      .search-form {
        width: 90%;
        .el-input__inner {
          height: 40px;
          line-height: 40px;
        }
      }
    }
    .el-carousel__arrow {
      display: none;
    }
    .latest-recommend {
      width: 100%;
      padding-left: 0 !important;
      padding-right: 0 !important;
      .el-card__body {
        padding: 15px;
        padding-bottom: 0;
      }
      .el-col-4 {
        width: 25%;
        padding-left: 7.5px !important;
        padding-right: 7.5px !important;
      }
      a {
        margin-bottom: 15px;
        .el-image {
          height: auto;
          width: 100%;
          border: 1px solid #e6e6e6;
        }
        div.el-link {
          font-size: 12px;
        }
      }
    }
    .right-at-recommend {
      display: none; // 屏蔽，影响整体美观
      width: 100%;
      margin-top: -20px;
      padding-left: 0 !important;
      padding-right: 0 !important;
      margin-bottom: 20px;
    }
    .categories {
      padding-bottom: 15px;
      .el-col-6 {
        width: 50%;
        .el-card__body {
          height: 75px;
          overflow: hidden;
        }
      }
    }

    .category-item {
      .el-col-12 {
        width: 100%;
        padding-left: 0 !important;
        padding-right: 0 !important;
        .card-body-right {
          padding-right: 0 !important;
          a {
            line-height: 35px !important;
            img {
              display: inline-block !important;
              height: 18px;
              width: 18px;
              position: relative;
              top: 3px;
              margin-right: 3px;
            }
          }
        }
      }
    }
  }
}
</style>
