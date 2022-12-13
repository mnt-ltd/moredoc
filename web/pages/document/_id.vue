<template>
  <div class="page page-document">
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <h1>
              <img :src="`/static/images/${document.ext}_24.png`" alt="" />
              {{ document.title }}
            </h1>
            <el-breadcrumb separator-class="el-icon-arrow-right">
              <el-breadcrumb-item>
                <nuxt-link to="/"
                  ><i class="el-icon-s-home"></i> 首页</nuxt-link
                >
              </el-breadcrumb-item>
              <el-breadcrumb-item
                v-for="breadcrumb in breadcrumbs"
                :key="'bread-' + breadcrumb.id"
              >
                <nuxt-link :to="`/category/${breadcrumb.id}`">{{
                  breadcrumb.title
                }}</nuxt-link>
              </el-breadcrumb-item>
              <el-breadcrumb-item>文档阅览</el-breadcrumb-item>
            </el-breadcrumb>
            <div class="float-right doc-info">
              <span
                ><i class="el-icon-files"></i>
                {{ document.pages || '-' }} 页</span
              >
              <span
                ><i class="el-icon-download"></i>
                {{ document.download_count || 0 }} 下载</span
              >
              <span
                ><i class="el-icon-view"></i>
                {{ document.view_count || 0 }} 浏览</span
              >
              <span
                ><i class="el-icon-chat-dot-square"></i>
                {{ document.comment_count || 0 }} 评论</span
              >
              <span
                ><i class="el-icon-star-off"></i>
                {{ document.favorite_count || 0 }} 收藏</span
              >
              <span
                ><el-rate
                  v-model="document.score"
                  disabled
                  show-score
                  text-color="#ff9900"
                  score-template="{value}"
                >
                </el-rate
              ></span>
            </div>
          </div>
          <div ref="docPages" class="doc-pages">
            <el-image
              v-for="page in pages"
              :key="page.src"
              :src="page.src"
              :alt="page.alt"
              lazy
              class="doc-page"
              :style="{
                width: pageWidth + 'px',
                height: pageHeight + 'px',
              }"
            ></el-image>
          </div>
          <div class="doc-page-more text-center">
            <div>下载文档到电脑，方便使用</div>
            <el-button type="primary" icon="el-icon-download">
              下载文档({{ formatBytes(document.size) }})</el-button
            >
            <div v-if="document.preview - pages.length > 0">
              还有 {{ document.preview - pages.length }} 页可预览，
              <span class="el-link el-link--primary" @click="continueRead"
                >继续阅读</span
              >
            </div>
          </div>
          <div>
            <div class="share-info">
              本文档由
              <nuxt-link
                :to="`/user/${document.user_id}`"
                class="el-link el-link--primary"
                >{{ document.user.username || '匿名用户' }}</nuxt-link
              >
              于
              <span class="text-muted">{{
                formatDatetime(document.created_at)
              }}</span>
              上传分享
            </div>
            <div class="btn-actions">
              <el-button type="primary" plain icon="el-icon-warning-outline"
                >举报</el-button
              >
              <el-button
                type="primary"
                icon="el-icon-download"
                class="float-right"
                >下载文档({{ formatBytes(document.size) }})</el-button
              >
              <el-button
                v-if="favorite.id > 0"
                type="primary"
                plain
                class="float-right"
                icon="el-icon-star-on"
                @click="deleteFavorite"
                >取消收藏</el-button
              >
              <el-button
                v-else
                type="primary"
                class="float-right"
                icon="el-icon-star-off"
                @click="createFavorite"
                >收藏</el-button
              >
            </div>
          </div>
        </el-card>
        <el-card
          v-if="document.id > 0"
          ref="commentBox"
          shadow="never"
          class="mgt-20px"
        >
          <FormComment :document-id="document.id" @success="commentSuccess" />
          <comment-list ref="commentList" :document-id="document.id" />
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never">
          <div slot="header">分享用户</div>
          <user-card :hide-actions="true" :user="document.user" />
        </el-card>
        <el-card shadow="never" class="mgt-20px relate-docs">
          <div slot="header">相关文档</div>
          <document-simple-list :docs="docs" />
        </el-card>
      </el-col>
    </el-row>
    <div class="fixed-buttons">
      <el-card shadow="never">
        <el-row>
          <el-col :span="18">
            <el-button-group class="btn-actions">
              <el-tooltip content="全屏阅读">
                <el-button icon="el-icon-full-screen"></el-button>
              </el-tooltip>
              <el-tooltip :content="favorite.id > 0 ? '取消收藏' : '收藏文档'">
                <el-button
                  v-if="favorite.id > 0"
                  icon="el-icon-star-on"
                  @click="deleteFavorite"
                ></el-button>
                <el-button
                  v-else
                  icon="el-icon-star-off"
                  @click="createFavorite"
                ></el-button>
              </el-tooltip>
              <el-tooltip content="缩小">
                <el-button icon="el-icon-zoom-out"></el-button>
              </el-tooltip>
              <el-tooltip content="放大">
                <el-button icon="el-icon-zoom-in"></el-button>
              </el-tooltip>
              <el-tooltip content="上一页">
                <el-button
                  icon="el-icon-arrow-up"
                  @click="prevPage"
                ></el-button>
              </el-tooltip>
              <el-tooltip content="当前页数/总页数">
                <el-button>{{ currentPage }}/{{ document.pages }}</el-button>
              </el-tooltip>
              <el-tooltip content="下一页">
                <el-button
                  icon="el-icon-arrow-down"
                  @click="nextPage"
                ></el-button>
              </el-tooltip>
            </el-button-group>
            <el-button
              class="btn-comment"
              icon="el-icon-chat-dot-square"
              @click="gotoComment"
              >文档点评</el-button
            >
            <el-button-group class="float-right">
              <el-button type="primary" icon="el-icon-coin" class="btn-coin"
                >{{ document.price || 0 }} 个魔豆</el-button
              >
              <el-button type="primary" icon="el-icon-download"
                >下载文档({{ formatBytes(document.size) }})</el-button
              >
            </el-button-group>
          </el-col>
          <el-col :span="6" class="text-right">
            <el-button icon="el-icon-top" @click="scrollTop"
              >回到顶部</el-button
            >
          </el-col>
        </el-row>
      </el-card>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import DocumentSimpleList from '~/components/DocumentSimpleList.vue'
