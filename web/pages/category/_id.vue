<template>
  <div class="page page-category">
    <el-row>
      <el-col :span="24">
        <el-card shadow="never" ref="breadcrumb">
          <div slot="header" class="clearfix">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item>
                <nuxt-link to="/"><i class="fa fa-home"></i> 首页</nuxt-link>
              </el-breadcrumb-item>
              <el-breadcrumb-item
                v-for="item in breadcrumbs"
                :key="'bread1-' + item.id"
              >
                <el-dropdown v-if="item.siblings.length > 0">
                  <span class="el-dropdown-link">
                    {{ item.title
                    }}<i class="el-icon-arrow-down el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item
                      v-for="ss in item.siblings"
                      :key="'s1-' + ss.id"
                    >
                      <nuxt-link
                        class="el-link el-link--default block"
                        :class="{
                          'el-link--primary': ss.id === item.id,
                        }"
                        :to="`/category/${ss.id}`"
                        >{{ ss.title }}</nuxt-link
                      ></el-dropdown-item
                    >
                  </el-dropdown-menu>
                </el-dropdown>
                <span v-else>{{ item.title }}</span>
              </el-breadcrumb-item>
            </el-breadcrumb>
          </div>
          <div
            v-if="
              breadcrumbs.length > 0 &&
              breadcrumbs[breadcrumbs.length - 1].show_description
            "
            class="category-description"
          >
            {{ breadcrumbs[breadcrumbs.length - 1].description }}
          </div>
          <div class="item-row" v-if="categoryChildren.length > 0">
            <div class="item-name">分类</div>
            <div class="item-content">
              <nuxt-link
                v-for="child in categoryChildren"
                :key="'tree-' + child.id"
                :to="`/category/${child.id}`"
                :title="child.title"
                class="el-link el-link--default"
                >{{ child.title }}</nuxt-link
              >
            </div>
          </div>
          <div class="item-row">
            <div class="item-name">类型</div>
            <div class="item-content">
              <nuxt-link
                v-for="item in exts"
                :key="item.value"
                :to="{ query: { ext: item.value } }"
                class="el-link"
                :class="
                  item.value === $route.query.ext ||
                  (!item.value && !$route.query.ext)
                    ? 'el-link--primary'
                    : 'el-link--default'
                "
                >{{ item.label }}</nuxt-link
              >
            </div>
          </div>
          <div class="item-row">
            <div class="item-name">费用</div>
            <div class="item-content">
              <nuxt-link
                v-for="item in feeTypeOptions"
                :key="item.value"
                :to="{ query: { ...$route.query, fee_type: item.value } }"
                class="el-link"
                :class="
                  item.value === $route.query.fee_type ||
                  (!item.value && !$route.query.fee_type)
                    ? 'el-link--primary'
                    : 'el-link--default'
                "
                >{{ item.label }}</nuxt-link
              >
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="mgt-20px">
      <el-col :span="18">
        <el-card shadow="never" ref="docList" class="doc-list">
          <div slot="header">
            <el-tabs v-model="query.sort" @tab-click="sortClick">
              <el-tab-pane name="default">
                <span slot="label"
                  ><i class="el-icon-coffee-cup"></i> 综合</span
                >
              </el-tab-pane>
              <el-tab-pane name="latest">
                <span slot="label"><i class="el-icon-date"></i> 最新</span>
              </el-tab-pane>
              <el-tab-pane name="view">
                <span slot="label"><i class="el-icon-view"></i> 浏览</span>
              </el-tab-pane>
              <el-tab-pane name="recommend">
                <span slot="label"
                  ><i class="el-icon-coordinate"></i> 推荐</span
                >
              </el-tab-pane>
              <el-tab-pane name="favorite">
                <span slot="label"><i class="el-icon-star-off"></i> 收藏</span>
              </el-tab-pane>
              <el-tab-pane name="download">
                <span slot="label"><i class="el-icon-download"></i> 下载</span>
              </el-tab-pane>
              <el-tab-pane name="pages">
                <span slot="label"><i class="el-icon-files"></i> 页数</span>
              </el-tab-pane>
            </el-tabs>
          </div>
          <div v-loading="loading" class="doc-list-data">
            <document-list v-if="documents.length > 0" :documents="documents" />
            <div v-if="empty && documents.length === 0" class="no-data">
              <el-empty description="暂无数据"></el-empty>
            </div>
          </div>
          <el-pagination
            v-if="total > 0"
            :current-page="query.page"
            :page-size="size"
            :layout="
              isMobile
                ? 'total, prev, pager, next'
                : 'total, prev, pager, next, jumper'
            "
            :pager-count="isMobile ? 5 : 7"
            :small="isMobile"
            :total="total"
            @current-change="pageChange"
          >
          </el-pagination>
        </el-card>
      </el-col>
      <el-col :span="6" class="hidden-xs-only">
        <el-card shadow="never" class="keywords" ref="keywords">
          <div slot="header">
            <el-row>
              <el-col :span="8" class="header-title">关键词</el-col>
            </el-row>
          </div>
          <div v-loading="loading">
            <nuxt-link
              v-for="keyword in keywords"
              :key="'kw' + keyword"
              :to="{ path: '/search', query: { wd: keyword } }"
            >
              <el-tag effect="plain"> {{ keyword }}</el-tag>
            </nuxt-link>
            <div v-if="keywords.length === 0">
              <el-empty description="暂无相关关键词"></el-empty>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import DocumentList from '~/components/DocumentList.vue'
