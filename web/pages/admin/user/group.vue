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
            <el-option
              v-for="item in groups"
              :key="'group_' + item.id"
              :label="item.title"
              :value="item.id"
            ></el-option>
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
            <el-option
              v-for="item in userStatusOptions"
              :key="'status_' + item.value"
              :label="item.label"
              :value="item.value"
            ></el-option>
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
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :table-data="groups"
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
import { listGroup } from '~/api/group'
import TableList from '~/components/TableList.vue'
export default {
  components: { TableList },
  layout: 'admin',
  data() {
    return {
      search: {
        wd: '',
        page: 1,
        status: [],
        group_id: [],
        size: 10,
      },
      groups: [],
      total: 0,
      fields: [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'icon',
          label: '图标',
          width: 80,
          type: 'avatar',
          fixed: 'left',
        },
        { prop: 'title', label: '名称', width: 150, fixed: 'left' },
        { prop: 'sort', label: '排序', width: 80, type: 'number' },
        { prop: 'user_count', label: '用户数', width: 80 },
        { prop: 'color', label: '颜色', width: 120 },
        { prop: 'is_default', label: '是否默认', width: 80, type: 'bool' },
        { prop: 'is_display', label: '是否展示', width: 80, type: 'bool' },
        { prop: 'description', label: '描述', width: 250 },
        { prop: 'created_at', label: '注册时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ],
      selectedIds: [],
    }
  },
  async created() {
    await this.listGroup()
  },
  methods: {
    async listGroup() {
      this.loading = true
      const res = await listGroup(this.search)
      if (res.status === 200) {
        this.groups = res.data.group
        this.total = res.data.total
      } else {
        this.$message.error(res.data.message)
      }
      this.loading = false
    },
    handleSizeChange(val) {
      this.search.size = val
      this.listGroup()
    },
    handlePageChange(val) {
      this.search.page = val
      this.listGroup()
    },
    onSearch() {
      this.search.page = 1
      this.listGroup()
    },
    batchDelete() {
      console.log('batchDelete')
    },
  },
}
</script>
<style></style>
