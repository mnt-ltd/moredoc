<template>
  <div class="com-document-simple-list">
    <ul>
      <li v-for="doc in docs" :key="'doc-' + doc.id">
        <nuxt-link
          :target="target"
          :to="`/document/${doc.id}`"
          class="el-link el-link--default"
        >
          <img
            :src="'/static/images/' + getIcon(doc.ext) + '_24.png'"
            :alt="getIcon(doc.ext) + '文档'"
          />
          {{ doc.title }}
        </nuxt-link>
      </li>
    </ul>
  </div>
</template>

<script>
import { getIcon } from '~/utils/utils'
export default {
  name: 'DocumentSimpleList',
  props: {
    docs: {
      type: Array,
      default: () => [],
    },
    target: {
      type: String,
      default: '',
    },
  },
  data() {
    return {}
  },
  async created() {},
  methods: {
    getIcon,
  },
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
