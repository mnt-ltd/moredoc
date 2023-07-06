<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="true"
        :show-delete="true"
        :disabled-delete="selectedRow.length === 0"
        :default-search="search"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      >
        <template slot="buttons">
          <el-form-item>
            <el-tooltip
              class="item"
              effect="dark"
              content="将转换失败的文档一键重置为待转换状态，以便重新转换"
              placement="top"
            >
              <el-button
                type="warning"
                @click="reconvertDocument"
                icon="el-icon-refresh"
                >失败重转</el-button
              >
            </el-tooltip>
          </el-form-item>
          <el-form-item>
            <el-tooltip
              class="item"
              effect="dark"
              content="批量修改选中的文档分类"
              placement="top"
            >
              <el-button
                type="success"
                @click="batchUpdateDocumentsCategory"
                :disabled="selectedRow.length === 0"
                icon="el-icon-edit"
                >批量分类</el-button
              >
            </el-tooltip>
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
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :actions-min-width="100"
        @editRow="editRow"
        @viewRow="viewRow"
        @selectRow="selectRow"
        @deleteRow="deleteRow"
      >
        <template slot="actions" slot-scope="scope">
          <el-tooltip
            v-if="scope.row.convert_error && scope.row.status === 3"
            class="item"
            effect="dark"
            placement="top"
          >
            <div slot="content">
              <div class="tooltip-box">
                {{ scope.row.convert_error }}
              </div>
            </div>
            <el-button
              type="text"
              size="small"
              class="text-warning"
              icon="el-icon-error"
              >转换失败原因</el-button
            >
          </el-tooltip>
          <el-button
            type="text"
            size="small"
            icon="el-icon-s-check"
            @click="recommendDocument(scope.row)"
            >推荐</el-button
          >
          <!-- <nuxt-link :to="`/document/${scope.row.id}`" target="_blank"
            ><el-button type="text" icon="el-icon-view" size="small"
              >查看</el-button
            ></nuxt-link
          > -->
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
    <el-dialog
      :close-on-click-modal="false"
      title="编辑文档"
      width="640px"
      :visible.sync="formVisible"
    >
      <FormUpdateDocument
        :category-trees="trees"
        :init-document="document"
        :is-admin="true"
        @success="formSuccess"
      />
    </el-dialog>
    <el-dialog
      :close-on-click-modal="false"
      title="批量分类"
      width="640px"
      :visible.sync="formDocumentsCategoryVisible"
    >
      <FormUpdateDocumentsCategory
        v-if="formDocumentsCategoryVisible"
        :category-trees="trees"
        :documents="categoryDocuments"
        @success="formSuccess"
      />
    </el-dialog>
    <el-dialog
      :close-on-click-modal="false"
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
import {
  deleteDocument,
  getDocument,
  listDocument,
  setDocumentReconvert,
} from '~/api/document'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import {
  categoryToTrees,
  parseQueryIntArray,
  parseQueryBoolArray,
  genLinkHTML,
} from '~/utils/utils'
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
      formDocumentsCategoryVisible: false,
      categoryDocuments: [],
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
  watch: {
    '$route.query': {
      immediate: true,
      async handler() {
        let search = { ...this.search, ...this.$route.query }
        search.page = parseInt(this.$route.query.page) || 1
        search.size = parseInt(this.$route.query.size) || 10
        search.wd = this.$route.query.wd || ''
        this.search = {
          ...search,
          ...parseQueryIntArray(this.$route.query, ['category_id', 'status']),
          ...parseQueryBoolArray(this.$route.query, ['is_recommend']),
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
      const res = await listDocument(search)
      if (res.status === 200) {
        const documents = res.data.document || []
        documents.forEach((item) => {
          // 对于转换中的文档，禁止删除
          item.disable_delete = item.status === 1
          ;(item.category_id || (item.category_id = [])).forEach((id) => {
            ;(item.category_name || (item.category_name = [])).push(
              this.categoryMap[id].title
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
        this.listDocument()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      // 新增，跳转到前台文档上传页面
      const routeUrl = this.$router.resolve({
        path: '/upload',
      })
      window.open(routeUrl.href, '_blank')
    },
    async reconvertDocument() {
      const res = await setDocumentReconvert()
      if (res.status === 200) {
        this.$message.success('提交成功，请耐心等待重新转换')
        this.listDocument()
      } else {
        this.$message.error(res.data.message || '操作失败')
      }
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
      this.formDocumentsCategoryVisible = false
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
    batchUpdateDocumentsCategory() {
      this.categoryDocuments = this.selectedRow
      this.formDocumentsCategoryVisible = true
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
        {
          prop: 'title_html',
          label: '文档',
          minWidth: 200,
          fixed: 'left',
          type: 'html',
        },
        { prop: 'ext', label: '扩展名', width: 70 },
        { prop: 'username_html', label: '上传者', width: 120, type: 'html' },
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
        { prop: 'size', label: '大小', width: 100, type: 'bytes' },
        { prop: 'pages', label: '页数', width: 80, type: 'number' },
        { prop: 'price', label: '价格', width: 80, type: 'number' },
        { prop: 'download_count', label: '下载', width: 70, type: 'number' },
        { prop: 'view_count', label: '浏览', width: 70, type: 'number' },
        { prop: 'favorite_count', label: '收藏', width: 70, type: 'number' },
        { prop: 'comment_count', label: '评论', width: 70, type: 'number' },
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
<style lang="scss" scoped>
.tooltip-box {
  max-width: 300px;
  word-break: break-all;
}
</style>
