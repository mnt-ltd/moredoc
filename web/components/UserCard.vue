<template>
  <div class="com-user-card">
    <div class="user-card-avatar">
      <!-- 如果是用户自己，则点击头像可以更换头像 -->
      <nuxt-link :to="'/user/' + user.id">
        <el-avatar :src="user.avatar">
          <!-- 默认 -->
          <img src="/static/images/avatar.png" alt="" />
        </el-avatar>
      </nuxt-link>
    </div>
    <div class="user-card-username">
      <strong>{{ user.username }}</strong>
      <small v-if="!hideActions" class="text-muted"
        >[
        <el-tooltip content="设置" placement="top">
          <a href="javascript:void(0);"><i class="el-icon-setting"></i></a>
        </el-tooltip>
        <span> | </span>
        <el-tooltip content="退出登录" placement="top">
          <a href="javascript:void(0);"
            ><i class="el-icon-switch-button"></i
          ></a>
        </el-tooltip>
        ]</small
      >
    </div>
    <div class="user-card-stat">
      <el-row class="help-block">
        <el-col :span="8">
          <div>文档</div>
          <div class="el-link el-link--primary">{{ user.doc_count || 0 }}</div>
        </el-col>
        <el-col :span="8">
          <div>收藏</div>
          <div class="el-link el-link--primary">
            {{ user.favorite_count || 0 }}
          </div>
        </el-col>
        <el-col :span="8"
          ><div>财富</div>
          <div class="el-link el-link--primary">
            {{ user.money || 0 }}
          </div>
        </el-col>
      </el-row>
    </div>
    <div v-if="!hideSignature" class="user-card-signature">
      <div>个性签名</div>
      <div class="help-block">
        {{ user.signature || '暂无个性签名' }}
      </div>
    </div>
  </div>
</template>
<script>
// 用户信息卡片
export default {
  name: 'UserCard',
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
          money: 0,
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
.com-user-card {
  text-align: center;
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
    margin: 10px 0;
    font-size: 16px;
    small {
      font-size: 13px;
    }
  }
  .user-card-stat {
    margin: 20px 0;
    .help-block {
      font-size: 12px;
      .el-link {
        font-size: 14px;
        margin-top: 5px;
      }
      .el-col {
        border-right: 1px solid #ddd;
        &:first-child {
          border-left: 1px solid #ddd;
        }
      }
    }
  }
  .user-card-signature {
    text-align: left;
    .help-block {
      margin-top: 10px;
      font-size: 13px;
      text-indent: 2em;
    }
  }
}
</style>