import { listDocument } from '~/api/document'
import { getIcon } from '~/utils/utils'
export default {
  components: { DocumentList },
  data() {
    return {
      filterText: '',
      defaultProps: {
        children: 'children',
        label: 'title',
      },
      query: {
        id: 0,
        sort: 'default',
        page: 1,
      },
      size: 10,
      breadcrumbs: [], // 面包屑
      trees: [],
      categoryChildren: [],
      documents: [],
      categoryId: parseInt(this.$route.params.id) || 0,
      total: 0,
      keywords: [],
      loading: false,
      empty: false,
      cardOffsetTop: 0,
      cardWidth: 0,
      title: '',
      hasExpand: false,
      exts: [
        { label: '不限', value: '' },
        { label: 'PDF', value: 'pdf' },
        { label: 'DOC', value: 'doc' },
        { label: 'PPT', value: 'ppt' },
        { label: 'XLS', value: 'xls' },
        { label: 'TXT', value: 'txt' },
        { label: '其它', value: 'other' },
      ],
      feeTypeOptions: [
        { label: '不限', value: '' },
        { label: '免费', value: 'free' },
        { label: '付费', value: 'charge' },
      ],
    }
  },
  head() {
    return {
      title: this.title + ' - ' + this.settings.system.sitename,
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: this.keywords.join(','),
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
    ...mapGetters('category', ['categories', 'categoryMap']),
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    },
    $route() {
      this.setQuery()
      this.loadData()
    },
  },
  async created() {
    if (this.categories.length === 0) {
      await this.$store.dispatch('category/getCategories')
    }
    const breadcrumbs = []
    let category = { siblings: [], ...this.categoryMap[this.categoryId] }
    if (category.id) {
      category.siblings =
        this.categories.filter((x) => x.parent_id === category.parent_id) || []
      breadcrumbs.push(category)
      while (category.parent_id) {
        category = { siblings: [], ...this.categoryMap[category.parent_id] }
        if (category.id) {
          category.siblings =
            this.categories.filter((x) => x.parent_id === category.parent_id) ||
            []
          breadcrumbs.splice(0, 0, category)
        }
      }
    }

    var titles = []
    breadcrumbs.forEach((x) => {
      titles.push(x.title)
    })
    this.title = titles.join(' · ')

    this.breadcrumbs = breadcrumbs

    // 查找当前最后一个面包屑导航的子分类
    let categoryChildren = []
    if (breadcrumbs.length > 0) {
      categoryChildren = this.categories.filter(
        (x) => x.parent_id === breadcrumbs[breadcrumbs.length - 1].id
      )
    }
    this.categoryChildren = categoryChildren

    this.setQuery()
    this.loadData()
  },
  mounted() {
    this.$nextTick(() => {
      try {
        this.cardOffsetTop = this.$refs.breadcrumb.$el.offsetHeight
      } catch (error) {}
    })
    window.addEventListener('scroll', this.handleScroll)
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    ...mapGetters('category', ['getCategories']),
    filterTree(value, data) {
      if (!value) return true
      return data.title.toLowerCase().includes(value.toLowerCase())
    },
    handleNodeClick(category) {
      this.$router.push({
        path: '/category/' + category.id,
      })
    },
    setQuery() {
      this.query.id = parseInt(this.$route.params.id) || 0
      this.query.sort = this.$route.query.sort || this.query.sort
      this.query.page = parseInt(this.$route.query.page) || 1
    },
    go2cate(id) {
      this.$router.push({
        path: '/category/' + id,
      })
    },
    handleScroll() {
      const scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop
      const keywords = this.$refs.keywords.$el
      if (keywords) {
        if (this.cardWidth === 0) {
          this.cardWidth = keywords.offsetWidth
        }

        if (scrollTop > this.cardOffsetTop) {
          keywords.style.position = 'fixed'
          keywords.style.top = '80px'
          keywords.style.zIndex = '1000'
          keywords.style.width = this.cardWidth + 'px'
        } else {
          keywords.style = null
        }
      }
    },
    sortClick(tab) {
      this.$router.push({
        path: `/category/${this.categoryId}`,
        query: {
          ...this.$route.query,
          sort: tab.name,
        },
      })
    },
    pageChange(page) {
      this.$router.push({
        path: `/category/${this.categoryId}`,
        query: {
          sort: this.query.sort,
          page,
        },
      })
    },
    async loadData() {
      this.loading = true
      let order = 'id desc'
      let status = []
      switch (this.query.sort) {
        case 'latest':
          order = 'id desc'
          break
        case 'view':
          order = 'view_count desc'
          break
        case 'favorite':
          order = 'favorite_count desc'
          break
        case 'comment':
          order = 'comment_count desc'
          break
        case 'pages':
          order = 'pages desc'
          break
        case 'recommend':
          order = 'recommend_at desc'
          break
        case 'download':
          order = 'download_count desc'
          break
        default:
          // 已转换完成的文档，基本有封面，展示的时候不会显得空落落的
          status = [2]
          break
      }
      const res = await listDocument({
        order,
        status,
        page: this.query.page,
        size: this.size,
        category_id: this.categoryId,
        ext: this.$route.query.ext,
        field: [
          'id',
          'title',
          'keywords',
          'description',
          'view_count',
          'favorite_count',
          'comment_count',
          'created_at',
          'size',
          'price',
          'pages',
          'ext',
          'score',
        ],
        fee_type: this.$route.query.fee_type,
      })
      if (res.status === 200) {
        this.total = res.data.total
        const documents = res.data.document || []
        const keywords = []
        this.documents = documents.map((x) => {
          x.icon = getIcon(x.ext)
          if (x.keywords) {
            x.keywords.split(',').forEach((keyword) => {
              keyword = keyword.trim()
              if (keyword && !keywords.includes(keyword)) {
                keywords.push(keyword)
              }
            })
          }
          x.score = parseFloat(x.score) / 100 || 4.0
          return x
        })
        this.keywords = keywords
      }
      this.loading = false
      if (this.query.page === 1 && this.documents.length === 0) {
        this.empty = true
      }
    },
  },
}
</script>
<style lang="scss">
.page-category {
  .el-breadcrumb__inner {
    cursor: pointer !important;
    a {
      font-weight: normal;
    }
  }
  .item-row {
    display: flex;
    .item-name {
      width: 60px;
      font-size: 15px;
      color: #bbb;
    }
    .item-content {
      flex: 1;
    }
    a {
      display: inline-block;
      margin-right: 20px;
      margin-bottom: 20px;
      font-weight: normal;
    }
    &:last-of-type {
      margin-bottom: -20px;
    }
  }
  .category-description {
    border: 1px dashed #ddd;
    margin-bottom: 20px;
    padding: 15px;
    border-radius: 4px;
    font-size: 14px;
    color: #888;
    line-height: 180%;
  }
  .doc-list-data {
    min-height: 200px;
    .no-data {
      text-align: center;
      font-size: 14px;
      color: #aaa;
    }
  }
  .doc-list {
    .el-card__header {
      padding: 0 20px;
      .el-tabs__header {
        margin: 0;
      }
      .el-tabs__item {
        line-height: 57px;
        height: 57px;
      }
    }
    .el-tabs__nav-wrap::after {
      background-color: transparent;
    }
  }
  .keywords {
    .el-card__body {
      padding-bottom: 10px;
      max-height: 480px;
      box-sizing: border-box;
      overflow: auto;
      /*定义滚动条高宽及背景 高宽分别对应横竖滚动条的尺寸*/
      &::-webkit-scrollbar {
        background-color: transparent;
        width: 6px;
        height: 6px;
      }
      &:hover::-webkit-scrollbar {
        background-color: rgb(241, 241, 241);
      }
      /*定义滚动条轨道 内阴影+圆角*/
      &::-webkit-scrollbar-track {
        background-color: transparent;
      }
      /*定义滑块 内阴影+圆角*/
      &::-webkit-scrollbar-thumb {
        background-color: transparent;
        border-radius: 3px;
      }
      &:hover::-webkit-scrollbar-thumb {
        background-color: rgb(193, 193, 193);
      }
      &:hover::-webkit-scrollbar-thumb::hover {
        background-color: rgb(168, 168, 168);
      }
    }
    a {
      margin-right: 10px;
      margin-bottom: 10px;
      display: inline-block;
      &:hover .el-tag--plain {
        background-color: #409eff;
        border-color: #409eff;
        color: #fff;
      }
    }
  }
}
@media screen and (max-width: $mobile-width) {
  .page-category {
    .el-col-18 {
      width: 100% !important;
    }
    .item-row {
      padding-bottom: 7px;
      .item-name {
        width: 50px;
      }
      a {
        margin-right: 10px;
        margin-bottom: 10px;
      }
    }
    .doc-list {
      .el-card__header {
        padding: 0 10px;
        .el-tabs__item {
          line-height: 40px;
          height: 40px;
          padding: 0 10px;
        }
      }
      .el-rate__icon {
        font-size: 15px;
      }
    }
  }
}
</style>
<style scoped lang="scss">
@media screen and (max-width: $mobile-width) {
  :deep(.com-document-list) {
    // h3 {
    //   a {
    //     white-space: inherit;
    //     overflow: auto;
    //     text-overflow: inherit;
    //   }
    // }
    .doc-cover {
      width: 25%;
      padding-right: 5px !important;
      .el-image {
        border: 1px solid #efefef;
      }
    }
    .el-col-20 {
      width: 75%;
    }
    .doc-desc {
      // display: none;
      width: 100%;
      font-size: 14px;
      -webkit-line-clamp: 2;
      height: 48px;
      line-height: 160%;
      padding-top: 8px;
    }
    .doc-info {
      font-size: 12px;
      .el-rate {
        float: right;
      }
      .float-right {
        float: left;
      }
    }
  }
}
</style>
