<template>
  <div>
    <user-favorite :user-id="user.id" />
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
