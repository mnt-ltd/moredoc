<template>
  <div class="page page-article-index">
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card shadow="never">
          <div slot="header">
            <el-breadcrumb separator="/">
              <el-breadcrumb-item>
                <nuxt-link to="/"><i class="fa fa-home"></i> 首页</nuxt-link>
              </el-breadcrumb-item>
              <el-breadcrumb-item>文章列表</el-breadcrumb-item>
            </el-breadcrumb>
          </div>
          <article-list :articles="articles" />
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
      <el-col :span="6">
        <el-card shadow="never" class="popular">
          <div slot="header">热门</div>
          <div v-if="populars.length > 0">
            <nuxt-link
              :to="`/article/${article.identifier}`"
              v-for="article in populars"
              :key="'article-' + article.id"
              :title="article.title"
              class="el-link el-link--default"
              >{{ article.title }}</nuxt-link
            >
          </div>
          <el-empty v-else> </el-empty>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
import { listArticle } from '~/api/article'
export default {
  name: 'PageArticleIndex',
  data() {
    return {
      total: 0,
      size: 10,
      query: {
        page: 1,
      },
      populars: [],
      articles: [],
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    $route: {
      handler() {
        let page = this.$route.query.page || 1
        this.query.page = parseInt(page) || 1
        this.getArticles()
      },
      immediate: true,
    },
  },
  async created() {
    await Promise.all([this.getPopulars()])
  },
  methods: {
    pageChange(page) {
      this.query.page = page
      this.$router.push({
        path: this.$route.path,
        query: this.query,
      })
    },
    async getArticles() {
      const res = await listArticle({ size: this.size, page: this.query.page })
      if (res.status !== 200) {
        this.$message.error(res.data.message || '获取文章列表失败')
        return
      }
      this.articles = res.data.article || []
      this.total = res.data.total || 0
    },
    async getPopulars() {
      const res = await listArticle({
        page: 1,
        size: 5,
        order: 'view_count desc',
      })
      if (res.status !== 200) {
        this.$message.error(res.data.message || '获取热门文章失败')
        return
      }
      this.populars = res.data.article || []
    },
  },
}
</script>
<style lang="scss">
.page-article-index {
  .el-card__body {
    padding-top: 0;
    padding-bottom: 0;
    .el-pagination {
      padding: 20px 0;
      border-top: 1px dashed #efefef;
    }
  }
  .popular {
    .el-card__header {
      padding: 14px 20px;
    }
    a {
      display: block;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      height: 40px;
      line-height: 40px;
      border-bottom: 1px dashed #efefef;
      &:last-of-type {
        border-bottom: 0;
      }
    }
  }
}

@media screen and (max-width: $mobile-width) {
  .page-article-index {
    .el-col {
      width: 100%;
    }
    .popular {
      margin-top: 15px;
    }
  }
}
</style>
