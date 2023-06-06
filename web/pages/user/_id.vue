<template>
  <div class="page page-user">
    <el-row>
      <el-col :span="24">
        <el-card shadow="never">
          <user-card2 :user="user" />
          <el-tabs
            v-model="activeTab"
            class="user-tabs mgt-20px"
            @tab-click="tabClick"
            type="card"
          >
            <el-tab-pane name="user-id">
              <span slot="label">
                <nuxt-link
                  class="el-link el-link--default"
                  :to="{
                    name: 'user-id',
                    params: { id: user.id },
                  }"
                  ><i class="el-icon-document"></i>&nbsp;文档</nuxt-link
                >
              </span>
            </el-tab-pane>
            <nuxt-child :user="user" />
          </el-tabs>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
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
      title: `[个人主页] ${this.user.username} - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: `个人主页,${this.settings.system.sitename},${this.settings.system.keywords}`,
        },
        {
          hid: 'description',
          name: 'description',
          content: `${this.settings.system.description}`,
        },
      ],
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
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
  .user-tabs {
    & > .el-tabs__content {
      min-height: calc(100vh - 280px);
    }
    .el-tabs__header {
      margin-bottom: 20px;
    }
  }
}

@media screen and (max-width: $mobile-width) {
  .page-user {
    .user-profile {
      width: 100% !important;
    }
    .user-right {
      margin-top: 15px;
      width: 100% !important;
      .user-tabs.el-tabs--border-card > .el-tabs__content {
        min-height: unset;
      }
    }
  }
}
</style>
