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
          </el-tab-pane>
          <nuxt-child />
        </el-tabs>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { getUser } from '~/api/user'
export default {
  data() {
    return {
      user: {
        id: 0,
      },
      activeTab: this.$route.name,
    }
  },
  head() {
    return {
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  created() {
    try {
      const id = parseInt(this.$route.params.id)
      this.user.id = id
      this.getUser()
    } catch (error) {}
  },
  methods: {
    tabClick(e) {
      this.$router.push({
        name: e.name,
        params: { id: this.user.id },
      })
    },
    async getUser() {
      const res = await getUser({ id: this.user.id })
      if (res.status === 200) {
        this.user = res.data || { id: 0 }
      }
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
        min-height: 805px;
      }
    }
  }
}
</style>
