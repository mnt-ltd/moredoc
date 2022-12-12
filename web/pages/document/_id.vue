<template>
  <div class="page page-document">
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card shadow="never">
          <div slot="header" class="clearfix">
            <h1>
              <img src="/static/images/pdf_24.png" alt="" />
              厦门才茂工业级CM520-82系列技术参数1
            </h1>
            <el-breadcrumb separator-class="el-icon-arrow-right">
              <el-breadcrumb-item>活动管理</el-breadcrumb-item>
              <el-breadcrumb-item>活动列表</el-breadcrumb-item>
              <el-breadcrumb-item>活动详情</el-breadcrumb-item>
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
          <div class="doc-pages">
            <div
              v-for="(page, index) in pages"
              :key="'page' + index"
              class="doc-page"
            >
              <img
                :src="page"
                :style="`min-height: ${document.height}px`"
                alt=""
              />
            </div>
          </div>
          <div class="doc-page-more text-center">
            <div>下载文档到电脑，方便使用</div>
            <el-button type="primary" icon="el-icon-download">
              下载文档</el-button
            >
            <div>
              还有 10 页可预览，
              <span class="el-link el-link--primary">继续阅读</span>
            </div>
          </div>
          <div>
            <div class="share-info">
              本文档由
              <nuxt-link to="/" class="el-link el-link--primary"
                >皇虫</nuxt-link
              >
              于 <span class="text-muted">2020-01-17 09:47:50</span> 上传分享
            </div>
            <div class="btn-actions">
              <el-button type="primary" plain icon="el-icon-warning-outline"
                >举报</el-button
              >
              <el-button
                type="primary"
                icon="el-icon-download"
                class="float-right"
                >下载文档(12.23MB)</el-button
              >
              <el-button
                type="primary"
                class="float-right"
                icon="el-icon-star-off"
                >收藏</el-button
              >
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="never">
          <div slot="header">分享用户</div>
          <user-card :hide-actions="true" :user="user" />
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
              <el-tooltip content="收藏/取消收藏文档">
                <el-button icon="el-icon-star-off"></el-button>
              </el-tooltip>
              <el-tooltip content="缩小">
                <el-button icon="el-icon-zoom-out"></el-button>
              </el-tooltip>
              <el-tooltip content="放大">
                <el-button icon="el-icon-zoom-in"></el-button>
              </el-tooltip>
              <el-tooltip content="上一页">
                <el-button icon="el-icon-arrow-up"></el-button>
              </el-tooltip>
              <el-tooltip content="当前页数/总页数">
                <el-button>5/12</el-button>
              </el-tooltip>
              <el-tooltip content="下一页">
                <el-button icon="el-icon-arrow-down"></el-button>
              </el-tooltip>
            </el-button-group>
            <el-button class="btn-comment" icon="el-icon-chat-dot-square"
              >文档点评</el-button
            >
            <el-button-group class="float-right">
              <el-button type="primary" icon="el-icon-coin" class="btn-coin"
                >0个金币</el-button
              >
              <el-button type="primary" icon="el-icon-download"
                >下载文档(12.23MB)</el-button
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
import DocumentSimpleList from '~/components/DocumentSimpleList.vue'
import { getDocument } from '~/api/document'
export default {
  name: 'PageDocument',
  components: { DocumentSimpleList },
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
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  created() {
    this.getDocument()
  },
  methods: {
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

        // 限定预览页数，拼装图片链接
        const pages = []
        for (let i = 1; i <= doc.preview; i++) {
          pages.push(`/view/page/${doc.attachment.hash}/${i}.gzip.svg`)
        }
        this.pages = pages
        this.document = doc
      } else {
        console.log(res)
        this.$message.error(res.data.message)
      }
    },
    scrollTop() {
      this.scrollToTop(300)
    },
    scrollToTop(duration) {
      // cancel if already on top
      if (document.scrollingElement.scrollTop === 0) return

      const cosParameter = document.scrollingElement.scrollTop / 2
      let scrollCount = 0
      let oldTimestamp = null

      function step(newTimestamp) {
        if (oldTimestamp !== null) {
          // if duration is 0 scrollCount will be Infinity
          scrollCount += (Math.PI * (newTimestamp - oldTimestamp)) / duration
          if (scrollCount >= Math.PI)
            return (document.scrollingElement.scrollTop = 0)
          document.scrollingElement.scrollTop =
            cosParameter + cosParameter * Math.cos(scrollCount)
        }
        oldTimestamp = newTimestamp
        window.requestAnimationFrame(step)
      }
      window.requestAnimationFrame(step)
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
      border: 5px solid $background-grey-light;
      border-bottom: 0;
      &:last-child {
        border-bottom: 5px solid $background-grey-light;
      }
      img {
        width: 100%;
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
