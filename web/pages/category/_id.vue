<template>
  <div class="page page-category">
    <el-row>
      <el-col :span="24">
        <el-card shadow="never">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item>
              <nuxt-link to="/"><i class="el-icon-s-home"></i> 首页</nuxt-link>
            </el-breadcrumb-item>
            <el-breadcrumb-item
              v-for="item in breadcrumbs"
              :key="'bread-' + item.id"
              :to="`/category/${item.id}`"
              >{{ item.title }}</el-breadcrumb-item
            >
          </el-breadcrumb>
        </el-card>
      </el-col>
    </el-row>
    <el-row :gutter="20" class="mgt-20px">
      <el-col :span="18">
        <el-card shadow="never" class="doc-list">
          <div slot="header">
            <el-tabs v-model="query.sort" @tab-click="sortClick">
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
            <document-list :documents="documents" />
          </div>
          <el-pagination
            v-if="total > 0"
            :current-page="query.page"
            :page-size="size"
            layout="total,  prev, pager, next, jumper"
            :total="total"
            @current-change="pageChange"
          >
          </el-pagination>
        </el-card>
      </el-col>

      <el-col :span="6">
        <el-card shadow="never" class="categories">
          <div slot="header">
            <el-row>
              <el-col :span="12" class="header-title">
                {{ breadcrumbs[0].title }}
              </el-col>
              <el-col :span="12">
                <el-input v-model="filterText" placeholder="分类过滤">
                </el-input>
              </el-col>
            </el-row>
          </div>
          <el-tree
            ref="tree"
            :data="trees"
            :props="defaultProps"
            :indent="8"
            node-key="id"
            :default-expanded-keys="defaultExpandedKeys"
            highlight-current
            :filter-node-method="filterTree"
            @node-click="handleNodeClick"
          ></el-tree>
        </el-card>
        <el-card shadow="never" class="mgt-20px keywords" ref="keywords">
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
      defaultExpandedKeys: [],
      defaultProps: {
        children: 'children',
        label: 'title',
      },
      query: {
        id: 0,
        sort: 'latest',
        page: 1,
      },
      size: 10,
      breadcrumbs: [], // 面包屑
      trees: [],
      documents: [],
      categoryId: parseInt(this.$route.params.id) || 0,
      total: 0,
      keywords: [],
      loading: false,
      cardOffsetTop: 0,
      cardWidth: 0,
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  computed: {
    ...mapGetters('category', ['categories', 'categoryTrees', 'categoryMap']),
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
  created() {
    const breadcrumbs = []
    let category = this.categoryMap[this.categoryId]
    if (category) {
      breadcrumbs.push(category)
      while (category.parent_id) {
        category = this.categoryMap[category.parent_id]
        if (category) {
          breadcrumbs.splice(0, 0, category)
        }
      }
    }
    this.breadcrumbs = breadcrumbs

    try {
      this.trees =
        this.categoryTrees.find((x) => x.id === breadcrumbs[0].id).children ||
        []
    } catch (error) {
      console.log(error)
    }

    this.setQuery()
    this.setDefaultExpandedKeys()
    this.loadData()
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
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
      this.query.sort = this.$route.query.sort || 'latest'
      this.query.page = parseInt(this.$route.query.page) || 1
    },
    setDefaultExpandedKeys() {
      const defaultExpandedKeys = []
      let category = this.breadcrumbs[this.breadcrumbs.length - 1] || {
        id: 0,
        title: '全部',
      }
      if (category) {
        defaultExpandedKeys.push(category.id)
        while (category.parent_id) {
          defaultExpandedKeys.unshift(category.parent_id)
          category = this.categoryMap[category.parent_id]
        }
      }
      this.defaultExpandedKeys = defaultExpandedKeys
    },
    handleScroll() {
      const scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop
      const keywords = this.$refs.keywords.$el
      if (keywords) {
        if (this.cardWidth === 0) {
          this.cardWidth = keywords.offsetWidth
          this.cardOffsetTop = keywords.offsetTop
        }

        if (scrollTop > this.cardOffsetTop) {
          keywords.style.position = 'fixed'
          keywords.style.top = '60px'
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
          break
      }
      const res = await listDocument({
        order,
        page: this.query.page,
        size: this.size,
        category_id: this.categoryId,
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
    },
  },
}
</script>
<style lang="scss">
.page-category {
  .categories {
    .el-card__header {
      padding-top: 0;
      padding-bottom: 0;
      .header-title {
        line-height: 56px;
      }
      .el-input {
        top: 10px;
        .el-input__inner {
          height: 35px;
          line-height: 35px;
        }
      }
    }
    .el-tree-node__content {
      height: 35px;
    }
    [role='treeitem'][aria-expanded='true'] > .el-tree-node__content {
      background-color: #f5f7fa;
      color: #409eff;
      font-weight: bold;
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
</style>
