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
        <el-tabs
          v-model="activeTab"
          class="user-tabs"
          type="border-card"
          @tab-click="tabClick"
        >
          <el-tab-pane name="user-id">
            <span slot="label">
              <nuxt-link
                class="el-link el-link--default"
                :to="{
                  name: 'user-id',
                  params: { id: user.id },
                }"
                ><i class="el-icon-magic-stick"></i>&nbsp;动态</nuxt-link
              >
            </span>
            动态 动态
          </el-tab-pane>
          <el-tab-pane name="user-id-document">
            <span slot="label">
              <nuxt-link
                class="el-link el-link--default"
                :to="{
                  name: 'user-id-document',
                  params: { id: user.id },
                }"
                ><i class="el-icon-document"></i>&nbsp;文档</nuxt-link
              >
            </span>
            <nuxt-child />
          </el-tab-pane>
          <el-tab-pane name="user-id-favorite">
            <span slot="label">
              <nuxt-link
                class="el-link el-link--default"
                :to="{
                  name: 'user-id-favorite',
                  params: { id: user.id },
                }"
                ><i class="el-icon-star-off"></i>&nbsp;收藏</nuxt-link
              >
            </span>
            <nuxt-child />
          </el-tab-pane>
          <!-- <el-tab-pane name="user-id-fans">
            <span slot="label">
              <nuxt-link
                class="el-link el-link--default"
                :to="{
                  name: 'user-id-fans',
                  params: { id: user.id },
                }"
                ><i class="fa fa-user-o"></i>&nbsp;粉丝</nuxt-link
              >
            </span>
            <nuxt-child />
          </el-tab-pane>
          <el-tab-pane name="user-id-follow">
            <span slot="label">
              <nuxt-link
                class="el-link el-link--default"
                :to="{
                  name: 'user-id-follow',
                  params: { id: user.id },
                }"
                ><i class="fa fa-user"></i>&nbsp;关注</nuxt-link
              >
            </span>
            <nuxt-child />
          </el-tab-pane> -->
        </el-tabs>
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
      activeTab: this.$route.name,
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  async created() {},
  methods: {
    tabClick(e) {
      this.$router.push({
        name: e.name,
        params: { id: this.user.id },
      })
    },
  },
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
    .user-tabs.el-tabs--border-card {
      box-shadow: none;
      border: 0;
      border-radius: 5px;
      & > .el-tabs__header {
        background-color: $background-grey-light;
        border-top-left-radius: 5px;
        border-top-right-radius: 5px;
      }
      .el-tabs__item.is-active {
        margin-top: 0;
        border-top: 1px solid #dcdfe6;
        border-top-left-radius: 5px;
        border-top-right-radius: 5px;
      }
      & > .el-tabs__content {
        padding-top: 5px;
      }
    }
  }
}
</style>
