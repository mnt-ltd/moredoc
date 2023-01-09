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
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :actions-min-width="160"
        @editRow="editRow"
        @viewRow="viewRow"
        @selectRow="selectRow"
        @deleteRow="deleteRow"
      >
        <template slot="actions" slot-scope="scope">
          <el-button
            type="text"
            size="small"
            icon="el-icon-s-check"
            @click="recommendDocument(scope.row)"
            >推荐</el-button
          >
          <nuxt-link :to="`/document/${scope.row.id}`" target="_blank"
            ><el-button type="text" icon="el-icon-view" size="small"
              >查看</el-button
            ></nuxt-link
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
    <el-dialog title="编辑文档" width="640px" :visible.sync="formVisible">
      <FormUpdateDocument
        :category-trees="trees"
        :init-document="document"
        :is-admin="true"
        @success="formSuccess"
      />
    </el-dialog>
    <el-dialog
      title="推荐设置"
      :visible.sync="formDocumentRecommendVisible"
      width="640px"
    >
      <FormDocumentRecommend :init-document="document" @success="formSuccess" />
    </el-dialog>
  </div>
</template>

<script>
import { listCategory } from '~/api/category'
import { deleteDocument, getDocument, listDocument } from '~/api/document'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import { categoryToTrees } from '~/utils/utils'
import { documentStatusOptions, boolOptions } from '~/utils/enum'
import FormUpdateDocument from '~/components/FormUpdateDocument.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormUpdateDocument },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formVisible: false,
      formDocumentRecommendVisible: false,
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
      boolOptions,
      document: { id: 0 },
    }
  },
  head() {
    return {
      title: `文档列表 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // 需要先加载分类数据
    await this.listCategory()
    await this.listDocument()
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
      const res = await listDocument(search)
      if (res.status === 200) {
        const documents = res.data.document || []
        documents.forEach((item) => {
          ;(item.category_id || (item.category_id = [])).forEach((id) => {
            ;(item.category_name || (item.category_name = [])).push(
              this.categoryMap[id].title
            )
          })
        })

        this.documents = documents
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
      // 新增，跳转到前台文档上传页面
      const routeUrl = this.$router.resolve({
        path: '/upload',
      })
      window.open(routeUrl.href, '_blank')
    },
    viewRow(row) {
      // 查看，跳转到前台文档详情页面
      const routeUrl = this.$router.resolve({
        path: '/document/' + row.uuid,
      })
      window.open(routeUrl.href, '_blank')
    },
    async editRow(row) {
      const res = await getDocument({ id: row.id })
      if (res.status === 200) {
        this.document = res.data
        this.formVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    async recommendDocument(row) {
      const res = await getDocument({ id: row.id })
      if (res.status === 200) {
        this.document = res.data
        this.formDocumentRecommendVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formSuccess() {
      this.formVisible = false
      this.formDocumentRecommendVisible = false
      this.listDocument()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}个】文档吗？删除之后将会进入到回收站，可以在回收站中恢复。`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteDocument({ id: ids })
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
        `您确定要删除文档【${row.title}】吗？删除之后将会进入到回收站，可以在回收站中恢复。`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteDocument({ id: row.id })
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
        {
          type: 'select',
          label: '推荐',
          name: 'is_recommend',
          placeholder: '请选择推荐状态',
          multiple: true,
          options: boolOptions,
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
        { prop: 'ext', label: '扩展名', width: 70 },
        { prop: 'username', label: '上传者', width: 120 },
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
        {
          prop: 'recommend_at',
          label: '推荐时间',
          width: 160,
          type: 'datetime',
        },
      ]
    },
  },
}
</script>
<style></style>
