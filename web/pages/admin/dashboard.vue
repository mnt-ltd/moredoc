<template>
  <div class="page-admin-dashboard">
    <el-card shadow="never">
      <div slot="header">数据统计</div>
      <el-descriptions class="margin-top" :column="2" border>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-tickets"></i>
            文档
          </template>
          {{ stats.document_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-chat-dot-square"></i>
            评论
          </template>
          {{ stats.comment_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-user"></i>
            用户
          </template>
          {{ stats.user_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-s-grid"></i>
            分类
          </template>
          {{ stats.category_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-warning"></i>
            举报
          </template>
          {{ stats.report_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-picture-outline"></i>
            横幅
          </template>
          {{ stats.banner_count || 0 }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label">
            <i class="el-icon-link"></i>
            友链
          </template>
          {{ stats.friendlink_count || 0 }}
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
    <el-card shadow="never" class="mgt-20px">
      <div slot="header">系统信息</div>
      <el-descriptions class="margin-top" :column="1" border>
        <el-descriptions-item>
          <template slot="label"> 操作系统 </template>
          {{ stats.os }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序名称 </template>
          moredoc · 魔豆文库
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序版本 </template>
          {{ stats.version }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序Hash </template>
          {{ stats.hash }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 构建时间 </template>
          {{ formatDatetime(stats.build_at) }}
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序开发 </template>
          深圳市摩枫网络科技有限公司 <b>M</b>orefun <b>N</b>etwork
          <b>T</b>echnology Co., <b>Ltd</b>.
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 服务支持邮箱 </template>
          <a href="mailto:truthhun@mnt.ltd" class="el-link el-link--primary"
            >truthhun@mnt.ltd</a
          >
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 服务支持官网 </template>
          <a
            href="https://mnt.ltd"
            class="el-link el-link--primary"
            target="_blank"
            title="摩枫网络科技 MNT.Ltd"
            >https://mnt.ltd</a
          >
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序使用手册 </template>
          <a
            href="https://www.bookstack.cn/books/moredoc"
            class="el-link el-link--primary"
            target="_blank"
            title="程序使用手册"
            >https://www.bookstack.cn/books/moredoc</a
          >
        </el-descriptions-item>
        <el-descriptions-item>
          <template slot="label"> 程序开源地址 </template>
          <ul class="opensource">
            <li>
              MNT：
              <a
                href="https://git.mnt.ltd/mnt/moredoc"
                class="el-link el-link--primary"
                target="_blank"
                title="摩枫Git"
                >https://git.mnt.ltd/mnt/moredoc</a
              >
            </li>
            <li>
              Gitee：
              <a
                href="https://git.mnt.ltd"
                class="el-link el-link--primary"
                target="_blank"
                title="Gitee"
                >https://gitee.com/mnt-ltd/moredoc</a
              >
            </li>
            <li>
              Github：
              <a
                href="https://github.com/mnt-ltd/moredoc"
                class="el-link el-link--primary"
                target="_blank"
                title="Github"
                >https://github.com/mnt-ltd/moredoc</a
              >
            </li>
          </ul>
        </el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>
</template>

<script>
import { getStats } from '~/api/config'
import { formatDatetime } from '~/utils/utils'

export default {
  layout: 'admin',
  head() {
    return {
      title: `面板 - ${this.settings.system.sitename}`,
    }
  },
  data() {
    return {
      stats: {
        admin_count: 0,
        student_count: 0,
        company_count: 0,
        category_count: 0,
        article_count: 0,
        article_pending_count: 0,
        comment_count: 0,
        comment_pending_count: 0,
        banner_count: 0,
        friendlink_count: 0,
        user_pending_count: 0,
        os: '-',
        version: '-',
        hash: '-',
        build_at: '',
      },
    }
  },
  computed: {
    settings() {
      return this.$store.state.setting.settings
    },
  },
  created() {
    this.getStats()
  },
  methods: {
    formatDatetime,
    async getStats() {
      const res = await getStats()
      if (res.status === 200) {
        this.stats = {
          ...this.stats,
          ...res.data,
        }
      }
    },
  },
}
</script>
<style lang="scss">
.page-admin-dashboard {
  .el-descriptions-item__label.is-bordered-label {
    width: 150px;
  }
  .systeminfo {
    b {
      color: crimson;
    }
  }
  .opensource .el-link {
    margin-top: -3px;
  }
}
</style>
