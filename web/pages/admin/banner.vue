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
        @onCreate="onCreate"
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
        :actions-min-width="80"
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
      width="640px"
      :title="banner.id > 0 ? '编辑横幅' : '新增横幅'"
      :visible.sync="formVisible"
    >
      <FormBanner
        ref="formBanner"
        :init-banner="banner"
        @success="formSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import { listBanner, deleteBanner, getBanner } from '~/api/banner'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormBanner from '~/components/FormBanner.vue'
import { bannerTypeOptions } from '~/utils/enum'
import { parseQueryIntArray } from '~/utils/utils'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormBanner },
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
      banner: {},
      bannerTypeOptions,
    }
  },
  head() {
    return {
      title: `横幅管理 - ${this.settings.system.sitename}`,
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
        this.listBanner()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
  },
  methods: {
    async listBanner() {
      this.loading = true
      const res = await listBanner(this.search)
      if (res.status === 200) {
        this.listData = res.data.banner
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
        this.listBanner()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.banner = {}
      this.formVisible = true
      this.$nextTick(() => {
        this.$refs.formBanner.reset()
      })
    },
    async editRow(row) {
      const res = await getBanner({ id: row.id })
      if (res.status === 200) {
        this.banner = res.data
        this.formVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formSuccess() {
      this.formVisible = false
      this.listBanner()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}个】横幅吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteBanner({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listBanner()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除横幅【${row.title}】吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteBanner({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listBanner()
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
          label: '类型',
          name: 'type',
          placeholder: '请选择横幅类型',
          multiple: true,
          options: this.bannerTypeOptions,
        },
        {
          type: 'select',
          label: '状态',
          name: 'enable',
          placeholder: '是否启用',
          multiple: true,
          options: [
            { label: '启用', value: 1 },
            { label: '禁用', value: 0 },
          ],
        },
      ]
    },
    initTableListFields() {
      const typeMap = {}
      this.bannerTypeOptions.forEach((item) => {
        typeMap[item.value] = item
      })
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number' },
        { prop: 'path', label: '横幅', width: 360, type: 'image' },
        {
          prop: 'type',
          label: '类型',
          width: 120,
          type: 'enum',
          enum: typeMap,
        },
        {
          prop: 'enable',
          label: '是否启用',
          width: 80,
          type: 'bool',
        },
        { prop: 'title', label: '名称', minWidth: 150 },
        { prop: 'url', label: '链接', minWidth: 150, type: 'link' },
        { prop: 'sort', label: '排序', width: 80, type: 'number' },
        { prop: 'description', label: '备注', width: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
