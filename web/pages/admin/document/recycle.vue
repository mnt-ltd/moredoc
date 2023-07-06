<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="false"
        :show-delete="false"
        :disabled-delete="selectedRow.length === 0"
        :default-search="search"
        @onSearch="onSearch"
      >
        <template slot="buttons">
          <el-form-item>
            <el-button
              type="success"
              icon="el-icon-refresh-left"
              :disabled="selectedRow.length === 0"
              @click="batchRecover"
              >恢复选中</el-button
            > </el-form-item
          ><el-form-item>
            <el-button
              type="warning"
              icon="el-icon-close"
              :disabled="selectedRow.length === 0"
              @click="batchDelete"
              >删除选中</el-button
            > </el-form-item
          ><el-form-item>
            <el-button
              type="danger"
              :disabled="selectedRow.length > 0"
              icon="el-icon-delete"
              @click="clearAll"
              >清空回收站</el-button
            >
          </el-form-item>
        </template>
      </FormSearch>
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
      >
        <template slot="actions" slot-scope="scope">
          <el-button
            type="text"
            icon="el-icon-refresh-left"
            size="small"
            @click="recoverRow(scope.row)"
            >恢复</el-button
          >
        </template>
      </TableList>
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
import {
  clearRecycleDocument,
  deleteRecycleDocument,
  listRecycleDocument,
  recoverRecycleDocument,
} from '~/api/document'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import { categoryToTrees, parseQueryIntArray, genLinkHTML } from '~/utils/utils'
import { documentStatusOptions } from '~/utils/enum'
import { mapGetters } from 'vuex'
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
      categoryMap: {},
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      documentStatusOptions,
      document: { id: 0 },
    }
  },
  head() {
    return {
      title: `回收站 - ${this.settings.system.sitename}`,
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
          page: parseInt(this.$route.query.page) || 1,
          size: parseInt(this.$route.query.size) || 10,
          ...parseQueryIntArray(this.$route.query, ['category_id', 'status']),
        }

        // 需要先加载分类数据
        if (this.trees.length === 0) {
          await this.listCategory()
        }
        await this.listDocument()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
  },
  methods: {
    async listCategory() {
      const res = await listCategory({ field: ['id', 'parent_id', 'title'] })
      if (res.status === 200) {
        let categories = res.data.category || []
        categories = categories.map((item) => {
          item.disable_delete = item.doc_count > 0
          return item
        })

        const categoryMap = {}
        categories.forEach((item) => {
          categoryMap[item.id] = item
        })
        this.categoryMap = categoryMap
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
      const res = await listRecycleDocument(search)
      this.loading = false
      if (res.status === 200) {
        const documents = res.data.document || []
        documents.forEach((item) => {
          ;(item.category_id || (item.category_id = [])).forEach((id) => {
            ;(item.category_name || (item.category_name = [])).push(
              this.categoryMap[id].title || '-' // 有可能分类已经被删除
            )
          })
          item.title_html = genLinkHTML(item.title, `/document/${item.id}`)
          item.username_html = genLinkHTML(
            item.username,
            `/user/${item.user_id}`
          )
        })

        this.documents = documents
        this.total = res.data.total
      } else {
        this.$message.error(res.data.message)
      }
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
        this.listDocument()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    recoverRow(row) {
      this.$confirm(`您确定要要恢复【${row.title}】吗？`, '温馨提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'info',
      }).then(async () => {
        const res = await recoverRecycleDocument({ id: [row.id] })
        if (res.status === 200) {
          this.$message.success('恢复成功')
          this.listDocument()
        } else {
          this.$message.error(res.data.message || '操作失败')
        }
      })
    },
    clearAll() {
      this.$confirm(
        '您确定要永久删除回收站中的所有文档吗？清空之后不可恢复，请慎重操作！',
        '风险提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'error',
        }
      ).then(async () => {
        const res = await clearRecycleDocument({ id: 0 })
        if (res.status === 200) {
          this.$message.success('清空成功')
          this.listDocument()
        } else {
          this.$message.error(res.data.message || '操作失败')
        }
      })
    },
    batchRecover() {
      this.$confirm(
        `您确定要从回收站中恢复选中的【${this.selectedRow.length}个】文档吗？`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'info',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await recoverRecycleDocument({ id: ids })
          if (res.status === 200) {
            this.$message.success('恢复成功')
            this.listDocument()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    batchDelete() {
      this.$confirm(
        `您确定要从回收站中删除选中的【${this.selectedRow.length}个】文档吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'info',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteRecycleDocument({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listDocument()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要从回收站中删除文档【${row.title}】吗？删除之后不可恢复！`,
        '告警',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteRecycleDocument({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listDocument()
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
        {
          prop: 'title_html',
          label: '名称',
          minWidth: 200,
          fixed: 'left',
          type: 'html',
        },
        { prop: 'username_html', label: '上传者', width: 120, type: 'html' },
        { prop: 'deleted_username', label: '删除者', width: 120 },
        { prop: 'deleted_at', label: '删除时间', width: 160, type: 'datetime' },
        {
          prop: 'status',
          label: '状态',
          width: 120,
          type: 'enum',
          enum: statusMap,
        },
        {
          prop: 'category_name',
          label: '分类',
          minWidth: 180,
          type: 'breadcrumb',
        },
        { prop: 'pages', label: '页数', width: 80, type: 'number' },
        { prop: 'price', label: '价格', width: 80, type: 'number' },
        { prop: 'download_count', label: '下载', width: 80, type: 'number' },
        { prop: 'view_count', label: '浏览', width: 80, type: 'number' },
        { prop: 'favorite_count', label: '收藏', width: 80, type: 'number' },
        { prop: 'comment_count', label: '评论', width: 80, type: 'number' },
        { prop: 'keywords', label: '关键字', minWidth: 200 },
        // { prop: 'description', label: '摘要', minWidth: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
