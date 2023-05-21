<template>
  <div class="com-user-download">
    <el-table v-loading="loading" :data="downloads" style="width: 100%">
      <el-table-column prop="title" label="文档" min-width="200">
        <template slot-scope="scope">
          <nuxt-link
            target="_blank"
            :to="{
              name: 'document-id',
              params: { id: scope.row.document_id },
            }"
            class="el-link el-link--default doc-title"
          >
            <img :src="`/static/images/${scope.row.icon}_24.png`" alt="" />
            {{ scope.row.title || '已被删除' }}
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
      <el-table-column prop="page" label="页数" width="70">
        <template slot-scope="scope">{{ scope.row.pages || '-' }}</template>
      </el-table-column>
      <el-table-column prop="size" label="大小" width="100">
        <template slot-scope="scope">{{
          formatBytes(scope.row.size)
        }}</template>
      </el-table-column>
      <el-table-column prop="created_at" label="下载时间" width="160">
        <template slot-scope="scope">
          <el-tooltip
            :content="formatDatetime(scope.row.created_at)"
            placement="top"
          >
            <span>{{ formatRelativeTime(scope.row.created_at) }}</span>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
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
      class="mgt-20px"
      @current-change="pageChange"
    >
    </el-pagination>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getUserDownloads } from '~/api/user'
import {
  formatDatetime,
  formatRelativeTime,
  formatBytes,
  getIcon,
} from '~/utils/utils'
export default {
  name: 'UserDownload',
  props: {
    userId: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      downloads: [],
      total: 0,
      loading: false,
      query: {
        page: parseInt(this.$route.query.page) || 1,
        size: 20,
      },
    }
  },
  watch: {
    '$route.query.page': {
      handler(val) {
        this.query.page = parseInt(val) || 1
        this.getUserDownloads()
      },
      immediate: true,
    },
  },
  computed: {
    ...mapGetters('user', ['user']),
  },
  methods: {
    formatDatetime,
    formatRelativeTime,
    formatBytes,
    async getUserDownloads() {
      this.loading = true
      const res = await getUserDownloads({
        page: this.query.page,
        size: this.query.size,
      })
      if (res.status === 200) {
        let downloads = res.data.download || []
        downloads = downloads.map((item) => {
          item.score = item.score / 100 || 3.0
          try {
            item.icon = getIcon(item.ext)
          } catch (error) {}
          return item
        })
        this.downloads = downloads
        this.total = res.data.total || 0
      }
      this.loading = false
    },
    pageChange(page) {
      this.$router.push({
        query: {
          page,
        },
      })
    },
  },
}
</script>

<style lang="scss">
.com-user-download {
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
