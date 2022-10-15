<template>
  <div>
    <el-card shadow="never" class="search-card">
      <el-form :inline="true" :model="search">
        <el-form-item label="关键字">
          <el-input
            v-model="search.wd"
            placeholder="请输入关键字"
            clearable
            @keydown.native.enter="onSearch"
          ></el-input>
        </el-form-item>
        <el-form-item label="用户组">
          <el-select
            v-model="search.group_id"
            placeholder="请选择用户组"
            multiple
            clearable
            filterable
          >
            <el-option label="区域一" value="1"></el-option>
            <el-option label="区域二" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select
            v-model="search.status"
            placeholder="请选择用户状态"
            multiple
            clearable
            filterable
          >
            <el-option label="区域一" value="1"></el-option>
            <el-option label="区域二" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="el-icon-search"
            :loading="loading"
            @click="onSearch"
            >查询</el-button
          >
        </el-form-item>
        <el-form-item>
          <el-button
            type="danger"
            icon="el-icon-delete"
            :disabled="selectedIds.length == 0"
            @click="batchDelete"
            >批量删除</el-button
          >
        </el-form-item>
      </el-form>
    </el-card>

    <el-card class="mgt-20px" shadow="never">
      <TableList
        :table-data="users"
        :fields="fields"
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
import TableList from '~/components/TableList.vue'
export default {
  components: { TableList },
  layout: 'admin',
  data() {
    return {
      loading: false,
      search: {
        wd: '',
        page: 1,
        status: [],
        group_id: [],
        size: 10,
      },
      users: [],
      total: 100,
      fields: [
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
      ],
      selectedIds: [],
    }
  },
  created() {
    this.listUser()
  },
  methods: {
    async listUser() {
      this.loading = true
      const res = await listUser(this.search)
      if (res.status === 200) {
        this.users = res.data.user
        this.total = res.data.total
      }
      this.loading = false
      console.log(res)
    },
    handleSizeChange(val) {
      this.search.size = val
      this.listUser()
    },
    handlePageChange(val) {
      this.search.page = val
      this.listUser()
    },
    onSearch() {
      this.search.page = 1
      this.listUser()
    },
    batchDelete() {
      console.log('batchDelete')
    },
  },
}
</script>
<style></style>
