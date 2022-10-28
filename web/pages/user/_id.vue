<template>
  <div class="page page-user">
    <el-row :gutter="20">
      <el-col :span="6">
        <el-card shadow="never">
          <div slot="header">个人主页</div>
          <user-card :user="user" />
        </el-card>
        <el-card shadow="never" class="mgt-20px shared-user">
          <div slot="header">分享达人</div>
          <user-list />
        </el-card>
      </el-col>
      <el-col :span="18" class="user-right">
        <el-card shadow="never">
          <el-tabs>
            <el-tab-pane name="doc">
              <span slot="label"><i class="el-icon-document"></i> 文档</span>
              用户上传的文档列表
            </el-tab-pane>
            <el-tab-pane name="favorite">
              <span slot="label"><i class="el-icon-star-off"></i> 收藏</span>
              用户收藏的文档列表
            </el-tab-pane>
            <el-tab-pane name="dynamic">
              <span slot="label"><i class="el-icon-magic-stick"></i> 动态</span>
              用户的动态列表
            </el-tab-pane>
            <!-- <el-tab-pane name="follow">
              <span slot="label"><i class="el-icon-user-solid"></i> 关注</span>
            </el-tab-pane>
            <el-tab-pane name="fans">
              <span slot="label"><i class="el-icon-user"></i> 粉丝</span>
            </el-tab-pane> -->
          </el-tabs>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'PageUser',
  data() {
    return {
      user: {
        id: 1,
        username: 'admin',
        avatar:
          'https://static.wenkuzhijia.cn/56f606c672f57d7700d4e7fcd4bdc5d3.png',
        signature: '',
        doc_count: 102,
        favorite_count: 12,
        money: 100,
      },
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔刀文库，开源文库系统',
    }
  },
  async created() {},
  methods: {},
}
</script>
<style lang="scss">
.page-user {
  .shared-user {
    .el-card__body {
      padding-top: 0;
    }
  }
  .user-right {
    .el-card__body {
      padding-top: 0;
      .el-tabs__item {
        height: 60px;
        line-height: 60px;
      }
    }
  }
}
</style>
