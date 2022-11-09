<template>
  <el-card shadow="never">
    <el-tabs v-model="activeName" type="card" @tab-click="handleClick">
      <el-tab-pane
        v-for="item in categories"
        :key="'category-' + item.value"
        :label="item.label"
        :name="item.value"
      >
      </el-tab-pane>
    </el-tabs>
    <FormConfig :init-configs="configs" />
  </el-card>
</template>

<script>
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
        // {
        //   label: '用户配置',
        //   value: 'user',
        // },
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
      ],
    }
  },
  created() {
    this.loadConfig()
  },
  methods: {
    handleClick(tab) {
      this.activeName = tab.name
      this.loadConfig()
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
