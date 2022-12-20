<template>
  <div>
    <user-dynamic :user-id="userId" />
  </div>
</template>
<script>
export default {
  data() {
    return {
      userId: parseInt(this.$route.params.id) || 0,
    }
  },
}
</script>
