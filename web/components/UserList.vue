<template>
  <div class="com-user-list">
    <el-row v-for="user in users" :key="'user-' + user.id" :gutter="20">
      <el-col :span="7">
        <div class="avatar">
          <nuxt-link :to="'/user/' + user.id">
            <el-avatar :src="user.avatar">
              <img src="/static/images/avatar.png" alt="" />
            </el-avatar>
          </nuxt-link>
        </div>
      </el-col>
      <el-col :span="17">
        <div class="info">
          <nuxt-link
            :to="'/user/' + user.id"
            class="el-link el-link--default"
            >{{ user.username }}</nuxt-link
          >
          <div class="doc-info">
            <span class="el-link el-link--primary">{{
              user.doc_count || 0
            }}</span>
            <span class="text-muted">篇文档</span>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { listUser } from '~/api/user'
export default {
  name: 'UserList',
  props: {
    limit: {
      type: Number,
      default: 5,
    },
    order: {
      type: String,
      default: 'doc_count desc',
    },
  },
  data() {
    return {
      users: [],
    }
  },
  watch: {
    limit() {
      this.getUsers()
    },
    order() {
      this.getUsers()
    },
  },
  created() {
    this.getUsers()
  },
  methods: {
    async getUsers() {
      const res = await listUser({
        limit: this.limit,
        sort: this.order,
      })
      if (res.status === 200) {
        this.users = res.data.user || []
      }
    },
  },
}
</script>

<style lang="scss">
.com-user-list {
  .el-row {
    border-bottom: 1px dashed #efefef;
    padding: 15px 0 10px;
    margin-left: 0 !important;
    margin-right: 0 !important;
    .text-muted {
      font-size: 13px;
    }
    .el-col:first-child {
      padding-left: 0 !important;
    }
    .doc-info {
      font-size: 13px;
      margin-top: 8px;
      .el-link {
        top: -2px;
      }
    }
    .el-avatar {
      border: 2px solid #ddd;
      padding: 3px;
      background-color: #fff;
      width: 55px;
      height: 55px;
      &:hover {
        border: 2px solid #409eff;
      }
      img {
        border-radius: 50%;
      }
    }
  }
}
</style>
