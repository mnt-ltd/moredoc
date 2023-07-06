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
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :table-data="listData"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :loading="loading"
        :show-edit="true"
        :show-delete="false"
        :show-select="false"
        :actions-min-width="80"
        @editRow="editRow"
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
      <FormPermission :init-permission="permission" @success="formSuccess" />
    </el-dialog>
  </div>
</template>

<script>
import { listPermission, getPermission } from '~/api/permission'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormPermission from '~/components/FormPermission.vue'
import { methodOptions } from '~/utils/enum'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormPermission },
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
      permission: {},
      methodOptions,
    }
  },
  head() {
    return {
      title: `权限管理 - ${this.settings.system.sitename}`,
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
        }
        this.listPermission()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
  },
  methods: {
    async listPermission() {
      this.loading = true
      const res = await listPermission(this.search)
      if (res.status === 200) {
        this.listData = res.data.permission
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
        this.listPermission()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    async editRow(row) {
      const res = await getPermission({ id: row.id })
      if (res.status === 200) {
        this.permission = res.data
        this.formVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formSuccess() {
      this.formVisible = false
      this.listPermission()
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
          type: 'text',
          label: 'API',
          name: 'path',
          placeholder: '请输入API',
        },
        {
          type: 'select',
          label: 'Method',
          name: 'method',
          placeholder: '请选择Method',
          multiple: true,
          options: this.methodOptions,
        },
      ]
    },
    initTableListFields() {
      const methodEnum = {
        GET: {
          label: 'GET',
          type: 'info',
        },
        POST: {
          label: 'POST',
          type: 'success',
        },
        GRPC: {
          label: 'GRPC',
          type: 'primary',
        },
        PUT: {
          label: 'PUT',
          type: 'warning',
        },
        DELETE: {
          label: 'DELETE',
          type: 'danger',
        },
      }
      this.tableListFields = [
        // { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        { prop: 'title', label: '名称', width: 240, fixed: 'left' },
        { prop: 'description', label: '描述', minWidth: 150 },
        {
          prop: 'method',
          label: 'Method',
          width: 80,
          type: 'enum',
          enum: methodEnum,
        },
        { prop: 'path', label: 'API', minWidth: 150 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
