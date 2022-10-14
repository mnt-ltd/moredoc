<template>
  <div>
    <div>{{ $route.name }}</div>
  </div>
</template>

<script>
export default {
  layout: 'admin',
  head() {
    return {
      title: `面板 - MOREDOC · 魔刀文库`,
    }
  },
}
</script>
