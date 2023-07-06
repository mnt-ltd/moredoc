<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="false"
        :show-delete="true"
        :disabled-delete="selectedRow.length === 0"
        :default-search="search"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      >
        <template slot="buttons">
          <el-dropdown
            :disabled="selectedRow.length === 0"
            @command="checkComment"
          >
            <el-button type="primary">
              批量审批 <i class="el-icon-arrow-down el-icon--right"></i>
            </el-button>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item :command="1">审核通过</el-dropdown-item>
              <el-dropdown-item :command="2">审核拒绝</el-dropdown-item>
              <el-dropdown-item :command="0">变为待审</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </FormSearch>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :loading="loading"
        :table-data="comments"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        :actions-min-width="160"
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
      v-if="comment.id > 0"
      title="评论编审"
      :visible.sync="formCommentVisible"
      width="640px"
    >
      <FormCommentCheck
        ref="formComment"
        :comment="comment"
        @success="formCommentSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import {
  listComment,
  deleteComment,
  getComment,
  checkComment,
} from '~/api/comment'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import { parseQueryIntArray, genLinkHTML } from '~/utils/utils'
export default {
  components: { TableList, FormSearch },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formCommentVisible: false,
      search: {
        wd: '',
        page: 1,
        status: [],
        size: 10,
        order: 'id desc',
      },
      comments: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      comment: { id: 0 },
    }
  },
  head() {
    return {
      title: `评论管理 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    settings() {
      return this.$store.state.setting.settings
    },
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
          ...parseQueryIntArray(this.$route.query, ['status']),
        }
        this.listComment()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
  },
  methods: {
    async listComment() {
      this.loading = true
      const res = await listComment(this.search)
      if (res.status === 200) {
        this.comments = (res.data.comment || []).map((item) => {
          item.username = item.user.username
          item.document_title_html = genLinkHTML(
            item.document_title,
            `/document/${item.document_id}`
          )
          item.username_html = genLinkHTML(
            item.username,
            `/user/${item.user_id}`
          )
          return item
        })
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
        this.listComment()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.comment = { id: 0 }
      this.formCommentVisible = true
      this.$nextTick(() => {
        this.$refs.commentForm.reset()
      })
    },
    async editRow(row) {
      const res = await getComment({ id: row.id })
      if (res.status === 200) {
        this.comment = res.data
        this.formCommentVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formCommentSuccess() {
      this.formCommentVisible = false
      this.listComment()
    },
    async checkComment(cmd) {
      const res = await checkComment({
        id: this.selectedRow.map((item) => item.id),
        status: cmd,
      })
      if (res.status === 200) {
        this.$message.success('审批成功')
        this.listComment()
        return
      }
      this.$message.error(res.data.message || '审批失败')
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}条】评论吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteComment({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listComment()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(`您确定要删除该评论吗？删除之后不可恢复！`, '温馨提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
        .then(async () => {
          const res = await deleteComment({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listComment()
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
          name: 'status',
          placeholder: '请选择状态',
          multiple: true,
          options: [
            { label: '审核拒绝', value: 2 },
            { label: '审核通过', value: 1 },
            { label: '待审核', value: 0 },
          ],
        },
      ]
    },
    initTableListFields() {
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'status',
          label: '状态',
          width: 80,
          type: 'enum',
          fixed: 'left',
          enum: {
            2: { label: '审核拒绝', value: 2, type: 'danger' },
            1: { label: '审核通过', value: 1, type: 'success' },
            0: { label: '待审核', value: 0 },
          },
        },
        {
          prop: 'document_title_html',
          label: '文档',
          minWidth: 150,
          type: 'html',
        },
        { prop: 'content', label: '评论内容', minWidth: 150 },
        { prop: 'username_html', label: '评论人', minWidth: 150, type: 'html' },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
