<template>
  <div class="com-form-comment">
    <el-form
      ref="form"
      :inline="true"
      :model="comment"
      :rules="rules"
      class="form-comment"
    >
      <el-form-item prop="content" class="comment-content">
        <el-input
          v-model="comment.content"
          type="textarea"
          :placeholder="placeholder"
          :autosize="{ minRows: 4, maxRows: 6 }"
        />
      </el-form-item>
      <el-form-item class="comment-btns">
        <el-row>
          <el-col :span="7"> 请文明评论，理性发言. </el-col>
          <el-col :span="17" class="text-right">
            <template v-if="captcha.enable">
              <el-form-item>
                <div class="captcha">
                  <div v-if="captcha.type == 'audio'">
                    <el-row :gutter="15">
                      <el-col :span="20">
                        <audio
                          controls="controls"
                          :src="captcha.captcha"
                        ></audio>
                      </el-col>
                      <el-col :span="4">
                        <el-tooltip placement="top" content="刷新语音验证码">
                          <el-button
                            icon="el-icon-refresh"
                            class="btn-audio-refresh"
                            @click="loadCaptcha"
                          ></el-button>
                        </el-tooltip>
                      </el-col>
                    </el-row>
                  </div>
                  <div v-else>
                    <el-tooltip placement="top" content="点击可刷新验证码">
                      <img
                        :src="captcha.captcha"
                        class="pointer"
                        @click="loadCaptcha"
                      />
                    </el-tooltip>
                  </div>
                </div>
              </el-form-item>
              <el-form-item
                prop="captcha"
                :rules="[
                  { required: true, trigger: 'blur', message: '请输入验证码' },
                ]"
              >
                <el-input
                  v-model="comment.captcha"
                  placeholder="请输入验证码"
                ></el-input>
              </el-form-item>
            </template>
            <el-form-item>
              <el-button
                type="primary"
                icon="el-icon-position"
                @click="submitForm('form')"
                >发布评论</el-button
              >
            </el-form-item>
          </el-col>
        </el-row>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { getUserCaptcha } from '~/api/user'
import { createComment } from '~/api/comment'
export default {
  name: 'FormComment',
  props: {
    articleId: {
      type: Number,
      default: 0,
    },
    parentId: {
      type: Number,
      default: 0,
    },
    placeholder: {
      type: String,
      default: '请输入评论内容',
    },
  },
  data() {
    return {
      comment: {
        article_id: this.articleId,
        parent_id: this.parentId,
        content: '',
        captcha: '',
        captcha_id: '',
      },
      captcha: {
        enable: false,
      },
      rules: {
        content: [
          { required: true, message: '请输入评论内容', trigger: 'blur' },
        ],
      },
    }
  },
  watch: {
    articleId: {
      handler(val) {
        this.comment.article_id = val
      },
      immediate: true,
    },
    parentId: {
      handler(val) {
        this.comment.parent_id = val
      },
      immediate: true,
    },
  },
  created() {
    this.loadCaptcha()
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(async (valid) => {
        if (valid) {
          const res = await createComment(this.comment)
          if (res.status === 200) {
            this.$message.success('评论成功')
            this.comment.content = ''
            this.comment.captcha = ''
            this.loadCaptcha()
            this.$emit('success')
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          return false
        }
      })
    },
    async loadCaptcha() {
      const res = await getUserCaptcha({ type: 'comment', t: Date.now() })
      if (res.data.enable) {
        // 启用了验证码
        this.comment = {
          ...this.comment,
          captcha_id: res.data.id,
        }
        this.captcha = res.data
      }
    },
  },
}
</script>
<style lang="scss">
.com-form-comment {
  .comment-content {
    width: 100%;
    .el-form-item__content {
      display: block;
    }
  }
  .comment-btns {
    width: 100%;
    img {
      height: 40px;
    }
    .el-form-item__content {
      display: block;
    }
    .captcha {
      float: left;
    }
  }
}
</style>
