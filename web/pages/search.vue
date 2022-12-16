<template>
  <div class="page page-search">
    <el-row class="header-links">
      <el-col :span="24">
        <nuxt-link to="/" class="el-link el-link--default">文库首页</nuxt-link>
        <nuxt-link
          v-for="category in categoryTrees"
          v-show="category.enable"
          :key="'cate-' + category.id"
          :to="`/category/${category.id}`"
          class="el-link el-link--default"
          >{{ category.title }}</nuxt-link
        >
        <span class="float-right">
          <nuxt-link to="/upload" class="el-link el-link--default"
            >上传文档</nuxt-link
          >
          <nuxt-link
            v-if="user.id > 0"
            :to="`/user/${user.id}`"
            class="el-link el-link--default"
            >会员中心</nuxt-link
          >
          <nuxt-link v-else to="/login">登录账户</nuxt-link>
        </span>
      </el-col>
    </el-row>
    <div class="search-box">
      <el-row :gutter="20">
        <el-col :span="4">
          <nuxt-link to="/" :title="settings.system.sitename"
            ><img
              :src="settings.system.logo"
              style="height: 56px; max-width: 100%"
              :alt="settings.system.sitename"
          /></nuxt-link>
        </el-col>
        <el-col :span="14">
          <el-input
            v-model="query.wd"
            class="search-input"
            size="large"
            placeholder="请输入关键词"
            @keyup.enter.native="onSearch"
          >
            <el-button slot="append" icon="el-icon-search" @click="onSearch" />
          </el-input>
        </el-col>
      </el-row>
    </div>
    <el-row :gutter="20" class="mgt-20px">
      <el-col :span="4" class="search-left">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <span>类型</span>
          </div>
          <nuxt-link
            v-for="item in searchTypes"
            :key="'st-' + item.value"
            :to="{
              path: '/search',
              query: {
                wd: query.wd,
                type: item.value,
                page: 1,
                size: 10,
              },
            }"
            :class="[
              'el-link',
              'el-link--default',
              item.value === query.type ? 'el-link-active' : '',
            ]"
            >{{ item.label }}</nuxt-link
          >
        </el-card>
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <span>排序</span>
          </div>
          <nuxt-link
            v-for="item in searchSorts"
            :key="'ss-' + item.value"
            :to="{
              path: '/search',
              query: {
                wd: query.wd,
                type: query.type,
                page: 1,
                size: 10,
                sort: item.value,
              },
            }"
            :class="[
              'el-link',
              'el-link--default',
              item.value === query.sort ? 'el-link-active' : '',
            ]"
            >{{ item.label }}</nuxt-link
          >
        </el-card>
      </el-col>
      <el-col :span="14" class="search-main">
        <el-card shadow="never">
          <div slot="header">
            本次搜索耗时
            <span class="el-link el-link--danger">0.001</span> 秒，在
            <span class="el-link el-link--primary">3235</span>
            篇文档中为您找到相关结果约
            <span class="el-link el-link--danger">4</span> 个.
          </div>
          <!-- <div class="search-result-none">没有搜索到内容...</div> -->
          <div class="search-result">
            <ul>
              <li v-for="i in 10" :key="'i-' + i">
                <h3 class="doc-title">
                  <a href="/document/" class="el-link el-link--primary">
                    <img
                      v-if="i === 1"
                      src="/static/images/pdf_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 10"
                      src="/static/images/epub_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 2"
                      src="/static/images/umd_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 3"
                      src="/static/images/mobi_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 4"
                      src="/static/images/chm_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 5"
                      src="/static/images/other_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 6"
                      src="/static/images/ppt_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 7"
                      src="/static/images/text_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 8"
                      src="/static/images/word_24.png"
                      alt=""
                    />
                    <img
                      v-if="i === 9"
                      src="/static/images/excel_24.png"
                      alt=""
                    />
                    张孝祥正在整理Java就业面试题大全
                  </a>
                </h3>
                <div class="doc-desc">
                  传智播客 ——IT 就业培训专家 http://www.itcast.cn
                  提示：本大全每半月更新一次，请持续保持关注！谢 谢！
                  索取网址：www.itcast.cn
                  从享受生活的角度上来说：“程序员并不是一种最好的职业，我认为两种人可以做
                </div>
                <div class="doc-info">
                  <el-rate
                    v-model="score"
                    disabled
                    show-score
                    text-color="#ff9900"
                    score-template="{value}"
                  >
                  </el-rate>
                  <span class="float-right"
                    >5 金币 | 141 页 | 786.00 KB
                    <span class="hidden-xs-only">| 2019-06-10 10:17</span></span
                  >
                </div>
              </li>
            </ul>
          </div>
        </el-card>
        <el-card shadow="never">
          <el-pagination
            :current-page="1"
            :page-size="10"
            layout="total,  prev, pager, next, jumper"
            :total="400"
          >
          </el-pagination>
        </el-card>
      </el-col>
      <el-col :span="6" class="search-right">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <span>大家在搜</span>
          </div>
          <nuxt-link to="/" class="el-link el-link--default"
            >JAVA语言</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >计算机图形学</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >程序设计</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"> 教材</nuxt-link>
          <nuxt-link to="/" class="el-link el-link--default"
            >面向对象</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >语言网页制作</nuxt-link
          >
        </el-card>
        <el-card shadow="never" class="mgt-20px">
          <img
            src="https://www.wenkuzhijia.cn/static/Home/default/img/cover.png"
            alt=""
          />
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
      score: 4.7,
      query: {
        wd: this.$route.query.wd || '',
        page: 1,
        size: 10,
        type: 'all', // 搜索类型
        sort: 'sort', // 排序
      },
      searchTypes: [
        { label: '不限', value: 'all' },
        { label: 'PDF', value: 'pdf' },
        { label: 'DOC', value: 'doc' },
        { label: 'PPT', value: 'ppt' },
        { label: 'XLS', value: 'xls' },
        { label: 'TXT', value: 'txt' },
        { label: '其他', value: 'other' },
      ],
      searchSorts: [
        { label: '默认排序', value: 'default' },
        { label: '页数排序', value: 'pages' },
        { label: '评分排序', value: 'score' },
        { label: '大小排序', value: 'size' },
        { label: '下载排序', value: 'download' },
        { label: '浏览排序', value: 'view' },
        { label: '收藏排序', value: 'favorite' },
      ],
    }
  },
  head() {
    return {
      bodyAttrs: {
        class: 'search-page',
      },
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
    ...mapGetters('category', ['categoryTrees']),
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    '$route.query': {
      handler(val) {
        const query = { ...this.query, ...val }
        query.page = parseInt(query.page) || 1
        query.size = parseInt(query.size) || 10
        this.query = query
        this.execSearch()
      },
      immediate: true,
      deep: true,
    },
  },
  async created() {},
  methods: {
    onSearch() {
      this.$router.push({
        path: '/search',
        query: {
          ...this.query,
          page: 1,
          size: 10,
          sort: 'default',
          type: 'all',
        },
      })
    },
    execSearch() {
      console.log('execSearch')
    },
  },
}
</script>
<style lang="scss">
.search-page {
  .layout-default {
    padding-top: 0;
    .el-main {
      padding-top: 0;
    }
  }
}
.page-search {
  width: $max-width;
  & > .el-row {
    width: $default-width;
    max-width: $max-width;
    margin: 0 auto !important;
  }
  .header-links {
    padding: 0 10px;
    .el-link {
      line-height: 35px;
      margin-right: 10px;
    }
  }
  .search-box {
    padding: 20px 0;
    margin-bottom: 20px;
    background-color: #fff;
    & > .el-row {
      margin: 0 auto !important;
      width: $default-width;
      max-width: $max-width;
    }
    .el-input-group__append,
    .el-input-group__prepend {
      background-color: #dcdfe6;
    }
    .el-input__inner {
      border-right: 0;
    }
    .search-input {
      margin-top: 5px;
    }
  }
  .search-left {
    .el-card:first-of-type {
      .el-card__body {
        padding-bottom: 0;
      }
    }
    .el-card__header {
      border-bottom: 0;
      padding: 15px 10px 0;
    }
    .el-card__body {
      padding: 15px 10px;
    }
    a {
      display: block;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      line-height: 36px;
      padding-left: 1em;
      &.el-link-active,
      &:hover {
        color: #409eff;
        background: $background-grey-light;
      }
    }
  }
  .search-right {
    a {
      display: inline-block;
      line-height: 30px;
      margin-right: 10px;
    }
  }
  .search-main {
    .el-card__header {
      font-weight: normal;
      font-size: 15px;
      line-height: 20px;
      color: #888;
      padding-bottom: 15px;
      .el-link {
        position: relative;
        top: -2px;
      }
    }

    .el-card__body {
      padding-top: 0;
      padding-bottom: 10px;
    }

    .search-result {
      ul,
      li {
        list-style: none;
        padding: 0;
      }
      li {
        padding-top: 20px;
        &:first-of-type {
          padding-top: 0;
        }
      }
    }
    h3 {
      margin-bottom: 10px;
      a {
        font-size: 18px;
        font-weight: normal;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        display: block;
        width: 100%;
        img {
          height: 18px;
          position: relative;
          top: 2px;
        }
      }
    }
    .doc-desc {
      font-size: 15px;
      color: #6b7a88;
      line-height: 180%;
      word-break: break-all;
      margin-bottom: 10px;
      display: -webkit-box;
      -webkit-line-clamp: 3;
      max-height: 81px;
      overflow: hidden;
      -webkit-box-orient: vertical;
      text-overflow: ellipsis;
    }
    .doc-info {
      color: #bdc3c7;
      font-size: 14px;
    }
  }
}
</style>
