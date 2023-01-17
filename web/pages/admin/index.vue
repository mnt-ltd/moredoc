<template>
  <div></div>
</template>

<script>
export default {
  layout: 'admin',
  created() {
    this.$router.push({ name: 'admin-dashboard' })
    return
  },
}
</script>
