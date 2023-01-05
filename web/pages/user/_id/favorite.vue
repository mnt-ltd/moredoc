<template>
  <div>
    <user-favorite :user-id="userId" />
  </div>
</template>

<script>
import UserFavorite from '~/components/UserFavorite.vue'
import { mapGetters } from 'vuex'
export default {
  components: { UserFavorite },
  props: {
    user: {
      type: Object,
      default: () => ({}),
    },
  },
  head() {
    return {
      title: `[个人主页]${this.user.username} · 收藏 - ${this.settings.system.sitename}`,
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: `${this.user.username},文档收藏,${this.settings.system.sitename}`,
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
