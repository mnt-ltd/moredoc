<template>
  <div class="menu-tree">
    <div v-for="item in items" :key="item.id" class="menu-tree__item">
      <div class="menu-tree__item__title">
        <span>{{ item.title }}</span>
        <span v-if="item.children" class="menu-tree__item__title__icon">
          <i class="fa fa-angle-down" aria-hidden="true"></i>
        </span>
      </div>
      <div v-if="item.children" class="menu-tree__item__children">
        <MenuTree :items="item.children" />
      </div>
    </div>
  </div>
</template>
<script>
import MenuTree from './MenuTree.vue'
export default {
  name: 'MenuTree',
  components: {
    MenuTree,
  },
  props: {
    items: {
      type: Array,
      required: true,
    },
  },
}
</script>
