<template>
  <div>
    <!-- <el-table :data="users" style="width: 100%">
      <el-table-column type="selection" width="55"> </el-table-column>
      <el-table-column prop="id" label="ID" width="80"> </el-table-column>
      <el-table-column prop="avatar" label="头像" width="75">
        <template slot-scope="scope">
          <el-avatar :size="45" :src="scope.row.avatar">
            <img src="/static/images/blank.png" />
          </el-avatar>
        </template>
      </el-table-column>
      <el-table-column prop="username" label="用户名" width="120">
      </el-table-column>
      <el-table-column prop="realname" label="姓名" width="120">
        <template slot-scope="scope">
          {{ scope.row.realname || '-' }}
        </template>
      </el-table-column>
      <el-table-column prop="email" label="邮箱" width="120">
        <template slot-scope="scope">
          {{ scope.row.email || '-' }}
        </template>
      </el-table-column>
      <el-table-column prop="mobile" label="电话" width="120">
        <template slot-scope="scope">
          {{ scope.row.mobile || '-' }}
        </template>
      </el-table-column>
      <el-table-column prop="address" label="地址" width="150">
        <template slot-scope="scope">
          {{ scope.row.address || '-' }}
        </template>
      </el-table-column>
      <el-table-column prop="signature" label="签名" width="150">
        <template slot-scope="scope">
          {{ scope.row.signature || '-' }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" min-width="100">
        <template slot-scope="scope">
          <el-button
            type="text"
            size="small"
            icon="el-icon-view"
            @click="handleClick(scope.row)"
            >查看</el-button
          >
          <el-button type="text" size="small" icon="el-icon-edit"
            >编辑</el-button
          >
          <el-button type="text" size="small" icon="el-icon-delete"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table> -->
    <TableList
      :table-data="users"
      :fields="fields"
      :show-actions="true"
      :show-view="true"
      :show-edit="true"
      :show-delete="true"
      :show-select="true"
    />
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
      users: [],
      total: 0,
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
      const res = await listUser()
      if (res.status === 200) {
        this.users = res.data.user
        this.total = res.data.total
      }
      console.log(res)
    },
  },
}
</script>
<style></style>
