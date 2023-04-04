<template>
  <div class="com-form-update-documents-category">
    <el-form ref="form" label-position="top" label-width="80px" :model="form">
      <el-form-item
        label="新文档分类"
        prop="category_id"
        :rules="[
          { required: true, trigger: 'blur', message: '请选择新的文档分类' },
        ]"
      >
        <el-cascader
          v-model="form.category_id"
          :options="categoryTrees"
          :filterable="true"
          :props="{
            checkStrictly: true,
            expandTrigger: 'hover',
            label: 'title',
            value: 'id',
          }"
          clearable
          placeholder="请选择新的文档分类"
        ></el-cascader>
      </el-form-item>
      <el-form-item label="文档列表" class="document-list">
        <ul>
          <li v-for="document in documents" :key="'doc-' + document.id">
            {{ document.title }}
          </li>
        </ul>
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          class="btn-block"
          icon="el-icon-check"
          @click="setDocumentsCategory"
          >提交</el-button
        >
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { setDocumentsCategory } from '~/api/document'
import { mapGetters } from 'vuex'
export default {
  name: 'FormUpdateDocumentsCategory',
  props: {
    // 文档分类
    categoryTrees: {
      type: Array,
      default: () => {
        return []
      },
    },
    documents: {
      type: Array,
      default: () => {
        return []
      },
    },
  },
  data() {
    return {
      form: {
        category_id: [],
        document_id: [],
      },
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  created() {},
  methods: {
    async setDocumentsCategory() {
      this.$refs.form.validate(async (valid) => {
        if (valid) {
          this.form.document_id = this.documents.map((item) => item.id)
          const res = await setDocumentsCategory(this.form)
          console.log(res)
          if (res.status === 200) {
            this.$message.success('修改成功')
            this.$emit('success', res.data)
          }
        }
      })
    },
  },
}
</script>
<style lang="scss">
.com-form-update-documents-category {
  .document-list {
    ul,
    li {
      list-style: none;
      margin: 0;
      padding: 0;
    }
    ul {
      max-height: 300px;
      overflow: auto;
    }
    li {
      line-height: 30px;
      color: #777;
    }
  }
}
</style>
