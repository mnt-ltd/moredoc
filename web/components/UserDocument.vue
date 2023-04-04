<template>
  <div class="com-user-document">
    <el-card shadow="never">
      <el-table v-loading="loading" :data="docs" style="width: 100%">
        <el-table-column prop="title" label="名称" min-width="300">
          <template slot-scope="scope">
            <nuxt-link
              target="_blank"
              :to="{
                name: 'document-id',
                params: { id: scope.row.id },
              }"
              class="el-link el-link--default doc-title"
            >
              <img :src="`/static/images/${scope.row.icon}_24.png`" alt="" />
              {{ scope.row.title }}
            </nuxt-link>
          </template>
        </el-table-column>
        <el-table-column prop="score" label="评分" width="110">
          <template slot-scope="scope">
            <el-rate
              :value="scope.row.score || 0.0"
              disabled
              score-template="{value}"
            ></el-rate>
          </template>
        </el-table-column>
        <el-table-column prop="view_count" label="浏览" width="70">
          <template slot-scope="scope">{{
            scope.row.view_count || 0
          }}</template>
        </el-table-column>
        <el-table-column prop="download_count" label="下载" width="70">
          <template slot-scope="scope">{{
            scope.row.download_count || 0
          }}</template>
        </el-table-column>
        <el-table-column prop="favorite_count" label="收藏" width="70">
          <template slot-scope="scope">{{
            scope.row.favorite_count || 0
          }}</template>
        </el-table-column>
        <el-table-column prop="page" label="页数" width="70">
          <template slot-scope="scope">{{ scope.row.pages || '-' }}</template>
        </el-table-column>
        <el-table-column prop="page" label="价格" width="70">
          <template slot-scope="scope">{{ scope.row.price || '0' }}</template>
        </el-table-column>
        <el-table-column prop="size" label="大小" width="100">
          <template slot-scope="scope">{{
            formatBytes(scope.row.size)
          }}</template>
        </el-table-column>
        <el-table-column prop="created_at" label="上传" width="100">
          <template slot-scope="scope">
            <el-tooltip
              :content="formatDatetime(scope.row.created_at)"
              placement="top"
            >
              <span>{{ formatRelativeTime(scope.row.created_at) }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column
          v-if="userId === user.id"
          label="操作"
          width="70"
          fixed="right"
        >
          <template slot-scope="scope">
            <el-tooltip content="编辑文档" placement="top">
              <el-button
                type="text"
                @click="updateDocument(scope.row)"
                icon="el-icon-edit"
              ></el-button>
            </el-tooltip>
            <el-tooltip content="删除文档" placement="top">
              <el-button
                type="text"
                icon="el-icon-delete"
                @click="deleteDocument(scope.row)"
              ></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-pagination
      v-if="total > 0"
      :current-page="query.page"
      :page-size="query.size"
      :layout="
        isMobile
          ? 'total, prev, pager, next'
          : 'total, prev, pager, next, jumper'
      "
      :pager-count="isMobile ? 5 : 7"
      :small="isMobile"
      :total="total"
      @current-change="pageChange"
    >
    </el-pagination>
    <el-dialog
      title="编辑文档"
      :visible.sync="updateDocumentVisible"
      :width="isMobile ? '95%' : '640px'"
    >
      <FormUpdateDocument
        :category-trees="categoryTrees"
        :init-document="document"
        :is-admin="false"
        @success="updateDocumentSuccess"
      />
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { deleteDocument, listDocument } from '~/api/document'
import {
  formatBytes,
  formatDatetime,
  formatRelativeTime,
  getIcon,
} from '~/utils/utils'

export default {
  name: 'UserDocument',
  props: {
    userId: {
      type: Number,
      required: true,
    },
  },
  data() {
    return {
      docs: [],
      total: 0,
      loading: false,
      query: {
        page: parseInt(this.$route.query.page) || 1,
        size: 15,
      },
      updateDocumentVisible: false,
      document: { id: 0 },
    }
  },
  computed: {
    ...mapGetters('user', ['user']),
    ...mapGetters('category', ['categoryTrees']),
    ...mapGetters('device', ['isMobile']),
  },
  watch: {
    '$route.query': {
      handler() {
        this.query.page = parseInt(this.$route.query.page) || 1
        this.getDocuments()
      },
      deep: true,
    },
  },
  created() {
    console.log('created')
    this.getDocuments()
  },
  methods: {
    formatBytes,
    formatDatetime,
    formatRelativeTime,
    updateDocument(row) {
      this.updateDocumentVisible = true
      const doc = { ...row }
      delete doc.icon
      this.document = doc
    },
    updateDocumentSuccess() {
      this.updateDocumentVisible = false
      this.getDocuments()
    },
    tabClick(tab) {
      this.activeTab = tab.name
    },
    async getDocuments() {
      if (this.userId === 0 || this.loading) return
      this.loading = true
      const res = await listDocument({
        ...this.query,
        user_id: this.userId,
      })
      if (res.status === 200) {
        const docs = res.data.document || []
        docs.map((item) => {
          item.score = item.score / 100 || 0.0
          try {
            item.icon = getIcon(item.ext)
          } catch (e) {
            console.log(e)
          }
          return item
        })
        this.docs = docs
        this.total = res.data.total || 0
      }
      this.loading = false
    },
    pageChange(page) {
      this.$router.push({
        query: { page },
      })
    },
    deleteDocument(row) {
      this.$confirm(`您确定要删除文档《${row.title}》吗？`, '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(async () => {
        const res = await deleteDocument({ id: row.id })
        if (res.status === 200) {
          this.$message({
            type: 'success',
            message: '删除成功!',
          })
          this.getDocuments()
        }
      })
    },
  },
}
</script>

<style lang="scss">
.com-user-document {
  .el-card {
    .el-card__header {
      padding-left: 0;
      font-weight: 400;
    }
    .el-card__body {
      padding: 15px 0 20px 0;
      a {
        line-height: 40px;
        display: inline-block;
        max-width: 100%;
      }
    }
  }
  .el-tabs__header {
    margin-bottom: 0;
  }
  .el-tabs__nav-wrap::after {
    background-color: transparent;
  }
  .el-tabs__item {
    height: 60px;
    line-height: 60px;
  }
  .el-tabs__item.is-active {
    border-top: 0 !important;
  }
  .doc-title {
    display: block;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    line-height: 180%;
    img {
      height: 18px;
      position: relative;
      top: 3px;
    }
  }
}
</style>
