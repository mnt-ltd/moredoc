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
        @onDelete="batchDelete"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :loading="loading"
        :table-data="listData"
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
      title="编辑附件"
      width="640px"
      :visible.sync="formVisible"
    >
      <FormAttachment :init-attachment="attachment" @success="formSuccess" />
    </el-dialog>
  </div>
</template>

<script>
import {
  listAttachment,
  deleteAttachment,
  getAttachment,
} from '~/api/attachment'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormAttachment from '~/components/FormAttachment.vue'
import { attachmentTypeOptions } from '~/utils/enum'
import { parseQueryIntArray } from '~/utils/utils'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormAttachment },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formVisible: false,
      search: {
        wd: '',
        page: 1,
        size: 10,
      },
      listData: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      attachment: {},
      attachmentTypeOptions,
    }
  },
  head() {
    return {
      title: `附件管理 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    '$route.query': {
      immediate: true,
      handler() {
        this.search = {
          ...this.search,
          ...this.$route.query,
          page: parseInt(this.$route.query.page) || 1,
          size: parseInt(this.$route.query.size) || 10,
          ...parseQueryIntArray(this.$route.query, ['enable', 'type']),
        }
        this.listAttachment()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // await this.listAttachment()
  },
  methods: {
    async listAttachment() {
      this.loading = true
      const res = await listAttachment(this.search)
      if (res.status === 200) {
        this.listData = res.data.attachment
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
      // this.listAttachment()
    },
    handlePageChange(val) {
      this.search.page = val
      this.$router.push({
        query: this.search,
      })
      // this.listAttachment()
    },
    onSearch(search) {
      this.search = { ...this.search, ...search, page: 1 }
      if (
        location.pathname + location.search ===
        this.$router.resolve({
          query: this.search,
        }).href
      ) {
        this.listAttachment()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    async editRow(row) {
      const res = await getAttachment({ id: row.id })
      if (res.status === 200) {
        this.attachment = res.data
        this.formVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formSuccess() {
      this.formVisible = false
      this.listAttachment()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}个】附件吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteAttachment({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listAttachment()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除附件【${row.name}】吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteAttachment({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listAttachment()
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
          label: '附件类型',
          name: 'type',
          placeholder: '请选择附件类型',
          multiple: true,
          options: this.attachmentTypeOptions,
        },
        {
          type: 'select',
          label: '是否合法',
          name: 'enable',
          placeholder: '请选择是否合法',
          multiple: true,
          options: [
            { label: '是', value: 1 },
            { label: '否', value: 0 },
          ],
        },
      ]
    },
    initTableListFields() {
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number' },
        { prop: 'type_name', label: '类型', width: 80 },
        { prop: 'name', label: '名称', minWidth: 200 },
        {
          prop: 'enable',
          label: '是否合法',
          width: 80,
          type: 'bool',
        },
        { prop: 'username', label: '上传者', width: 120 },
        { prop: 'ip', label: 'IP', width: 120 },
        { prop: 'size', label: '大小', width: 90, type: 'bytes' },
        { prop: 'width', label: '宽', width: 90 },
        { prop: 'height', label: '高', width: 90 },
        { prop: 'ext', label: '扩展', width: 90 },
        { prop: 'hash', label: 'HASH', width: 290 },
        { prop: 'path', label: '存储路径', minWidth: 300 },
        { prop: 'description', label: '备注', width: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
