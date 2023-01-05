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
          <nuxt-link class="el-link el-link--default" v-else to="/login"
            >登录账户</nuxt-link
          >
        </span>
      </el-col>
    </el-row>
    <div class="search-box">
      <el-row :gutter="20">
        <el-col :span="4">
          <nuxt-link to="/" :title="settings.system.sitename"
            ><img
              :src="settings.system.logo"
              style="max-width: 100%"
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
            v-for="item in searchExts"
            :key="'st-' + item.value"
            :to="{
              path: '/search',
              query: {
                wd: query.wd,
                ext: item.value,
                page: 1,
                size: 10,
              },
            }"
            :class="[
              'el-link',
              'el-link--default',
              item.value === query.ext ? 'el-link-active' : '',
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
                ext: query.ext,
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
        <el-card v-loading="loading" shadow="never">
          <div slot="header">
            本次搜索耗时
            <span class="el-link el-link--danger">{{ spend || '0.000' }}</span>
            秒，在
            <span class="el-link el-link--primary">{{
              stats.document_count || '0'
            }}</span>
            篇文档中为您找到相关结果约
            <span class="el-link el-link--danger">{{ total || 0 }}</span> 个.
          </div>
          <!-- <div class="search-result-none">没有搜索到内容...</div> -->
          <div class="search-result">
            <ul>
              <li v-if="docs.length === 0">
                <div class="noresult">暂无搜索结果</div>
              </li>
              <li v-for="doc in docs" :key="'doc-' + doc.id">
                <h3 class="doc-title">
                  <a
                    :href="`/document/${doc.id}`"
                    class="el-link el-link--primary"
                  >
                    <img
                      :src="`/static/images/${doc.icon}_24.png`"
                      :alt="`${doc.icon}文档`"
                    />
                    {{ doc.title }}
                  </a>
                </h3>
                <div class="doc-desc">{{ doc.description }}</div>
                <div class="doc-info">
                  <el-rate
                    v-model="doc.score"
                    disabled
                    show-score
                    text-color="#ff9900"
                    score-template="{value}"
                  >
                  </el-rate>
                  <span class="float-right"
                    >{{ doc.price || 0 }} 魔豆 | {{ doc.pages || '-' }} 页 |
                    {{ formatBytes(doc.size) }}
                    <span class="hidden-xs-only">| 2019-06-10 10:17</span></span
                  >
                </div>
              </li>
            </ul>
          </div>
          <el-pagination
            v-if="total > 0"
            :current-page="query.page"
            :page-size="query.size"
            layout="total,  prev, pager, next, jumper"
            :total="total"
            @current-change="onPageChange"
          >
          </el-pagination>
        </el-card>
      </el-col>
      <el-col v-if="keywords.length > 0" :span="6" class="search-right">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <span>相关搜索词</span>
          </div>
          <nuxt-link
            v-for="keyword in keywords"
            :key="'kw-' + keyword"
            :to="{
              path: '/search',
              query: {
                wd: keyword,
                page: 1,
                size: 10,
              },
            }"
            class="el-link el-link--default"
            >{{ keyword }}</nuxt-link
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
import { getStats } from '~/api/config'
import { searchDocument } from '~/api/document'
import { formatBytes, getIcon } from '~/utils/utils'
export default {
  data() {
    return {
      loading: false,
      query: {
        wd: this.$route.query.wd || '',
        page: 1,
        size: 10,
        ext: 'all', // 搜索类型
        sort: 'default', // 排序
      },
      searchExts: [
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
        { label: '下载排序', value: 'download_count' },
        { label: '浏览排序', value: 'view_count' },
        { label: '收藏排序', value: 'favorite_count' },
      ],
      docs: [],
      total: 0,
      spend: '',
      keywords: [],
      stats: {
        document_count: '-',
      },
    }
  },
  head() {
    return {
      bodyAttrs: {
        class: 'search-page',
      },
      title: this.settings.system.title || 'MOREDOC · 魔豆文库',
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: `${query.wd},文档搜索,${this.settings.system.sitename},${this.settings.system.keywords}`,
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
    },
  },
  created() {
    const query = { ...this.query, ...this.$route.query }
    query.page = parseInt(query.page) || 1
    query.size = parseInt(query.size) || 10
    this.query = query
    this.getStats()
  },
  methods: {
    formatBytes,
    onSearch() {
      this.$router.push({
        path: '/search',
        query: {
          ...this.query,
          page: 1,
          size: 10,
          sort: 'default',
          ext: 'all',
        },
      })
    },
    async getStats() {
      const res = await getStats()
      if (res.status === 200) {
        this.stats = res.data
      }
    },
    async execSearch() {
      this.loading = true
      const res = await searchDocument(this.query)
      if (res.status === 200) {
        this.total = res.data.total
        this.spend = res.data.spend
        const docs = res.data.document || []
        const keywords = []
        this.docs = docs.map((doc) => {
          doc.score = doc.score || 300
          doc.score = doc.score / 100
          doc.icon = getIcon(doc.ext)
          try {
            doc.keywords.split(',').map((keyword) => {
              keyword = keyword.trim()
              if (keyword && !keywords.includes(keyword)) {
                keywords.push(keyword)
              }
              return keyword
            })
          } catch (error) {}
          this.keywords = keywords
          return doc
        })
      }
      this.loading = false
    },
    onPageChange(page) {
      this.$router.push({
        path: '/search',
        query: {
          ...this.query,
          page,
        },
      })
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
      min-height: 540px;
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
      .noresult {
        text-align: center;
        font-size: 14px;
        line-height: 200px;
        color: #999;
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
        display: inline-block;
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