import { getDocument } from '~/api/document'
import { getFavorite, createFavorite, deleteFavorite } from '~/api/favorite'
import { formatDatetime, formatBytes } from '~/utils/utils'
import FormComment from '~/components/FormComment.vue'
import CommentList from '~/components/CommentList.vue'
export default {
  name: 'PageDocument',
  components: { DocumentSimpleList, FormComment, CommentList },
  data() {
    return {
      docs: [],
      user: {
        id: 0,
      },
      document: {
        id: 0,
        score: 4.0,
        user: {
          id: 0,
        },
        attachment: {
          hash: '',
        },
      },
      documentId: parseInt(this.$route.params.id) || 0,
      pages: [],
      pageHeight: 0,
      pageWidth: 0,
      currentPage: 1,
      breadcrumbs: [],
      favorite: {
        id: 0,
      },
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  computed: {
    ...mapGetters('category', ['categoryMap']),
  },
  created() {
    Promise.all([this.getDocument(), this.getFavorite()])
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
  },
  destroyed() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    formatDatetime,
    formatBytes,
    async getDocument() {
      const res = await getDocument({
        id: this.documentId,
        with_author: true,
      })
      if (res.status === 200) {
        const doc = res.data || {}
        doc.score = parseFloat(doc.score) / 100 || 4.0

        if (!doc.preview) {
          doc.preview = doc.pages
        }

        // 限定每次预览页数
        let preview = 2
        if (doc.preview < preview) {
          preview = doc.preview
        }

        // 限定预览页数，拼装图片链接
        const pages = []
        for (let i = 1; i <= preview; i++) {
          pages.push({
            src: `/view/page/${doc.attachment.hash}/${i}.gzip.svg`,
            alt: `${doc.title} 第${i + 1}页`,
          })
        }

        this.breadcrumbs = (doc.category_id || []).map((id) => {
          return this.categoryMap[id]
        })

        doc.ext = doc.ext.replace('.', '')
        this.pages = pages
        this.document = doc
        this.pageWidth = this.$refs.docPages.offsetWidth
        this.pageHeight =
          (this.$refs.docPages.offsetWidth / doc.width) * doc.height
      } else {
        console.log(res)
        this.$message.error(res.data.message)
      }
    },
    handleScroll() {
      const scrollTop =
        document.documentElement.scrollTop || document.body.scrollTop
      // 还有5像素的border
      let currentPage = Math.round(scrollTop / (this.pageHeight + 5)) + 1
      if (currentPage > this.pages.length) {
        currentPage = this.pages.length
      }
      this.currentPage = currentPage
    },
    scrollTop() {
      this.scrollTo(0)
    },
    gotoComment() {
      try {
        this.scrollTo(this.$refs.commentBox.$el.offsetTop)
      } catch (error) {
        console.log('gotoComment', error)
      }
    },
    commentSuccess() {
      this.$refs.commentList.getComments()
    },
    prevPage() {
      if (this.currentPage > 1) {
        const currentPage = this.currentPage - 1
        this.scrollToPage(currentPage)
      }
    },
    nextPage() {
      if (this.currentPage < this.document.preview) {
        const currentPage = this.currentPage + 1
        if (currentPage > this.pages.length) {
          this.continueRead()
        }
        this.scrollToPage(currentPage)
      }
    },
    scrollToPage(page) {
      const scrollTop = (page - 1) * this.pageHeight
      this.scrollTo(scrollTop)
    },
    scrollTo(position) {
      document.scrollingElement.scrollTo({
        top: position,
        behavior: 'smooth',
      })
    },
    async getFavorite() {
      const res = await getFavorite({
        document_id: this.documentId,
      })
      if (res.status === 200) {
        this.favorite = res.data || { id: 0 }
      }
    },
    // 取消收藏
    async deleteFavorite() {
      const res = await deleteFavorite({ id: this.favorite.id })
      if (res.status === 200) {
        this.$message.success('取消收藏成功')
        this.favorite = { id: 0 }
      } else {
        this.$message.error(res.data.message)
      }
    },
    // 添加收藏
    async createFavorite() {
      const res = await createFavorite({
        document_id: this.documentId,
      })
      if (res.status === 200) {
        this.$message.success('收藏成功')
        this.favorite = res.data
      } else {
        this.$message.error(res.data.message)
      }
    },
    continueRead() {
      let end = this.pages.length + 5
      if (end > this.document.preview) {
        end = this.document.preview
      }
      for (let i = this.pages.length + 1; i <= end; i++) {
        this.pages.push({
          src: `/view/page/${this.document.attachment.hash}/${i}.gzip.svg`,
          alt: `${this.document.title} 第${i + 1}页`,
        })
      }
    },
  },
}
</script>
<style lang="scss">
.page-document {
  .relate-docs {
    .el-card__body {
      padding-top: 10px;
    }
  }
  h1 {
    margin: 0;
    img {
      position: relative;
      top: 3px;
    }
  }
  .el-breadcrumb {
    font-weight: normal;
    margin-top: 12px;
    color: #565656;
    .el-breadcrumb__inner a,
    .el-breadcrumb__inner.is-link {
      font-weight: normal;
    }
    .el-breadcrumb__separator[class*='icon'] {
      margin: 0 3px;
    }
    .el-breadcrumb__inner {
      color: #666;
    }
  }
  .doc-info {
    font-weight: normal;
    position: relative;
    top: -16px;
    font-size: 14px;
    color: #bbb;
    & > span {
      margin-left: 8px;
    }
    .el-rate {
      position: relative;
      top: -2px;
    }
  }
  .doc-pages {
    .doc-page {
      display: block;
      width: 100%;
      box-sizing: border-box;
      border: 5px solid $background-grey-light;
      border-bottom: 0;
      min-height: 500px;
      &:last-child {
        border-bottom: 5px solid $background-grey-light;
      }
      img {
        width: 100%;
        background-color: #fff;
      }
    }
  }
  .doc-page-more {
    padding: 30px 0;
    border: 5px solid $background-grey-light;
    border-top: 0;
    color: #565656;
    .el-button {
      margin: 10px 0;
    }
  }
  .share-info {
    font-size: 15px;
    color: #666;
    margin: 15px 0;
    .el-link {
      top: -2px;
    }
  }

  .fixed-buttons {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    z-index: 100;
    width: 100%;
    background-color: #ecf0f1;
    height: 50px;
    [class*=' el-icon-'],
    [class^='el-icon-'] {
      font-weight: bold;
    }
    .el-card {
      border-radius: 0;
      background-color: transparent;
      width: $default-width;
      max-width: $max-width;
      margin: 0 auto;
      .el-card__body {
        padding: 0;
      }
      .el-button {
        border: 0;
        border-radius: 0;
        padding: 18px 20px;
      }
      .btn-comment {
        top: 1px;
        position: relative;
        background-color: transparent;
        &:hover {
          background-color: #ecf5ff;
        }
      }
      .btn-actions .el-button {
        background-color: transparent;
        &:hover {
          background-color: #ecf5ff;
        }
      }
      .btn-coin {
        background-color: transparent;
        color: #606266;
        cursor: auto;
      }
    }
  }
}
</style>
