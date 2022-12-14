<template>
  <div class="com-document-simple-list">
    <ul>
      <li v-for="doc in docs" :key="'doc-' + doc.id">
        <nuxt-link to="/document/" class="el-link el-link--default">
          <img :src="'/static/images/' + doc.type + '_24.png'" alt="" />
          {{ doc.title }}
        </nuxt-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'DocumentSimpleList',
  props: {
    docs: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {}
  },
  head() {
    return {
      title: 'MOREDOC · 魔豆文库，开源文库系统',
    }
  },
  async created() {},
  methods: {},
}
</script>
<style lang="scss">
.com-document-simple-list {
  ul,
  li {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  li {
    a {
      display: block;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      line-height: 40px;
      height: 40px;
      border-bottom: 1px dashed #efefef;
      img {
        height: 18px;
        position: relative;
        top: 3px;
        margin-right: 5px;
      }
    }
  }
}
</style>
