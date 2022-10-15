<template>
  <div>
    <!-- <el-form :inline="true" :model="search" class="demo-form-inline">
      <el-form-item label="审批人">
        <el-input v-model="formInline.user" placeholder="审批人"></el-input>
      </el-form-item>
      <el-form-item label="活动区域">
        <el-select v-model="formInline.region" placeholder="活动区域">
          <el-option label="区域一" value="shanghai"></el-option>
          <el-option label="区域二" value="beijing"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">查询</el-button>
      </el-form-item>
    </el-form> -->
    <TableList
      :table-data="users"
      :fields="fields"
      :show-actions="true"
      :show-view="true"
      :show-edit="true"
      :show-delete="true"
      :show-select="true"
    />
    <div class="text-right">
      <el-pagination
        class="mgt-20px"
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
      search: {
        wd: '',
        status: [],
        group_id: [],
        page: 1,
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
    }
  },
  created() {
    this.listUser()
  },
  methods: {
    async listUser() {
      const res = await listUser(this.search)
      if (res.status === 200) {
        this.users = res.data.user
        this.total = res.data.total
      }
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
  },
}
</script>
<style></style>
