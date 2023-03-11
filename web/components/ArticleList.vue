<template>
  <div class="com-article-list">
    <ul v-if="articles.length > 0">
      <li v-for="article in articles" :key="'a' + article.id">
        <nuxt-link
          :to="`/article/${article.identifier}`"
          class="el-link el-link--default"
          :title="article.title"
        >
          <span v-if="withHtml" v-html="article.title"></span>
          <template v-else>{{ article.title }}</template>
        </nuxt-link>
        <div v-if="withDescription" class="description">
          <span v-if="withHtml" v-html="article.description"></span>
          <template v-else>{{ article.description }}</template>
        </div>
        <div class="help-block">
          <el-tooltip
            :content="'发布时间:' + formatDatetime(article.created_at)"
          >
            <span title="发布时间">
              <i class="el-icon-time"></i>
              {{ formatRelativeTime(article.created_at) }}
            </span>
          </el-tooltip>
          <span
            ><i class="el-icon-view"></i>
            {{ article.view_count || 0 }} 浏览</span
          >
          <span
            ><i class="el-icon-user"></i>
            {{ article.autor || settings.system.sitename || '-' }}</span
          >
          <!-- <span
            ><i class="el-icon-chat-dot-square"></i>
            {{ article.comment_count || 0 }} 评论</span
          >
          <span
            ><i class="el-icon-star-off"></i>
            {{ article.favorite_count || 0 }} 收藏</span
          > -->
        </div>
      </li>
    </ul>
    <div v-else>
      <el-empty description="暂无数据"></el-empty>
    </div>
  </div>
</template>
<script>
import { formatDatetime, formatRelativeTime } from '~/utils/utils'
import { mapGetters } from 'vuex'
export default {
  props: {
    articles: {
      type: Array,
      default: () => [],
    },
    withDescription: {
      type: Boolean,
      default: false,
    },
    withHtml: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  methods: {
    formatRelativeTime,
    formatDatetime,
  },
}
</script>
<style lang="scss">
.com-article-list {
  mark {
    background-color: transparent;
    color: red;
  }
}
</style>
<style lang="scss" scoped>
.com-article-list {
  ul,
  li {
    margin: 0;
    padding: 0;
    list-style: none;
  }
  li {
    border-bottom: 1px dashed #efefef;
    padding: 20px 0;
    &:last-child {
      border-bottom: none;
    }
    .help-block {
      color: #999;
      font-size: 13px;
      margin-top: 10px;
      span {
        margin-right: 20px;
      }
    }
    a {
      font-size: 18px;
      color: #222;
      display: inline-block;
      max-width: 100%;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
  .no-articles {
    text-align: center;
    font-size: 15px;
    color: #ccc;
    padding: 40px 0;
  }
  .description {
    margin-top: 10px;
    color: #777;
    font-size: 15px;
    line-height: 24px;
    max-height: 72px;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
  }
}
</style>
