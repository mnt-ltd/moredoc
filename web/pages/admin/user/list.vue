<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :disabled-delete="selectedIds.length == 0"
        :loading="loading"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      />
    </el-card>

    <el-card class="mgt-20px" shadow="never">
      <TableList
        :table-data="users"
        :fields="listFields"
        :show-actions="true"
        :show-view="true"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
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
  </div>
</template>

<script>
import { listUser } from '~/api/user'
import { listGroup } from '~/api/group'
import { userStatusOptions } from '~/utils/enum'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
export default {
  components: { TableList, FormSearch },
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
      groups: [],
      users: [],
      total: 100,
      searchFormFields: [],
      listFields: [],
      selectedIds: [],
    }
  },
  async created() {
    this.initTableListFields()
    await this.listGroup()
    await this.initSearchForm()
    await this.listUser()
  },
  methods: {
    async listUser() {
      this.loading = true
      const res = await listUser(this.search)
      if (res.status === 200) {
        this.users = res.data.user
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
      this.listUser()
    },
    handlePageChange(val) {
      this.search.page = val
      this.listUser()
    },
    onSearch(search) {
      this.search = { ...this.search, ...search }
      this.search.page = 1
      this.listUser()
    },
    onCreate() {
      console.log('onCreate')
    },
    batchDelete() {
      console.log('batchDelete')
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
        {
          type: 'select',
          label: '状态',
          name: 'status',
          placeholder: '请选择用户状态',
          multiple: true,
          options: this.userStatusOptions,
        },
      ]
    },
    initTableListFields() {
      this.listFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'avatar',
          label: '头像',
          width: 80,
          type: 'avatar',
          fixed: 'left',
        },
        { prop: 'username', label: '用户名', width: 150, fixed: 'left' },
        { prop: 'doc_count', label: '文档', width: 80, type: 'number' },
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
