<template>
  <div class="com-user-card2">
    <el-row :gutter="20">
      <el-col :span="16" :xs="24">
        <div class="user">
          <div class="user-card-avatar">
            <nuxt-link :to="'/user/' + user.id">
              <UserAvatar :size="64" :user="user" />
            </nuxt-link>
          </div>
          <div class="user-profile">
            <h2 class="user-card-username">{{ user.username }}</h2>
            <div class="help-block signature">
              {{ user.signature || '暂无个性签名' }}
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="8" :xs="24">
        <div class="user-card-stat">
          <el-row class="help-block">
            <el-col :span="8">
              <div>文档</div>
              <div class="el-link el-link--primary">
                {{ user.doc_count || 0 }}
              </div>
            </el-col>
            <el-col :span="8">
              <div>收藏</div>
              <div class="el-link el-link--primary">
                {{ user.favorite_count || 0 }}
              </div>
            </el-col>
            <el-col :span="8"
              ><div>{{ settings.system.credit_name || '魔豆' }}</div>
              <div class="el-link el-link--primary">
                {{ user.credit_count || 0 }}
              </div>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { mapGetters } from 'vuex'
// 用户信息卡片
export default {
  name: 'UserCard2',
  props: {
    user: {
      type: Object,
      default: () => {
        return {
          id: 0,
          name: '',
          avatar: '',
          signature: '',
          doc_count: 0,
          favorite_count: 0,
          credit_count: 0,
        }
      },
    },
    hideActions: {
      type: Boolean,
      default: false,
    },
    hideSignature: {
      type: Boolean,
      default: false,
    },
  },
  data() {
    return {}
  },
  async created() {},
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  methods: {},
}
</script>
<style lang="scss" scoped>
.com-user-card2 {
  .user {
    display: flex;
    .user-card-avatar {
      width: 70px;
    }
    .user-profile {
      flex: 1;
      margin-left: 10px;
    }
  }
  .el-avatar {
    border: 2px solid #ddd;
    padding: 3px;
    background-color: #fff;
    width: 80px;
    height: 80px;
    &:hover {
      border: 2px solid #409eff;
    }
    img {
      border-radius: 50%;
    }
  }
  .user-card-username {
    margin: 0 0 10px;
    font-size: 25px;
  }
  .user-card-stat {
    text-align: center;
    .help-block {
      font-size: 14px;
      .el-link {
        font-size: 18px;
        margin-top: 8px;
      }
      .el-col {
        border-right: 1px solid #eee;
        &:first-child {
          border-left: 1px solid #eee;
        }
      }
    }
  }
  .signature {
    margin-top: 10px;
    font-size: 13px;
  }
}

@media screen and (max-width: $mobile-width) {
  .user-card-stat {
    margin-top: 20px;
    .help-block {
      font-size: 12px;
      .el-link {
        font-size: 14px;
        margin-top: 8px;
      }
    }
  }
}
</style>
