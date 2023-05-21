export default {
  data() {
    return {
      isMobile: false,
      isPad: false,
      isPC: true,
    }
  },
  mounted() {
    this.handleScreenResize()
    window.addEventListener('resize', this.handleScreenResize)
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.handleScreenResize)
  },
  methods: {
    handleScreenResize() {
      const width = document.body.clientWidth
      if (width < 768) {
        this.isMobile = true
        this.isPad = false
        this.isPC = false
      } else if (width < 992) {
        this.isMobile = false
        this.isPad = true
        this.isPC = false
      } else {
        this.isMobile = false
        this.isPad = false
        this.isPC = true
      }
    },
  },
}
