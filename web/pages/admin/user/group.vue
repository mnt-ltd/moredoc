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
        @onDelete="batchDelete"
        @onSearch="onSearch"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :table-data="groups"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :actions-min-width="200"
        @selectRow="selectRow"
        @deleteRow="deleteRow"
        @editRow="editRow"
        @permission="setGroupPermission"
      >
        <template slot="actions" slot-scope="scope">
          <el-button
            type="text"
            icon="el-icon-coordinate"
            size="small"
            @click="setGroupPermission(scope.row)"
            >后台授权</el-button
          >
        </template>
      </TableList>
    </el-card>
    <el-card v-if="total > 0" shadow="never" class="mgt-20px">
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
      :title="group.id ? '编辑分组' : '新增分组'"
      :visible.sync="formGroupVisible"
      width="640px"
    >
      <FormGroup :init-group="group" @success="success" />
    </el-dialog>
    <el-drawer
      :title="`【${group.title}】角色授权`"
      :visible.sync="formGroupPermissionVisible"
    >
      <div style="padding: 0 20px">
        <FormGroupPermission
          ref="groupPermission"
          :group-id="group.id"
          @success="updateGroupPermissionSuccess"
        />
      </div>
    </el-drawer>
  </div>
</template>

<script>
import { listGroup, deleteGroup, getGroup } from '~/api/group'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormGroup from '~/components/FormGroup.vue'
import FormGroupPermission from '~/components/FormGroupPermission.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormGroup, FormGroupPermission },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formGroupVisible: false,
      formGroupPermissionVisible: false,
      search: {
        wd: '',
        page: 1,
        size: 10,
      },
      groups: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      group: {},
    }
  },
  head() {
    return {
      title: `角色管理 - ${this.settings.system.sitename}`,
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
        this.listGroup()
      },
    },
  },
  async created() {
    this.initGroup()
    this.initSearchForm()
    this.initTableListFields()
    // await this.listGroup()
  },
  methods: {
    async listGroup() {
      this.loading = true
      const res = await listGroup(this.search)
      if (res.status === 200) {
        const groups = res.data.group
        try {
          for (let i = 0; i < groups.length; i++) {
            groups[i].disable_delete =
              groups[i].user_count > 0 || groups[i].is_default
          }
        } catch (error) {}
        this.groups = groups || []
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
    updateGroupPermissionSuccess() {
      // 权限设置成功，需要：
      // 1. 隐藏设置功能
      this.formGroupPermissionVisible = false
      // 2. vuex重载用户权限
      // 3. 刷新页面，以便使设置的权限生效
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
        this.listGroup()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.initGroup()
      this.formGroupVisible = true
    },
    setGroupPermission(row) {
      this.formGroupPermissionVisible = true
      this.$nextTick(() => {
        this.$refs.groupPermission.resetChecked()
        this.group = row
      })
    },
    async editRow(row) {
      const res = await getGroup({ id: row.id })
      if (res.status === 200) {
        this.group = res.data
        this.formGroupVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    success() {
      this.formGroupVisible = false
      this.listGroup()
    },
    setGroup() {
      this.formGroupVisible = false
    },
    deleteRow(row) {
      this.$confirm(
        `您是否要删除【${row.title}】分组？删除之后不可恢复！`,
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      ).then(async () => {
        const res = await deleteGroup({ id: row.id })
        if (res.status === 200) {
          this.$message.success('删除成功')
          this.listGroup()
        } else {
          this.$message.error(res.data.message)
        }
      })
    },
    batchDelete() {
      this.$confirm(
        `您是否要删除选择的【${this.selectedRow.length}个】分组?删除之后不可恢复！`,
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      ).then(async () => {
        const ids = this.selectedRow.map((item) => item.id)
        const res = await deleteGroup({ id: ids })
        if (res.status === 200) {
          this.$message.success('删除成功')
          this.listGroup()
        } else {
          this.$message.error(res.data.message)
        }
      })
    },
    selectRow(rows) {
      this.selectedRow = rows
    },
    initGroup() {
      this.group = {
        id: 0,
        sort: 0,
        color: '#000000',
        title: '',
        is_display: true,
        enable_upload: false,
        enable_comment_approval: false,
        is_default: false,
      }
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
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        { prop: 'title', label: '名称', width: 150, fixed: 'left' },
        { prop: 'sort', label: '排序', width: 80, type: 'number' },
        { prop: 'user_count', label: '用户数', width: 80, type: 'number' },
        { prop: 'color', label: '颜色', width: 120, type: 'color' },
        { prop: 'is_default', label: '是否默认', width: 80, type: 'bool' },
        {
          prop: 'enable_upload',
          label: '允许上传文档',
          width: 120,
          type: 'bool',
        },
        {
          prop: 'enable_comment_approval',
          label: '评论需审核',
          width: 120,
          type: 'bool',
        },
        { prop: 'is_display', label: '是否展示', width: 80, type: 'bool' },
        { prop: 'description', label: '描述', width: 250 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
