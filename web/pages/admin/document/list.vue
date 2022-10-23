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
        :loading="loading"
        :table-data="documents"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="false"
        :show-delete="true"
        :show-select="true"
        @selectRow="selectRow"
        @deleteRow="deleteRow"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div class="text-right">
        <el-pagination
          background
          :current-page="search.page"
          :page-sizes="[10, 20, 50, 100]"
          :page-size="search.size"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>
  </div>
</template>

<script>
import { listCategory } from '~/api/category'
import { listDocument } from '~/api/document'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import { categoryToTrees } from '~/utils/utils'
import { documentStatusOptions } from '~/utils/enum'
export default {
  components: { TableList, FormSearch },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formVisible: false,
      search: {
        page: 1,
        size: 10,
      },
      documents: [],
      trees: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      documentStatusOptions,
      document: { id: 0 },
    }
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    this.listCategory()
    await this.listDocument()
  },
  methods: {
    async listCategory() {
      const res = await listCategory()
      if (res.status === 200) {
        let categories = res.data.category || []
        categories = categories.map((item) => {
          item.disable_delete = item.doc_count > 0
          return item
        })
        this.categories = categories

        this.trees = categoryToTrees(categories, false)
        this.total = res.data.total
        this.initSearchForm()
      } else {
        this.$message.error(res.data.message)
      }
    },
    async listDocument() {
      this.loading = true
      const search = { ...this.search }
      if (search.category_id && typeof search.category_id === 'object') {
        search.category_id = search.category_id[search.category_id.length - 1]
      }
      const res = await listDocument(search)
      if (res.status === 200) {
        this.documents = res.data.document || []
        this.total = res.data.total
      } else {
        this.$message.error(res.data.message)
      }
      this.loading = false
    },
    handleSizeChange(val) {
      this.search.size = val
      this.listDocument()
    },
    handlePageChange(val) {
      this.search.page = val
      this.listDocument()
    },
    onSearch(search) {
      this.search = { ...this.search, page: 1, ...search }
      this.listDocument()
    },
    onCreate() {
      this.document = { id: 0, enable: true }
      this.formVisible = true
    },
    formSuccess() {
      this.formVisible = false
      this.listDocument()
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
          // const ids = this.selectedRow.map((item) => item.id)
          // const res = await deleteCategory({ id: ids })
          // if (res.status === 200) {
          //   this.$message.success('删除成功')
          //   this.listCategory()
          // } else {
          //   this.$message.error(res.data.message)
          // }
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
          // const res = await deleteCategory({ id: row.id })
          // if (res.status === 200) {
          //   this.$message.success('删除成功')
          //   this.listCategory()
          // } else {
          //   this.$message.error(res.data.message)
          // }
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
          name: 'status',
          placeholder: '请选择状态',
          multiple: true,
          options: documentStatusOptions,
        },
        // 级联
        {
          type: 'cascader',
          label: '分类',
          name: 'category_id',
          placeholder: '请选择分类',
          trees: this.trees,
        },
      ]
    },
    initTableListFields() {
      const statusMap = {}
      this.documentStatusOptions.forEach((item) => {
        statusMap[item.value] = item
      })
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        { prop: 'title', label: '名称', minWidth: 200, fixed: 'left' },
        { prop: 'username', label: '上传者', width: 120 },
        {
          prop: 'status',
          label: '状态',
          width: 120,
          type: 'enum',
          enum: statusMap,
        },
        { prop: 'pages', label: '页数', width: 80, type: 'number' },
        { prop: 'price', label: '价格', width: 80, type: 'number' },
        { prop: 'download_count', label: '下载', width: 80, type: 'number' },
        { prop: 'view_count', label: '浏览', width: 80, type: 'number' },
        { prop: 'favorite_count', label: '收藏', width: 80, type: 'number' },
        { prop: 'comment_count', label: '评论', width: 80, type: 'number' },
        { prop: 'keywords', label: '关键字', minWidth: 200 },
        { prop: 'description', label: '摘要', minWidth: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
