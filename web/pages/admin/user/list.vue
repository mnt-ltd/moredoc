<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :disabled-delete="selectedRows.length == 0"
        :loading="loading"
        :default-search="search"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      />
    </el-card>

    <el-card class="mgt-20px" shadow="never">
      <TableList
        :table-data="users"
        :loading="loading"
        :fields="listFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :actions-min-width="100"
        @editRow="editRow"
        @viewRow="viewRow"
        @deleteRow="deleteRow"
        @selectRow="selectRow"
      >
        <template slot="actions" slot-scope="scope">
          <el-button
            type="text"
            size="small"
            icon="el-icon-setting"
            @click="setUser(scope.row)"
            >设置</el-button
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

    <el-dialog
      :close-on-click-modal="false"
      :title="user.id ? '设置用户' : '新增用户'"
      :visible.sync="formUserVisible"
      width="640px"
    >
      <FormUser
        ref="formUser"
        :init-user="user"
        :groups="groups"
        @success="success"
      />
    </el-dialog>
    <el-dialog
      :close-on-click-modal="false"
      title="编辑用户"
      :visible.sync="formUserProfileVisible"
      width="640px"
    >
      <FormUserProfile
        ref="formUserProfile"
        :init-user="user"
        @success="successProfile"
      />
    </el-dialog>
  </div>
</template>

<script>
import { deleteUser, getUser, listUser } from '~/api/user'
import { listGroup } from '~/api/group'
import { userStatusOptions } from '~/utils/enum'
import { parseQueryIntArray, genLinkHTML } from '~/utils/utils'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormUser from '~/components/FormUser.vue'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormUser },
  layout: 'admin',
  data() {
    return {
      userStatusOptions,
      loading: false,
      search: {
        wd: '',
        page: 1,
        status: [],
        group_id: [],
        size: 10,
      },
      formUserVisible: false,
      formUserProfileVisible: false,
      formRechargeVisible: false,
      groups: [],
      users: [],
      user: { id: 0 },
      total: 0,
      searchFormFields: [],
      listFields: [],
      selectedRows: [],
    }
  },
  head() {
    return {
      title: `用户管理 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    '$route.query': {
      immediate: true,
      async handler() {
        this.search = {
          ...this.search,
          ...this.$route.query,
          page: parseInt(this.$route.query.page) || 1,
          size: parseInt(this.$route.query.size) || 10,
          ...parseQueryIntArray(this.$route.query, ['group_id', 'status']),
        }
        if (this.groups.length === 0) {
          await this.listGroup()
        }
        // 这里要执行下初始化，避免数据请求回来了，但是表格字段还没初始化，导致列表布局错乱
        await this.initTableListFields()
        this.listUser()
      },
    },
  },
  async created() {
    await this.initSearchForm()
    await this.initTableListFields()
    await this.initSearchForm() // 请求完成用户组数据之后再初始化下搜索表单，因为下拉枚举需要用到用户组数据
  },
  methods: {
    async listUser() {
      this.loading = true
      const res = await listUser(this.search)
      if (res.status === 200) {
        let users = res.data.user || []
        users.map((item) => {
          item.username_html = genLinkHTML(item.username, `/user/${item.id}`)
          let groups = (item.group_id || []).map((id) => {
            let group = this.groups.find((group) => group.id === id)
            return group ? group.title : ''
          })
          item.group = groups.join(', ')
        })
        this.users = users
        this.total = res.data.total
      } else {
        this.$message.error(res.data.message)
      }
      this.loading = false
    },
    async listGroup() {
      const res = await listGroup({ field: ['id', 'title'] })
      if (res.status === 200) {
        this.groups = res.data.group
      } else {
        this.$message.error(res.data.message)
      }
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
        this.listUser()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.formUserVisible = true
      this.$nextTick(() => {
        this.$refs.formUser.reset()
      })
    },
    setUser(row) {
      this.formUserVisible = true
      this.$nextTick(() => {
        this.$refs.formUser.reset()
        this.user = { ...row }
      })
    },
    async editRow(row) {
      const res = await getUser({ id: row.id })
      if (res.status !== 200) {
        this.$message.error(res.data.message)
        return
      }
      this.formUserProfileVisible = true
      this.$nextTick(() => {
        this.$refs.formUserProfile.reset()
        this.user = { ...res.data }
      })
    },
    viewRow(row) {
      this.$router.push(`/user/${row.id}`)
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除用户【${row.username}】吗？删除之后不可恢复，请慎重操作！`,
        '告警',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      ).then(async () => {
        const res = await deleteUser({ id: row.id })
        if (res.status === 200) {
          this.$message.success('删除成功')
          this.listUser()
        } else {
          this.$message.error(res.data.message)
        }
      })
    },
    selectRow(rows) {
      this.selectedRows = rows
    },
    success() {
      this.formUserVisible = false
      this.listUser()
    },
    successProfile() {
      this.formUserProfileVisible = false
      this.listUser()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除【${this.selectedRows.length}个】用户？删除之后不可恢复，请慎重操作！`,
        '告警',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      ).then(async () => {
        const ids = this.selectedRows.map((item) => item.id)
        const res = await deleteUser({ id: ids })
        if (res.status === 200) {
          this.$message.success('删除成功')
          this.listUser()
        } else {
          this.$message.error(res.data.message)
        }
      })
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
          label: '用户组',
          name: 'group_id',
          placeholder: '请选择用户组',
          multiple: true,
          options: this.groups.map((item) => {
            return {
              label: item.title,
              value: item.id,
            }
          }),
        },
        // {
        //   type: 'select',
        //   label: '状态',
        //   name: 'status',
        //   placeholder: '请选择用户状态',
        //   multiple: true,
        //   options: this.userStatusOptions,
        // },
      ]
    },
    initTableListFields() {
      if (this.listFields.length > 0) return
      this.listFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'avatar',
          label: '头像',
          width: 80,
          type: 'avatar',
          fixed: 'left',
        },
        {
          prop: 'username_html',
          label: '用户名',
          width: 150,
          fixed: 'left',
          type: 'html',
        },
        {
          prop: 'group',
          label: '用户组',
          width: 150,
        },
        { prop: 'doc_count', label: '文档', width: 80, type: 'number' },
        { prop: 'credit_count', label: '积分', width: 100, type: 'number' },
        { prop: 'follow_count', label: '关注', width: 80, type: 'number' },
        { prop: 'fans_count', label: '粉丝', width: 80, type: 'number' },
        { prop: 'favorite_count', label: '收藏', width: 80, type: 'number' },
        { prop: 'comment_count', label: '评论', width: 80, type: 'number' },
        { prop: 'realname', label: '姓名', width: 150 },
        { prop: 'email', label: '邮箱', width: 200 },
        { prop: 'mobile', label: '电话', width: 200 },
        { prop: 'identity', label: '身份证', width: 250 },
        { prop: 'address', label: '地址', width: 250 },
        { prop: 'signature', label: '签名', width: 250 },
        { prop: 'created_at', label: '注册时间', width: 160, type: 'datetime' },
        { prop: 'register_ip', label: '注册IP', width: 160 },
        { prop: 'login_at', label: '最后登录', width: 160, type: 'datetime' },
        {
          prop: 'last_login_ip',
          label: '最后登录IP',
          width: 160,
        },
      ]
    },
  },
}
</script>
<style></style>
