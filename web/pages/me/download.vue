<template>
  <div>
    <user-download :user-id="user.id" />
  </div>
</template>
<script>
export default {
  props: {
    user: {
      type: Object,
      default: () => ({}),
    },
  },
}
</script>
