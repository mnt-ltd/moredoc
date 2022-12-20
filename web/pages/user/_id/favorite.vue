<template>
  <div>
    <user-favorite :user-id="userId" />
  </div>
</template>

<script>
import UserFavorite from '~/components/UserFavorite.vue'
export default {
  components: { UserFavorite },
  data() {
    return {
      userId: parseInt(this.$route.params.id) || 0,
    }
  },
}
</script>

<style></style>
