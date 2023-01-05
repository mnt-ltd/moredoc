<template>
  <div>
    <user-document :user-id="userId" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  props: {
    user: {
      type: Object,
      default: () => ({}),
    },
  },
  head() {
    return {
      title: `[个人主页]${this.user.username} · 文档 - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: `${this.user.username},文档,${this.settings.system.sitename}`,
        },
        {
          hid: 'keywords',
          name: 'keywords',
          content: `${this.settings.system.description}`,
        },
      ],
    }
  },
  data() {
    return {
      userId: parseInt(this.$route.params.id) || 0,
    }
  },
  computed: {
    ...mapGetters('setting', ['settings']),
  },
}
</script>

<style></style>
