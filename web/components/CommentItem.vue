<template>
  <el-row :class="`com-comment-item com-comment-item-${size}`">
    <el-col :span="isMobile ? 3 : 2">
      <nuxt-link :to="{ name: 'user-id', params: { id: user.id } }">
        <user-avatar
          v-if="isMobile"
          :size="size == 'small' ? 32 : 36"
          :user="comment.user"
        />
        <user-avatar
          v-else
          :size="size == 'small' ? 40 : 48"
          :user="comment.user"
        />
      </nuxt-link>
    </el-col>
    <el-col :span="isMobile ? 21 : 22">
      <div class="username">
        <nuxt-link
          class="el-link el-link--default"
          :to="{ name: 'user-id', params: { id: comment.user_id } }"
          >{{ comment.user.username }}</nuxt-link
        >
      </div>
      <div class="comment-content">
        <!-- eslint-disable-next-line vue/no-v-html -->
        <span v-html="comment.reply_user" />
        {{ comment.content }}
      </div>
      <div class="comment-action">
        <el-row class="help-block">
          <el-col :span="12">
            <el-tooltip
              :content="formatDatetime(comment.created_at)"
              placement="right"
            >
              <small class="text-muted">
                <i class="el-icon-time"></i>
                {{ formatRelativeTime(comment.created_at) }}
              </small>
            </el-tooltip>
          </el-col>
          <el-col :span="12" class="text-right">
            <el-button
              type="text"
              size="small"
              icon="el-icon-chat-dot-square"
              @click="reply"
              >回复</el-button
            >
          </el-col>
        </el-row>
      </div>
      <form-comment
        v-if="replyComment"
        :document-id="comment.document_id"
        :parent-id="comment.id"
        :placeholder="`回复 ${comment.user.username}`"
        @success="commentSuccess"
      />
      <slot></slot>
    </el-col>
  </el-row>
</template>
<script>
import { mapGetters } from 'vuex'
import UserAvatar from '~/components/UserAvatar.vue'
import { formatRelativeTime, formatDatetime } from '~/utils/utils'

export default {
  name: 'CommentItem',
  components: { UserAvatar },
  props: {
    size: {
      type: String,
      default: 'default', // default、small
    },
    comment: {
      type: Object,
      default: () => ({
        id: 0,
        parent_id: 0,
        user_id: 0,
        username: '匿名',
        avatar: '',
        group_id: 0,
        verify_status: 0,
        content: '内容加载中...',
        created_at: '0000-00-00',
      }),
    },
  },
  data() {
    return {
      replyComment: false,
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
  },
  methods: {
    formatRelativeTime,
    formatDatetime,
    reply() {
      this.replyComment = !this.replyComment
    },
    commentSuccess() {
      this.$emit('success')
    },
  },
}
</script>
<style lang="scss" scoped>
.com-comment-item {
  font-size: 14px;
  .comment-content {
    margin-top: 10px;
    margin-bottom: 10px;
    background-color: #f5f7f8;
    border-radius: 4px;
    padding: 20px;
    box-sizing: border-box;
    color: #565656;
    span {
      position: relative;
      top: -2px;
    }
  }
  .username a {
    font-weight: 400;
    font-size: 1.2em;
  }
}
.com-comment-item-small {
  font-size: 13px;
  .el-col-2 {
    width: 7%;
  }
  .el-col-22 {
    width: 93%;
  }
  .comment-content {
    padding: 15px;
  }
}
@media screen and (max-width: $mobile-width) {
}
</style>
