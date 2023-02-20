<template>
  <div class="com-user-dynamic">
    <el-card shadow="never">
      <el-table v-loading="loading" :data="dynamics" style="width: 100%">
        <el-table-column
          prop="created_at"
          label="时间"
          :width="isMobile ? 90 : 160"
        >
          <template slot-scope="scope">
            <el-tooltip
              :content="formatDatetime(scope.row.created_at)"
              placement="top"
            >
              <span>{{ formatRelativeTime(scope.row.created_at) }}</span>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column prop="title" label="内容">
          <template slot-scope="scope">
            <!-- eslint-disable-next-line vue/no-v-html -->
            <span v-html="scope.row.content"></span>
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
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getDynamics } from '~/api/user'
import { formatRelativeTime, formatDatetime } from '~/utils/utils'
export default {
  name: 'UserDynamic',
  props: {
    userId: {
      type: Number,
      default: 0,
    },
  },
  data() {
    return {
      loading: false,
      query: {
        page: parseInt(this.$route.query.page) || 1,
        size: 20,
      },
      dynamics: [],
      total: 0,
    }
  },
  computed: {
    ...mapGetters('device', ['isMobile']),
  },
  watch: {
    '$route.query': {
      handler(val) {
        this.query.page = parseInt(val.page) || 1
        this.getDynamics()
      },
      immediate: true,
    },
  },
  methods: {
    formatDatetime,
    formatRelativeTime,
    tabClick(tab) {
      this.activeTab = tab.name
    },
    pageChange(page) {
      this.query.page = page
      this.$router.push({
        name: 'user-id',
        params: {
          id: this.userId,
        },
        query: {
          page,
        },
      })
    },
    async getDynamics() {
      if (this.loading) return
      this.loading = true
      const res = await getDynamics({ ...this.query, id: this.userId })
      if (res.status === 200) {
        this.dynamics = res.data.dynamic || []
        this.total = res.data.total || 0
      }
      this.loading = false
    },
  },
}
</script>

<style lang="scss">
.com-user-dynamic {
  .el-card {
    .el-card__header {
      padding-left: 0;
      font-weight: 400;
    }
    .el-card__body {
      padding: 18px 0 20px;
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
