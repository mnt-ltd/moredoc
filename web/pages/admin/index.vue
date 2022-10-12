<template>
  <div></div>
</template>

<script>
export default {
  name: 'AdminIndex',
  layout: 'admin',
  created() {
    this.$router.push('/admin/dashboard')
  },
}
</script>
