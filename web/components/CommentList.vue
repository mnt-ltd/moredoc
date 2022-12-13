<template>
  <div class="com-comment-list">
    <comment-item
      v-for="comment in comments"
      :key="'comment-' + comment.id"
      :comment="comment"
      @success="commentSuccess"
    >
      <comment-item
        v-for="child in comment.children"
        :key="'comment-' + child.id"
        class="comment-child"
        :comment="child"
        :size="'small'"
        @success="commentSuccess"
      ></comment-item>
    </comment-item>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
import CommentItem from '~/components/CommentItem.vue'
import { listComment } from '~/api/comment'

export default {
  name: 'CommentList',
  components: { CommentItem },
  props: {
    documentId: {
      type: Number,
      default: 0,
    },
    parentId: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      comments: [],
      req: {
        document_id: this.documentId,
        parent_id: this.parentId,
      },
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
  },
  watch: {
    documentId: {
      handler(val) {
        this.req.document_id = val
      },
      immediate: true,
    },
    parentId: {
      handler(val) {
        this.req.parent_id = val
      },
      immediate: true,
    },
  },
  created() {
    this.getComments()
  },
  methods: {
    //  获取文章评论列表
    async getComments() {
      if (!this.req.document_id) return
      const res = await listComment({
        document_id: this.documentId,
        order: 'id asc',
      })
      if (res.status === 200) {
        this.comments = this.comments2tree(res.data.comment || [])
      }
    },
    commentSuccess() {
      this.getComments()
    },
    comments2tree(comments) {
      const tree = []
      const map = {}
      comments.forEach((comment) => {
        map[comment.id] = comment
      })
      comments.forEach((comment) => {
        // 寻找最顶层父级
        let parent = map[comment.parent_id]
        let replyUser = ''
        if (comment.parent_id && parent) {
          try {
            replyUser = `<a href="/user/${parent.user.id}" class="el-link el-link--primary" target="blank">@${parent.user.username}</a>`
          } catch (error) {}
        }
        while (parent && parent.parent_id) {
          parent = map[parent.parent_id]
        }
        comment.reply_user = replyUser
        if (parent) {
          ;(parent.children || (parent.children = [])).push(comment)
        } else {
          tree.push(comment)
        }
      })
      return tree
    },
  },
}
</script>
<style lang="scss" scoped>
.com-comment-list {
  & > .el-row {
    margin-top: 20px;
    border-bottom: 1px solid #efefef;
    padding-bottom: 10px;
  }
  & > .el-row:first-of-type {
    margin-top: 0;
  }
}
</style>
