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
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :loading="loading"
        :table-data="friendlinks"
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
      :title="friendlink.id ? '编辑友链' : '新增友链'"
      :visible.sync="formFriendlinkVisible"
      width="640px"
    >
      <FormFriendlink
        ref="friendlinkForm"
        :init-friendlink="friendlink"
        @success="formFriendlinkSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import {
  listFriendlink,
  deleteFriendlink,
  getFriendlink,
} from '~/api/friendlink'
import { genLinkHTML, parseQueryIntArray } from '~/utils/utils'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormFriendlink from '~/components/FormFriendlink.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormFriendlink },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formFriendlinkVisible: false,
      search: {
        wd: '',
        page: 1,
        enable: [],
        size: 10,
      },
      friendlinks: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      friendlink: { id: 0 },
    }
  },
  head() {
    return {
      title: `友链管理 - ${this.settings.system.sitename}`,
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
          ...parseQueryIntArray(this.$route.query, ['enable']),
        }
        this.listFriendlink()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // await this.listFriendlink()
  },
  methods: {
    async listFriendlink() {
      this.loading = true
      const res = await listFriendlink(this.search)
      if (res.status === 200) {
        let friendlinks = res.data.friendlink || []
        friendlinks.map((item) => {
          item.title_html = genLinkHTML(item.title, item.link)
        })
        this.friendlinks = friendlinks
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
        this.$router.push({
          query: this.search,
        })
      } else {
        this.listFriendlink()
      }
    },
    onCreate() {
      this.friendlink = { id: 0 }
      this.formFriendlinkVisible = true
      this.$nextTick(() => {
        this.$refs.friendlinkForm.reset()
      })
    },
    async editRow(row) {
      const res = await getFriendlink({ id: row.id })
      if (res.status === 200) {
        this.friendlink = res.data
        this.formFriendlinkVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formFriendlinkSuccess() {
      this.formFriendlinkVisible = false
      this.listFriendlink()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}条】友链吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteFriendlink({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listFriendlink()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除友链【${row.title}】吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteFriendlink({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listFriendlink()
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
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'enable',
          label: '状态',
          width: 80,
          type: 'bool',
          fixed: 'left',
        },
        {
          prop: 'title_html',
          label: '名称',
          minWidth: 150,
          fixed: 'left',
          type: 'html',
        },
        { prop: 'link', label: '链接', minWidth: 250 },
        { prop: 'sort', label: '排序', width: 80, type: 'number' },
        { prop: 'description', label: '描述', minWidth: 250 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
