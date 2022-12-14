<template>
  <div class="com-user-document">
    <el-card shadow="never">
      <div slot="header">
        <span>文档列表</span>
        <el-tabs class="float-right">
          <el-tab-pane name="latest">
            <span slot="label"><i class="el-icon-date"></i> 最新</span>
          </el-tab-pane>
          <el-tab-pane name="hot">
            <span slot="label"><i class="el-icon-view"></i> 浏览</span>
          </el-tab-pane>
          <el-tab-pane name="favorite">
            <span slot="label"><i class="el-icon-date"></i> 收藏</span>
          </el-tab-pane>
          <el-tab-pane name="download">
            <span slot="label"><i class="el-icon-download"></i> 下载</span>
          </el-tab-pane>
          <el-tab-pane name="pages">
            <span slot="label"><i class="el-icon-files"></i> 页数</span>
          </el-tab-pane>
        </el-tabs>
      </div>
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
              <img
                :src="'/static/images/' + scope.row.type + '_24.png'"
                alt=""
              />
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
        </el-table-column>
        <el-table-column prop="download_count" label="下载" width="70">
        </el-table-column>
        <el-table-column prop="favorite_count" label="收藏" width="70">
          <template slot-scope="scope">{{
            scope.row.favorite_count || '-'
          }}</template>
        </el-table-column>
        <el-table-column prop="page" label="页数" width="70">
          <template slot-scope="scope">{{ scope.row.page || '-' }}</template>
        </el-table-column>
        <el-table-column prop="size" label="大小" width="70">
          <template slot-scope="scope">{{ scope.row.size || '-' }}</template>
        </el-table-column>
        <el-table-column prop="created_at" label="上传" width="100">
        </el-table-column>
        <el-table-column label="操作" width="70" fixed="right">
          <template slot-scope="scope">
            <el-tooltip content="编辑文档" placement="top">
              <el-button type="text" icon="el-icon-edit"></el-button>
            </el-tooltip>
            <el-tooltip content="删除文档" placement="top">
              <el-button type="text" icon="el-icon-delete"></el-button>
            </el-tooltip>
          </template>
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
  name: 'UserDocument',
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
      title: 'MOREDOC · 魔豆文库，开源文库系统',
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
.com-user-document {
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
  .el-table {
    margin-top: -11px;
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
