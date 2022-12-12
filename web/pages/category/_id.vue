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
              <el-tab-pane name="hot">
                <span slot="label"><i class="el-icon-view"></i> 浏览</span>
              </el-tab-pane>
              <el-tab-pane name="recommend">
                <span slot="label"><i class="el-icon-date"></i> 推荐</span>
              </el-tab-pane>
              <el-tab-pane name="favorite">
                <span slot="label"><i class="el-icon-date"></i> 收藏</span>
              </el-tab-pane>
              <el-tab-pane name="download">
                <span slot="label"><i class="el-icon-download"></i> 下载</span>
              </el-tab-pane>
              <el-tab-pane name="pages">
                <span slot="label"><i class="el-icon-files"></i> 页数</span>
              </el-tab-pane>
            </el-tabs>
          </div>
          <div class="doc-list-data">
            <ul>
              <li v-for="x in 10" :key="'x-' + x">
                <el-row :gutter="20">
                  <el-col :span="4" class="doc-cover">
                    <nuxt-link to="/document/">
                      <img
                        v-if="x % 3 == 0"
                        src="https://static.wenkuzhijia.cn/fe5642195d3060c51d12fccaa46f4c61.jpg"
                        alt=""
                        style="width: 100%"
                      />
                      <img
                        v-if="x % 3 == 1"
                        src="https://static.sitestack.cn/projects/entgo-0.11-zh/uploads/202210/171f825ac77e9e82.png/cover"
                        alt=""
                        style="width: 100%"
                      />
                      <img
                        v-if="x % 3 == 2"
                        src="https://static.sitestack.cn/projects/learn-go-with-tests-14.0-en/uploads/202210/171edb9f6a91aa73.png/cover"
                        alt=""
                        style="width: 100%"
                      />
                    </nuxt-link>
                  </el-col>
                  <el-col :span="20">
                    <h3>
                      <nuxt-link
                        to="/document/"
                        class="el-link el-link--primary"
                        ><img src="/static/images/pdf_24.png" alt="" />
                        啊实打实大</nuxt-link
                      >
                    </h3>
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
                        <span class="hidden-xs-only"
                          >| 2019-06-10 10:17</span
                        ></span
                      >
                    </div>
                    <div class="doc-desc">
                      GoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame不是一款WEBGoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame不是一款WEBGoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame不是一款WEBGoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame不是一款WEBGoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame不是一款WEBGoFrame是一款模块化、高性能、企业级的Go基础开发框架。GoFrame不是一款WEB/RPC框架，而是一款通用性的基础开发框架，是Golang标准库的一个增强扩展级，包含通用核心的基础开发组件，优点是实战化、模块化、文档全面、模块丰富、易用性高、通用性强、面向团队。
                    </div>
                  </el-col>
                </el-row>
              </li>
            </ul>
          </div>
          <el-pagination
            v-if="total > 0"
            :current-page="query.page"
            :page-size="10"
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
              <el-col :span="8" class="header-title">
                {{ breadcrumbs[0].title }}
              </el-col>
              <el-col :span="16">
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
        <el-card shadow="never" class="mgt-20px keywords">
          <div slot="header">
            <el-row>
              <el-col :span="8" class="header-title">关键词</el-col>
            </el-row>
          </div>
          <nuxt-link
            v-for="(keyword, index) in keywords"
            :key="'kw' + keyword"
            :to="{ path: '/search', query: { wd: keyword } }"
            :class="['tag-' + (index % 5)]"
          >
            <el-tag effect="plain"> {{ keyword }}</el-tag>
          </nuxt-link>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'PageCategory',
  data() {
    return {
      filterText: '',
      score: 4.5,
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
      breadcrumbs: [], // 面包屑
      trees: [],
      categoryId: parseInt(this.$route.params.id) || 0,
      total: 0,
      keywords: ['Java', '书栈网', 'PHP教程', 'MySQL'],
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
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
    loadData() {
      // TODO: 加载数据
      console.log('loadData')
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
  .doc-list-data {
    ul,
    li {
      list-style: none;
      padding: 0;
      margin: 0;
    }
    li {
      margin-bottom: 20px;
      padding-bottom: 18px;
      border-bottom: 1px dashed #ddd;
    }
    h3 {
      margin: 0;
      a {
        font-size: 18px;
        font-weight: 400;
        display: block;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        img {
          height: 18px;
          position: relative;
          top: 3px;
        }
      }
    }
    .doc-cover {
      img {
        width: 100%;
        border: 1px solid #efefef;
        border-radius: 3px;
      }
    }
    .doc-info {
      margin: 10px 0;
      font-size: 13px;
      color: #bbb;
      .float-right {
        position: relative;
        top: 16px;
      }
    }
    .doc-desc {
      font-size: 15px;
      color: #888;
      line-height: 180%;
      overflow: hidden;
      text-overflow: ellipsis;
      display: -webkit-box;
      -webkit-line-clamp: 3;
      max-height: 81px;
      -webkit-box-orient: vertical;
    }
  }
}
</style>
