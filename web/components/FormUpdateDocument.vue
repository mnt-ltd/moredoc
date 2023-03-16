<template>
  <div class="com-form-update-document">
    <el-form
      ref="document"
      label-position="top"
      label-width="80px"
      :model="document"
    >
      <el-form-item
        label="名称"
        prop="title"
        :rules="[
          { required: true, message: '请输入文档名称', trigger: 'blur' },
        ]"
      >
        <el-input
          v-model="document.title"
          placeholder="请输入文档名称"
        ></el-input>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12" :xs="24">
          <el-form-item
            label="分类"
            prop="category_id"
            :rules="[
              { required: true, trigger: 'blur', message: '请选择文档分类' },
            ]"
          >
            <el-cascader
              v-model="document.category_id"
              :options="categoryTrees"
              :filterable="true"
              :props="{
                checkStrictly: true,
                expandTrigger: 'hover',
                label: 'title',
                value: 'id',
              }"
              clearable
              placeholder="请选择文档分类"
            ></el-cascader>
          </el-form-item>
        </el-col>
        <el-col :span="12" :xs="24">
          <el-form-item
            :label="`价格(${settings.system.credit_name || '魔豆'})`"
            prop="price"
          >
            <el-input-number
              v-model.number="document.price"
              placeholder="文档价格"
              clearable
              :min="0"
              :step="1"
            ></el-input-number> </el-form-item
        ></el-col>
      </el-row>
      <el-form-item
        v-if="isAdmin"
        label="状态"
        prop="status"
        :rules="[
          { required: true, message: '请选择文档状态', trigger: 'change' },
        ]"
      >
        <el-select
          v-model="document.status"
          filterable
          placeholder="请选择文档状态"
        >
          <el-option
            v-for="item in documentStatusOptions"
            :key="'status-' + item.value"
            :value="item.value"
            :label="item.label"
            :disabled="item.disabled"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="关键字">
        <el-input
          v-model="document.keywords"
          placeholder="请输入文档关键字，多个关键字用英文逗号分隔"
        ></el-input>
      </el-form-item>
      <el-form-item label="摘要">
        <el-input
          v-model="document.description"
          placeholder="请输入文档摘要"
          type="textarea"
          rows="5"
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="setDocument"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { updateDocument } from '~/api/document'
import { documentStatusOptions } from '~/utils/enum'
import { mapGetters } from 'vuex'
export default {
  name: 'FormUpdateDocument',
  props: {
    // 是否是管理员。如果是管理员，则可以更新文档的状态，如禁用等
    isAdmin: {
      type: Boolean,
      default: false,
    },
    // 文档分类
    categoryTrees: {
      type: Array,
      default: () => {
        return []
      },
    },
    initDocument: {
      type: Object,
      default: () => {
        return this.getInitialDocumentData()
      },
    },
  },
  data() {
    return {
      documentStatusOptions,
      document: this.getInitialDocumentData(),
    }
  },
  watch: {
    initDocument: {
      handler(val) {
        this.document = { price: 0, ...val }
      },
      immediate: true,
    },
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  created() {
    this.documentStatusOptions = documentStatusOptions.map((item) => {
      if (item.value === 1) {
        // 转换中 这个状态不能选中
        item.disabled = true
      }
      return item
    })
    this.document = this.initDocument
  },
  methods: {
    getInitialDocumentData() {
      return {
        id: 0,
        title: '',
        keywords: '',
        description: '',
        category_id: [],
        price: 0,
        status: 0,
      }
    },
    reset() {
      this.document = this.getInitialDocumentData()
      this.$refs.document.clearValidate()
    },
    setDocument() {
      this.$refs.document.validate(async (valid) => {
        if (valid) {
          const res = await updateDocument(this.document)
          if (res.status === 200) {
            this.$message.success('更新成功')
            this.$emit('success')
          } else {
            this.$message.error(res.data.message || '更新失败')
          }
        }
      })
    },
  },
}
</script>
<style lang="scss">
.com-form-update-document {
  .el-select {
    width: 100%;
  }
}
</style>
