<template>
  <div class="com-user-dynamic">
    <el-card shadow="never">
      <el-table v-loading="loading" :data="docs" style="width: 100%">
        <el-table-column prop="type" label="类型" width="100">
        </el-table-column>
        <el-table-column prop="created_at" label="时间" width="160">
        </el-table-column>
        <el-table-column prop="title" label="内容">
          <template slot-scope="scope">{{ scope.row.title || '-' }}</template>
        </el-table-column>
      </el-table>
    </el-card>
    <el-pagination
      :current-page="1"
      :page-size="10"
      layout="total,  prev, pager, next, jumper"
      :total="400"
    >
    </el-pagination>
  </div>
</template>

<script>
export default {
  name: 'UserDynamic',
  data() {
    return {
      docs: [
        {
          id: 1,
          title: 'Docker — 从入门到实战-BookStack.CN',
          type: 'pdf',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 3.5,
          created_at: '2022-01-02',
        },
        {
          id: 2,
          title: 'MongoDB简明教程',
          type: 'word',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 1.5,
          created_at: '2022-01-02',
        },
        {
          id: 3,
          title: 'TypeScript 官方文档',
          type: 'excel',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 4,
          created_at: '2022-01-02',
        },
        {
          id: 4,
          title:
            'DolphinPHP1.3.0完全开发手册-基于ThinkPHP5.0.20的快速开发框架-05221135',
          type: 'text',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 4.5,
          created_at: '2022-01-02',
        },
        {
          id: 1,
          title: 'Docker — 从入门到实战-BookStack.CN',
          type: 'pdf',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 3.5,
          created_at: '2022-01-02',
        },
        {
          id: 2,
          title: 'MongoDB简明教程',
          type: 'word',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 1.5,
          created_at: '2022-01-02',
        },
        {
          id: 3,
          title: 'TypeScript 官方文档',
          type: 'excel',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 4,
          created_at: '2022-01-02',
        },
        {
          id: 4,
          title:
            'DolphinPHP1.3.0完全开发手册-基于ThinkPHP5.0.20的快速开发框架-05221135',
          type: 'text',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 4.5,
          created_at: '2022-01-02',
        },
        {
          id: 1,
          title: 'Docker — 从入门到实战-BookStack.CN',
          type: 'pdf',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 3.5,
          created_at: '2022-01-02',
        },
        {
          id: 2,
          title: 'MongoDB简明教程',
          type: 'word',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 1.5,
          created_at: '2022-01-02',
        },
        {
          id: 3,
          title: 'TypeScript 官方文档',
          type: 'excel',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 4,
          created_at: '2022-01-02',
        },
        {
          id: 4,
          title:
            'DolphinPHP1.3.0完全开发手册-基于ThinkPHP5.0.20的快速开发框架-05221135',
          type: 'text',
          cover: '',
          view_count: 0,
          download_count: 1,
          score: 4.5,
          created_at: '2022-01-02',
        },
      ],
      total: 100,
      loading: false,
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  methods: {
    tabClick(tab) {
      this.activeTab = tab.name
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
      padding: 0;
      padding-bottom: 20px;
    }
  }
  .float-right {
    position: relative;
    top: -20px;
  }
  .el-table {
    margin-top: -1px;
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
