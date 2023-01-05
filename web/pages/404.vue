<template>
  <div class="page page-notfound">
    <div class="notfound">
      <img src="/static/images/404.png" alt="404" />
      <h3 class="mgt-15px">Page Not Found</h3>
      <nuxt-link to="/">
        <el-button type="primary"
          ><i class="el-icon-s-home"></i> 返回首页</el-button
        >
      </nuxt-link>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  data() {
    return {}
  },
  head() {
    return {
      title: `404 - Not Found - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'keywords',
          name: 'keywords',
          content: '404,Not Found',
        },
        {
          hid: 'description',
          name: 'description',
          content: '内容不存在',
        },
      ],
    }
  },
  watch: {},
  async created() {},
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  methods: {},
}
</script>
<style>
.page-notfound {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  text-align: center;
  padding: 200px 0;
}
</style>
