<template>
  <div class="page page-article">
    <el-row :gutter="20">
      <el-col :span="18">
        <el-card shadow="never">
          <div slot="header"><h1>关于我们</h1></div>
          <div class="help-block text-muted article-info">
            <span><i class="el-icon-user"></i> 文库之家</span>
            <span><i class="el-icon-view"></i> 32 阅读</span>
            <span class="float-right"
              ><i class="el-icon-time"></i> 2022-01-02</span
            >
          </div>
          <article class="mgt-20px">
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
            <p>文章内容123254234234234</p>
          </article>
        </el-card>
      </el-col>
      <el-col :span="6" class="article-list">
        <el-card shadow="never">
          <div slot="header">相关链接</div>
          <nuxt-link to="/" class="el-link el-link--default"
            >关于我们</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >联系我们</nuxt-link
          >
          <nuxt-link to="/" class="el-link el-link--default"
            >免责声明</nuxt-link
          >
          <!-- 如果文章数不多，不显示。这里根据最后更新时间排序 -->
          <el-pagination
            class="mgt-20px"
            :current-page="1"
            :page-size="10"
            layout="total, prev, next"
            :total="400"
          >
          </el-pagination>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
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
.page-article {
  h1 {
    margin: 0;
    font-size: 16px;
  }
  .article-info {
    span {
      margin-right: 10px;
    }
  }
  .article-list {
    .el-card__body {
      padding-top: 10px;
    }
    a {
      display: block;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      height: 40px;
      line-height: 40px;
      border-bottom: 1px dashed #efefef;
    }
  }
}
</style>
