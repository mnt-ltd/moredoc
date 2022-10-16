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
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :table-data="friendlinks"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        @selectRow="selectRow"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div class="text-right">
        <el-pagination
          background
          :current-page="search.page"
          :page-sizes="[10, 20, 50, 100, 200]"
          :page-size="search.size"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        >
        </el-pagination>
      </div>
    </el-card>

    <!-- <el-dialog
      :title="group.id ? '编辑分组' : '新增分组'"
      :init-group="group"
      :visible.sync="formGroupVisible"
    >
      <FormGroup />
    </el-dialog> -->
  </div>
</template>

<script>
import { listFriendlink } from '~/api/friendlink'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
export default {
  components: { TableList, FormSearch },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formGroupVisible: false,
      search: {
        wd: '',
        page: 1,
        status: [],
        size: 10,
      },
      friendlinks: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      group: {},
    }
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    await this.listFriendlink()
  },
  methods: {
    async listFriendlink() {
      this.loading = true
      const res = await listFriendlink(this.search)
      if (res.status === 200) {
        this.friendlinks = res.data.friendlink
        this.total = res.data.total
      } else {
        this.$message.error(res.data.message)
      }
      this.loading = false
    },
    handleSizeChange(val) {
      this.search.size = val
      this.listFriendlink()
    },
    handlePageChange(val) {
      this.search.page = val
      this.listFriendlink()
    },
    onSearch(search) {
      this.search = { ...this.search, page: 1, ...search }
      this.listFriendlink()
    },
    onCreate() {
      this.formGroupVisible = true
    },
    setGroup() {
      this.formGroupVisible = false
    },
    batchDelete() {
      console.log('batchDelete')
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
            { label: '启用', value: 0 },
            { label: '禁用', value: 1 },
          ],
        },
      ]
    },
    initTableListFields() {
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        { prop: 'title', label: '名称', width: 150, fixed: 'left' },
        {
          prop: 'status',
          label: '状态',
          width: 80,
          type: 'enum',
          enum: { 1: '禁用', 0: '启用' },
        },
        { prop: 'sort', label: '排序', width: 80, type: 'number' },
        { prop: 'description', label: '描述', width: 250 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
