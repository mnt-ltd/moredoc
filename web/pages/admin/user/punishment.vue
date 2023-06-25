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
      />
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
import { listPunishment, getPunishment } from '~/api/punishment'
import { genLinkHTML, parseQueryIntArray } from '~/utils/utils'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormPunishment from '~/components/FormPunishment.vue'
import { punishmentTypeOptions } from '~/utils/enum'
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
      },
      punishments: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      punishment: { id: 0 },
    }
  },
  head() {
    return {
      title: `惩罚管理 - ${this.settings.system.sitename}`,
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
        this.listPunishment()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // await this.listPunishment()
  },
  methods: {
    async listPunishment() {
      this.loading = true
      const res = await listPunishment(this.search)
      if (res.status === 200) {
        let punishments = res.data.punishment || []
        punishments.map((item) => {
          item.title_html = genLinkHTML(item.title, item.link)
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
      this.search = { ...this.search, ...search, page: 1 }
      if (
        location.href.lastIndexOf(
          this.$router.resolve({
            query: this.search,
          }).href
        ) > -1
      ) {
        this.$router.push({
          query: this.search,
        })
      } else {
        this.listPunishment()
      }
    },
    onCreate() {
      this.punishment = { id: 0 }
      this.formPunishmentVisible = true
      this.$nextTick(() => {
        this.$refs.punishmentForm.reset()
      })
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
          type: 'text',
          label: '关键字',
          name: 'wd',
          placeholder: '请输入关键字',
        },
        {
          type: 'select',
          label: '类型',
          name: 'type',
          placeholder: '请选择惩罚类型',
          multiple: true,
          options: this.punishmentTypeOptions,
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
          prop: 'type',
          label: '类型',
          minWidth: 120,
          type: 'enum',
          enum: enumOptions,
        },
        {
          prop: 'enable',
          label: '是否启用',
          width: 80,
          type: 'bool',
        },
        {
          prop: 'user_id',
          label: '用户',
          minWidth: 150,
        },
        { prop: 'reason', label: '原因', minWidth: 250 },
        { prop: 'remark', label: '备注', minWidth: 250 },
        { prop: 'start_time', label: '开始时间', width: 160, type: 'datetime' },
        { prop: 'end_time', label: '结束时间', width: 160, type: 'datetime' },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style></style>
