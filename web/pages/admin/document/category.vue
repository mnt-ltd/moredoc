<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="true"
        :show-delete="true"
        :disabled-delete="selectedRow.length === 0"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :table-data="trees"
        :loading="loading"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :tree-props="{ children: 'children' }"
        @selectRow="selectRow"
        @editRow="editRow"
        @deleteRow="deleteRow"
      />
    </el-card>
    <el-dialog
      :close-on-click-modal="false"
      :title="category.id ? '编辑分类' : '新增分类'"
      :visible.sync="formVisible"
      :width="'640px'"
    >
      <FormCategory
        ref="categoryForm"
        :init-category="category"
        :trees="trees"
        @success="formCategorySuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import { listCategory, deleteCategory, getCategory } from '~/api/category'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormCategory from '~/components/FormCategory.vue'
import { categoryToTrees, parseQueryIntArray } from '~/utils/utils'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormCategory },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formVisible: false,
      search: {
        wd: '',
        status: [],
      },
      categories: [],
      trees: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      category: { id: 0, title: '', cover: '', sort: '', icon: '' },
    }
  },
  head() {
    return {
      title: `分类管理 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    '$route.query': {
      immediate: true,
      async handler() {
        this.search = {
          ...this.search,
          ...this.$route.query,
          ...parseQueryIntArray(this.$route.query, ['enable']),
        }
        await this.initTableListFields()
        this.listCategory()
      },
    },
  },
  async created() {
    this.initSearchForm()
  },
  methods: {
    async listCategory() {
      this.loading = true
      const res = await listCategory(this.search)
      if (res.status === 200) {
        let categories = res.data.category || []
        categories = categories.map((item) => {
          item.disable_delete = item.doc_count > 0
          return item
        })
        this.categories = categories

        this.trees = categoryToTrees(categories)
        this.total = res.data.total
      } else {
        this.$message.error(res.data.message)
      }
      this.loading = false
    },
    handleSizeChange(val) {
      this.search.size = val
      this.$router.push({
        query: this.search,
      })
    },
    handlePageChange(val) {
      this.search.page = val
      this.$router.push({
        query: this.search,
      })
    },
    onSearch(search) {
      this.search = { ...this.search, ...search, page: 1 }
      if (
        location.pathname + location.search ===
        this.$router.resolve({
          query: this.search,
        }).href
      ) {
        this.listCategory()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.category = {
        id: 0,
        enable: true,
        title: '',
        cover: '',
        icon: '',
        parent_id: 0,
        sort: 0,
      }
      this.formVisible = true
    },
    async editRow(row) {
      const res = await getCategory({ id: row.id })
      if (res.status === 200) {
        this.category = { cover: '', icon: '', ...res.data }
        this.formVisible = true
      } else {
        this.$message.error(res.data.message || '查询失败')
      }
    },
    formCategorySuccess() {
      this.formVisible = false
      this.listCategory()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}个】分类吗？本次删除会连同子分类一起删除，删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteCategory({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listCategory()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除分类【${row.title}】吗？本次删除会连同子分类一起删除，删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteCategory({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listCategory()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    selectRow(rows) {
      this.selectedRow = rows
    },
    initSearchForm() {
      this.searchFormFields = [
        {
          type: 'text',
          label: '关键字',
          name: 'wd',
          placeholder: '请输入关键字',
        },
        {
          type: 'select',
          label: '状态',
          name: 'enable',
          placeholder: '请选择状态',
          multiple: true,
          options: [
            { label: '启用', value: 1 },
            { label: '禁用', value: 0 },
          ],
        },
      ]
    },
    initTableListFields() {
      if (this.tableListFields.length > 0) {
        return
      }
      this.tableListFields = [
        { prop: 'title', label: '名称', minWidth: 120, fixed: 'left' },
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'enable',
          label: '是否启用',
          width: 80,
          type: 'bool',
        },
        {
          prop: 'show_description',
          label: '显示描述',
          width: 80,
          type: 'bool',
        },
        {
          prop: 'sort',
          label: '排序',
          width: 80,
          type: 'number',
        },
        { prop: 'icon', label: '图标', width: 48, type: 'image' },
        { prop: 'cover', label: '封面', width: 100, type: 'image' },
        { prop: 'doc_count', label: '文档数', width: 80, type: 'number' },
        { prop: 'description', label: '分类描述', minWidth: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
