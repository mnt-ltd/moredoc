<template>
  <el-menu-item
    :class="hiddenXS ? 'hidden-xs-only':''"
    :key="'nav-' + navigation.id"
    :index="`nav-${navigation.id}`"
  >
    <template v-if="navigation.href">
      <!-- 判断是否存在 https:// 或者 http:// -->
      <template
        v-if="
          navigation.href.indexOf('https://') === 0 ||
          navigation.href.indexOf('http://') === 0
        "
      >
        <a
          :href="navigation.href"
          :style="`color:${navigation.color}`"
          :target="navigation.target"
          class="el-link el-link--default"
          >{{ navigation.title }}</a
        >
      </template>
      <nuxt-link
        v-else
        :style="`color:${navigation.color}`"
        :to="navigation.href"
        :target="navigation.target"
        class="el-link el-link--default"
        >{{ navigation.title }}</nuxt-link
      >
    </template>
    <template v-else>
      <span :style="`color:${navigation.color}`">{{ navigation.title }}</span>
    </template>
  </el-menu-item>
</template>
<script>
export default {
  name: 'NavigationLink',
  props: {
    navigation: {
      type: Object,
      default: () => {},
    },
    hiddenXS: {
      type: Boolean,
      default: false,
    },
  },
}
</script>
