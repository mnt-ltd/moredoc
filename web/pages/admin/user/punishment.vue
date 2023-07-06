<template>
  <div>
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="true"
        :show-delete="false"
        :disabled-delete="selectedRow.length === 0"
        :default-search="search"
        @onSearch="onSearch"
        @onCreate="onCreate"
      >
        <template slot="inputs">
          <el-form-item label="用户">
            <el-select
              v-model="search.user_id"
              filterable
              multiple
              remote
              reserve-keyword
              placeholder="请输入用户名"
              :remote-method="remoteSearchUser"
              :loading="loading"
            >
              <el-option
                v-for="user in users"
                :key="'userid' + user.id"
                :label="user.username"
                :value="user.id"
              >
              </el-option>
            </el-select>
          </el-form-item>
        </template>
        <template slot="buttons">
          <el-form-item>
            <el-tooltip
              class="item"
              effect="dark"
              content="批量取消处罚"
              placement="top"
            >
              <el-button
                type="warning"
                @click="batchCancelPunishment"
                :disabled="selectedRow.length === 0"
                icon="el-icon-edit"
                >批量取消</el-button
              >
            </el-tooltip>
          </el-form-item>
        </template>
      </FormSearch>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :loading="loading"
        :table-data="punishments"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="false"
        :show-select="true"
        :actions-min-width="80"
        @selectRow="selectRow"
        @editRow="editRow"
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
      :title="punishment.id ? '编辑惩罚' : '新增惩罚'"
      :visible.sync="formPunishmentVisible"
      width="640px"
    >
      <FormPunishment
        ref="punishmentForm"
        :init-punishment="punishment"
        @success="formPunishmentSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import {
  listPunishment,
  getPunishment,
  cancelPunishment,
} from '~/api/punishment'
import { genLinkHTML, parseQueryIntArray } from '~/utils/utils'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormPunishment from '~/components/FormPunishment.vue'
import { punishmentTypeOptions } from '~/utils/enum'
import { listUser } from '~/api/user'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormPunishment },
  layout: 'admin',
  data() {
    return {
      punishmentTypeOptions,
      loading: false,
      formPunishmentVisible: false,
      search: {
        wd: '',
        page: 1,
        enable: [],
        size: 10,
        user_id: [],
      },
      punishments: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      punishment: { id: 0 },
      users: [],
    }
  },
  head() {
    return {
      title: `处罚管理 - ${this.settings.system.sitename}`,
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
          ...parseQueryIntArray(this.$route.query, ['user_id']),
        }
        this.listPunishment()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // await this.listPunishment()
    if ((this.search.user_id || []).length > 0) {
      this.searchUser('', this.search.user_id)
    }
  },
  methods: {
    async remoteSearchUser(wd) {
      this.searchUser(wd)
    },
    async searchUser(wd, userId = []) {
      const res = await listUser({
        page: 1,
        size: 10,
        wd: wd,
        id: userId || [],
        field: ['id', 'username'],
      })
      if (res.status === 200) {
        this.users = res.data.user || []
      }
    },
    async batchCancelPunishment() {
      let res = await this.$confirm(
        `您确定要取消选中的${this.selectedRow.length}条处罚吗？`,
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
      if (res) {
        const ids = this.selectedRow.map((item) => item.id)
        const res = await cancelPunishment({ id: ids })
        if (res.status === 200) {
          this.$message.success('批量取消成功')
          this.listPunishment()
        } else {
          this.$message.error(res.data.message)
        }
      }
    },
    async listPunishment() {
      this.loading = true
      const res = await listPunishment(this.search)
      if (res.status === 200) {
        let punishments = res.data.punishment || []
        punishments.map((item) => {
          item.user_html = genLinkHTML(item.username, `/user/${item.id}`)
        })
        this.punishments = punishments
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
      this.search = {
        ...this.search,
        ...search,
        user_id: this.search.user_id,
        page: 1,
      }
      if (
        location.pathname + location.search ===
        this.$router.resolve({
          query: this.search,
        }).href
      ) {
        this.listPunishment()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.punishment = { id: 0, type: [], enable: true }
      this.formPunishmentVisible = true
    },
    async editRow(row) {
      const res = await getPunishment({ id: row.id })
      if (res.status === 200) {
        this.punishment = res.data
        this.formPunishmentVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formPunishmentSuccess() {
      this.formPunishmentVisible = false
      this.listPunishment()
    },
    selectRow(rows) {
      this.selectedRow = rows
    },
    initSearchForm() {
      this.searchFormFields = [
        {
          type: 'select',
          label: '类型',
          name: 'type',
          placeholder: '请选择惩罚类型',
          multiple: true,
          options: this.punishmentTypeOptions,
        },
        {
          type: 'text',
          label: '关键字',
          name: 'wd',
          placeholder: '请输入关键字',
        },
      ]
    },
    initTableListFields() {
      const enumOptions = {}
      this.punishmentTypeOptions.map((item) => {
        enumOptions[item.value] = item
      })

      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'enable',
          label: '启用处罚',
          width: 80,
          type: 'bool',
        },
        {
          prop: 'type',
          label: '类型',
          minWidth: 120,
          type: 'enum',
          enum: enumOptions,
        },
        {
          prop: 'user_html',
          label: '用户',
          minWidth: 150,
          type: 'html',
        },
        { prop: 'end_time', label: '截止时间', width: 160, type: 'datetime' },
        { prop: 'reason', label: '原因', minWidth: 250 },
        { prop: 'remark', label: '备注', minWidth: 250 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
