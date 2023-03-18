<template>
  <el-card shadow="never">
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane
        v-for="item in categories"
        :key="'category-' + item.value"
        :label="item.label"
        :name="item.value"
      >
      </el-tab-pane>
    </el-tabs>
    <FormConfig :init-configs="configs" />
    <el-alert
      v-if="activeName == 'converter'"
      title="同时启用GZIP和SVGO，相对直接的SVG文件，总体可以节省85%左右的存储空间。（启用SVGO，需要全局安装node.js的SVGO模块）"
      type="info"
      :closable="false"
      show-icon
    >
    </el-alert>
  </el-card>
</template>

<script>
import { mapGetters } from 'vuex'
import { listConfig } from '~/api/config'
import FormConfig from '~/components/FormConfig.vue'
export default {
  components: { FormConfig },
  layout: 'admin',
  data() {
    return {
      activeName: 'system',
      configs: [],
      categories: [
        {
          label: '系统配置',
          value: 'system',
        },
        {
          label: '展示配置',
          value: 'display',
        },
        {
          label: '底链配置',
          value: 'footer',
        },
        {
          label: '验证码配置',
          value: 'captcha',
        },
        {
          label: '安全配置',
          value: 'security',
        },
        {
          label: '转换配置',
          value: 'converter',
        },
        {
          label: '下载配置',
          value: 'download',
        },
        {
          label: '积分配置',
          value: 'score',
        },
        {
          label: '邮箱配置',
          value: 'email',
        },
      ],
    }
  },
  head() {
    return {
      title: `系统设置 - ${this.settings.system.sitename}`,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  watch: {
    '$route.query': {
      handler() {
        this.activeName = this.$route.query.tab || 'system'
        this.loadConfig()
      },
      immediate: true,
    },
  },
  methods: {
    handleClick(tab) {
      this.activeName = tab.name
      this.$router.push({
        query: {
          tab: tab.name,
        },
      })
    },
    async loadConfig() {
      const res = await listConfig({ category: [this.activeName] })
      if (res.status === 200) {
        this.configs = res.data.config || []
      } else {
        this.configs = []
        this.$message.error(res.data.message)
      }
    },
  },
}
</script>
