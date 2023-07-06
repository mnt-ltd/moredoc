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
        :table-data="navigations"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :tree-props="{ children: 'children' }"
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
      :title="navigation.id ? '编辑导航' : '新增导航'"
      :visible.sync="formNavigationVisible"
      width="640px"
    >
      <FormNavigation
        ref="navigationForm"
        :init-navigation="navigation"
        :trees="navigations"
        @success="formNavigationSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import {
  listNavigation,
  deleteNavigation,
  getNavigation,
} from '~/api/navigation'
import { genLinkHTML, parseQueryIntArray, categoryToTrees } from '~/utils/utils'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormNavigation from '~/components/FormNavigation.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormNavigation },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formNavigationVisible: false,
      search: {
        wd: '',
        page: 1,
        enable: [],
        size: 10,
      },
      navigations: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      navigation: { id: 0 },
    }
  },
  head() {
    return {
      title: `导航管理 - ${this.settings.system.sitename}`,
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
        this.listNavigation()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // await this.listNavigation()
  },
  methods: {
    async listNavigation() {
      this.loading = true
      const res = await listNavigation(this.search)
      if (res.status === 200) {
        let navigations = res.data.navigation || []
        navigations.map((item) => {
          item.title_html = genLinkHTML(item.title, item.href)
          if (item.color) {
            // 增加链接颜色
            item.title_html = item.title_html.replace(
              '<a',
              `<a style="color:${item.color}" `
            )
          }
        })

        let trees = categoryToTrees(navigations, false)
        this.navigations = trees
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
        this.listNavigation()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.navigation = { id: 0 }
      this.formNavigationVisible = true
      this.$nextTick(() => {
        this.$refs.navigationForm.reset()
      })
    },
    async editRow(row) {
      const res = await getNavigation({ id: row.id })
      if (res.status === 200) {
        this.navigation = res.data
        this.formNavigationVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formNavigationSuccess() {
      this.formNavigationVisible = false
      this.listNavigation()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}条】导航吗？会连带着子导航一起删除，删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteNavigation({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listNavigation()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除导航【${row.title}】吗？会连带着子导航一起删除，删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteNavigation({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listNavigation()
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
      ]
    },
    initTableListFields() {
      this.tableListFields = [
        // { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'title_html',
          label: '名称',
          minWidth: 200,
          fixed: 'left',
          type: 'html',
        },
        {
          prop: 'enable',
          label: '启用',
          width: 80,
          type: 'bool',
        },
        { prop: 'href', label: '链接', minWidth: 200 },
        { prop: 'target', label: '打开方式', width: 80 },
        { prop: 'sort', label: '排序', width: 80, type: 'number' },
        { prop: 'description', label: '描述', minWidth: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
