<template>
  <div class="com-form-article">
    <el-form
      ref="formArticle"
      label-position="top"
      label-width="80px"
      :model="article"
    >
      <el-form-item
        label="标题"
        prop="title"
        :rules="[
          { required: true, trigger: 'blur', message: '请输入文章标题' },
        ]"
      >
        <el-input
          v-model="article.title"
          placeholder="请输入文章标题"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item
        label="标识"
        prop="identifier"
        :rules="[
          {
            required: true,
            trigger: 'blur',
            message: '请输入文章标识，建议为字母和数字组合',
          },
        ]"
      >
        <!-- 如果是编辑文章，不允许修改文章标识 -->
        <el-input
          v-model="article.identifier"
          placeholder="请输入文章标识，建议为字母和数字组合"
          :disabled="article.id > 0"
          clearable
        ></el-input>
      </el-form-item>
      <el-form-item label="作者">
        <el-input
          v-model="article.author"
          placeholder="请输入文章作者，可为空，默认为当前登录用户"
        ></el-input>
      </el-form-item>
      <el-form-item label="关键字">
        <el-input
          v-model="article.keywords"
          placeholder="请输入文章关键字，多个关键字用英文逗号分隔"
        ></el-input>
      </el-form-item>
      <el-form-item label="描述">
        <el-input
          v-model="article.description"
          placeholder="请输入文章描述，可为空"
          type="textarea"
          rows="3"
        ></el-input>
      </el-form-item>
      <el-form-item label="内容" class="editor-item">
        <Toolbar
          style="border-bottom: 1px solid #ccc"
          :editor="editor"
          :default-config="toolbarConfig"
          :mode="mode"
        />
        <Editor
          v-model="article.content"
          style="height: 800px; overflow-y: hidden"
          :default-config="editorConfig"
          :mode="mode"
          @onCreated="onCreated"
        />
      </el-form-item>
      <el-form-item v-if="!isEditorFullScreen">
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          :loading="loading"
          @click="onSubmit"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { Boot } from '@wangeditor/editor'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import markdownModule from '@wangeditor/plugin-md'
import { createArticle, updateArticle } from '~/api/article'

export default {
  name: 'FormArticle',
  components: {
    Editor,
    Toolbar,
  },
  props: {
    initArticle: {
      type: Object,
      default: () => {
        return {}
      },
    },
  },
  data() {
    return {
      loading: false,
      article: {},
      editor: null,
      toolbarConfig: {},
      isEditorFullScreen: false,
      editorConfig: {
        placeholder: '请输入内容...',
        MENU_CONF: {
          uploadImage: {
            server: '/api/v1/upload/article?type=image',
            fieldName: 'file',
            maxFileSize: 20 * 1024 * 1024, // 20M
            headers: {
              Authorization: 'Bearer ' + this.$store.getters['user/token'],
            },
            timeout: 30 * 1000, // 30s
            withCredentials: false,
            onFailed: (file, res) => {
              this.$message.error(`${file.name}上传失败：${res.msg}`)
            },
          },
          uploadVideo: {
            server: '/api/v1/upload/article?type=video',
            fieldName: 'file',
            maxFileSize: 1024 * 1024 * 1024, // 1GB
            headers: {
              Authorization: 'Bearer ' + this.$store.getters['user/token'],
            },
            timeout: 600 * 1000, // 10min
            withCredentials: false,
            onFailed: (file, res) => {
              this.$message.error(`${file.name}上传失败：${res.msg}`)
            },
          },
        },
      },
      mode: 'default', // 'default' or 'simple'
    }
  },
  watch: {
    initArticle: {
      handler(val) {
        this.article = { ...val }
      },
      immediate: true,
    },
  },
  created() {
    Boot.registerModule(markdownModule)
    this.article = { ...this.initArticle }
  },
  beforeDestroy() {
    const editor = this.editor
    if (editor == null) return
    editor.destroy() // 组件销毁时，及时销毁编辑器
  },
  methods: {
    onCreated(editor) {
      this.editor = Object.seal(editor) // 一定要用 Object.seal() ，否则会报错
      this.editor.on('fullScreen', () => {
        this.isEditorFullScreen = true
      })
      this.editor.on('unFullScreen', () => {
        this.isEditorFullScreen = false
      })
    },
    onSubmit() {
      this.$refs.formArticle.validate(async (valid) => {
        if (!valid) {
          return
        }
        this.loading = true
        const article = { ...this.article }
        if (this.article.id > 0) {
          const res = await updateArticle(article)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        } else {
          const res = await createArticle(article)
          if (res.status === 200) {
            this.$message.success('新增成功')
            this.resetFields()
            this.$emit('success', res.data)
          } else {
            this.$message.error(res.data.message)
          }
        }
        this.loading = false
      })
    },
    clearValidate() {
      this.$refs.formArticle.clearValidate()
    },
    resetFields() {
      this.article = {
        id: 0,
        title: '',
        identifier: '',
        keywords: '',
        description: '',
        content: '',
      }
    },
    reset() {
      this.resetFields()
      this.clearValidate()
    },
  },
}
</script>
<style src="@wangeditor/editor/dist/css/style.css"></style>
<style lang="scss">
.com-form-article {
  .editor-item {
    .el-form-item__content {
      overflow: hidden;
      border: 1px solid #ccc;
      border-radius: 4px;
      line-height: unset;
    }
    h1 {
      font-size: 1.7em;
    }
  }
  .w-e-text-container [data-slate-editor] blockquote {
    border-left-width: 4px;
  }
  .w-e-text-container [data-slate-editor] table th {
    border-right-width: 1px !important;
  }
}
</style>
