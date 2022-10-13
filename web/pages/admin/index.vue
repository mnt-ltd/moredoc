<template>
  <div>{{ $route.name }}</div>
</template>

<script>
export default {
  layout: 'admin',
  created() {
    this.$router.push({ name: 'admin-dashboard' })
  },
}
</script>
