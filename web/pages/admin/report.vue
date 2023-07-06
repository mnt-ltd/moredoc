<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="false"
        :show-delete="true"
        :disabled-delete="selectedRow.length === 0"
        :default-search="search"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :loading="loading"
        :table-data="reports"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        @selectRow="selectRow"
        @editRow="editRow"
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

    <el-dialog
      :close-on-click-modal="false"
      :title="report.id ? '编辑举报' : '新增举报'"
      :visible.sync="formReportVisible"
      width="640px"
    >
      <FormReport
        ref="reportForm"
        :init-report="report"
        :is-admin="true"
        @success="formReportSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import { listReport, deleteReport } from '~/api/report'
import { reportOptions } from '~/utils/enum'
import { parseQueryIntArray, genLinkHTML } from '~/utils/utils'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormReport from '~/components/FormReport.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormReport },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formReportVisible: false,
      search: {
        wd: '',
        page: 1,
        status: [],
        size: 10,
      },
      reports: [],
      reportOptions,
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      report: { id: 0 },
    }
  },
  head() {
    return {
      title: `举报管理 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
  },
  watch: {
    '$route.query': {
      handler() {
        this.search = {
          ...this.search,
          ...this.$route.query,
          page: parseInt(this.$route.query.page) || 1,
          size: parseInt(this.$route.query.size) || 10,
          ...parseQueryIntArray(this.$route.query, ['status']),
        }
        this.listReport()
      },
      immediate: true,
    },
  },
  methods: {
    async listReport() {
      this.loading = true
      const res = await listReport(this.search)
      if (res.status === 200) {
        let reports = res.data.report || []
        reports.map((item) => {
          item.username_html = genLinkHTML(
            item.username,
            `/user/${item.user_id}`
          )
          item.document_title_html = genLinkHTML(
            item.document_title,
            `/document/${item.document_id}`
          )
        })
        this.reports = reports
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
        this.listReport()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.report = { id: 0 }
      this.formReportVisible = true
      this.$nextTick(() => {
        this.$refs.reportForm.reset()
      })
    },
    async editRow(row) {
      this.report = row
      this.formReportVisible = true
    },
    formReportSuccess() {
      this.formReportVisible = false
      this.listReport()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}条】举报吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteReport({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listReport()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除对文档【${row.document_title}】的举报吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteReport({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listReport()
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
          options: [
            { label: '已处理', value: 1 },
            { label: '未处理', value: 0 },
          ],
        },
      ]
    },
    initTableListFields() {
      const reasonEnum = {}
      this.reportOptions.forEach((item) => {
        reasonEnum[item.value] = item
      })

      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'status',
          label: '是否已处理',
          width: 100,
          type: 'bool',
          fixed: 'left',
        },
        {
          prop: 'document_title_html',
          label: '文档',
          minWidth: 150,
          fixed: 'left',
          type: 'html',
        },
        {
          prop: 'reason',
          label: '举报原因',
          width: 80,
          type: 'enum',
          enum: reasonEnum,
        },
        {
          prop: 'username_html',
          label: '举报人',
          width: 100,
          type: 'html',
        },
        { prop: 'remark', label: '处理描述', minWidth: 150 },
        { prop: 'created_at', label: '举报时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
