<template>
  <div class="page page-article">
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card shadow="never">
          <div slot="header">
            <h1>{{ article.title }}</h1>
          </div>
          <div class="help-block text-muted article-info">
            <!-- 如果没有作者，则默认显示网站名称 -->
            <span><i class="el-icon-user"></i> {{ article.author }}</span>
            <span
              ><i class="el-icon-view"></i>
              {{ article.view_count || 0 }} 阅读</span
            >
            <span class="float-right"
              ><i class="el-icon-time"></i> 最后更新:
              {{ formatDate(article.updated_at) }}
            </span>
          </div>
          <article class="mgt-20px markdown-body">
            <!-- eslint-disable-next-line vue/no-v-html -->
            <div data-slate-editor v-html="article.content"></div>
          </article>
        </el-card>
      </el-col>
      <el-col :span="6" class="article-list">
        <el-card shadow="never">
          <div slot="header">相关链接</div>
          <nuxt-link to="/" class="el-link el-link--default"
            >关于我们</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >联系我们</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >免责声明</nuxt-link
          >
          <!-- 如果文章数不多，不显示。这里根据最后更新时间排序 -->
          <el-pagination
            class="mgt-20px"
            :current-page="1"
            :page-size="10"
            layout="total, prev, next"
            :total="400"
          >
          </el-pagination>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getArticle } from '~/api/article'
import { formatDate } from '~/utils/utils'
export default {
  name: 'PageArticle',
  components: {},
  data() {
    return {
      id: this.$route.params.id,
      article: {},
      editor: null,
      editorConfig: {
        readOnly: true,
      },
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  async created() {
    const res = await getArticle({ identifier: this.$route.params.id })
    if (res.status === 200) {
      this.article = res.data
    } else {
      this.$message.error(res.data.message || '查询失败')
      this.$router.push('/404')
    }
  },
  methods: {
    formatDate,
    onCreated(editor) {
      this.editor = Object.seal(editor) // 一定要用 Object.seal() ，否则会报错
      this.editor.on('fullScreen', () => {
        this.isEditorFullScreen = true
      })
      this.editor.on('unFullScreen', () => {
        this.isEditorFullScreen = false
      })
    },
  },
}
</script>
<style lang="scss">
.page-article {
  .el-card__header {
    h1 {
      font-size: 16px;
      margin: 0;
    }
  }
  [data-w-e-type='todo'] {
    input {
      margin-right: 5px;
    }
  }
  .article-info {
    span {
      margin-right: 10px;
    }
  }
  .article-list {
    .el-card__body {
      padding-top: 10px;
    }
    a {
      display: block;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      height: 40px;
      line-height: 40px;
      border-bottom: 1px dashed #efefef;
    }
  }
  article {
    img {
      max-width: 100%;
    }
    .w-e-text-container [data-slate-editor] blockquote {
      border-left-width: 4px !important;
    }
    line-height: 180%;
    blockquote {
      padding: 10px;
      color: #777;
      font-size: 0.95em;
      background-color: #f6f8fa;
    }
  }
}
</style>
