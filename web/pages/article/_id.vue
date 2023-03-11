<template>
  <div class="page page-article">
    <el-row :gutter="20">
      <el-col :span="24">
        <el-card shadow="never">
          <div slot="header">
            <h1>{{ article.title }}</h1>
            <el-breadcrumb separator="/">
              <el-breadcrumb-item>
                <nuxt-link to="/"><i class="fa fa-home"></i> 首页</nuxt-link>
              </el-breadcrumb-item>
              <el-breadcrumb-item>
                <nuxt-link to="/article">文章列表</nuxt-link>
              </el-breadcrumb-item>
              <el-breadcrumb-item>文章详情</el-breadcrumb-item>
            </el-breadcrumb>
          </div>
          <div class="help-block text-muted article-info">
            <!-- 如果没有作者，则默认显示网站名称 -->
            <span
              ><i class="el-icon-user"></i>
              {{ article.author || settings.system.sitename || '-' }}</span
            >
            <span
              ><i class="el-icon-view"></i>
              {{ article.view_count || 0 }} 阅读</span
            >
            <span class="float-right"
              ><i class="el-icon-time"></i>
              <span class="hidden-xs-only">最后更新:</span>
              {{ formatDate(article.updated_at) }}
            </span>
          </div>
          <article class="mgt-20px markdown-body">
            <!-- eslint-disable-next-line vue/no-v-html -->
            <div data-slate-editor v-html="article.content"></div>
          </article>
        </el-card>
      </el-col>
      <!-- <el-col :span="6" class="article-list">
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
        </el-card>
      </el-col> -->
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getArticle } from '~/api/article'
import { formatDate } from '~/utils/utils'
export default {
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
      title: `${this.article.title} - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: this.article.keywords,
        },
        {
          hid: 'description',
          name: 'description',
          content: this.article.description,
        },
      ],
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
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
  .el-breadcrumb {
    margin-top: 15px;
    .el-breadcrumb__inner a,
    .el-breadcrumb__inner.is-link {
      color: #666;
    }
    .el-breadcrumb__item:last-child .el-breadcrumb__inner {
      color: #777;
    }
    .el-breadcrumb__inner a:hover,
    .el-breadcrumb__inner.is-link:hover {
      color: #409eff;
    }
  }
  .el-card__header {
    h1 {
      font-size: 24px;
      font-weight: 400;
      margin: 0;
      color: #111;
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
      &.float-right {
        margin-right: 0;
      }
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
