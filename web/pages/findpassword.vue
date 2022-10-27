<template>
  <div class="page-index">
    {{ $route.name }}
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
