<template>
  <div class="page-admin-article">
    <el-card shadow="never" class="search-card">
      <FormSearch
        :fields="searchFormFields"
        :loading="loading"
        :show-create="true"
        :show-delete="true"
        :disabled-delete="selectedRow.length === 0"
        :default-search="search"
        @onSearch="onSearch"
        @onCreate="onCreate"
        @onDelete="batchDelete"
      />
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <TableList
        :loading="loading"
        :table-data="articles"
        :fields="tableListFields"
        :show-actions="true"
        :show-view="false"
        :show-edit="true"
        :show-delete="true"
        :show-select="true"
        @selectRow="selectRow"
        @editRow="editRow"
        @deleteRow="deleteRow"
      >
        <!-- 查看文章 -->
        <!-- <template slot="actions" slot-scope="scope">
          <nuxt-link
            target="_blank"
            :to="{
              name: 'article-id',
              params: { id: scope.row.identifier },
            }"
          >
            <el-button type="text" size="mini" icon="el-icon-view"
              >查看</el-button
            >
          </nuxt-link>
        </template> -->
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
    <el-drawer
      :title="article.id ? '编辑文章' : '新增文章'"
      :visible.sync="formArticleVisible"
      :size="'80%'"
      :wrapper-closable="true"
    >
      <FormArticle
        ref="articleForm"
        :init-article="article"
        @success="formSuccess"
      />
    </el-drawer>
  </div>
</template>

<script>
import { listArticle, deleteArticle, getArticle } from '~/api/article'
import TableList from '~/components/TableList.vue'
import FormSearch from '~/components/FormSearch.vue'
import FormArticle from '~/components/FormArticle.vue'
import { genLinkHTML } from '~/utils/utils'
import { mapGetters } from 'vuex'
export default {
  components: { TableList, FormSearch, FormArticle },
  layout: 'admin',
  data() {
    return {
      loading: false,
      formArticleVisible: false,
      search: {
        wd: '',
        page: 1,
        status: [],
        size: 10,
      },
      articles: [],
      total: 0,
      searchFormFields: [],
      tableListFields: [],
      selectedRow: [],
      article: {
        id: 0,
        title: '',
        identifier: '',
        keywords: '',
        description: '',
        content: '',
      },
    }
  },
  head() {
    return {
      title: `文章管理 - ${this.settings.system.sitename}`,
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
        this.listArticle()
      },
    },
  },
  async created() {
    this.initSearchForm()
    this.initTableListFields()
    // await this.listArticle()
  },
  methods: {
    async listArticle() {
      this.loading = true
      const res = await listArticle(this.search)
      if (res.status === 200) {
        let articles = res.data.article || []
        articles.map((item) => {
          item.title_html = genLinkHTML(
            item.title,
            `/article/${item.identifier}`
          )
        })
        this.articles = articles
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
      // this.listArticle()
    },
    handlePageChange(val) {
      this.search.page = val
      this.$router.push({
        query: this.search,
      })
      // this.listArticle()
    },
    onSearch(search) {
      this.search = { ...this.search, ...search, page: 1 }
      if (
        location.pathname + location.search ===
        this.$router.resolve({
          query: this.search,
        }).href
      ) {
        this.listArticle()
      } else {
        this.$router.push({
          query: this.search,
        })
      }
    },
    onCreate() {
      this.article = { id: 0 }
      this.formArticleVisible = true
      this.$nextTick(() => {
        try {
          this.$refs.articleForm.reset()
        } catch (error) {
          console.log(error)
        }
      })
    },
    async editRow(row) {
      const res = await getArticle({ id: row.id })
      if (res.status === 200) {
        this.article = res.data
        this.formArticleVisible = true
      } else {
        this.$message.error(res.data.message)
      }
    },
    formSuccess() {
      this.formArticleVisible = false
      this.listArticle()
    },
    batchDelete() {
      this.$confirm(
        `您确定要删除选中的【${this.selectedRow.length}条】文章吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const ids = this.selectedRow.map((item) => item.id)
          const res = await deleteArticle({ id: ids })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listArticle()
          } else {
            this.$message.error(res.data.message)
          }
        })
        .catch(() => {})
    },
    deleteRow(row) {
      this.$confirm(
        `您确定要删除文章【${row.title}】吗？删除之后不可恢复！`,
        '温馨提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
        }
      )
        .then(async () => {
          const res = await deleteArticle({ id: row.id })
          if (res.status === 200) {
            this.$message.success('删除成功')
            this.listArticle()
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
      ]
    },
    initTableListFields() {
      this.tableListFields = [
        { prop: 'id', label: 'ID', width: 80, type: 'number', fixed: 'left' },
        {
          prop: 'title_html',
          label: '标题',
          minWidth: 150,
          fixed: 'left',
          type: 'html',
        },
        { prop: 'identifier', label: '标识', width: 200 },
        { prop: 'view_count', label: '浏览', width: 80, type: 'number' },
        { prop: 'keywords', label: '关键字', width: 200 },
        { prop: 'description', label: '摘要', minWidth: 200 },
        { prop: 'created_at', label: '创建时间', width: 160, type: 'datetime' },
        { prop: 'updated_at', label: '更新时间', width: 160, type: 'datetime' },
      ]
    },
  },
}
</script>
<style lang="scss">
.page-admin-article {
  .el-drawer__body {
    padding: 0 20px;
  }
}
</style>
