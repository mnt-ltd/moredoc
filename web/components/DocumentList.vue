<template>
  <div class="com-document-list">
    <ul>
      <li v-for="doc in documents" :key="'doc-' + doc.id">
        <el-row :gutter="20">
          <el-col :span="4" class="doc-cover">
            <nuxt-link :to="`/document/${doc.id}`">
              <document-cover :document="doc" />
            </nuxt-link>
          </el-col>
          <el-col :span="20">
            <h3>
              <nuxt-link
                :to="`/document/${doc.id}`"
                class="el-link el-link--primary"
                ><img
                  :src="`/static/images/${getIcon(doc.ext)}_24.png`"
                  :alt="`${getIcon(doc.ext)} 文档`"
                />
                {{ doc.title }}</nuxt-link
              >
            </h3>
            <div class="doc-info">
              <el-rate
                v-model="doc.score"
                disabled
                show-score
                text-color="#ff9900"
                score-template="{value}"
              >
              </el-rate>
              <span class="float-right">
                {{ doc.price || 0 }}
                {{ settings.system.credit_name || '魔豆' }}
                <span class="line">|</span> {{ doc.pages || '-' }} 页
                <span class="line">|</span>
                {{ formatBytes(doc.size) || '-' }}
                <span class="hidden-xs-only"
                  ><span class="line">|</span>
                  {{ formatDatetime(doc.created_at) }}</span
                ></span
              >
            </div>
            <div class="doc-desc">
              {{ doc.description }}
            </div>
          </el-col>
        </el-row>
      </li>
    </ul>
  </div>
</template>
<script>
import DocumentCover from './DocumentCover.vue'
import { formatBytes, formatDatetime, getIcon } from '~/utils/utils'
import { mapGetters } from 'vuex'
export default {
  name: 'DocumentList',
  components: { DocumentCover },
  props: {
    documents: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {}
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
  methods: {
    formatBytes,
    formatDatetime,
    getIcon,
  },
}
</script>
<style lang="scss" scoped>
.com-document-list {
  ul,
  li {
    list-style: none;
    padding: 0;
    margin: 0;
  }
  li {
    margin-bottom: 20px;
    padding-bottom: 18px;
    border-bottom: 1px dashed #ddd;
  }
  h3 {
    margin: 0;
    a {
      font-size: 18px;
      font-weight: 400;
      display: inline-block;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      max-width: 100%;
      img {
        height: 18px;
        position: relative;
        top: 3px;
      }
    }
  }
  .doc-cover {
    .el-image {
      width: 100%;
      border: 2px solid #efefef;
      border-radius: 5px;
      &:hover img {
        transform: scale(1.2);
      }
    }
  }
  .doc-info {
    margin: 10px 0;
    font-size: 13px;
    color: #888;
    .line {
      color: #ddd;
      margin: 0 5px;
    }
  }
  .doc-desc {
    font-size: 14px;
    color: #999;
    line-height: 180%;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 3;
    max-height: 78px;
    -webkit-box-orient: vertical;
  }
}
</style>
<style lang="scss">
.com-document-list {
  .el-image {
    img {
      transition: transform 0.3s ease 0s;
      &:hover {
        transform: scale(1.2);
      }
    }
  }
}
</style>
